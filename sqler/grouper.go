package sqler

import "strings"

type Grouper struct {
	keys []string
}

func (g *Grouper) Reset() {
	g.keys = nil
}

func (g *Grouper) SetKeys(keys ...string) {
	g.keys = keys
}

func (g *Grouper) AddKeys(keys ...string) {
	if g.keys == nil {
		g.keys = keys
	} else {
		g.keys = append(g.keys, keys...)
	}
}

func (g *Grouper) ToSqlAndArgs() (string, []interface{}) {
	keys := g.keys
	if keys == nil {
		return "", nil
	}
	count := len(keys)
	if count == 0 {
		return "", nil
	}
	var build strings.Builder
	index := 0
	for _, value := range keys {
		if index > 0 {
			build.WriteString(",")
		} else {
			build.WriteString(" group by ")
			index++
		}
		build.WriteString(value)
	}
	return build.String(), nil
}

func GetGrouper(keys ...string) *Grouper {
	return &Grouper{keys: keys}
}
