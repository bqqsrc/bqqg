//  Copyright (C) 晓白齐齐,版权所有.

package loger

import (
	"bytes"
	"os"
	"regexp"
	"testing"
)

// 使用Println("hello", 23, "world")、Print("hello", 23, "world\n")、Printf("hello %d world\n", 23)3个函数来测试
// 以及5个等级输出日志的对应函数来测试
func testPrint(t *testing.T, flag int, tag string, logerLevel, contextLevel LogLevelType, pattern string, useFormat bool, useln bool, isDefault bool) {
	buf := new(bytes.Buffer)
	l := New(tag, logerLevel, flag, buf)
	var funcMap = map[LogLevelType]logFuncStruct{
		LogDebug:    {l.Debug, l.Debugf, l.Debugln},
		LogInfo:     {l.Info, l.Infof, l.Infoln},
		LogWarn:     {l.Warn, l.Warnf, l.Warnln},
		LogError:    {l.Error, l.Errorf, l.Errorln},
		LogCritical: {l.Critical, l.Criticalf, l.Criticalln},
		LogNone:     {l.Print, l.Printf, l.Println},
	}
	if isDefault {
		SetWriter(buf)
		defer SetWriter(os.Stderr)
		SetFlag(flag)
		SetTag(tag)
		SetLevel(logerLevel)
		funcMap = map[LogLevelType]logFuncStruct{
			LogDebug:    {Debug, Debugf, Debugln},
			LogInfo:     {Info, Infof, Infoln},
			LogWarn:     {Warn, Warnf, Warnln},
			LogError:    {Error, Errorf, Errorln},
			LogCritical: {Critical, Criticalf, Criticalln},
			LogNone:     {Print, Printf, Println},
		}
	}
	logFuncs := funcMap[contextLevel]
	if useFormat {
		logFuncs.funcFormat("hello %d world", 23)
	} else if useln {
		logFuncs.funcLine("hello", 23, "world")
	} else {
		logFuncs.funcCommon("hello ", 23, " world\n")
	}
	expect := buf.String()
	if pattern == RNoData {
		pattern = ""
	} else {
		pattern = "^" + pattern + "hello 23 world\n$"
	}
	matched, err := regexp.MatchString(pattern, expect)
	if err != nil {
		t.Fatal("pattern did not compile:", err)
	}
	if !matched {
		t.Errorf("log output should match %q, but is %q", pattern, expect)
	}
}

func testPrintIgnore(t *testing.T, flag int, tag string, logerLevel, contextLevel LogLevelType, pattern string, useFormat bool, useln bool, isDefault bool) {
	buf := new(bytes.Buffer)
	l := New(tag, logerLevel, flag, buf)
	var ignoreFuncs = ignoreFuncStruct{
		l.PrintIgnoreLevel,
		l.PrintfIgnoreLevel,
		l.PrintlnIgnoreLevel,
	}
	if isDefault {
		SetWriter(buf)
		defer SetWriter(os.Stderr)
		SetFlag(flag)
		SetTag(tag)
		SetLevel(logerLevel)
		ignoreFuncs = ignoreFuncStruct{
			PrintIgnoreLevel,
			PrintfIgnoreLevel,
			PrintlnIgnoreLevel,
		}
	}
	if useFormat {
		ignoreFuncs.funcFormat(contextLevel, "hello %d world", 23)
	} else if useln {
		ignoreFuncs.funcLine(contextLevel, "hello", 23, "world")
	} else {
		ignoreFuncs.funcCommon(contextLevel, "hello ", 23, " world\n")
	}
	expect := buf.String()
	pattern = "^" + pattern + "hello 23 world\n$"
	matched, err := regexp.MatchString(pattern, expect)
	if err != nil {
		t.Fatal("pattern did not compile:", err)
	}
	if !matched {
		t.Errorf("log output should match %q, but is %q", pattern, expect)
	}
}
