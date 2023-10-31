package databasehelper

import (
	//	"bqqgc/errer"
	// "database/sql"
	// "io/ioutil"
	// "os"
	// "strings"
	"github.com/bqqsrc/bqqg/database"
)

var defaultRegister = ""

func SetDefaultRegister(name string) {
	defaultRegister = name
}

func RegistController(name, driverName string, sqlName string) error {
	_, err := database.RegistController(name, driverName, sqlName)
	return err
}
