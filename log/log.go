//  Copyright (C) 晓白齐齐,版权所有.

package log

// 日志模块
import (
	"fmt"
	"io"
	"os"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

//TODO 添加calldepth的测试样例

type LogLevelType = int32

const (
	LogNone LogLevelType = iota
	LogDebug
	LogInfo
	LogWarn
	LogError
	LogCritical
)

const (
	LDate         = 1 << iota // 输出当地的时区的日期
	LTime                     // 输出当地的时区的时间
	LMicroseconds             // 输出当地时间精确到微秒
	LNanosceonds              // 输出当前时间戳的纳秒数
	LLongFile                 // 输出日志所在长文件名称
	LShortFile                // 输出日志所在段文件名称
	LUTC                      // 输出当地的UTC时间
	LTag                      // 将tag输出
	LPreTag                   // 将tag输出到最前面
	LLevel                    // 输出日志level登记
	LStdFlags     = LDate | LTime
)

type Log struct {
	mu        sync.Mutex
	tag       string
	level     LogLevelType
	flag      int
	out       io.Writer
	calldepth int
	buf       []byte
	isDiscard int32
}

func New(tag string, level LogLevelType, flag int, out io.Writer) *Log {
	l := &Log{tag: tag, level: level, flag: flag, out: out}
	l.isDiscard = 0
	if out == io.Discard {
		l.isDiscard = 1
	}
	return l
}

func (l *Log) SetWriter(w io.Writer) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.out = w
	l.isDiscard = 0
	if w == io.Discard {
		l.isDiscard = 1
	}
}

func (l *Log) Writer() io.Writer {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.out
}

func (l *Log) SetTag(tag string) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.tag = tag
}

func (l *Log) Tag() string {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.tag
}

func (l *Log) SetCallDepth(calldepth int) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.calldepth = calldepth
}

func (l *Log) CallDepth() int {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.calldepth
}

func (l *Log) SetLevel(level LogLevelType) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.level = level
}

func (l *Log) Level() LogLevelType {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.level
}

func (l *Log) SetFlag(flag int) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.flag = flag
}

func (l *Log) Flag() int {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.flag
}

func itoa(buf *[]byte, i int, wid int) {
	var b [20]byte
	bp := len(b) - 1
	for i >= 10 || wid > 1 {
		wid--
		q := i / 10
		b[bp] = byte('0' + i - q*10)
		bp--
		i = q
	}
	b[bp] = byte('0' + i)
	*buf = append(*buf, b[bp:]...)
}

func i64toa(buf *[]byte, i int64, wid int) {
	var b [20]byte
	bp := len(b) - 1
	for i >= 10 || wid > 1 {
		wid--
		q := i / 10
		b[bp] = byte('0' + i - q*10)
		bp--
		i = q
	}
	b[bp] = byte('0' + i)
	*buf = append(*buf, b[bp:]...)
}

func (l *Log) formatHeader(buf *[]byte, t time.Time, file string, line int, level LogLevelType) {
	if l.flag&LPreTag != 0 {
		*buf = append(*buf, l.tag...)
		*buf = append(*buf, ' ')
	}
	if l.flag&(LDate|LTime|LMicroseconds|LNanosceonds|LUTC) != 0 {
		if l.flag&LUTC != 0 {
			t = t.UTC()
		}
		if l.flag&LDate != 0 {
			year, month, day := t.Date()
			itoa(buf, year, 4)
			*buf = append(*buf, '/')
			itoa(buf, int(month), 2)
			*buf = append(*buf, '/')
			itoa(buf, day, 2)
			if l.flag&(LTime|LMicroseconds) != 0 {
				*buf = append(*buf, '-')
			} else {
				*buf = append(*buf, ' ')
			}
		}
		if l.flag&(LTime|LMicroseconds) != 0 {
			hour, min, sec := t.Clock()
			itoa(buf, hour, 2)
			*buf = append(*buf, ':')
			itoa(buf, min, 2)
			*buf = append(*buf, ':')
			itoa(buf, sec, 2)
			if l.flag&LMicroseconds != 0 {
				*buf = append(*buf, '.')
				itoa(buf, t.Nanosecond()/1e3, 6)
			}
			*buf = append(*buf, ' ')
		}
		if l.flag&LNanosceonds != 0 {
			*buf = append(*buf, '(')
			i64toa(buf, t.UnixNano(), 9)
			*buf = append(*buf, ')', ' ')
		}
	}
	if level != LogNone && l.flag&LLevel != 0 {
		switch level {
		case LogDebug:
			*buf = append(*buf, "[Debug]"...)
		case LogInfo:
			*buf = append(*buf, "[Info]"...)
		case LogWarn:
			*buf = append(*buf, "[Warn]"...)
		case LogError:
			*buf = append(*buf, "[Error]"...)
		case LogCritical:
			*buf = append(*buf, "[Critical]"...)
		}
		*buf = append(*buf, ' ')
	}
	if l.flag&(LLongFile|LShortFile) != 0 {
		if l.flag&LShortFile != 0 && l.flag&LLongFile == 0 {
			short := file
			for i := len(file) - 1; i > 0; i-- {
				if file[i] == '/' {
					short = file[i+1:]
					break
				}
			}
			file = short
		}
		*buf = append(*buf, file...)
		*buf = append(*buf, ':')
		itoa(buf, line, -1)
		*buf = append(*buf, ':', ' ')
	}
	if l.flag&LTag != 0 && l.flag&LPreTag == 0 {
		*buf = append(*buf, l.tag...)
		*buf = append(*buf, ' ')
	}
}

func (l *Log) Output(calldepth int, level LogLevelType, s string) error {
	var file string
	var line int
	now := time.Now()
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.flag&(LLongFile|LShortFile) != 0 {
		l.mu.Unlock()
		var ok bool
		_, file, line, ok = runtime.Caller(calldepth)
		if !ok {
			file = "???"
			line = 0
		}
		l.mu.Lock()
	}
	l.buf = l.buf[:0]
	l.formatHeader(&l.buf, now, file, line, level)
	l.buf = append(l.buf, s...)
	if len(s) == 0 { // } || s[len(s) - 1] != '\n' {
		l.buf = append(l.buf, '\n')
	}
	_, err := l.out.Write(l.buf)
	return err
}

func (l *Log) Printf(format string, v ...any) {
	if atomic.LoadInt32(&l.isDiscard) != 0 {
		return
	}
	l.Output(2+l.calldepth, l.level, fmt.Sprintf(format+"\n", v...))
}

func (l *Log) Print(v ...any) {
	if atomic.LoadInt32(&l.isDiscard) != 0 {
		return
	}
	l.Output(2+l.calldepth, l.level, fmt.Sprint(v...))
}

func (l *Log) Println(v ...any) {
	if atomic.LoadInt32(&l.isDiscard) != 0 {
		return
	}
	l.Output(2+l.calldepth, l.level, fmt.Sprintln(v...))
}

func (l *Log) Fatalf(format string, v ...any) {
	l.Output(2+l.calldepth, l.level, fmt.Sprintf(format+"\n", v...))
	os.Exit(1)
}

func (l *Log) Fatal(v ...any) {
	l.Output(2+l.calldepth, l.level, fmt.Sprint(v...))
	os.Exit(1)
}

func (l *Log) Fatalln(v ...any) {
	l.Output(2+l.calldepth, l.level, fmt.Sprintln(v...))
	os.Exit(1)
}

func (l *Log) Panicf(format string, v ...any) {
	s := fmt.Sprintf(format+"\n", v...)
	l.Output(2+l.calldepth, l.level, s)
	panic(s)
}

func (l *Log) Panic(v ...any) {
	s := fmt.Sprint(v...)
	l.Output(2+l.calldepth, l.level, s)
	panic(s)
}

func (l *Log) Panicln(v ...any) {
	s := fmt.Sprintln(v...)
	l.Output(2+l.calldepth, l.level, s)
	panic(s)
}

func (l *Log) Debugf(format string, v ...any) {
	printfByLevel(l, LogDebug, format, v...)
}

func (l *Log) Debug(v ...any) {
	printByLevel(l, LogDebug, v...)
}

func (l *Log) Debugln(v ...any) {
	printlnByLevel(l, LogDebug, v...)
}

func (l *Log) Infof(format string, v ...any) {
	printfByLevel(l, LogInfo, format, v...)
}

func (l *Log) Info(v ...any) {
	printByLevel(l, LogInfo, v...)
}

func (l *Log) Infoln(v ...any) {
	printlnByLevel(l, LogInfo, v...)
}

func (l *Log) Warnf(format string, v ...any) {
	printfByLevel(l, LogWarn, format, v...)
}

func (l *Log) Warn(v ...any) {
	printByLevel(l, LogWarn, v...)
}

func (l *Log) Warnln(v ...any) {
	printlnByLevel(l, LogWarn, v...)
}

func (l *Log) Errorf(format string, v ...any) {
	printfByLevel(l, LogError, format, v...)
}

func (l *Log) Error(v ...any) {
	printByLevel(l, LogError, v...)
}

func (l *Log) Errorln(v ...any) {
	printlnByLevel(l, LogError, v...)
}

func (l *Log) Criticalf(format string, v ...any) {
	printfByLevel(l, LogCritical, format, v...)
}

func (l *Log) Critical(v ...any) {
	printByLevel(l, LogCritical, v...)
}

func (l *Log) Criticalln(v ...any) {
	printlnByLevel(l, LogCritical, v...)
}

func (l *Log) PrintfIgnoreLevel(level LogLevelType, format string, v ...any) {
	if atomic.LoadInt32(&l.isDiscard) != 0 {
		return
	}
	l.Output(2+l.calldepth, level, fmt.Sprintf(format+"\n", v...))
}

func (l *Log) PrintIgnoreLevel(level LogLevelType, v ...any) {
	if atomic.LoadInt32(&l.isDiscard) != 0 {
		return
	}
	l.Output(2+l.calldepth, level, fmt.Sprint(v...))
}

func (l *Log) PrintlnIgnoreLevel(level LogLevelType, v ...any) {
	if atomic.LoadInt32(&l.isDiscard) != 0 {
		return
	}
	l.Output(2+l.calldepth, level, fmt.Sprintln(v...))
}

var std = New("", LogNone, LStdFlags, os.Stderr)

func Default() *Log { return std }

func SetWriter(w io.Writer) {
	std.SetWriter(w)
}
func Writer() io.Writer {
	return std.Writer()
}

func SetTag(tag string) {
	std.SetTag(tag)
}
func Tag() string {
	return std.Tag()
}

func SetLevel(level LogLevelType) {
	std.SetLevel(level)
}
func Level() LogLevelType {
	return std.Level()
}

func SetFlag(flag int) {
	std.SetFlag(flag)
}
func Flag() int {
	return std.Flag()
}

func SetCallDepth(calldepth int) {
	std.SetCallDepth(calldepth)
}

func CallDepth() int {
	return std.CallDepth()
}

func Output(calldepth int, level LogLevelType, s string) error {
	return std.Output(calldepth+1, level, s)
}

func Printf(format string, v ...any) {
	if atomic.LoadInt32(&std.isDiscard) != 0 {
		return
	}
	std.Output(2+std.calldepth, std.level, fmt.Sprintf(format+"\n", v...))
}

func Print(v ...any) {
	if atomic.LoadInt32(&std.isDiscard) != 0 {
		return
	}
	std.Output(2+std.calldepth, std.level, fmt.Sprint(v...))
}

func Println(v ...any) {
	if atomic.LoadInt32(&std.isDiscard) != 0 {
		return
	}
	std.Output(2+std.calldepth, std.level, fmt.Sprintln(v...))
}

func Fatalf(format string, v ...any) {
	std.Output(2+std.calldepth, std.level, fmt.Sprintf(format+"\n", v...))
	os.Exit(1)
}

func Fatal(v ...any) {
	std.Output(2+std.calldepth, std.level, fmt.Sprint(v...))
	os.Exit(1)
}

func Fatalln(v ...any) {
	std.Output(2+std.calldepth, std.level, fmt.Sprintln(v...))
	os.Exit(1)
}

func Panicf(format string, v ...any) {
	s := fmt.Sprintf(format+"\n", v...)
	std.Output(2+std.calldepth, std.level, s)
	panic(s)
}

func Panic(v ...any) {
	s := fmt.Sprint(v...)
	std.Output(2+std.calldepth, std.level, s)
	panic(s)
}

func Panicln(v ...any) {
	s := fmt.Sprintln(v...)
	std.Output(2+std.calldepth, std.level, s)
	panic(s)
}

func Debugf(format string, v ...any) {
	printfByLevel(std, LogDebug, format, v...)
}

func Debug(v ...any) {
	printByLevel(std, LogDebug, v...)
}

func Debugln(v ...any) {
	printlnByLevel(std, LogDebug, v...)
}

func Infof(format string, v ...any) {
	printfByLevel(std, LogInfo, format, v...)
}

func Info(v ...any) {
	printByLevel(std, LogInfo, v...)
}

func Infoln(v ...any) {
	printlnByLevel(std, LogInfo, v...)
}

func Warnf(format string, v ...any) {
	printfByLevel(std, LogWarn, format, v...)
}

func Warn(v ...any) {
	printByLevel(std, LogWarn, v...)
}

func Warnln(v ...any) {
	printlnByLevel(std, LogWarn, v...)
}

func Errorf(format string, v ...any) {
	printfByLevel(std, LogError, format, v...)
}

func Error(v ...any) {
	printByLevel(std, LogError, v...)
}

func Errorln(v ...any) {
	printlnByLevel(std, LogError, v...)
}

func Criticalf(format string, v ...any) {
	printfByLevel(std, LogCritical, format, v...)
}

func Critical(v ...any) {
	printByLevel(std, LogCritical, v...)
}

func Criticalln(v ...any) {
	printlnByLevel(std, LogCritical, v...)
}

func PrintfIgnoreLevel(level LogLevelType, format string, v ...any) {
	if atomic.LoadInt32(&std.isDiscard) != 0 {
		return
	}
	std.Output(2+std.calldepth, level, fmt.Sprintf(format+"\n", v...))
}

func PrintIgnoreLevel(level LogLevelType, v ...any) {
	if atomic.LoadInt32(&std.isDiscard) != 0 {
		return
	}
	std.Output(2+std.calldepth, level, fmt.Sprint(v...))
}

func PrintlnIgnoreLevel(level LogLevelType, v ...any) {
	if atomic.LoadInt32(&std.isDiscard) != 0 {
		return
	}
	std.Output(2+std.calldepth, level, fmt.Sprintln(v...))
}

func printfByLevel(l *Log, level LogLevelType, format string, v ...any) {
	if atomic.LoadInt32(&l.isDiscard) != 0 {
		return
	}
	if atomic.LoadInt32(&l.level) <= level {
		l.Output(3+l.calldepth, level, fmt.Sprintf(format+"\n", v...))
	}
}

func printByLevel(l *Log, level LogLevelType, v ...any) {
	if atomic.LoadInt32(&l.isDiscard) != 0 {
		return
	}
	if atomic.LoadInt32(&l.level) <= level {
		l.Output(3+l.calldepth, level, fmt.Sprint(v...))
	}
}

func printlnByLevel(l *Log, level LogLevelType, v ...any) {
	if atomic.LoadInt32(&l.isDiscard) != 0 {
		return
	}
	if atomic.LoadInt32(&l.level) <= level {
		l.Output(3+l.calldepth, level, fmt.Sprintln(v...))
	}
}
