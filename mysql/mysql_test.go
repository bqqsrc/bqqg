package mysql

import (
	"fmt"
	"testing"
)

func Test_InitDB_ExceSql(t *testing.T) {
	user := "root"
	passwd := "123456"
	ip := "localhost"
	port := "3306"
	dbName := "test"
	result, db := InitDB(user, passwd, ip, port, dbName)
	if !result {
		t.Fatalf("Test_InitDB_ExceSql error, init da error")
	}
	if db == nil {
		t.Fatalf("Test_InitDB_ExceSql error, DB is null")
	}

	mysqlStr := "insert into users_list(account, passwd) values (\"test_account\", \"test_password\");"
	result = ExceSql(mysqlStr, db)
	if !result {
		t.Fatalf("Test_InitDB_ExceSql error, ExceSql error")
	}
	CloseDB(db)
}

func Test_InitDB_InsertData(t *testing.T) {
	user := "root"
	passwd := "123456"
	ip := "localhost"
	port := "3306"
	dbName := "test"
	result, db := InitDB(user, passwd, ip, port, dbName)
	if !result {
		t.Fatalf("Test_InitDB_ExceSql error, init da error")
	}
	if db == nil {
		t.Fatalf("Test_InitDB_ExceSql error, DB is null")
	}

	var data = make(DataField)
	data["account"] = "test_account2"
	data["passwd"] = "test_password2"
	dataType := map[string]string{
		"account": "string",
		"passwd":  "string",
	}
	result = InsertData("users_list", data, dataType, db)
	if !result {
		t.Fatalf("Test_InitDB_InsertData error, InsertData error")
	}

	newData := DataField{
		"user_id":          "10",
		"formulation_name": "GX-5888",
		"customer":         "小明",
		"color":            "黑白色",
		"material":         "PP",
		"stock_num":        "8",
		"sell_price":       "78.9",
		"create_time":      "2020-8-14",
		"price_time":       "2020-9-7",
		"total_cost":       "14",
		"has_price":        "1",
		"statement":        "一句测试的备注",
		"product_weight":   "33",
		"manual_cost":      "26.7",
		"put_time":         "2020-9-14",
		"pigment_cost":     "23",
		"total_weight":     "32g + 26g = 58g",
	}
	dataType = map[string]string{
		"user_id":          "int",
		"formulation_name": "string",
		"customer":         "string",
		"color":            "string",
		"material":         "string",
		"stock_num":        "float",
		"sell_price":       "float",
		"create_time":      "date",
		"price_time":       "date",
		"total_cost":       "float",
		"has_price":        "int",
		"statement":        "string",
		"product_weight":   "float",
		"manual_cost":      "float",
		"put_time":         "date",
		"pigment_cost":     "float",
		"total_weight":     "string",
	}
	result = InsertData("formulation_list", newData, dataType, db)
	if !result {
		t.Fatalf("Test_InitDB_InsertData error, InsertData error")
	}

	newData = DataField{
		"user_id":          "10",
		"formulation_name": "GX-5888",
		"customer":         "小明",
		"color":            "黑白色",
		"material":         "PP",
		"stock_num":        "8",
		"sell_price":       "78.9",
		"create_time":      "2020-8-14",
		"price_time":       "2020-9-7",
		"total_cost":       "14",
		"has_price":        "1",
		"statement":        "一句测试的备注",
		"product_weight":   "33",
		"manual_cost":      "26.7",
		"put_time":         "now()",
		"pigment_cost":     "23",
		"total_weight":     "32g + 26g = 58g",
	}
	result = InsertData("formulation_list", newData, dataType, db)
	if !result {
		t.Fatalf("Test_InitDB_InsertData error, InsertData error")
	}
	CloseDB(db)
}

func Test_InitDB_DeleteData(t *testing.T) {
	user := "root"
	passwd := "123456"
	ip := "localhost"
	port := "3306"
	dbName := "test"
	result, db := InitDB(user, passwd, ip, port, dbName)
	if !result {
		t.Fatalf("Test_InitDB_ExceSql error, init da error")
	}
	if db == nil {
		t.Fatalf("Test_InitDB_ExceSql error, DB is null")
	}

	option := []string{"account=\"test_account2\"", "passwd=\"test_password2\""}
	result = DeleteData("users_list", option, db)
	if !result {
		t.Fatalf("Test_InitDB_DeleteData error, DeleteData error")
	}
	CloseDB(db)
}

func Test_InitDB_UpdateData(t *testing.T) {
	user := "root"
	passwd := "123456"
	ip := "localhost"
	port := "3306"
	dbName := "test"
	result, db := InitDB(user, passwd, ip, port, dbName)
	if !result {
		t.Fatalf("Test_InitDB_ExceSql error, init da error")
	}
	if db == nil {
		t.Fatalf("Test_InitDB_ExceSql error, DB is null")
	}

	var data = make(DataField)
	data["account"] = "test_account3"
	data["passwd"] = "test_password3"
	dataType := map[string]string{
		"account": "string",
		"passwd":  "string",
	}
	option := []string{"account=\"test_account\"", "passwd=\"test_password\""}

	result = UpdateData("users_list", data, dataType, option, db)
	if !result {
		t.Fatalf("Test_InitDB_UpdateData error, UpdateData error")
	}
	CloseDB(db)
}

func Test_InitDB_SelectData(t *testing.T) {
	user := "root"
	passwd := "123456"
	ip := "localhost"
	port := "3306"
	dbName := "test"
	result, db := InitDB(user, passwd, ip, port, dbName)
	if !result {
		t.Fatalf("Test_InitDB_ExceSql error, init da error")
	}
	if db == nil {
		t.Fatalf("Test_InitDB_ExceSql error, DB is null")
	}

	var key = []string{"account", "passwd", "type"}
	dataType := map[string]string{
		"account": "string",
		"passwd":  "string",
		"type":    "int",
	}
	option := []string{"id=4"}
	result, resultData := SelectData("users_list", key, dataType, option, db)
	fmt.Printf("resultData is %v\n", resultData)
	fmt.Println("=============================================================")
	if !result {
		t.Fatalf("Test_InitDB_UpdateData error, UpdateData error")
	}
	if fmt.Sprintf("%T", resultData["type"]) != "int" {
		t.Fatal("Test_InitDB_UpdateData, type of type is not int")
	}

	dataType = map[string]string{
		"user_id":          "int",
		"formulation_name": "string",
		"customer":         "string",
		"color":            "string",
		"material":         "string",
		"stock_num":        "float",
		"sell_price":       "float",
		"create_time":      "date",
		"price_time":       "date",
		"total_cost":       "float",
		"has_price":        "int",
		"statement":        "string",
		"product_weight":   "float",
		"manual_cost":      "float",
		"put_time":         "date",
		"pigment_cost":     "float",
		"total_weight":     "string",
	}
	key = []string{"user_id", "formulation_name", "customer", "color", "material", "stock_num", "sell_price", "create_time", "price_time", "total_cost", "has_price", "statement", "product_weight", "manual_cost", "put_time", "pigment_cost", "total_weight"}
	option = []string{"id=3"}
	result, resultData = SelectData("formulation_list", key, dataType, option, db)
	fmt.Printf("resultData is %v\n", resultData)
	fmt.Println("=============================================================")
	if !result {
		t.Fatalf("Test_InitDB_UpdateData error, UpdateData error")
	}
	if fmt.Sprintf("%T", resultData["user_id"]) != "int" {
		t.Fatal("Test_InitDB_UpdateData, type of user_id is not int")
	}
	if fmt.Sprintf("%T", resultData["pigment_cost"]) != "float64" {
		t.Fatal("Test_InitDB_UpdateData, type of pigment_cost is not float64")
	}
	CloseDB(db)
}
