package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"

	_ "modernc.org/sqlite"
)

var database *sql.DB

const ROBOTTABLE = "robots"

// Open the database. If it doesn't exist, create it. Return an error if there is a problem.
func OpenDB() error {
	var err error
	fn := "robots.db"
	database, err = sql.Open("sqlite", fn)
	if err != nil {
		return err
	}
	database.SetMaxOpenConns(1)
	return nil
}

// Create a new DB. Remove the old one if it exists.
func CreateDB() error {

	DropTable(ROBOTTABLE)
	// DropTable(GAMEDBTABLE)
	// DropTable(BIRTHRECORDTABLE)
	// DropTable(LOGGINGTABLE)
	// DropTable(USERSTABLE)
	// DropTable(SAVTOSAV)

	var s string
	s = "BEGIN;\n"
	s += CreateTableFromStruct(ROBOTTABLE, therobots{})
	s += "\n"
	//s = strings.ToLower(s)
	//tx.MustExec(s)

	// s += CreateTableFromStruct(GAMEDBTABLE, gamedb{})
	// s += "\n"
	//s = strings.ToLower(s)
	//tx.MustExec(s)

	// s += CreateTableFromStruct(BIRTHRECORDTABLE, birthrecord{})
	// s += "\n"
	//s = strings.ToLower(s)
	//tx.MustExec(s)

	// s += CreateTableFromStruct(LOGGINGTABLE, tlog{})
	// s += "\n"
	//s = strings.ToLower(s)
	//tx.MustExec(s)

	// s += CreateTableFromStruct(USERSTABLE, user{})
	// s += "\n"
	//s = strings.ToLower(s)
	//tx.MustExec(s)

	// s += CreateTableFromStruct(SAVTOSAV, savtosav{})
	// s += "\n"

	s += "COMMIT;\n"
	fmt.Println(s)
	statement, err := database.Prepare(s)
	if err != nil {
		log.Println(err, s)
		return err
	}

	_, err = statement.Exec()
	if err != nil {
		log.Println(err)
		return err
	}

	// initGame()

	fmt.Println("Created a new database.")
	os.Exit(0)
	return nil
}

// Drop a table if it exists.
func DropTable(table string) {
	statement := fmt.Sprintf("DROP TABLE IF EXISTS %s;", table)
	database.Exec(statement)
}

// Create table based on struct.
// Retuns the sql statement as a string.
// This is a work in progress.
func CreateTableFromStruct(table string, s interface{}) string {

	var reflectedValue reflect.Value = reflect.ValueOf(s) // reflect the struct (interface)

	var sqlstatement string

	//os.Remove("db/savages.db")

	sqlstatement1 := "CREATE TABLE " + table + " ("
	for i := 0; i < reflectedValue.NumField(); i++ {
		var vt string
		varName := reflectedValue.Type().Field(i).Name // get the name of the field
		sqlstatement += "," + varName + " "
		varType := reflectedValue.Type().Field(i).Type // get the type of the field

		// Did this differnt than the other reflect code. This is a work in progress.
		switch varType.Kind() {
		case reflect.Int:
			if varName == "ID" { // detect if the field is the ID field
				vt = "INTEGER NOT NULL PRIMARY KEY"
			} else {
				vt = "INTEGER"
			}
		case reflect.Int8:
			vt = "INTEGER"
		case reflect.Int16:
			vt = "INTEGER"
		case reflect.Int32:
			vt = "INTEGER"
		case reflect.Int64:
			vt = "INTEGER"
		case reflect.Uint:
			vt = "INTEGER"
		case reflect.Uint8:
			vt = "INTEGER"
		case reflect.Uint16:
			vt = "INTEGER"
		case reflect.Uint32:
			vt = "INTEGER"
		case reflect.Uint64:
			vt = "INTEGER"
		case reflect.String:
			vt = "TEXT"
		case reflect.Float64:
			vt = "REAL"
		case reflect.Float32:
			vt = "REAL"
		case reflect.Bool:
			vt = "INTEGER"
		}
		sqlstatement += vt
	}

	// such a crappy way to do this. Return to this at a later date.
	//sqlstatement = sqlstatement[1:] // remove the first comma
	sqlstatement = strings.TrimPrefix(sqlstatement, ",")
	sqlstatement += ");"
	sqlstatement = sqlstatement1 + sqlstatement

	return sqlstatement
}

func InsertIntoTable(table string, s interface{}) string {

	var middlesql1 string
	var middlesql2 string

	var reflectedValue reflect.Value = reflect.ValueOf(s)

	middlesql1 = "INSERT INTO " + table + " ("
	middlesql2 = ")VALUES("
	for i := 0; i < reflectedValue.NumField(); i++ {

		varName := reflectedValue.Type().Field(i).Name
		varType := reflectedValue.Type().Field(i).Type
		varValue := reflectedValue.Field(i).Interface()

		middlesql1 += varName + ","

		// This is my normal way of working with reflect. Strings may be slower but easier to read.
		switch varType.Kind() {
		case reflect.Int:
			if varName == "ID" {
				middlesql2 += fmt.Sprintf("NULL") + ","
			} else {
				middlesql2 += fmt.Sprintf("%d", varValue.(int)) + ","
			}
			//middlesql2 += fmt.Sprintf("%d", varValue.(int)) + ","
		case reflect.Int8:
			middlesql2 += fmt.Sprintf("%d", varValue.(int8)) + ","
		case reflect.Int16:
			middlesql2 += fmt.Sprintf("%d", varValue.(int16)) + ","
		case reflect.Int32:
			middlesql2 += fmt.Sprintf("%d", varValue.(int32)) + ","
		case reflect.Int64:
			middlesql2 += fmt.Sprintf("%d", varValue.(int64)) + ","
		case reflect.Uint:
			middlesql2 += fmt.Sprintf("%d", varValue.(uint)) + ","
		case reflect.Uint8:
			middlesql2 += fmt.Sprintf("%d", varValue.(uint8)) + ","
		case reflect.Uint16:
			middlesql2 += fmt.Sprintf("%d", varValue.(uint16)) + ","
		case reflect.Uint32:
			middlesql2 += fmt.Sprintf("%d", varValue.(uint32)) + ","
		case reflect.Uint64:
			middlesql2 += fmt.Sprintf("%d", varValue.(uint64)) + ","
		case reflect.String:
			middlesql2 += "'" + varValue.(string) + "',"
		case reflect.Float32:
			middlesql2 += fmt.Sprintf("%f", varValue.(float64)) + ","
		case reflect.Float64:
			middlesql2 += fmt.Sprintf("%f", varValue.(float64)) + ","
		case reflect.Bool:
			middlesql2 += fmt.Sprintf("%v", varValue.(bool)) + ","
		default:
			return ""
		}
	}

	middlesql1 = middlesql1[:len(middlesql1)-1]
	middlesql2 = middlesql2[:len(middlesql2)-1] + ");"
	yyy := middlesql1 + middlesql2
	return yyy
}

func RemoveRow(table string, where string) string {

	var sql1, sql2 string
	sql1 = "DELETE FROM " + table + " WHERE "
	sql2 = sql1 + where + ";"
	return sql2

}

func UpdateRow(table string, v map[string]any, where string) string {

	var sql1, sql2, sql3 string
	sql1 = "UPDATE " + table + " SET "

	for p, c := range v {
		switch c.(type) {
		case int:
			sql2 += p + " = " + fmt.Sprintf("%d", c) + ","
		case int8:
			sql2 += p + " = " + fmt.Sprintf("%d", c) + ","
		case int16:
			sql2 += p + " = " + fmt.Sprintf("%d", c) + ","
		case int32:
			sql2 += p + " = " + fmt.Sprintf("%d", c) + ","
		case int64:
			sql2 += p + " = " + fmt.Sprintf("%d", c) + ","
		case uint:
			sql2 += p + " = " + fmt.Sprintf("%d", c) + ","
		case uint8:
			sql2 += p + " = " + fmt.Sprintf("%d", c) + ","
		case uint16:
			sql2 += p + " = " + fmt.Sprintf("%d", c) + ","
		case uint32:
			sql2 += p + " = " + fmt.Sprintf("%d", c) + ","
		case uint64:
			sql2 += p + " = " + fmt.Sprintf("%d", c) + ","
		case string:
			sql2 += p + " = " + "'" + fmt.Sprintf("%s", c) + "'" + ","
		case float32:
			sql2 += p + " = " + fmt.Sprintf("%f", c) + ","
		case float64:
			sql2 += p + " = " + fmt.Sprintf("%f", c) + ","
		case bool:
			sql2 += p + " = " + fmt.Sprintf("%v", c) + ","
		default:
			return ""
		}

	}

	sql3 = sql1 + sql2 + " WHERE " + where + ";"
	return sql3

}
