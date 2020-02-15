package control

import (
	"bytes"
	ksql "database/sql"
	"fmt"
	kinit "gorm_generate/initialize"
	kbase "gorm_generate/work/control/base"
	kdao "gorm_generate/work/dao"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type gorm struct {
}

func Newgorm() *gorm {
	return &gorm{}
}
func (ts *gorm) Load() []kbase.RouteWrapStruct {
	m := make([]kbase.RouteWrapStruct, 0)
	m = append(m, kbase.Wrap("GET", "/gorm", ts.gormindex, 0))
	m = append(m, kbase.Wrap("POST", "/gorm", ts.gorm, 0))
	return m
}

// 127.0.0.1:8089/gorm
func (ts *gorm) gormindex(c *gin.Context) {
	c.HTML(http.StatusOK, "gorm.html", gin.H{})
}

// 127.0.0.1:8089/gorm
func (ts *gorm) gorm(c *gin.Context) {
	tablename := kbase.GetParam(c, "tablename")
	sql := kbase.GetParam(c, "sql")

	kdao.DropTable(nil, tablename)
	kdao.Create(nil, sql)

	defer kdao.DropTable(nil, tablename)

	mysql_mysqldb, _ := kinit.Conf.GetString("mysql.mysqldb")
	ts.run(c, mysql_mysqldb, tablename)
}

func (ts *gorm) run(c *gin.Context, table_schema, table_name string) {
	sqlStr := "select column_name,data_type,column_key from information_schema.columns where table_schema=? and table_name=?;"
	rows, err := kinit.Gorm.Raw(sqlStr, table_schema, table_name).Rows()
	if err != nil {
		kinit.LogError.Println("get  fail :", err)
		return
	}
	createStr := ts.createTable(table_name, rows)

	rows, err = kinit.Gorm.Raw(sqlStr, table_schema, table_name).Rows()
	insertStr := ts.insertTable(table_name, rows)

	sqlStr = "select column_name from information_schema.key_column_usage where table_schema=? and table_name=?;"
	rows, err = kinit.Gorm.Raw(sqlStr, table_schema, table_name).Rows()
	getStr := ts.getTable(table_name, rows)

	rows, err = kinit.Gorm.Raw(sqlStr, table_schema, table_name).Rows()
	updateStr := ts.updateTable(table_name, rows)

	fmt.Fprintln(c.Writer, "\n", createStr, "\n", insertStr, "\n", getStr, "\n", updateStr)

}

func (ts *gorm) createTable(table_name string, rows *ksql.Rows) string {
	var sql string
	sql += "package model\n\n\n\n type " + ts.converUpper(table_name) + " struct {\n"
	for rows.Next() {
		column_name := ""
		data_type := ""
		column_key := ""
		err := rows.Scan(&column_name, &data_type, &column_key)
		if err != nil {
			kinit.LogError.Println("scan fail :", err)
			return ""
		}
		if "ID" == strings.ToUpper(column_name) {
			sql += "\tID     int64  `gorm:\"primary_key\" json:\"-\"`\n"
		} else {
			sql += "\t" + ts.converUpper(column_name) + " " + ts.getType(data_type, column_name) + " `gorm:\"column:" + column_name + "\" json:\"" + column_name + "\"`\n"
		}
	}
	sql += "}\n"
	sql += "func (" + ts.converUpper(table_name) + ") TableName() string {\n"
	sql += "\treturn \"" + table_name + "\"\n}"

	return sql
}

//---------------------------------------------------------------------------
func (ts *gorm) insertTable(table_name string, rows *ksql.Rows) string {
	var bf bytes.Buffer
	var bt bytes.Buffer
	//func InsertCoinTransformLog(tx *jgorm.DB, user_id, dst_coin_id, src_coin_id int, dst_coins, src_coins, trans float64, order_id int) (kmodel.FmCoinTransformLog, error) {
	fmt.Fprintf(&bf, "func Insert%s(tx *jgorm.DB", ts.converUpper(table_name))
	fmt.Fprintf(&bt, ` { 
	if tx == nil {  
		tx = kinit.Gorm
	}`)
	fmt.Fprintf(&bt, "\n\tobj := kmodel.%s{\n", ts.converUpper(table_name))
	for rows.Next() {
		column_name := ""
		data_type := ""
		column_key := ""
		rows.Scan(&column_name, &data_type, &column_key)

		if column_key != "PRI" {
			fmt.Fprintf(&bf, ", %s %s", column_name, ts.getType(data_type, column_name))
			fmt.Fprintf(&bt, "\t\t%s:%s,\n", ts.converUpper(column_name), column_name)
		}
	}
	fmt.Fprintf(&bf, ")(kmodel.%s,error)", ts.converUpper(table_name))
	fmt.Fprintf(&bt, `	}
	if err := tx.Create(&obj).Error; err != nil {
		kinit.LogError.Println(err)
		return obj, err
	}
	return obj, nil
}`)

	return bf.String() + bt.String()
}

func (ts *gorm) getTable(table_name string, rows *ksql.Rows) string {
	var sql string
	sql += ""
	for rows.Next() {
		column_name := ""
		err := rows.Scan(&column_name)
		if err != nil {
			kinit.LogError.Println("scan fail :", err)
			return ""
		}
		sql += fmt.Sprintf(`
func Get%sBy%s(tx *jgorm.DB, %s xxxx) kmodel.%s {
	if tx == nil {
		tx = kinit.Gorm
	}
	var objs kmodel.%s
	tx.Where("%s=? ", %s).First(&objs)
	return objs
}`, ts.converUpper(table_name), ts.converUpper(column_name), column_name, ts.converUpper(table_name),
			ts.converUpper(table_name),
			column_name, column_name)
	}

	return sql
}

//---------------------------------------------------------------------------
func (ts *gorm) updateTable(table_name string, rows *ksql.Rows) string {
	var sql string
	sql += ""
	for rows.Next() {
		column_name := ""
		err := rows.Scan(&column_name)
		if err != nil {
			kinit.LogError.Println("scan fail :", err)
			return ""
		}
		sql += fmt.Sprintf(`
func Update%sBy%s(tx *jgorm.DB, %s xxx) error {
	if tx == nil {
		tx = kinit.Gorm
	}

	if err := tx.Model(kmodel.%s{}).Where("%s=?", %s).Updates(map[string]interface{}{"xxx": xxxx}).Error; err != nil {
		kinit.LogError.Println(err)
		return err
	}
	return nil
}`, ts.converUpper(table_name), ts.converUpper(column_name), column_name,
			ts.converUpper(table_name), column_name, column_name)
	}

	return sql
}

//---------------------------------------------------------------------------

//---------------------------------------------------------------------------

func (ts *gorm) getType(s, name string) string {
	if s == "smallint" {
		return "int64"
	}
	if s == "varchar" {
		return "string"
	}
	if s == "tinyint" {
		return "int64"
	}
	if s == "mediumint" {
		return "int64"
	}
	if s == "int" {
		if strings.Index(name, "time") != -1 {
			return "int64"
		}
		return "int64"
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
		return "string"
	}
	if s == "blob" {
		return "string"
	}
	if s == "varbinary" {
		return "string"
	}
	if s == "timestamp" {
		return "string"
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

func (ts *gorm) converUpper(s string) string {
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
