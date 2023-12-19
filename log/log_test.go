//  Copyright (C) 晓白齐齐,版权所有.

package log

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
	"time"
)

func Test_Seter_Geter(t *testing.T) {
	var b bytes.Buffer
	l := New("", LogNone, 0, nil)
	l.SetWriter(&b)
	l.SetLevel(LogError)
	l.SetFlag(LLongFile | LLevel)
	l.SetTag("Test")
	if got := l.Writer(); got != &b {
		t.Errorf("log Writer should match %q, got %q", &b, got)
	}
	if got := l.Tag(); got != "Test" {
		t.Errorf("log Tag should match Test, got %s", got)
	}
	if got := l.Level(); got != LogError {
		t.Errorf("log Level should match %d, got %d", LogError, got)
	}
	if got := l.Flag(); got != LLongFile|LLevel {
		t.Errorf("log Flag should match %d, got %d", LLongFile|LLevel, got)
	}
}

func Test_Output(t *testing.T) {
	const testString = "test"
	var b bytes.Buffer
	l := New("", LogDebug, 0, &b)
	l.Println(testString)
	if expected := testString + "\n"; b.String() != expected {
		t.Errorf("log output should match %q, got %q", expected, b.String())
	}
}

func Test_Default(t *testing.T) {
	if got := Default(); got != std {
		t.Errorf("Default() should be [%p], got [%p]", std, got)
	}
	if got := Writer(); got != os.Stderr {
		t.Errorf("Writer() should be [%p], got [%p]", os.Stderr, got)
	}
	if got := Tag(); got != "" {
		t.Errorf("log Tag should match empty string, got %s", got)
	}
	if got := Level(); got != LogNone {
		t.Errorf("log Level should match %d, got %d", LogNone, got)
	}
	if got := Flag(); got != LStdFlags {
		t.Errorf("log Flag should match %d, got %d", LStdFlags, got)
	}
	var b bytes.Buffer
	SetWriter(&b)
	defer SetWriter(os.Stderr)
	SetTag("TestDefault")
	SetLevel(LogDebug)
	SetFlag(LLevel)
	if got := Writer(); got != &b {
		t.Errorf("Writer() should be [%p], got [%p]", &b, got)
	}
	if got := Tag(); got != "TestDefault" {
		t.Errorf("log Tag should match TestDefault, got %s", got)
	}
	if got := Level(); got != LogDebug {
		t.Errorf("log Level should match %d, got %d", LogDebug, got)
	}
	if got := Flag(); got != LLevel {
		t.Errorf("log Flag should match %d, got %d", LLevel, got)
	}
	const testString = "Test a default log"
	err := Output(3, Level(), testString)
	if err != nil {
		t.Errorf("log Output err should be nil, got %s", err)
	}
	if expect := fmt.Sprintf("[Debug] %s", testString); expect != b.String() {
		t.Errorf("log Output expect %s, got %s", expect, b.String())
	}
}

type logFuncStruct struct {
	funcCommon func(...any)
	funcFormat func(string, ...any)
	funcLine   func(...any)
}
type ignoreFuncStruct struct {
	funcCommon func(LogLevelType, ...any)
	funcFormat func(LogLevelType, string, ...any)
	funcLine   func(LogLevelType, ...any)
}

const (
	RNoData          = `!!!NoneData`
	Rdate            = `[0-9][0-9][0-9][0-9]/[0-9][0-9]/[0-9][0-9]`                                                                // 匹配一个日期
	Rtime            = `[0-9][0-9]:[0-9][0-9]:[0-9][0-9]`                                                                          // 匹配一个时间
	Rmicroseconds    = `\.[0-9][0-9][0-9][0-9][0-9][0-9]`                                                                          // 匹配一个微妙
	Rnanoseconds     = `\([0-9][0-9][0-9][0-9][0-9][0-9][0-9][0-9][0-9][0-9][0-9][0-9][0-9][0-9][0-9][0-9][0-9][0-9][0-9][0-9]*\)` // 匹配一个纳秒
	Rline            = `(41|43|45):`                                                                                               // 匹配函数行号
	Rlongfile        = `.*/[A-Za-z0-9_\-]+\.go:` + Rline                                                                           // 匹配长文件名
	Rshortfile       = `[A-Za-z0-9_\-]+\.go:` + Rline                                                                              // 匹配短文件名
	RlineIgnore      = `(83|85|87):`                                                                                               // 匹配函数行号
	RlongfileIgnore  = `.*/[A-Za-z0-9_\-]+\.go:` + RlineIgnore                                                                     // 匹配长文件名
	RshortfileIgnore = `[A-Za-z0-9_\-]+\.go:` + RlineIgnore                                                                        // 匹配短文件名
)

type tester struct {
	flag         int
	tag          string
	logerLevel   LogLevelType
	contextLevel LogLevelType
	pattern      string
}

var tests = []tester{
	// 单个的flag测试
	{0, "", LogNone, LogNone, ""},
	{0, "XXX", LogNone, LogNone, ""},
	{LDate, "XXX", LogNone, LogNone, Rdate + " "},
	{LTime, "XXX", LogNone, LogNone, Rtime + " "},
	{LMicroseconds, "XXX", LogNone, LogNone, Rtime + Rmicroseconds + " "},
	{LNanosceonds, "XXX", LogNone, LogNone, Rnanoseconds + " "},
	{LLongFile, "XXX", LogNone, LogNone, Rlongfile + " "},
	{LShortFile, "XXX", LogNone, LogNone, Rshortfile + " "},
	{LTag, "XXX", LogNone, LogNone, "XXX "},
	{LPreTag, "XXX", LogNone, LogNone, "XXX "},
	{LStdFlags, "XXX", LogNone, LogNone, Rdate + "-" + Rtime + " "},
	{LLevel, "XXX", LogNone, LogNone, ""},
	{LLevel, "XXX", LogNone, LogDebug, `\[Debug\] `},
	// 特殊的2个组合flag的测试样例
	{LDate | LTime, "XXX", LogNone, LogNone, Rdate + "-" + Rtime + " "},
	{LTime | LMicroseconds, "XXX", LogNone, LogNone, Rtime + Rmicroseconds + " "},
	{LDate | LMicroseconds, "XXX", LogNone, LogNone, Rdate + "-" + Rtime + Rmicroseconds + " "},
	{LStdFlags | LMicroseconds, "XXX", LogNone, LogNone, Rdate + "-" + Rtime + Rmicroseconds + " "},
	{LLongFile | LShortFile, "XXX", LogNone, LogNone, Rlongfile + " "},
	{LTag | LPreTag, "XXX", LogNone, LogNone, "XXX "},
	{LLevel | LTag, "XXX", LogNone, LogWarn, `\[Warn\] XXX `},
	{LLevel | LPreTag, "XXX", LogNone, LogError, `XXX \[Error\] `},
	// 特殊的3个以上组合flag的测试样例
	{LDate | LTime | LMicroseconds, "XXX", LogNone, LogNone, Rdate + "-" + Rtime + Rmicroseconds + " "},
	{LDate | LTime | LMicroseconds | LNanosceonds, "XXX", LogNone, LogNone, Rdate + "-" + Rtime + Rmicroseconds + " " + Rnanoseconds + " "},
	{LDate | LTime | LMicroseconds | LNanosceonds | LLevel | LLongFile | LTag, "XXX", LogNone, LogNone, Rdate + "-" + Rtime + Rmicroseconds + " " + Rnanoseconds + " " + Rlongfile + " XXX "},
	{LDate | LTime | LMicroseconds | LNanosceonds | LLevel | LLongFile | LTag, "XXX", LogNone, LogInfo, Rdate + "-" + Rtime + Rmicroseconds + " " + Rnanoseconds + ` \[Info\] ` + Rlongfile + " XXX "},
	{LDate | LTime | LMicroseconds | LNanosceonds | LLevel | LLongFile | LPreTag, "XXX", LogNone, LogInfo, "XXX " + Rdate + "-" + Rtime + Rmicroseconds + " " + Rnanoseconds + ` \[Info\] ` + Rlongfile + " "},
	{LDate | LTime | LMicroseconds | LNanosceonds | LLevel | LLongFile | LShortFile | LTag | LPreTag, "XXX", LogNone, LogInfo, "XXX " + Rdate + "-" + Rtime + Rmicroseconds + " " + Rnanoseconds + ` \[Info\] ` + Rlongfile + " "},
	// 测试各种日志输出接口的测试样例
	{LDate | LMicroseconds | LNanosceonds | LLevel | LLongFile | LTag, "XXX", LogDebug, LogNone, Rdate + "-" + Rtime + Rmicroseconds + " " + Rnanoseconds + " " + Rlongfile + " XXX "},
	{LDate | LMicroseconds | LNanosceonds | LLevel | LLongFile | LTag, "XXX", LogDebug, LogDebug, Rdate + "-" + Rtime + Rmicroseconds + " " + Rnanoseconds + ` \[Debug\] ` + Rlongfile + " XXX "},
	{LDate | LMicroseconds | LNanosceonds | LLevel | LLongFile | LTag, "XXX", LogDebug, LogInfo, Rdate + "-" + Rtime + Rmicroseconds + " " + Rnanoseconds + ` \[Info\] ` + Rlongfile + " XXX "},
	{LDate | LMicroseconds | LNanosceonds | LLevel | LLongFile | LTag, "XXX", LogInfo, LogNone, Rdate + "-" + Rtime + Rmicroseconds + " " + Rnanoseconds + " " + Rlongfile + " XXX "},
	{LDate | LMicroseconds | LNanosceonds | LLevel | LLongFile | LTag, "XXX", LogInfo, LogDebug, RNoData},
	{LDate | LMicroseconds | LNanosceonds | LLevel | LLongFile | LTag, "XXX", LogInfo, LogInfo, Rdate + "-" + Rtime + Rmicroseconds + " " + Rnanoseconds + ` \[Info\] ` + Rlongfile + " XXX "},
	{LDate | LMicroseconds | LNanosceonds | LLevel | LLongFile | LTag, "XXX", LogInfo, LogWarn, Rdate + "-" + Rtime + Rmicroseconds + " " + Rnanoseconds + ` \[Warn\] ` + Rlongfile + " XXX "},
	{LDate | LMicroseconds | LNanosceonds | LLevel | LLongFile | LTag, "XXX", LogWarn, LogNone, Rdate + "-" + Rtime + Rmicroseconds + " " + Rnanoseconds + " " + Rlongfile + " XXX "},
	{LDate | LMicroseconds | LNanosceonds | LLevel | LLongFile | LTag, "XXX", LogWarn, LogInfo, RNoData},
	{LDate | LMicroseconds | LNanosceonds | LLevel | LLongFile | LTag, "XXX", LogWarn, LogWarn, Rdate + "-" + Rtime + Rmicroseconds + " " + Rnanoseconds + ` \[Warn\] ` + Rlongfile + " XXX "},
	{LDate | LMicroseconds | LNanosceonds | LLevel | LLongFile | LTag, "XXX", LogWarn, LogError, Rdate + "-" + Rtime + Rmicroseconds + " " + Rnanoseconds + ` \[Error\] ` + Rlongfile + " XXX "},
	{LDate | LMicroseconds | LNanosceonds | LLevel | LLongFile | LTag, "XXX", LogError, LogNone, Rdate + "-" + Rtime + Rmicroseconds + " " + Rnanoseconds + " " + Rlongfile + " XXX "},
	{LDate | LMicroseconds | LNanosceonds | LLevel | LLongFile | LTag, "XXX", LogError, LogWarn, RNoData},
	{LDate | LMicroseconds | LNanosceonds | LLevel | LLongFile | LTag, "XXX", LogError, LogError, Rdate + "-" + Rtime + Rmicroseconds + " " + Rnanoseconds + ` \[Error\] ` + Rlongfile + " XXX "},
	{LDate | LMicroseconds | LNanosceonds | LLevel | LLongFile | LTag, "XXX", LogError, LogCritical, Rdate + "-" + Rtime + Rmicroseconds + " " + Rnanoseconds + ` \[Critical\] ` + Rlongfile + " XXX "},
	{LDate | LMicroseconds | LNanosceonds | LLevel | LLongFile | LTag, "XXX", LogCritical, LogNone, Rdate + "-" + Rtime + Rmicroseconds + " " + Rnanoseconds + " " + Rlongfile + " XXX "},
	{LDate | LMicroseconds | LNanosceonds | LLevel | LLongFile | LTag, "XXX", LogCritical, LogError, RNoData},
	{LDate | LMicroseconds | LNanosceonds | LLevel | LLongFile | LTag, "XXX", LogCritical, LogCritical, Rdate + "-" + Rtime + Rmicroseconds + " " + Rnanoseconds + ` \[Critical\] ` + Rlongfile + " XXX "},
	{LDate | LMicroseconds | LNanosceonds | LLevel | LLongFile | LTag, "XXX", LogCritical + 1, LogCritical, RNoData},
}

var testIgnore = []tester{
	{LDate | LMicroseconds | LNanosceonds | LLevel | LLongFile | LTag, "XXX", LogDebug, LogNone, Rdate + "-" + Rtime + Rmicroseconds + " " + Rnanoseconds + " " + RlongfileIgnore + " XXX "},
	{LDate | LMicroseconds | LNanosceonds | LLevel | LLongFile | LTag, "XXX", LogDebug, LogDebug, Rdate + "-" + Rtime + Rmicroseconds + " " + Rnanoseconds + ` \[Debug\] ` + RlongfileIgnore + " XXX "},
	{LDate | LMicroseconds | LNanosceonds | LLevel | LLongFile | LTag, "XXX", LogDebug, LogInfo, Rdate + "-" + Rtime + Rmicroseconds + " " + Rnanoseconds + ` \[Info\] ` + RlongfileIgnore + " XXX "},
	{LDate | LMicroseconds | LNanosceonds | LLevel | LLongFile | LTag, "XXX", LogInfo, LogNone, Rdate + "-" + Rtime + Rmicroseconds + " " + Rnanoseconds + " " + RlongfileIgnore + " XXX "},
	{LDate | LMicroseconds | LNanosceonds | LLevel | LLongFile | LTag, "XXX", LogInfo, LogDebug, Rdate + "-" + Rtime + Rmicroseconds + " " + Rnanoseconds + ` \[Debug\] ` + RlongfileIgnore + " XXX "},
	{LDate | LMicroseconds | LNanosceonds | LLevel | LLongFile | LTag, "XXX", LogInfo, LogInfo, Rdate + "-" + Rtime + Rmicroseconds + " " + Rnanoseconds + ` \[Info\] ` + RlongfileIgnore + " XXX "},
	{LDate | LMicroseconds | LNanosceonds | LLevel | LLongFile | LTag, "XXX", LogInfo, LogWarn, Rdate + "-" + Rtime + Rmicroseconds + " " + Rnanoseconds + ` \[Warn\] ` + RlongfileIgnore + " XXX "},
	{LDate | LMicroseconds | LNanosceonds | LLevel | LLongFile | LTag, "XXX", LogWarn, LogNone, Rdate + "-" + Rtime + Rmicroseconds + " " + Rnanoseconds + " " + RlongfileIgnore + " XXX "},
	{LDate | LMicroseconds | LNanosceonds | LLevel | LLongFile | LTag, "XXX", LogWarn, LogInfo, Rdate + "-" + Rtime + Rmicroseconds + " " + Rnanoseconds + ` \[Info\] ` + RlongfileIgnore + " XXX "},
	{LDate | LMicroseconds | LNanosceonds | LLevel | LLongFile | LTag, "XXX", LogWarn, LogWarn, Rdate + "-" + Rtime + Rmicroseconds + " " + Rnanoseconds + ` \[Warn\] ` + RlongfileIgnore + " XXX "},
	{LDate | LMicroseconds | LNanosceonds | LLevel | LLongFile | LTag, "XXX", LogWarn, LogError, Rdate + "-" + Rtime + Rmicroseconds + " " + Rnanoseconds + ` \[Error\] ` + RlongfileIgnore + " XXX "},
	{LDate | LMicroseconds | LNanosceonds | LLevel | LLongFile | LTag, "XXX", LogError, LogNone, Rdate + "-" + Rtime + Rmicroseconds + " " + Rnanoseconds + " " + RlongfileIgnore + " XXX "},
	{LDate | LMicroseconds | LNanosceonds | LLevel | LLongFile | LTag, "XXX", LogError, LogWarn, Rdate + "-" + Rtime + Rmicroseconds + " " + Rnanoseconds + ` \[Warn\] ` + RlongfileIgnore + " XXX "},
	{LDate | LMicroseconds | LNanosceonds | LLevel | LLongFile | LTag, "XXX", LogError, LogError, Rdate + "-" + Rtime + Rmicroseconds + " " + Rnanoseconds + ` \[Error\] ` + RlongfileIgnore + " XXX "},
	{LDate | LMicroseconds | LNanosceonds | LLevel | LLongFile | LTag, "XXX", LogError, LogCritical, Rdate + "-" + Rtime + Rmicroseconds + " " + Rnanoseconds + ` \[Critical\] ` + RlongfileIgnore + " XXX "},
	{LDate | LMicroseconds | LNanosceonds | LLevel | LLongFile | LTag, "XXX", LogCritical, LogNone, Rdate + "-" + Rtime + Rmicroseconds + " " + Rnanoseconds + " " + RlongfileIgnore + " XXX "},
	{LDate | LMicroseconds | LNanosceonds | LLevel | LLongFile | LTag, "XXX", LogCritical, LogError, Rdate + "-" + Rtime + Rmicroseconds + " " + Rnanoseconds + ` \[Error\] ` + RlongfileIgnore + " XXX "},
	{LDate | LMicroseconds | LNanosceonds | LLevel | LLongFile | LTag, "XXX", LogCritical, LogCritical, Rdate + "-" + Rtime + Rmicroseconds + " " + Rnanoseconds + ` \[Critical\] ` + RlongfileIgnore + " XXX "},
	{LDate | LMicroseconds | LNanosceonds | LLevel | LLongFile | LTag, "XXX", LogCritical + 1, LogCritical, Rdate + "-" + Rtime + Rmicroseconds + " " + Rnanoseconds + ` \[Critical\] ` + RlongfileIgnore + " XXX "},
}

func TestAll(t *testing.T) {
	for _, testcase := range tests {
		testPrint(t, testcase.flag, testcase.tag, testcase.logerLevel, testcase.contextLevel, testcase.pattern, true, false, true)
		testPrint(t, testcase.flag, testcase.tag, testcase.logerLevel, testcase.contextLevel, testcase.pattern, false, true, true)
		testPrint(t, testcase.flag, testcase.tag, testcase.logerLevel, testcase.contextLevel, testcase.pattern, false, false, true)
		testPrint(t, testcase.flag, testcase.tag, testcase.logerLevel, testcase.contextLevel, testcase.pattern, true, false, false)
		testPrint(t, testcase.flag, testcase.tag, testcase.logerLevel, testcase.contextLevel, testcase.pattern, false, true, false)
		testPrint(t, testcase.flag, testcase.tag, testcase.logerLevel, testcase.contextLevel, testcase.pattern, false, false, false)
	}
	for _, testcase := range testIgnore {
		testPrintIgnore(t, testcase.flag, testcase.tag, testcase.logerLevel, testcase.contextLevel, testcase.pattern, true, false, true)
		testPrintIgnore(t, testcase.flag, testcase.tag, testcase.logerLevel, testcase.contextLevel, testcase.pattern, false, true, true)
		testPrintIgnore(t, testcase.flag, testcase.tag, testcase.logerLevel, testcase.contextLevel, testcase.pattern, false, false, true)
		testPrintIgnore(t, testcase.flag, testcase.tag, testcase.logerLevel, testcase.contextLevel, testcase.pattern, true, false, false)
		testPrintIgnore(t, testcase.flag, testcase.tag, testcase.logerLevel, testcase.contextLevel, testcase.pattern, false, true, false)
		testPrintIgnore(t, testcase.flag, testcase.tag, testcase.logerLevel, testcase.contextLevel, testcase.pattern, false, false, false)
	}
}

func Test_PrintEmptyCreatesLine(t *testing.T) {
	var b bytes.Buffer
	l := New("", LogNone, 0, &b)
	l.Print()
	output := b.String()
	if output != "\n" {
		t.Errorf("output expected a \\n, got %s", output)
	}
}

func TestDiscard(t *testing.T) {
	l := New("", LogNone, 0, io.Discard)
	s := strings.Repeat("a", 102400)
	c := testing.AllocsPerRun(100, func() { l.Printf("%s", s) })
	if c > 1 {
		t.Errorf("got %v allocs, want at most 1", c)
	}
}

func Test_UTCFlag(t *testing.T) {
	var b bytes.Buffer
	l := New("", LogNone, LStdFlags, &b)
	l.SetFlag(LStdFlags | LUTC) // | Ltime | LUTC)
	now := time.Now().UTC()
	l.Print("hello")
	want := fmt.Sprintf("%d/%.2d/%.2d-%.2d:%.2d:%.2d hello",
		now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	got := b.String()
	if got == want {
		return
	}
	// 有可能是在执行获取now的时候和输出日志时跳了秒数变了
	now = now.Add(time.Second)
	want = fmt.Sprintf("%d/%.2d/%.2d-%.2d:%.2d:%.2d hello",
		now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	if got == want {
		return
	}
	t.Errorf("log wan %q; bug got %q", want, got)
}

func Test_OutputRace(t *testing.T) {
	var b bytes.Buffer
	l := New("", LogNone, LStdFlags, &b) // New(&b, "", 0)
	for i := 0; i < 100; i++ {
		go func() {
			l.SetFlag(0)
		}()
		l.Output(0, LogNone, "")
	}
}

func Benchmark_Itoa(b *testing.B) {
	dst := make([]byte, 0, 64)
	for i := 0; i < b.N; i++ {
		dst = dst[0:0]
		itoa(&dst, 2015, 4)   // year
		itoa(&dst, 1, 2)      // month
		itoa(&dst, 30, 2)     // day
		itoa(&dst, 12, 2)     // hour
		itoa(&dst, 56, 2)     // minute
		itoa(&dst, 0, 2)      // second
		itoa(&dst, 987654, 6) // microsecond
	}
}

func Benchmark_Println(b *testing.B) {
	const testString = "test"
	var buf bytes.Buffer
	l := New("", LogNone, LStdFlags, &buf)
	for i := 0; i < b.N; i++ {
		buf.Reset()
		l.Println(testString)
	}
}

func Benchmark_Debugln(b *testing.B) {
	const testString = "test"
	var buf bytes.Buffer
	l := New("", LogDebug, LStdFlags, &buf)
	for i := 0; i < b.N; i++ {
		buf.Reset()
		l.Debugln(testString)
	}
}

func Benchmark_PrintlnNoFlags(b *testing.B) {
	const testString = "test"
	var buf bytes.Buffer
	l := New("", LogDebug, 0, &buf)
	for i := 0; i < b.N; i++ {
		buf.Reset()
		l.Println(testString)
	}
}

func Benchmark_PrintlnNoOutput(b *testing.B) {
	const testString = "test"
	var buf bytes.Buffer
	l := New("", LogError, 0, &buf)
	for i := 0; i < b.N; i++ {
		buf.Reset()
		l.Infoln(testString)
	}
}
