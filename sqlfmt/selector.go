package sqlfmt

import (
	"strings"
)

// 一个select语句的格式化工具
//
// 可以转化为select的sql选择语句
//
// # Keys设置所有的要选择的键，Tables设置选择目标的表
//
// 如果需要join语句，则设置Joins；
// 如果需要group，则设置Groups；
// 如果需要limit，则设置Limit，Limit默认大于0才会生效；
// 如果需要offset，则设置Offset，offset必须大于等于0才会生效；
// 如果需要having语句，则设置Having
type Selector struct {
	Keys       *Fields       // 查找的键
	Tables     *Fields       // 查找的表
	Joins      []*Join       // join条件
	Groups     *Fields       // group条件
	Conditions *AndCondBatch // where条件
	Orders     []*Order      // order排序条件
	Limit      int           // limit，限制
	Offset     int           // offset，偏移量
	Having     *AndCondBatch // having条件
}

func (s *Selector) Release() {
	s.Keys = nil
	s.Tables = nil
	s.Joins = nil
	s.Groups = nil
	s.Conditions = nil
	s.Orders = nil
	s.Limit = 0
	s.Offset = 0
	s.Having = nil
}

func (s *Selector) Reset() {
	s.Keys.Reset()
	s.Tables.Reset()
	s.Joins = nil
	s.Groups.Reset()
	s.Conditions.Reset()
	s.Orders = nil
	s.Limit = 0
	s.Offset = 0
	s.Having.Reset()
}

func (s *Selector) AddSelectKeys(keys ...string) {
	if s.Keys == nil {
		s.Keys = &Fields{}
	}
	s.Keys.AddFields(keys...)
}

func (s *Selector) AddSelectKeysBySqlFmt(keys ...SqlFmt) {
	if s.Keys == nil {
		s.Keys = &Fields{}
	}
	s.Keys.AddSqlFmt(keys...)
}

func (s *Selector) AddTables(tables ...string) {
	if s.Tables == nil {
		s.Keys = &Fields{}
	}
	s.Tables.AddFields(tables...)
}

func (s *Selector) AddTablesBySqlFmt(tables ...SqlFmt) {
	if s.Tables == nil {
		s.Keys = &Fields{}
	}
	s.Tables.AddSqlFmt(tables...)
}

func (s *Selector) AddJoins(joins ...*Join) {
	if s.Joins == nil {
		s.Joins = joins
	} else {
		s.Joins = append(s.Joins, joins...)
	}
}

func (s *Selector) AddGroupKeys(keys ...string) {
	if s.Groups == nil {
		s.Groups = &Fields{}
	}
	s.Groups.AddFields(keys...)
}

func (s *Selector) AddGroupKeysBySqlFmt(keys ...SqlFmt) {
	if s.Groups == nil {
		s.Groups = &Fields{}
	}
	s.Groups.AddSqlFmt(keys...)
}

func (s *Selector) AddConditions(conds ...SqlFmt) {
	if s.Conditions == nil {
		s.Conditions = &AndCondBatch{}
	}
	s.Conditions.AddConditions(conds...)
}

func (s *Selector) AddOrders(orders ...*Order) {
	if s.Orders == nil {
		s.Orders = orders
	} else {
		s.Orders = append(s.Orders, orders...)
	}
}

func (s *Selector) SetLimit(limit int) {
	s.Limit = limit
}

func (s *Selector) SetOffset(offset int) {
	s.Offset = offset
}

func (s *Selector) AddHavingConditions(conds ...SqlFmt) {
	if s.Having == nil {
		s.Having = &AndCondBatch{}
	}
	s.Having.AddConditions(conds...)
}

func (s *Selector) ToSqlAndArgs() (string, []any, error) {
	selectorName := "Selector"
	if s.Tables == nil {
		return "", nil, SqlFmtError(selectorName, "Tables is nil")
	}
	var build strings.Builder
	var sqlStr string
	var err error
	var args []any
	finalArgs := make([]any, 0)
	build.WriteString("select ")
	if s.Keys == nil {
		build.WriteString("* ")
	} else {
		if sqlStr, args, err = s.Keys.ToSqlAndArgs(); err != nil {
			return "", nil, err
		} else {
			build.WriteString(sqlStr)
			finalArgs = append(finalArgs, args...)
		}
	}
	if sqlStr, args, err = s.Tables.ToSqlAndArgs(); err != nil {
		return "", nil, err
	}
	build.WriteString(" from (")
	build.WriteString(sqlStr)
	build.WriteString(")")
	finalArgs = append(finalArgs, args...)
	if s.Joins != nil {
		count := len(s.Joins)
		if count > 0 {
			for _, value := range s.Joins {
				if sqlStr, args, err = value.ToSqlAndArgs(); err != nil {
					return "", nil, err
				} else {
					build.WriteString(" ")
					build.WriteString(sqlStr)
					finalArgs = append(finalArgs, args...)
				}
			}
		}
	}
	if s.Conditions != nil {
		if sqlStr, args, err = cond2Where(s.Conditions); err != nil {
			return "", nil, err
		} else {
			build.WriteString(sqlStr)
			finalArgs = append(finalArgs, args...)
		}
	}
	if s.Groups != nil {
		if sqlStr, args, err = s.Groups.ToSqlAndArgs(); err != nil {
			return "", nil, err
		} else {
			build.WriteString(" group by ")
			build.WriteString(sqlStr)
			finalArgs = append(finalArgs, args...)
		}
	}
	if s.Having != nil {
		if sqlStr, args, err = cond2Having(s.Having); err != nil {
			return "", nil, err
		} else {
			build.WriteString(sqlStr)
			finalArgs = append(finalArgs, args...)
		}
	}
	if s.Orders != nil {
		count := len(s.Orders)
		if count > 0 {
			build.WriteString(" order by ")
			for index, value := range s.Joins {
				if sqlStr, args, err = value.ToSqlAndArgs(); err != nil {
					return "", nil, err
				} else {
					if index > 0 {
						build.WriteString(",")
					}
					build.WriteString(sqlStr)
					finalArgs = append(finalArgs, args...)
				}
			}
		}
	}
	if s.Limit > 0 {
		build.WriteString(" limit ?")
		finalArgs = append(finalArgs, s.Limit)
	}
	if s.Offset > 0 {
		build.WriteString(" offset ?")
		finalArgs = append(finalArgs, s.Offset)
	}
	return build.String(), finalArgs, nil
}
