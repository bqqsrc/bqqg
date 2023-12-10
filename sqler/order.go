package sqler

import (
	"fmt"
	"strings"

	"github.com/bqqsrc/loger"
)

type OrderBy struct {
	Key string
	Asc bool
}

type Order struct {
	key string
	asc bool
}

func (o *Order) Set(key string, isAsc bool) {
	o.key = key
	o.asc = isAsc
}

func (o *Order) Reset() {
	o.key = ""
	o.asc = false
}

func (o *Order) ToSqlAndArgs() (string, []interface{}) {
	funcName := "Order.ToSqlAndArgs"
	loger.Debugf("%s, o.key: %s\no.asc: %t\n", funcName, o.key, o.asc)
	if o.key == "" {
		return "", nil
	}
	ascStr := "desc"
	if o.asc {
		ascStr = "asc"
	}
	sql := fmt.Sprintf("%s %s", o.key, ascStr)
	loger.Debugf("%s, sql: %s\n", funcName, sql)
	return sql, nil
}

func GetOrder(key string, isAsc bool) *Order {
	return &Order{key: key, asc: isAsc}
}

type OrderBatch struct {
	orders []*Order
}

func (o *OrderBatch) AddOrder(key string, isAsc bool) {
	if o.orders == nil {
		o.orders = make([]*Order, 0)
	}
	for _, value := range o.orders {
		if value.key == key {
			loger.Errorf("redeclare OrderBy: %s", key)
			return
		}
	}
	o.orders = append(o.orders, &Order{key: key, asc: isAsc})
}
func (o *OrderBatch) SetOrder(key string, isAsc bool) {
	orders := o.orders
	if orders == nil {
		orders = make([]*Order, 0)
	}
	count := len(orders)
	change := false
	for i := 0; i < count; i++ {
		if orders[i].key == key {
			orders[i].asc = isAsc
			change = true
			break
		}
	}
	if !change {
		orders = append(orders, &Order{key: key, asc: isAsc})
	}
	o.orders = orders
}
func (o *OrderBatch) AddOrders(orders ...*Order) {
	if o.orders == nil {
		o.orders = orders
	} else {
		o.orders = append(o.orders, orders...)
	}
}

func (o *OrderBatch) SetOrders(orders ...*Order) {
	o.orders = orders
}
func (o *OrderBatch) Reset() {
	o.orders = nil
}

func (o *OrderBatch) ToSqlAndArgs() (string, []interface{}) {
	orders := o.orders
	if orders == nil {
		return "", nil
	}
	count := len(orders)
	if count == 0 {
		return "", nil
	}
	var build strings.Builder
	index := 0
	args := make([]interface{}, 0)
	for _, value := range orders {
		sql, arg := value.ToSqlAndArgs()
		if sql != "" {
			if index > 0 {
				build.WriteString(",")
			} else {
				build.WriteString(" order by ")
				index++
			}
			build.WriteString(sql)
			args = append(args, arg...)
		}
	}
	return build.String(), args
}

func GetOrderBatch() *OrderBatch {
	return &OrderBatch{}
}
