package main

import (
	"flag"
	"strings"

	kconf "github.com/heyuanlong/go-utils/common/conf"
	klog "github.com/heyuanlong/go-utils/common/log"
	kmysql "github.com/heyuanlong/go-utils/db/mysql"
)

func init() {
	kconf.SetFile("conf/config.cfg")
	kmysql.InitMysql()
}

//gorm_generate.exe -s beegoblog -t users

var sql string

func main() {
	table_schema := flag.String("s", "test", "table_schema")
	table_name := flag.String("t", "test", "table_name")
	flag.Parse()

	header(*table_name)

	sqlStr := "select column_name,data_type from information_schema.columns where table_schema=? and table_name=?;"
	rows, err := kmysql.MysqlClient.Query(sqlStr, *table_schema, *table_name)
	if err != nil {
		klog.Error.Println("get  fail :", err)
		return
	}
	for rows.Next() {
		column_name := ""
		data_type := ""
		err = rows.Scan(&column_name, &data_type)
		if err != nil {
			klog.Error.Println("scan fail :", err)
			return
		}
		//klog.Info.Println("column_name:", converUpper(column_name), " data_type:", data_type)
		row(column_name, data_type)
	}
	tail(*table_name)
	klog.Info.Println("ok\n")
	klog.Info.Println(sql)
}

func header(name string) {
	sql += "type " + converUpper(name) + " struct {\n"
}
func row(name, types string) {
	if "ID" == strings.ToUpper(name) {
		sql += "\tID     uint64  `gorm:\"primary_key\" json:\"-\"`\n"
	} else {
		sql += "\t" + converUpper(name) + " " + getType(types) + "`gorm:\"column:" + name + "\" json:\"" + name + "\"`\n"
	}
}
func tail(name string) {
	sql += "}\n"
	sql += "func (" + converUpper(name) + ") TableName() string {\n"
	sql += "\treturn \"" + name + "\"\n}"
}

func getType(s string) string {
	if s == "smallint" {
		return "int"
	}
	if s == "varchar" {
		return "string"
	}
	if s == "tinyint" {
		return "int"
	}
	if s == "mediumint" {
		return "int64"
	}
	if s == "int" {
		return "int"
	}
	if s == "text" {
		return "string"
	}
	if s == "mediumtext" {
		return "string"
	}
	if s == "char" {
		return "string"
	}
	if s == "mediumblob" {
		return "string"
	}
	if s == "enum" {
		return "string"
	}
	if s == "float" {
		return "string"
	}
	if s == "date" {
		return "string"
	}
	if s == "decimal" {
		return "float64"
	}
	if s == "double" {
		return "float64"
	}
	if s == "longtext" {
		return "string"
	}
	if s == "bigint" {
		return "int64"
	}
	if s == "datetime" {
		return "int64"
	}
	if s == "blob" {
		return "string"
	}
	if s == "varbinary" {
		return "string"
	}
	if s == "timestamp" {
		return "int64"
	}
	if s == "set" {
		return "string"
	}
	if s == "longblob" {
		return "string"
	}
	if s == "time" {
		return "string"
	}

	return ""
}

func converUpper(s string) string {
	tmp := strings.Split(s, "_")
	var res string
	for i := 0; i < len(tmp); i++ {
		v := []rune(tmp[i])
		for y := 0; y < len(v); y++ {
			if y == 0 {
				if v[y] >= 97 && v[y] <= 122 {
					v[y] -= 32
				}
				res += string(v[y])
			} else {
				res += string(v[y])
			}
		}
	}
	return res
}
