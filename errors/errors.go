// Copyright (C) 晓白齐齐,版权所有.

package errors

import (
	"fmt"
	"strings"
)

type ErrorGroup []error

func (eg ErrorGroup) Error() string {
	if eg != nil && len(eg) > 0 {
		var build strings.Builder
		build.WriteString("some error hanppened: \n")
		for index, err := range eg {
			build.WriteString(fmt.Sprintf("%d: %s\n", index+1, err.Error()))
		}
		return build.String()
	}
	return ""
}

func (eg ErrorGroup) AddErrors(errs ...error) ErrorGroup {
	if eg == nil {
		eg = make(ErrorGroup, 0)
	}
	for _, err := range errs {
		if errGroup, ok := err.(ErrorGroup); ok {
			eg = append(eg, errGroup...)
		} else {
			eg = append(eg, err)
		}
	}
	return eg
}

func (eg ErrorGroup) AddErrorf(format string, args ...any) ErrorGroup {
	err := fmt.Errorf(format, args...)
	if eg == nil {
		return ErrorGroup{err}
	}
	return append(eg, err)
}
