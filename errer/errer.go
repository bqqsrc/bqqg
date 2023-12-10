package errer

import "fmt"

type callerErr struct {
	caller string
	format string
	args   []interface{}
}

func (e *callerErr) Error() string {
	fmtStr := fmt.Sprintf("call by: %s,\nerror is : %s\n", e.caller, e.format)
	return fmt.Sprintf(fmtStr, e.args...)
}

func CallerErr(caller, format string, args ...interface{}) error {
	return &callerErr{caller, format, args}
}

type callerErr_new struct {
	pack string 
	caller string
	format string
	args   []interface{}
}

func (e *callerErr_new) Error() string {
	fmtStr := fmt.Sprintf("call by: %s.%s,\nerror is : %s\n", e.pack, e.caller, e.format)
	return fmt.Sprintf(fmtStr, e.args...)
}

func CallerErrNew(pack, caller, format string, args ...interface{}) error {
	return &callerErr_new{pack: pack, caller: caller, format: format, args: args}
}
