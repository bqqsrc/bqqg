//  Copyright (C) 晓白齐齐,版权所有.

package log

// // 日志模块
// import (
// //	"io"
// 	"os"
// 	"log"
// 	"fmt"
// 	"github.com/bqqsrc/goper/kernel"
// )

// type MutilLoger struct {
// 	prefix string
// 	level int
// 	flag int
// 	outs []io.Writer
// 	logers []*Loger
// 	buf []byte
// }

// func NewMutilLoger(prefix string, level, flag int, outs []io.Writer, logers []*Loger) *MutilLoger {
// 	l := &MutilLoger{prefix: prefix, level: level, flag: flag, outs: outs, logers: logers}
// 	return l
// }

// func (ml *MutilLoger) AddOutput(ws ...io.Writer) {
// 	if ml.outs == nil {
// 		ml.outs = ws
// 	} else {
// 		ml.outs = append(ml.outs, ws...)
// 	}
// }

// func (ml *MutilLoger) SetOutput(ws ...io.Writer) {
// 	ml.outs = ws
// }

// func (ml *MutilLoger) AddLogers(logs ...*Loger) {
// 	if ml.logers == nil {
// 		ml.logers = logs
// 	} else {
// 		ml.logers = append(ml.logers, logs...)
// 	}
// }

// func (ml *MutilLoger) SetOutput(logs ...*Loger) {
// 	ml.logers = logs
// 	ml.logers = logs
// }
