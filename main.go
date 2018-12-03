package main

import (
	"bytes"
	ksql "database/sql"
	"flag"
	"fmt"
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

func main() {
	table_schema := flag.String("s", "test", "table_schema")
	table_name := flag.String("t", "test", "table_name")
	flag.Parse()

	sqlStr := "select column_name,data_type,column_key from information_schema.columns where table_schema=? and table_name=?;"
	rows, err := kmysql.MysqlClient.Query(sqlStr, *table_schema, *table_name)
	if err != nil {
		klog.Error.Println("get  fail :", err)
		return
	}
	createStr := createTable(*table_name, rows)
	rows, err = kmysql.MysqlClient.Query(sqlStr, *table_schema, *table_name)
	insertStr := insertTable(*table_name, rows)

	klog.Info.Println("\n", createStr, "\n", insertStr)
}

//---------------------------------------------------------------------------
func createTable(table_name string, rows *ksql.Rows) string {
	var sql string
	sql += "type " + converUpper(table_name) + " struct {\n"
	for rows.Next() {
		column_name := ""
		data_type := ""
		column_key := ""
		err := rows.Scan(&column_name, &data_type, &column_key)
		if err != nil {
			klog.Error.Println("scan fail :", err)
			return ""
		}
		if "ID" == strings.ToUpper(column_name) {
			sql += "\tID     uint64  `gorm:\"primary_key\" json:\"-\"`\n"
		} else {
			sql += "\t" + converUpper(column_name) + " " + getType(data_type, column_name) + " `gorm:\"column:" + column_name + "\" json:\"" + column_name + "\"`\n"
		}
	}
	sql += "}\n"
	sql += "func (" + converUpper(table_name) + ") TableName() string {\n"
	sql += "\treturn \"" + table_name + "\"\n}"

	return sql
}

//---------------------------------------------------------------------------
func insertTable(table_name string, rows *ksql.Rows) string {
	var bf bytes.Buffer
	var bt bytes.Buffer
	//func InsertCoinTransformLog(tx *jgorm.DB, user_id, dst_coin_id, src_coin_id int, dst_coins, src_coins, trans float64, order_id int) (kmodel.FmCoinTransformLog, error) {
	fmt.Fprintf(&bf, "func Insert%s(tx *jgorm.DB", converUpper(table_name))
	fmt.Fprintf(&bt, ` { 
	if tx == nil {  
		tx = gorm.GormDB 
	}`)
	fmt.Fprintf(&bt, "\n\tobj := kmodel.%s{\n", converUpper(table_name))
	for rows.Next() {
		column_name := ""
		data_type := ""
		column_key := ""
		rows.Scan(&column_name, &data_type, &column_key)
		fmt.Fprintf(&bf, ", %s %s", column_name, getType(data_type, column_name))
		if column_key != "PRI" {
			fmt.Fprintf(&bt, "\t\t%s:%s,\n", converUpper(column_name), column_name)
		}
	}
	fmt.Fprintf(&bf, ")(kmodel.%s,error)", converUpper(table_name))
	fmt.Fprintf(&bt, `	}
	if err := tx.Create(&obj).Error; err != nil {
		klog.Error.Println(err)
		return obj, err
	}
	return obj, nil
}`)

	return bf.String() + bt.String()
}

//---------------------------------------------------------------------------

//---------------------------------------------------------------------------

func getType(s, name string) string {
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
		if strings.Index(name, "time") != -1 {
			return "int64"
		}
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
