package sqler

import "strings"

const (
	Left = iota
	Right
	Inner
	Outer
)

type Joiner struct {
	leftOrRight   string
	innerOrOuter  string
	table         string
	conditionList *ConditionBatch
}

func (j *Joiner) Reset() {
	j.leftOrRight = ""
	j.innerOrOuter = ""
	j.table = ""
	j.conditionList = nil
}
func (j *Joiner) Set(leftOrRight, innerOrOuter int, table string) {
	if leftOrRight == Left {
		j.SetLeft()
	} else if leftOrRight == Right {
		j.SetRight()
	}
	if innerOrOuter == Inner {
		j.SetInner()
	} else if innerOrOuter == Outer {
		j.SetOuter()
	}
	j.SetTable(table)
}

func (j *Joiner) SetLeft() {
	j.leftOrRight = "left"
}

func (j *Joiner) SetRight() {
	j.leftOrRight = "right"
}

func (j *Joiner) SetInner() {
	j.innerOrOuter = "inner"
}

func (j *Joiner) SetOuter() {
	j.innerOrOuter = "outer"
}

func (j *Joiner) SetTable(table string) {
	j.table = table
}

func (j *Joiner) SetConditions(conditions *ConditionBatch) {
	j.conditionList = conditions
}

func (j *Joiner) AddConditions(conditions ...SqlAndArgs) {
	if j.conditionList == nil {
		j.conditionList = GetConditionBatch()
	}
	j.conditionList.AddConditions(conditions...)
}

func (j *Joiner) ToSqlAndArgs() (string, []interface{}) {
	if j.table == "" {
		return "", nil
	}
	var build strings.Builder
	if j.leftOrRight != "" {
		build.WriteString(" ")
		build.WriteString(j.leftOrRight)
	}
	if j.innerOrOuter != "" {
		build.WriteString(" ")
		build.WriteString(j.innerOrOuter)
	}
	build.WriteString(" join ")
	build.WriteString(j.table)
	if j.conditionList != nil {
		sqlStr, args := j.conditionList.toOn()
		build.WriteString(sqlStr)
		return build.String(), args
	}
	return build.String(), nil
}

func GetJoiner(leftOrRight, innerOrOuter int, table string) *Joiner {
	joiner := &Joiner{}
	joiner.Set(leftOrRight, innerOrOuter, table)
	return joiner
}
