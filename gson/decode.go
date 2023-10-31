//  Copyright (C) 晓白齐齐,版权所有.

package gson

import (
	"fmt"
	"reflect"
	"sync"
)

type targetInfo struct {
	isKey         bool          // 是否根据下标确定值
	exclusive     bool          // 当前目标是否具有排他性
	isUnmarshaler bool          // 是否是unmarshaler对象
	unmarshaler   Unmarshaler   // unmarshaler对象，如果isUnmarshaler为false，该值无效
	anyValue      reflect.Value // 其他任何值对象，如果isUnmarshaler为true，该值无效
}

var targetInfoPool = sync.Pool{
	New: func() any {
		return &targetInfo{}
	},
}

func newTargetInfo(isKey, exclusive, isUnmarshaler bool, unmarshaler Unmarshaler, anyValue reflect.Value) *targetInfo {
	info := targetInfoPool.Get().(*targetInfo)
	//TODO 要解决的bug
	if info == nil {
		for i := 0; i < 10; i++ {
			info = targetInfoPool.Get().(*targetInfo)
			if info != nil {
				break
			}
		}
	}
	info.set(isKey, exclusive, isUnmarshaler, unmarshaler, anyValue)
	return info
}

func freeTargetInfo(info *targetInfo) {
	targetInfoPool.Put(info)
}

func (t *targetInfo) reset() {
	t.isKey = false
	t.exclusive = false
	t.isUnmarshaler = false
	t.unmarshaler = nil
	t.anyValue = reflect.ValueOf(nil)
}

func (t *targetInfo) set(isKey, exclusive, isUnmarshaler bool, unmarshaler Unmarshaler, anyValue reflect.Value) {
	t.isUnmarshaler = isUnmarshaler
	t.isKey = isKey
	t.exclusive = exclusive
	t.unmarshaler = unmarshaler
	t.anyValue = anyValue
}

type parseInfo struct {
	currentIndex int    // 已解析的下标
	currentKey   string // 当前解析的
}

var parseInfoPool = sync.Pool{
	New: func() any {
		return &parseInfo{}
	},
}

func newParseInfo(currentIndex int, currentKey string) *parseInfo {
	info := parseInfoPool.Get().(*parseInfo)
	// TODO 这里当两次new太快时，就会导致获取到一个空的info
	if info == nil {
		for i := 0; i < 10; i++ {
			info = parseInfoPool.Get().(*parseInfo)
			if info != nil {
				break
			}
		}
	}
	info.currentIndex = currentIndex
	info.currentKey = currentKey
	return info
}

func freeParseInfo(info *parseInfo) {
	parseInfoPool.Put(info)
}

func (p *parseInfo) reset() {
	p.currentIndex = -1
	p.currentKey = ""
}

// 键事件，发现一个键或者结束一个键
// 参数：所在的Decoder
//
//	记录键值的Lexer
//	如果为true表示发现一个键，如果为false表示结束一个键
type KeyHandler func(*Decoder, *Lexer, bool) bool

// 值事件，发现一个值
// 参数：如果当前的值是一组数值的话，表示第几个值，如果新发现一个键，则会重新从0开始，-1表示当前是以键的形式回调
//
//	如果当前的值是一个对象的话，返回当前找到的键
//	所在的Decoder
//	保存值数据的Lexer
type ValueHandler func(int, string, *Decoder, *Lexer) bool

// 对象事件，发现一个对象的开头，或者是对象的结尾
// 参数：键
//
//	所在的Decoder
//	如果为true表示发现了对象的开头，如果为false表示发现了对象的结尾
type ObjectHandler func(string, *Decoder, bool) bool
type Decoder struct {
	scan             *scanner   // 扫描器
	lexer            *Lexer     // 保存字符流
	currentParseInfo *parseInfo // 当前的解析信息
	currentTarget    *targetInfo
	err              error        // 错误
	parseStack       []*parseInfo // 解析信息栈
	targetStack      []*targetInfo
	KeyEventCall     KeyHandler    // 键事件的回调，发现或者结束一个键
	ValueEventCall   ValueHandler  // 值事件的回调
	ObjectEventCall  ObjectHandler // 对象事件的回调，发现一个对象开头或对象结尾
	tempKey          string        // 临时记录的解析的键
}

var decoderPool = sync.Pool{
	New: func() any {
		return &Decoder{}
	},
}

func NewDecoder() *Decoder {
	l := NewLexer()
	s := newScanner()
	d := decoderPool.Get().(*Decoder)
	l.Reset()
	s.reset(true)
	d.Reset()
	s.lexer = l
	s.dataFoundHandler = d.dataFoundHandler
	d.scan = s
	d.lexer = l
	return d
}

func FreeDecoder(d *Decoder) {
	freeScanner(d.scan)
	FreeLexer(d.lexer)
	freeParseInfo(d.currentParseInfo)
	freeTargetInfo(d.currentTarget)
	for _, info := range d.parseStack {
		freeParseInfo(info)
	}
	for _, info := range d.targetStack {
		freeTargetInfo(info)
	}
	d.parseStack = nil
	d.targetStack = nil
	decoderPool.Put(d)
}

func (d *Decoder) Reset() {
	if d.scan != nil {
		d.scan.reset(false)
	}
	if d.lexer != nil {
		d.lexer.Reset()
	}
	d.err = nil
	if d.currentParseInfo != nil {
		d.currentParseInfo.reset()
	}
	if d.currentTarget != nil {
		d.currentTarget.reset()
	}
	d.parseStack = d.parseStack[0:0]
	d.targetStack = d.targetStack[0:0]
	d.KeyEventCall = nil
	d.ValueEventCall = nil
	d.ObjectEventCall = nil
	d.tempKey = ""
}

const maxDecodeNestingDepth = 10000 // 目标栈最大的嵌套深度，一个结构体（map）嵌套另一个结构体（map），最多只能嵌套的层数
func (d *Decoder) pushTarget(target *targetInfo) {
	if d.currentTarget != nil {
		d.targetStack = append(d.targetStack, d.currentTarget)
	}
	if len(d.targetStack) > maxDecodeNestingDepth {
		d.decodeError("nest-depth of targetInfo is too large")
	}
	d.currentTarget = target
}

func (d *Decoder) popTarget() *targetInfo {
	ret := d.currentTarget
	n := len(d.targetStack)
	if n == 0 {
		d.currentTarget = nil
		return ret
	}
	d.currentTarget = d.targetStack[n-1]
	n = n - 1
	d.targetStack = d.targetStack[0:n]
	return ret
}

func (d *Decoder) pushParseInfo(parse *parseInfo) {
	if d.currentParseInfo != nil {
		d.parseStack = append(d.parseStack, d.currentParseInfo)
	}
	if len(d.parseStack) > maxDecodeNestingDepth {
		d.decodeError("nest-depth of parseInfo is too large")
	}
	d.currentParseInfo = parse
}

func (d *Decoder) popParseInfo() *parseInfo {
	ret := d.currentParseInfo
	n := len(d.parseStack)
	if n == 0 {
		d.currentParseInfo = nil
		return ret
	}
	d.currentParseInfo = d.parseStack[n-1]
	n = n - 1
	d.parseStack = d.parseStack[0:n]
	return ret
}

func (d *Decoder) isExclusive() bool {
	if d.currentTarget != nil {
		return d.currentTarget.exclusive
	}
	return false
}

func (d *Decoder) printDecoder(tag string) {
	fmt.Println(tag, "is::", "parseinfo:", d.currentParseInfo,
		"; parseStack:", d.parseStack,
		"; currentTarget:", d.currentTarget,
		"; targetStack:", d.targetStack,
		"; tempKey:", d.tempKey,
		"; err:", d.err, "; ")
	if d.currentTarget != nil {
		fmt.Println("currentTarget exclusive", d.currentTarget.exclusive)
	} else {
		fmt.Println("currentTarget is nil")
	}
}

func (d *Decoder) decodeError(context string) {
	currentK := d.tempKey
	if currentK == "" && d.currentParseInfo != nil {
		currentK = d.currentParseInfo.currentKey
	}
	d.err = &DecodeError{currentK, context}
}

// 在这里进行解析
// 找到了一个键、值等的回调
func (d *Decoder) dataFoundHandler(dataKind int, l *Lexer, line int64, offset int64) (error, bool) {
	if d.err != nil {
		return d.err, false
	}
	ret := true
	switch dataKind {
	case dataKey: // 键
		ret = d.findKey(l)
	case dataValue: // 值
		ret = d.findValue(l)
	case dataObjBegin: // 对象开始
		ret = d.findObjBegin()
	case dataObjEnd: // 对象结束
		ret = d.findObjEnd()
	}
	return d.err, ret
}

func (d *Decoder) unmarshalOrConvert(l *Lexer) {
	if d.currentTarget != nil {
		if d.currentTarget.isUnmarshaler {
			if d.currentTarget.isKey {
				d.err = d.currentTarget.unmarshaler.UnmarshalByKey(d.currentParseInfo.currentKey, l)
			} else {
				d.err = d.currentTarget.unmarshaler.UnmarshalByIndex(d.currentParseInfo.currentIndex, l)
			}
		} else {
			key := ""
			if d.currentTarget.isKey {
				key = d.currentParseInfo.currentKey
			}
			d.err = convertValue(d.currentTarget.anyValue, key, d.currentParseInfo.currentIndex, l, d, true, nil)
		}
	}
}

func (d *Decoder) findObjBegin() bool {
	if d.currentParseInfo == nil { // 不应该出现的bug，所以不用添加测试样例
		d.decodeError("there must be a parseInfo while find a value, but it is nil")
		return true
	}
	isNewKey := false
	if d.tempKey != "" {
		d.currentParseInfo.currentKey = d.tempKey
		d.tempKey = ""
		isNewKey = true
	}
	d.currentParseInfo.currentIndex++
	isCurrentTargetNil := false
	if d.currentTarget != nil {
		d.unmarshalOrConvert(nil)
	} else {
		isCurrentTargetNil = true
	}
	if isNewKey && !d.isExclusive() && d.KeyEventCall != nil {
		newKeyL := NewLexer()
		newKeyL.Write([]byte(d.currentParseInfo.currentKey))
		defer FreeLexer(newKeyL)
		if !d.KeyEventCall(d, newKeyL, true) {
			return false
		}
	}
	if isCurrentTargetNil && d.currentTarget != nil && !d.currentTarget.isKey {
		d.unmarshalOrConvert(nil)
	}
	if d.err == nil && !d.isExclusive() && d.ObjectEventCall != nil && !d.ObjectEventCall(d.currentParseInfo.currentKey, d, true) {
		return false
	}
	d.pushParseInfo(newParseInfo(-1, ""))
	return true
}

func (d *Decoder) findObjEnd() bool {
	if !d.keyDoneHandler() {
		return false
	}
	if d.err == nil && !d.isExclusive() && d.ObjectEventCall != nil && !d.ObjectEventCall(d.currentParseInfo.currentKey, d, false) {
		return false
	}
	if d.currentTarget != nil && d.currentTarget.isKey {
		d.popParseInfo()
		oldTarget := d.popTarget()
		if oldTarget != nil && d.currentTarget != nil && !d.currentTarget.isUnmarshaler {
			key := ""
			if d.currentTarget != nil && d.currentTarget.isKey {
				key = d.currentParseInfo.currentKey
			}
			d.err = convertValue(d.currentTarget.anyValue, key, d.currentParseInfo.currentIndex, nil, d, true, oldTarget)
		}
	} else {
		d.popParseInfo()
	}
	return true
}

// 查找到一个值的回调
func (d *Decoder) findKey(l *Lexer) bool {
	if !d.keyDoneHandler() {
		return false
	}
	// 如果currentParseInfo不为nil，重置，否则push一个ParseInfo
	if d.currentParseInfo != nil {
		d.currentParseInfo.reset()
	} else {
		d.pushParseInfo(newParseInfo(-1, ""))
	}
	d.tempKey = l.String()
	return true
}

// 一个键结束了的回调，当发现一个新的键，或者到了解析结束的时候，就要调用这个函数
func (d *Decoder) keyDoneHandler() bool {
	// 将相关的target pop出来
	var oldTarget *targetInfo
	for {
		if d.currentTarget != nil && !d.currentTarget.isKey {
			oldTarget = d.popTarget()
		} else {
			break
		}
	}
	// 回调结束键事件
	if d.currentParseInfo != nil && d.currentParseInfo.currentKey != "" {
		if oldTarget != nil && d.currentTarget != nil && !d.currentTarget.isUnmarshaler {
			key := ""
			if d.currentTarget != nil && d.currentTarget.isKey {
				key = d.currentParseInfo.currentKey
			}
			d.err = convertValue(d.currentTarget.anyValue, key, d.currentParseInfo.currentIndex, nil, d, true, oldTarget)
		}
		if d.err == nil && !d.isExclusive() && d.KeyEventCall != nil {
			oldKeyL := NewLexer()
			oldKeyL.Write([]byte(d.currentParseInfo.currentKey))
			defer FreeLexer(oldKeyL)
			if !d.KeyEventCall(d, oldKeyL, false) {
				return false
			}
		}
	}
	return true
}

// 查找到一个值
func (d *Decoder) findValue(l *Lexer) bool {
	if d.currentParseInfo == nil { // 不应该出现的bug，所以不用添加测试样例
		d.decodeError("there must be a parseInfo while find a value, but it is nil")
		return true
	}
	// 如果临时的键不为空，回调发现键事件
	if d.tempKey != "" {
		d.currentParseInfo.currentKey = d.tempKey
		d.tempKey = ""
		if !d.isExclusive() && d.KeyEventCall != nil {
			newKeyL := NewLexer()
			newKeyL.Write([]byte(d.currentParseInfo.currentKey))
			defer FreeLexer(newKeyL)
			if !d.KeyEventCall(d, newKeyL, true) {
				return false
			}
		}
	}
	d.currentParseInfo.currentIndex++
	d.unmarshalOrConvert(l)
	if d.err == nil && !d.isExclusive() && d.ValueEventCall != nil && !d.ValueEventCall(d.currentParseInfo.currentIndex, d.currentParseInfo.currentKey, d, l) {
		return false
	}
	return true
}

func convertValue(value reflect.Value, key string, index int, l *Lexer, d *Decoder, autoSub bool, oldTarget *targetInfo) error {
	elemKind := value.Kind()
	valElem := value
	if elemKind != reflect.Pointer {
		if !autoSub {
			return fmt.Errorf("gson can't convert a non-pointer type: %s", elemKind)
		}
	} else {
		valElem = value.Elem()
		elemKind = valElem.Kind()
	}
	if elemKind == reflect.Pointer {
		ok := false
		if valElem, elemKind, ok = getElemAndKind(valElem, elemKind); !ok {
			return fmt.Errorf("gson can't convert a Pointer-Type to a nil-value")
		}
	}
	// 以下是对空指针New一个值得方法，但是有问题
	// for {
	// 	if elemKind == reflect.Pointer {
	// 		elemType := valElem.Type()                // *bool
	// 		subElemType := elemType.Elem()  // 指针指向的类型            // bool
	// 		// subElemKind := subElemType.Kind()
	// 		if valElem.IsNil() {
	// 			subElemNewVal := reflect.New(subElemType)  // *bool
	// 			// elemNewVal := reflect.New(elemType)         // **bool
	// 			valElem.Set(subElemNewVal)
	// 			valElem = valElem.Elem()
	// 			elemKind = valElem.Kind()
	// 			// break
	// 		}
	// 	} else {
	// 		break
	// 	}
	// }

	if key != "" {
		if elemKind == reflect.Array || elemKind == reflect.Slice || isBasicKind(elemKind) {
			return fmt.Errorf("a %s-type can't has a key", elemKind)
		}
	} else {
		if elemKind == reflect.Struct || elemKind == reflect.Map {
			return fmt.Errorf("a %s-type must has keys, not found a key", elemKind)
		}
	}
	if isBasicKind(elemKind) {
		if index > 0 {
			return fmt.Errorf("too many field got for a %s-type, a %s-type can only has a field", elemKind, elemKind)
		}
		if l == nil {
			return fmt.Errorf("a %s-type can't be a {", elemKind)
		}
		return convertBasic(valElem, l)
	}
	switch elemKind {
	case reflect.Array, reflect.Slice:
		return convertArray(valElem, l, index, d, autoSub, oldTarget)
	case reflect.Struct:
		return convertStruct(valElem, l, key, d, autoSub, oldTarget)
	case reflect.Map:
		return convertMap(valElem, l, key, d, autoSub, oldTarget)
	}
	return nil
}

func isBasicKind(k reflect.Kind) bool {
	return k == reflect.Bool || k == reflect.Int || k == reflect.Int8 || k == reflect.Int16 || k == reflect.Int32 || k == reflect.Int64 ||
		k == reflect.Uint || k == reflect.Uint8 || k == reflect.Uint16 || k == reflect.Uint32 || k == reflect.Uint64 || k == reflect.Float32 ||
		k == reflect.Float64 || k == reflect.String
}

func convertBasic(value reflect.Value, l *Lexer) error {
	var err error
	var v any
	if value.CanSet() {
		switch value.Type().Kind() {
		case reflect.Bool:
			v, err = l.Bool()
		case reflect.Int:
			v, err = l.Int()
		case reflect.Int8:
			v, err = l.Int8()
		case reflect.Int16:
			v, err = l.Int16()
		case reflect.Int32:
			v, err = l.Int32()
		case reflect.Int64:
			v, err = l.Int64()
		case reflect.Uint:
			v, err = l.Uint()
		case reflect.Uint8:
			v, err = l.Uint8()
		case reflect.Uint16:
			v, err = l.Uint16()
		case reflect.Uint32:
			v, err = l.Uint32()
		case reflect.Uint64:
			v, err = l.Uint64()
		case reflect.Float32:
			v, err = l.Float32()
		case reflect.Float64:
			v, err = l.Float64()
		case reflect.String:
			v = l.String()
		}
		if err == nil {
			value.Set(reflect.ValueOf(v))
		}
		return err
	}
	return nil
}

func convertArray(value reflect.Value, l *Lexer, index int, d *Decoder, autoSub bool, oldTarget *targetInfo) error {
	if oldTarget != nil {
		return nil
	}
	if !value.CanSet() {
		return nil
	}
	len := value.Len()
	cap := value.Cap()
	elemKind := value.Kind()
	if cap <= index { // 元素的容量比当前索引要小，扩容，扩到原来的容量的2倍
		if elemKind == reflect.Array {
			return fmt.Errorf("array-index of value is out of range, array len is %d, index is %d", len, index)
		}
		newCap := cap + cap/2
		if newCap < 4 {
			newCap = 4
		}
		newSlice := reflect.MakeSlice(value.Type(), len, newCap)
		reflect.Copy(newSlice, value)
		value.Set(newSlice)
	}
	if elemKind == reflect.Slice {
		subElemType := value.Type().Elem()
		subElemKind := subElemType.Kind()
		if subElemKind == reflect.Map {
			subElemValue := reflect.MakeMap(subElemType)
			value.Set(reflect.Append(value, subElemValue))
		} else {
			value.SetLen(index + 1)
		}
	}
	subValue := value.Index(index)
	subKind := subValue.Kind()
	if subKind == reflect.Pointer {
		ok := false
		if subValue, subKind, ok = getElemAndKind(subValue, subKind); !ok {
			return fmt.Errorf("gson can't convert a Pointer-Type to a nil-value in a %s, index is %d", value.Type(), index)
		}
	}
	if isBasicKind(subKind) {
		return convertBasic(subValue, l)
	} else {
		if autoSub {
			switch subKind {
			case reflect.Struct, reflect.Map:
				if l != nil {
					return fmt.Errorf("a %s-type must be begin with a {", subKind)
				}
				d.pushTarget(newTargetInfo(true, d.currentTarget.exclusive, false, nil, subValue))
				return nil
			}
		}
		return fmt.Errorf("gson can't convert a arry/slice its elem is %s", subKind)
	}
	return nil
}

func convertStruct(value reflect.Value, l *Lexer, key string, d *Decoder, autoSub bool, oldTarget *targetInfo) error {
	if oldTarget != nil {
		return nil
	}
	valType := value.Type()
	num := value.NumField()
	var val reflect.Value
	for i := 0; i < num; i++ {
		tf := valType.Field(i)
		if k := tf.Tag.Get("gson"); k == key {
			val = value.Field(i)
			break
		}
	}
	if !val.IsValid() {
		if _, ok := valType.FieldByName(key); ok {
			val = value.FieldByName(key)
		}
	}
	if val.IsValid() {
		fieldKind := val.Type().Kind()
		if fieldKind == reflect.Pointer {
			ok := false
			if val, fieldKind, ok = getElemAndKind(val, fieldKind); !ok {
				return fmt.Errorf("gson can't convert a Pointer-Type to a nil-value in a %s, key is %s", val.Type(), key)
			}
		}
		if isBasicKind(fieldKind) {
			if l == nil {
				return fmt.Errorf("a %s-type can't be a {", fieldKind)
			}
			return convertBasic(val, l)
		} else {
			if autoSub {
				switch fieldKind {
				case reflect.Struct, reflect.Map:
					if l != nil {
						return fmt.Errorf("a %s-type must be begin with a {", fieldKind)
					}
					d.pushTarget(newTargetInfo(true, d.currentTarget.exclusive, false, nil, val))
					return nil
				case reflect.Array, reflect.Slice:
					d.pushTarget(newTargetInfo(false, d.currentTarget.exclusive, false, nil, val))
					d.currentParseInfo.currentIndex--
					d.findValue(l)
					return nil
				}
			}
			return fmt.Errorf("gson can't convert a struct its elem is %s", fieldKind)
		}
	}
	return nil
}

func convertMap(value reflect.Value, l *Lexer, key string, d *Decoder, autoSub bool, oldTarget *targetInfo) error {
	// 判断键的类型是否合法
	valType := value.Type()
	elemType := valType.Elem()
	keyType := valType.Key()
	keyKind := keyType.Kind()
	if !isBasicKind(keyKind) {
		return fmt.Errorf("a map's key can only be Int, Bool, Float, Uint, String, can't be %s", keyKind)
	}
	// 创建key的值
	keyValue := reflect.New(keyType)
	keyElem := keyValue.Elem()
	if keyKind != reflect.String {
		tmpl := NewLexer()
		tmpl.Write([]byte(key))
		if err := convertBasic(keyElem, tmpl); err != nil {
			return fmt.Errorf("convert key to map error: %s", err)
		}
	} else {
		keyElem.SetString(key)
	}
	// 创建值
	elemValue := reflect.New(elemType)
	valueElem := elemValue.Elem()
	elemKind := valueElem.Kind()
	if elemKind == reflect.Pointer {
		ok := false
		if valueElem, elemKind, ok = getElemAndKind(valueElem, elemKind); !ok {
			return fmt.Errorf("gson can't convert a data which elem is a Pointer-Type to a nil-value in a %s, key is %s", valueElem.Type(), key)
		}
	}
	if value.IsNil() {
		tmpValue := reflect.MakeMap(valType)
		value.Set(tmpValue)
	}
	if isBasicKind(elemKind) {
		if err := convertBasic(valueElem, l); err != nil {
			return err
		}
	} else {
		if autoSub {
			switch elemKind {
			case reflect.Struct, reflect.Map, reflect.Array, reflect.Slice:
				if oldTarget != nil {
					if valueElem.CanSet() {
						if !oldTarget.isUnmarshaler {
							valueElem.Set(oldTarget.anyValue)
						}
					}
					value.SetMapIndex(keyValue.Elem(), elemValue.Elem())
					return nil
				} else {
					if elemKind == reflect.Struct || elemKind == reflect.Map {
						if l != nil {
							return fmt.Errorf("a %s-type must be begin with a {", elemKind)
						}
						d.pushTarget(newTargetInfo(true, d.currentTarget.exclusive, false, nil, valueElem))
					} else {
						d.pushTarget(newTargetInfo(false, d.currentTarget.exclusive, false, nil, valueElem))
						d.currentParseInfo.currentIndex--
						d.findValue(l)
					}
				}
				return nil
			}
		}
		return fmt.Errorf("gson can't convert a map its elem is %s", elemKind)
	}
	value.SetMapIndex(keyValue.Elem(), elemValue.Elem())
	return nil
}

func getElemAndKind(val reflect.Value, kind reflect.Kind) (reflect.Value, reflect.Kind, bool) {
	for {
		if kind == reflect.Pointer {
			if val.IsNil() {
				return val, kind, false
			}
			val = val.Elem()
			kind = val.Kind()
		} else {
			return val, kind, true
		}
	}
	return val, kind, false
}

// 实现：根据key的值，将一个l转为target中的某个字段的值
func ConvertByKey(target any, key string, l *Lexer) error {
	if err := checkTarget(target); err != nil {
		return err
	}
	if tmpshaler, ok := target.(Unmarshaler); ok {
		return tmpshaler.UnmarshalByKey(key, l)
	}
	return convertValue(reflect.ValueOf(target), key, -1, l, nil, false, nil)
}

// 实现：根据index下标，将一个l转为target中的某个字段的值
func ConvertByIndex(target any, index int, l *Lexer) error {
	if err := checkTarget(target); err != nil {
		return err
	}
	if tmpshaler, ok := target.(Unmarshaler); ok {
		return tmpshaler.UnmarshalByIndex(index, l)
	}
	return convertValue(reflect.ValueOf(target), "", index, l, nil, false, nil)
}

// 设置一个Unmarshaler作为转换目标
func (d *Decoder) SetUnmarshaler(u Unmarshaler, exclusive bool) error {
	return d.setUnmarshaler(u, exclusive, false)
}

func (d *Decoder) AddSubUnmarshaler(u Unmarshaler) error {
	return d.setUnmarshaler(u, d.currentTarget.exclusive, true)
}

func (d *Decoder) setUnmarshaler(u Unmarshaler, exclusive, isSub bool) error {
	if !isSub && d.currentTarget != nil {
		d.decodeError("a target has already set, can't set another target twice")
		return d.err
	}
	if d.currentParseInfo == nil {
		if isSub {
			d.decodeError("a sub-unmarshaler must has a parent target, but there is no a parent target")
			return d.err
		}
		d.pushParseInfo(newParseInfo(-1, ""))
	}
	d.pushTarget(newTargetInfo(true, false, exclusive, u, reflect.ValueOf(nil)))
	return nil
}

// 设置一个任意值（指针类型的）作为转换目标
func (d *Decoder) SetAnyTarget(v any, exclusive bool) error {
	return d.setAny(v, exclusive, false)
}

func (d *Decoder) AddSubTarget(v any) error {
	return d.setAny(v, d.currentTarget.exclusive, true)
}

func (d *Decoder) setAny(v any, exclusive, isSub bool) error {
	if err := checkTarget(v); err != nil {
		d.err = err
		return err
	}
	if tmpshaler, ok := v.(Unmarshaler); ok {
		return d.setUnmarshaler(tmpshaler, exclusive, isSub)
	}
	if !isSub && d.currentTarget != nil {
		d.decodeError("a target has already set, can't set another target twice")
		return d.err
	}
	if isSub && d.currentTarget == nil {
		d.decodeError("a parent target must be set while AddSubTarget, but there is no a parent target")
		return d.err
	}
	if d.currentParseInfo == nil {
		d.pushParseInfo(newParseInfo(-1, ""))
	}
	isKey := false
	t := reflect.TypeOf(v).Elem().Kind()
	if t == reflect.Struct || t == reflect.Map {
		isKey = true
	}
	d.pushTarget(newTargetInfo(isKey, exclusive, false, nil, reflect.ValueOf(v)))
	return nil
}

func finishDecoder(d *Decoder) bool {
	if d != nil && !d.keyDoneHandler() {
		return false
	}
	d.popParseInfo()
	if d.currentParseInfo != nil || (d.parseStack != nil && len(d.parseStack) > 0) {
		d.decodeError(fmt.Sprintf("parseStack must be empty and currentParseInfo must be nil while finishDecoder, currentParseInfo is %v, parseStack is %v", d.currentParseInfo, d.parseStack))
		return true
	}
	if d.currentTarget != nil || (d.targetStack != nil && len(d.targetStack) > 0) {
		d.decodeError(fmt.Sprintf("targetStack must be empty and currentTarget must be nil while finishDecoder, currentTarget is %v, targetStack is %v", d.currentTarget, d.targetStack))
		return true
	}
	return true
}

// 检查目标转换目标是否合格
func checkTarget(v any) error {
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Pointer || rv.IsNil() {
		return &InvalidUnmarshalError{reflect.TypeOf(v)}
	}
	return nil
}
func UnmarshalAny(data []byte, v any) error {
	if err := checkTarget(v); err != nil {
		return err
	}
	decoder := NewDecoder()
	defer FreeDecoder(decoder)
	decoder.SetAnyTarget(v, true)
	return decodeAndFinish(data, decoder)
}

func Unmarshal(data []byte, u Unmarshaler) error {
	decoder := NewDecoder()
	defer FreeDecoder(decoder)
	decoder.SetUnmarshaler(u, true)
	return decodeAndFinish(data, decoder)
}

func DecodeData(data []byte, keyEventCall KeyHandler, valueEventCall ValueHandler, objEventCall ObjectHandler) error {
	decoder := NewDecoder()
	defer FreeDecoder(decoder)
	decoder.KeyEventCall = keyEventCall
	decoder.ValueEventCall = valueEventCall
	decoder.ObjectEventCall = objEventCall
	return decodeAndFinish(data, decoder)
}

func decodeAndFinish(data []byte, d *Decoder) error {
	if err := work(data, d.scan); err == nil {
		finishDecoder(d)
		if d.err != nil {
			return d.err
		}
	} else {
		return err
	}
	return nil
}
