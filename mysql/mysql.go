package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"strconv"
)

type DataField map[string]interface{}

func InitDB(user, passwd, ip, port, dbName string) (bool, *sql.DB) {
	sqlStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, passwd, ip, port, dbName)
	db, _ := sql.Open("mysql", sqlStr)
	db.SetConnMaxLifetime(100)
	db.SetMaxIdleConns(10)
	if err := db.Ping(); err != nil {
		log.Fatalf("Open sql %s error, error is %s", sqlStr, error(err))
		return false, nil
	} else {
		return true, db
	}
}

func CloseDB(db *sql.DB) {
	db.Close()
}

func ExceSql(sqlStr string, db *sql.DB) bool {
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
		return false
	}
	stmt, err := tx.Prepare(sqlStr)
	if err != nil {
		log.Fatal(err)
		return false
	}
	_, err = stmt.Exec()
	if err != nil {
		log.Fatal(err)
		return false
	}
	tx.Commit()
	return true
}

func InsertData(table string, data DataField, dataType map[string]string, db *sql.DB) bool {
	execSql := makeInsertSql(table, data, dataType)
	return ExceSql(execSql, db)
}

func DeleteData(table string, option []string, db *sql.DB) bool {
	execSql := makeDeleteSql(table, option)
	return ExceSql(execSql, db)
}

func UpdateData(table string, data DataField, dataType map[string]string, option []string, db *sql.DB) bool {
	execSql := makeUpdateSql(table, data, dataType, option)
	return ExceSql(execSql, db)
}

func SelectData(table string, keys []string, dataType map[string]string, option []string, db *sql.DB) (bool, DataField) {
	execSql := makeSelectSql(table, keys, option)

	rows, err := db.Query(execSql)
	if db == nil {
		log.Fatalf("db is null")
	}
	if err != nil {
		log.Fatalf("Query error %s", error(err))
	}

	columns, err := rows.Columns()
	if err != nil {
		log.Fatal("get rolumn name error")
	}

	values := make([]sql.RawBytes, len(columns))

	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	rowMap := make(map[string]string)
	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			log.Fatal("rows Scan error")
		}

		var value string
		for i, col := range values {
			if col != nil {
				value = string(col)
				rowMap[columns[i]] = value
			}
		}
	}
	if len(rowMap) > 0 {
		var result = make(DataField)
		for i := 0; i < len(keys); i++ {
			key := keys[i]
			if value, ok := rowMap[key]; ok {
				valueType := dataType[key]
				switch valueType {
				case "int":
					if tempValue, err := strconv.Atoi(value); err == nil {
						result[key] = tempValue
					} else {
						log.Fatal("SelectData error")
					}
				case "float":
					if tempValue, err := strconv.ParseFloat(value, 64); err == nil {
						result[key] = tempValue
					} else {
						log.Fatal("SelectData error")
					}
				default:
					result[key] = value
				}
			} else {
				result[key] = nil
			}
		}
		return true, result
	} else {
		return true, nil
	}
}

func SelectDatas(table string, keys []string, dataType map[string]string, option []string, db *sql.DB) (bool, []DataField) {
	execSql := makeSelectSql(table, keys, option)

	rows, err := db.Query(execSql)
	if db == nil {
		log.Fatalf("db is null")
	}
	if err != nil {
		log.Fatalf("Query error %s", error(err))
	}

	columns, err := rows.Columns()
	if err != nil {
		log.Fatal("get rolumn name error")
	}

	values := make([]sql.RawBytes, len(columns))

	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	var allRowMaps []map[string]string // make(map[string]string)
	for rows.Next() {
		rowMap := make(map[string]string)
		err = rows.Scan(scanArgs...)
		if err != nil {
			log.Fatal("rows Scan error")
		}

		var value string
		for i, col := range values {
			if col != nil {
				value = string(col)
				rowMap[columns[i]] = value
			}
		}
		allRowMaps = append(allRowMaps, rowMap)
	}
	if len(allRowMaps) > 0 && len(allRowMaps[0]) > 0 {
		var result []DataField // = make(DataField)
		for i := 0; i < len(allRowMaps); i++ {
			rowMap := allRowMaps[i]
			rowField := make(DataField)
			for i := 0; i < len(keys); i++ {
				key := keys[i]
				if value, ok := rowMap[key]; ok {
					valueType := dataType[key]
					switch valueType {
					case "int":
						if tempValue, err := strconv.Atoi(value); err == nil {
							rowField[key] = tempValue
						} else {
							log.Fatal("SelectData error")
						}
					case "float":
						if tempValue, err := strconv.ParseFloat(value, 64); err == nil {
							rowField[key] = tempValue
						} else {
							log.Fatal("SelectData error")
						}
					default:
						rowField[key] = value
					}
				} else {
					rowField[key] = nil
				}
			}
			result = append(result, rowField)
		}
		return true, result
	} else {
		return true, nil
	}
}

func makeInsertSql(table string, data DataField, dataType map[string]string) string {
	valueStr := ""
	keyStr := ""
	for key, value := range data {
		keyStr = keyStr + "`" + key + "`,"
		if tempStr, ok := value.(string); ok {
			switch dataType[key] {
			case "string":
				valueStr = valueStr + "\"" + tempStr + "\","
			case "date":
				if tempStr == "now()" {
					valueStr = valueStr + tempStr + ","
				} else {
					valueStr = valueStr + "\"" + tempStr + "\","
				}
			default:
				valueStr = valueStr + tempStr + ","
			}
		}
	}
	keyStr = keyStr[:len(keyStr)-1]
	valueStr = valueStr[:len(valueStr)-1]
	sqlStr := "insert into `" + table + "`(" + keyStr + ") values (" + valueStr + ");"
	return sqlStr
}

func makeDeleteSql(table string, option []string) string {
	sqlStr := "delete from `" + table + "`"
	optionStr := ""
	if len(option) > 0 {
		for i := 0; i < len(option); i++ {
			optionStr = optionStr + option[i] + " and "
		}
		optionStr = optionStr[:len(optionStr)-5]
		sqlStr = sqlStr + " where " + optionStr
	}
	sqlStr = sqlStr + ";"
	return sqlStr
}

func makeUpdateSql(table string, data DataField, dataType map[string]string, option []string) string {
	optionStr := ""
	keyValueStr := ""
	for key, value := range data {
		keyValueStr = keyValueStr + "`" + key + "`="
		if tempStr, ok := value.(string); ok {
			switch dataType[key] {
			case "string":
				keyValueStr = keyValueStr + "\"" + tempStr + "\""
			case "date":
				if tempStr == "now()" {
					keyValueStr = keyValueStr + tempStr
				} else {
					keyValueStr = keyValueStr + "\"" + tempStr + "\""
				}
			default:
				keyValueStr = keyValueStr + tempStr
			}
		} else {
			keyValueStr = keyValueStr + tempStr
		}
		keyValueStr = keyValueStr + ","
	}
	keyValueStr = keyValueStr[:len(keyValueStr)-1]
	sqlStr := "update " + table + " set " + keyValueStr
	if len(option) > 0 {
		for i := 0; i < len(option); i++ {
			optionStr = optionStr + option[i] + " and "
		}
		optionStr = optionStr[:len(optionStr)-5]
		sqlStr = sqlStr + " where " + optionStr
	}
	sqlStr = sqlStr + ";"
	return sqlStr
}

func makeSelectSql(table string, key []string, option []string) string {
	keyStr := ""
	for i := 0; i < len(key); i++ {
		keyStr = keyStr + "`" + key[i] + "`,"
	}
	keyStr = keyStr[:len(keyStr)-1]

	sqlStr := "select " + keyStr + " from " + table
	if len(option) > 0 {
		optionStr := ""
		for i := 0; i < len(option); i++ {
			optionStr = optionStr + option[i] + " and "
		}
		optionStr = optionStr[:len(optionStr)-5]
		sqlStr = sqlStr + " where " + optionStr
	}
	sqlStr = sqlStr + ";"
	return sqlStr
}
