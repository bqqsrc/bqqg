//  Copyright (C) 晓白齐齐,版权所有.

package gson

import (
	"strconv"
	"sync"
)

func Valid(data []byte) bool {
	scan := newScanner()
	defer freeScanner(scan)
	return work(data, scan) == nil
}

func work(data []byte, scan *scanner) error {
	scan.reset(false)
	scan.data = data
	if scan.lexer == nil {
		lexer := NewLexer()
		defer FreeLexer(lexer)
		scan.lexer = lexer
	}
	for _, c := range data {
		if scan.step(scan, c) == scanError {
			return scan.err
		}

		scan.bytes++
		// 先判断再过行，因为回调会给出具体的行数和偏移量，如果先过行再判断，可能出现给出的行数是不准确的
		if isLine(c) {
			scan.line++
			scan.lineOffset = 1
		} else {
			scan.lineOffset++
		}
	}
	if eof(scan) == scanError {
		return scan.err
	}
	return nil
}

// 找到了键、值、标签的回调处理
// 参数：数据种类
//     词义解析器
//     数组结尾所在行数
//     数组结尾相对所在行数的偏移量
// 返回参数：错误（没有错误时为nil）
//         是否继续遍历（返回true时继续遍历）
type foundHandler func(int, *Lexer, int64, int64) (error, bool)

// 解析的状态
type parseKind byte

// 状态栈的状态，gson需要成对：强制字符串；对象；存在花括号的占位符；存在花括号的标签替换符号，只有这4种情况才需要压入栈，且gson有可能存在解析栈为空的情况
const (
	parseStateEmpty parseKind = iota // 解析栈为空
	parseString                      // 解析强制字符串中
	parseObject                      // 解析字典中
	parseHolder                      // 解析占位符
	parseReplace                     // 解析标签替换符号
)

// 找到的数据种类
const (
	dataNone     = iota // 没有任何值
	dataKey             // 只有键
	dataValue           // 只有值
	dataVar             // 只有变量
	dataObjBegin        // 只有对象起始
	dataObjEnd          // 只有对象结束
)

const maxNestingDepth = 10000 // 解析栈最大数量，即最大的嵌套深度
const posNotFound = -1

const (
	scanNone             = iota // 无状态
	scanContinue                // 继续扫描
	scanSkipSpace               // 跳过空白字符
	scanBeginObject             // 开始对象
	scanEndObject               // 结束对象
	scanBeginForceString        // 强制字符串起始
	scanEndForceString          // 强制字符串结束
	scanBeginComment            // 注释起始
	scanEndComment              // 注释结束
	scanBeginLiteral            // 按字面意思继续扫描
	scanEscape                  // 转义字符
	scanKey                     // 一个键
	scanVar                     // 一个变量
	scanEndValue                // 结束一个值
	scanBeginHolder             // 开始占位符
	scanEndHolder               // 结束占位符
	scanBeginReplace            // 开始标签替换符
	scanEndReplace              // 结束标签替换符
	scanManualEnd               // 回调返回结束符
	scanEnd                     // 扫描结束
	scanError                   // 扫描失败
)

type scanner struct {
	data             []byte
	parseState       []parseKind              // 解析的状态栈
	step             func(*scanner, byte) int // 下一步扫描的处理函数
	dataFoundHandler foundHandler             // 找到了键、值、标签、占位符、标签替换符、变量、对象起始和终止点的回调
	// 全局扫描标志记录
	line       int64 // 扫描到第几行
	bytes      int64 // 扫描到第几个字符
	lineOffset int64 // 扫描到该行字符的偏移量（即第几行的第几个字符）
	// 当前扫描的数据的标志位记录
	posBegin int64 // 当前数据的起始下标
	posEnd   int64 // 当前数据的结束下标
	// 上一个数据的行号，偏移量
	lastBytes      int64 // 上一个扫描到的数据的结尾所在偏移量
	lastLine       int64 // 上一个扫描到的数据的尾字符所在的行数
	lastLineOffset int64 // 上一个扫描到的数据的尾字符在其行的第几个字符
	// 强制字符串相关
	forceStrBeginNum int   // 强制字符串的起始构造字符数量
	forceStrEndNum   int   // 强制字符串的结尾构造字符数量
	err              error // 错误
	// 与数据解析相关的成员
	lexer *Lexer
}

var scannerPool = sync.Pool{
	New: func() any {
		return &scanner{}
	},
}

func newScanner() *scanner {
	scan := scannerPool.Get().(*scanner)
	scan.reset(true)
	return scan
}

func freeScanner(scan *scanner) {
	scan.data = nil
	if len(scan.parseState) > 1024 {
		scan.parseState = nil
	}
	scannerPool.Put(scan)
}

func (s *scanner) reset(resetHandler bool) {
	s.data = nil
	s.parseState = s.parseState[0:0]
	s.step = stateBeginValue
	if resetHandler {
		s.dataFoundHandler = nil
	}
	s.line = 1
	s.bytes = 0
	s.lineOffset = 1
	s.err = nil
	s.forceStrBeginNum = 0
	s.forceStrEndNum = 0
	s.posBegin = posNotFound
	s.posEnd = posNotFound
	s.lastBytes = posNotFound
	s.lastLine = posNotFound
	s.lastLineOffset = posNotFound
	if s.lexer != nil {
		s.lexer.Reset()
	}
}

// 确定一个数据的具体类型后的重置操作
func (s *scanner) resetAfterHandler() {
	s.forceStrBeginNum = 0
	s.forceStrEndNum = 0
	s.posBegin = posNotFound
	s.posEnd = posNotFound
	s.lastBytes = posNotFound
	s.lastLine = posNotFound
	s.lastLineOffset = posNotFound
	s.lexer.Reset()

	s.posBegin = s.bytes
	s.lexer.SetValueType(stringV)
}

func (s *scanner) pushParseState(c byte, newParseState parseKind, successState int) int {
	s.parseState = append(s.parseState, newParseState)
	if len(s.parseState) <= maxNestingDepth {
		return successState
	}
	return s.error(c, "exceeded max depth")
}

func (s *scanner) popParseState() parseKind {
	n := len(s.parseState)
	if n == 0 {
		return parseStateEmpty
	}
	ret := s.parseState[n-1]
	n = n - 1
	s.parseState = s.parseState[0:n]
	return ret
}

func (s *scanner) error(c byte, context string) int {
	s.step = stateError
	s.err = &SyntaxError{"invalid character " + quoteChar(c) + ", " + context, s.line, s.lineOffset}
	return scanError
}

//TODO Test
func (s *scanner) handlerError(err error) int {
	s.step = stateError
	s.err = &SyntaxError{"parse error character " + quoteChar(s.data[s.lastBytes]) + ", " + err.Error(), s.lastLine, s.lastLineOffset}
	return scanError
}

func hasFlushData(s *scanner) bool { // 是否有未写入缓存的数据
	return s.posBegin != posNotFound && s.posEnd != posNotFound && s.posEnd >= s.posBegin
}

func hasData(s *scanner) bool { // 是否有缓存数据
	return hasFlushData(s) || s.lexer.Len() > 0
}

func foundHandlerCallback(s *scanner, dataKind int) int {
	if hasFlushData(s) {
		s.lexer.Write(s.data[s.posBegin : s.posEnd+1])
	}
	if hasData(s) || dataKind == dataObjBegin || dataKind == dataObjEnd {
		if s.dataFoundHandler != nil {
			err, ret := s.dataFoundHandler(dataKind, s.lexer, s.lastLine, s.lastLineOffset)
			if err != nil {
				return s.handlerError(err)
			}
			if !ret {
				return scanManualEnd
			}
		}
	}
	s.resetAfterHandler()
	return scanNone
}

// 开始解析任何值
func stateBeginValue(s *scanner, c byte) int {
	if isSpace(c) {
		return scanSkipSpace
	}
	// 下面这句注释了的代码添加了一个'='判断，这是为了后面新增占位符等功能而做的，当前是1.0.0版，不把=当作构造字符，而是将其当成一个普通字符来处理
	if hasData(s) && c != ':' && c != '#' && (version <= 100 || c != '=') {
		// :按键处理
		// =按变量处理
		// #按注释处理，即不处理，等待下一个有效值再处理
		// !、$、{、}按值处理即其他字符按值处理
		if scanS := foundHandlerCallback(s, dataValue); scanS != scanNone {
			return scanS
		}
	} else if s.posBegin == posNotFound && c != '#' { // 由于注释是要跳过的，所以注释符号不能记录它的下标
		s.posBegin = s.bytes
	}
	// 当非空时，对特殊字符串做出特殊处理，其余的都不变，继续扫描
	switch c {
	case '{': // 一个字典的开始
		return stateBeginObject(s, c)
	case '`': // 一个强制字符串的开始
		s.step = stateBeginForceString
		return stateBeginForceString(s, c)
	case '#': // 一个注释的开始
		s.step = stateBeginComment
		return scanBeginComment
	case '%': // 一个转义字符
		s.step = stateEscape
		return scanEscape
	case '-': // 一个负数的开始
		s.lexer.SetValueType(intV)
		s.step = stateNeg
		return scanBeginLiteral
	case '0': // 一个小数点的开始
		s.lexer.SetValueType(uintV)
		s.step = state0
		return scanBeginLiteral
	case 't': // 一个true的开始
		s.step = stateT
		return scanBeginLiteral
	case 'f': // 一个false的开始
		s.step = stateF
		return scanBeginLiteral
	case 'n': // 一个nil的开始
		s.step = stateN
		return scanBeginLiteral
	case ':': // 前面的值是一个键
		return stateBeginKey(s, c)
	// 这是一个变量定义的构造字符，当前是第1.0.0版本，这个功能先去除，后面版本再新增
	case '=': // 前面的值是一个变量
		if version > 100 {
			return stateBeginVar(s, c)
		} else {
			break
		}
	case '}': // 一个字典结束了
		return stateEndObject(s, c)
	}

	if '1' <= c && c <= '9' { // 一个数字的开始
		s.step = state1
		s.lexer.SetValueType(uintV)
		return scanBeginLiteral
	}
	// 其他的字符，进入字符串处理函数，返回继续扫描状态
	s.step = stateInString
	return scanContinue
}

// 将当前值往前n位记录位有效值的结束位置
func setPrePos(s *scanner, n int64) {
	s.posEnd = s.bytes - n
	s.lastBytes = s.bytes - n
	s.lastLine = s.line // 由于正常都是先处理字节再判断是否过行，因此处理数据时的行即有效值的行
	s.lastLineOffset = s.lineOffset - n
}

// 解析字符串中
func stateInString(s *scanner, c byte) int {
	if c != '{' && c != '}' { // 假如是一个字典的开始或者结束，有可能是一段非普通字符串（如数值字符串）然后马上紧邻{或者}，此时是不能将它的数值类型设置位字符的
		s.lexer.SetValueType(stringV)
	}
	if isSpace(c) {
		return stateEndValue(s, c)
	}
	switch c {
	case '%': // 转义字符
		s.step = stateEscape
		return scanEscape
	case '{':
		setPrePos(s, 1) // 因为这里的情况是{紧邻了字符串，所以要往前1位记录结束下标，后面的相同的道理
		if stateEndValue(s, c) == scanError {
			return scanError
		}
		return stateBeginObject(s, c)
	case ':':
		setPrePos(s, 1)
		return stateBeginKey(s, c)
	// 变量构造字符在1.0.0版本不添加，等到后面新版本再添加
	case '=':
		if version > 100 {
			setPrePos(s, 1)
			return stateBeginVar(s, c)
		} else {
			break
		}
	case '#':
		setPrePos(s, 1)
		s.step = stateBeginComment
		return scanBeginComment
	case '}':
		setPrePos(s, 1)
		if stateEndValue(s, c) == scanError {
			return scanError
		}
		return stateEndObject(s, c)
	}
	s.step = stateInString
	return scanContinue
}

// 有效转义字符后将转义字符前的数据推入缓存
func flushEscapeData(s *scanner) {
	s.posEnd = s.bytes - 2
	if hasFlushData(s) { // }
		s.lexer.Write(s.data[s.posBegin : s.posEnd+1])
	}
	s.posBegin = s.bytes
	s.posEnd = posNotFound
}

// 解析转义字符
func stateEscape(s *scanner, c byte) int {
	// printMessage("stateEscape", s, c)
	if isSpace(c) { // 如果转义字符后面一位直接是空白符，那么转义字符直接也当作普通字符计入有效数据中
		setPrePos(s, 1)
		return stateEndValue(s, c)
	}
	// 转义字符后一位为非空字符，转义字符有效，往前2位记录结束点，然后将数据推入缓存
	// 重新以当前位置位初始位置，继续扫描数据
	flushEscapeData(s)
	s.step = stateInString
	return scanContinue
}

// 解析强制字符串内的转义字符
func stateForceEscape(s *scanner, c byte) int {
	// printMessage("stateForceEscape", s, c)
	if c == '`' { // 强制字符串里的转义字符只对`生效，所以转义字符后如果不是`，是不处理的
		flushEscapeData(s)
		s.step = stateInForceString
		return scanContinue
	}
	return stateInForceString(s, c)
}

// 解析一个键
func stateBeginKey(s *scanner, c byte) int {
	if !hasData(s) { // 如果在一个:之前没有缓存里没有数值，那么这个:就找不到对应的key，是异常的，var也一样
		return s.error(c, "an error key-define without a key-name")
	}
	if scanS := foundHandlerCallback(s, dataKey); scanS != scanNone {
		return scanS
	}
	s.posBegin = posNotFound // 解析完一个键后，所有有效数据都清空了，所以s.posBegin不能是当前位置，得设置为未找到
	s.step = stateBeginValue
	// printMessage("after stateBeginKey", s, c)
	return scanKey
}

// 解析一个常量
func stateBeginVar(s *scanner, c byte) int {
	if !hasData(s) {
		return s.error(c, "an error var-define without a var-name")
	}
	if scanS := foundHandlerCallback(s, dataVar); scanS != scanNone {
		return scanS
	}
	s.posBegin = posNotFound
	s.step = stateBeginValue
	return scanVar
}

// 解析一个强制字符
func stateBeginForceString(s *scanner, c byte) int {
	if c != '`' { // 如果不是`，那么进入有效字符了
		s.posBegin = s.bytes
		if c == '%' { // 如果是%，开始转义，这是强制字符串首字符就为%得情况
			s.step = stateForceEscape
		} else {
			s.step = stateInForceString
		}
		return s.pushParseState(c, parseString, scanBeginForceString)
	}
	s.forceStrBeginNum++
	if s.forceStrBeginNum > 3 {
		return s.error(c, "too many "+quoteChar(c)+" before a string")
	}
	return scanContinue
}

func stateInForceString(s *scanner, c byte) int {
	s.step = stateInForceString
	if c != '`' {
		s.forceStrEndNum = 0
	}
	switch c {
	case '%':
		s.step = stateForceEscape
		return scanEscape
	case '`':
		return stateEndForceString(s, c)
	}
	return scanContinue
}

func stateEndForceString(s *scanner, c byte) int {
	s.forceStrEndNum++
	if s.forceStrBeginNum == s.forceStrEndNum {
		s.step = stateFinishForceString
	}
	return scanContinue
}

func stateFinishForceString(s *scanner, c byte) int {
	if !isSpace(c) { // 如果结尾遇到了同样数量得`，还得`之后为空白符才算结束，非空白符得情况继续强制字符串中
		return stateInForceString(s, c)
	}
	setPrePos(s, int64(s.forceStrEndNum+1)) // 由于此时已经又多读了一个空白字符，加上结尾几个`，所以往前s.forceStrEndNum + 1位
	s.forceStrBeginNum = 0
	s.forceStrEndNum = 0
	if sta := s.popParseState(); sta != parseString {
		return s.error(c, `an error token, want a "string end character", got a `+tokenEndName(sta))
	}
	return stateEndValue(s, c)
}

// 解析注释
func stateBeginComment(s *scanner, c byte) int {
	if c != '`' { // 不带`的情况得注释，马上跳到进入注释状态中
		s.step = stateInComment
		if s.forceStrBeginNum > 0 {
			return s.pushParseState(c, parseString, scanBeginForceString)
		} else {
			return stateInComment(s, c)
		}
	}
	s.forceStrBeginNum++
	if s.forceStrBeginNum > 3 {
		return s.error(c, "too many "+quoteChar(c)+" before a string")
	}
	return scanContinue
}

func stateInComment(s *scanner, c byte) int {
	if s.forceStrBeginNum <= 0 && isLine(c) { // 不带`的情况下，换行就马上结束注释
		return stateFinishComment(s, c)
	}
	if c == '`' && s.forceStrBeginNum > 0 { // 带`的情况，发现`则尝试结束注释
		return stateEndComment(s, c)
	}
	if c != '`' {
		s.forceStrEndNum = 0
	}
	return scanContinue
}

func stateEndComment(s *scanner, c byte) int {
	s.forceStrEndNum++
	if s.forceStrBeginNum == s.forceStrEndNum {
		s.step = stateFinishComment
	}
	return scanContinue
}

func stateFinishComment(s *scanner, c byte) int {
	if s.forceStrBeginNum > 0 && c == '`' {
		s.step = stateInComment
		return scanContinue
	}
	if s.forceStrBeginNum > 0 {
		s.forceStrBeginNum = 0
		s.forceStrEndNum = 0
		if sta := s.popParseState(); sta != parseString {
			return s.error(c, `an error token, want a "string end character", got a `+tokenEndName(sta))
		}
	}
	s.step = stateBeginValue
	if !isSpace(c) { // 如果注释之后不是非空白字符，这个字符要进入读取数值阶段
		return stateBeginValue(s, c)
	}
	return scanEndComment
}

// 字典的开始和结束
func stateBeginObject(s *scanner, c byte) int {
	if hasData(s) { // 如果字典结束符号出现时，还有缓存，那么将这些缓存处理未设置
		if scanS := foundHandlerCallback(s, dataValue); scanS != scanNone {
			return scanS
		}
	}
	if scanS := foundHandlerCallback(s, dataObjBegin); scanS != scanNone {
		return scanS
	}
	s.step = stateBeginValue
	s.posBegin = posNotFound // 字典开始和结束后的有效值初始位置都要设置位未找到
	return s.pushParseState(c, parseObject, scanBeginObject)
}

func stateEndObject(s *scanner, c byte) int {
	sta := s.popParseState()
	if sta == parseStateEmpty {
		return s.error(c, `an error '}', not found its begin '{'`)
	}
	if sta != parseObject {
		return s.error(c, `an error token, want a "object end character", got a `+tokenEndName(sta))
	}
	if hasData(s) { // 如果字典结束符号出现时，还有缓存，那么将这些缓存处理未设置
		if scanS := foundHandlerCallback(s, dataValue); scanS != scanNone {
			return scanS
		}
	}
	if scanS := foundHandlerCallback(s, dataObjEnd); scanS != scanNone {
		return scanS
	}
	s.step = stateBeginValue
	s.posBegin = posNotFound
	return scanEndObject
}

// 解析负数
func stateNeg(s *scanner, c byte) int {
	if '0' <= c && c <= '9' {
		s.step = state1
		return scanContinue
	}
	return stateInString(s, c)
}

// 解析非小数
func state1(s *scanner, c byte) int {
	if '0' <= c && c <= '9' {
		if s.lexer.GetValueType() != intV {
			s.lexer.SetValueType(uintV)
		}
		return scanContinue
	}
	if isSpace(c) {
		return stateEndValue(s, c)
	}
	s.step = state0
	return state0(s, c)
}

// 解析小数
func state0(s *scanner, c byte) int {
	if c == '.' {
		s.lexer.SetValueType(floatV)
		s.step = stateDot
		return scanContinue
	}
	if c == 'e' || c == 'E' {
		s.lexer.SetValueType(floatV)
		s.step = stateE
		return scanContinue
	}
	if '0' <= c && c <= '9' {
		return scanContinue
	}
	if isSpace(c) {
		return stateEndValue(s, c)
	}
	return stateInString(s, c)
}

// 解析非小数后的科学计数法
func stateDot(s *scanner, c byte) int {
	if '0' <= c && c <= '9' {
		s.step = stateDot0
		return scanContinue
	}
	return stateInString(s, c)
}

// 解析小数点后的科学计数法
func stateDot0(s *scanner, c byte) int {
	if '0' <= c && c <= '9' {
		return scanContinue
	}
	if c == 'e' || c == 'E' {
		s.step = stateE
		s.lexer.SetValueType(floatV)
		return scanContinue
	}
	if isSpace(c) {
		return stateEndValue(s, c)
	}
	return stateInString(s, c)
}

// 解析科学计数法
func stateE(s *scanner, c byte) int {
	if c == '+' || c == '-' {
		s.step = stateESign
		return scanContinue
	}
	if '0' <= c && c <= '9' {
		s.step = stateE0
		return scanContinue
	}
	return stateInString(s, c)
}

// 解析科学计数法的结尾
func stateESign(s *scanner, c byte) int {
	if '0' <= c && c <= '9' {
		s.step = stateE0
		return scanContinue
	}
	return stateInString(s, c)
}

func stateE0(s *scanner, c byte) int {
	if '0' <= c && c <= '9' {
		return scanContinue
	}
	if isSpace(c) {
		return stateEndValue(s, c)
	}
	return stateInString(s, c)
}

// 解析true
func stateT(s *scanner, c byte) int {
	if c == 'r' {
		s.step = stateTr
		return scanContinue
	}
	return stateInString(s, c)
}

func stateTr(s *scanner, c byte) int {
	if c == 'u' {
		s.step = stateTru
		return scanContinue
	}
	return stateInString(s, c)
}

func stateTru(s *scanner, c byte) int {
	if c == 'e' {
		s.step = stateTrue
		return scanContinue
	}
	return stateInString(s, c)
}

func stateTrue(s *scanner, c byte) int {
	if isSpace(c) {
		s.lexer.SetValueType(boolV)
		return stateEndValue(s, c)
	}
	return stateInString(s, c)
}

// 解析false
func stateF(s *scanner, c byte) int {
	if c == 'a' {
		s.step = stateFa
		return scanContinue
	}
	return stateInString(s, c)
}

func stateFa(s *scanner, c byte) int {
	if c == 'l' {
		s.step = stateFal
		return scanContinue
	}
	return stateInString(s, c)
}

func stateFal(s *scanner, c byte) int {
	if c == 's' {
		s.step = stateFals
		return scanContinue
	}
	return stateInString(s, c)
}

func stateFals(s *scanner, c byte) int {
	if c == 'e' {
		s.step = stateFalse
		return scanContinue
	}
	return stateInString(s, c)
}

func stateFalse(s *scanner, c byte) int {
	if isSpace(c) {
		s.lexer.SetValueType(boolV)
		return stateEndValue(s, c)
	}
	return stateInString(s, c)
}

// 解析nil
func stateN(s *scanner, c byte) int {
	if c == 'i' {
		s.step = stateNi
		return scanContinue
	}
	return stateInString(s, c)
}

func stateNi(s *scanner, c byte) int {
	if c == 'l' {
		s.step = stateNil
		return scanContinue
	}
	return stateInString(s, c)
}

func stateNil(s *scanner, c byte) int {
	if isSpace(c) {
		s.lexer.SetValueType(nilV)
		return stateEndValue(s, c)
	}
	return stateInString(s, c)
}

// 结束一个值的解析，当遇到空白符的时候的状态
func stateEndValue(s *scanner, c byte) int {
	if s.posEnd == posNotFound {
		setPrePos(s, 1)
	}
	s.step = stateBeginValue // 继续读取一个值
	return scanEndValue
}

// 文件结束了的处理
func eof(s *scanner) int {
	// 文件结尾不是空白符时，posEnd来不及记录的情况
	if s.posBegin != posNotFound && s.posEnd == posNotFound {
		if s.forceStrBeginNum > 0 && s.forceStrBeginNum == s.forceStrEndNum { // 如果是强制字符，按强制字符串处理
			setPrePos(s, int64(s.forceStrBeginNum+1))
		} else {
			setPrePos(s, 1)
		}
	}
	if s.forceStrBeginNum > 0 && s.forceStrBeginNum == s.forceStrEndNum { // 如果最后是强制字符串，推出栈
		if sta := s.popParseState(); sta != parseString {
			return s.error(0, `an error token, want a "string end character", got a `+tokenEndName(sta))
		}
	}
	// 如果有缓存数据未处理，处理为值
	if hasFlushData(s) {
		if scanS := foundHandlerCallback(s, dataValue); scanS != scanNone {
			return scanS
		}
	}
	// 判断解析栈是否为空
	if sta := s.popParseState(); sta != parseStateEmpty {
		s.err = &SyntaxError{`unexpected end of gson input, want a ` + tokenEndName(sta) + `, but not found`, s.line, s.lineOffset}
		return scanError
	} else {
		return scanEnd
	}
}

func stateError(s *scanner, c byte) int {
	return scanError
}

func isBlank(c byte) bool {
	return c <= ' ' && (c == ' ' || c == '\t' || c == '\r')
}
func isLine(c byte) bool {
	// 这里解决windows系统的换行是\r\n的问题：解决方案，windows系统的换行为\r\n，这里面有2个字符，因此，将\r按照空白字符处理即可
	return c <= ' ' && c == '\n'
}

// 空格、换行、回车、制表符都是空白符
func isSpace(c byte) bool {
	return isBlank(c) || isLine(c)
}

func quoteChar(c byte) string {
	if c == '\'' {
		return `'\''`
	}
	if c == '"' {
		return `'"'`
	}
	s := strconv.Quote(string(c))
	return "'" + s[1:len(s)-1] + "'"
}

func tokenEndName(state parseKind) string {
	switch state {
	case parseString:
		return `"string end character"`
	case parseObject:
		return `"object end character"`
	case parseHolder:
		return `"placeholder end character"`
	case parseReplace:
		return `"replace symbol end character"`
	default:
		return ""
	}
}
