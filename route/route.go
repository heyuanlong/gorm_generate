package route

import (
	"fmt"
	"log"
	"path/filepath"

	kinit "gorm_generate/initialize"
	kbase "gorm_generate/work/control/base"

	"github.com/gin-gonic/gin"
)

type ControlInterface interface {
	Load() []kbase.RouteWrapStruct
}

//-------------------------------------------------------------------------

type RouteStruct struct {
	engine *gin.Engine
	port   int
}

func NewRouteStruct(port int) *RouteStruct {
	ts := &RouteStruct{
		port: port,
	}
	ts.engine = gin.New()
	ts.engine.Use(gin.Logger(), gin.Recovery())

	ts.engine.Static("/assets", "./view/assets")
	ts.engine.Static("/upload", "./view/upload")

	ts.LoadHTML()
	return ts
}

func (ts *RouteStruct) LoadHTML() {
	templates := make([]string, 0)
	firstTemplates, err := filepath.Glob("view/**/*.html")
	if nil != err {
		log.Fatal("load theme templates failed: " + err.Error())
	}

	rootTemplates, err := filepath.Glob("view/*.html")
	if nil != err {
		log.Fatal("load theme templates failed: " + err.Error())
	}

	twoTemplates, err := filepath.Glob("view/**/**/*.html")
	if nil != err {
		log.Fatal("load theme templates failed: " + err.Error())
	}

	templates = append(templates, firstTemplates...)
	templates = append(templates, rootTemplates...)
	templates = append(templates, twoTemplates...)

	ts.engine.LoadHTMLFiles(templates...)
}

func (ts *RouteStruct) Load(control ControlInterface) {
	wps := control.Load()
	for _, v := range wps {
		switch v.Method {
		case "GET":
			ts.engine.GET(v.Path, v.F)
		case "POST":
			ts.engine.POST(v.Path, v.F)
		case "PUT":
			ts.engine.PUT(v.Path, v.F)
		case "PATCH":
			ts.engine.PATCH(v.Path, v.F)
		case "OPTIONS":
			ts.engine.OPTIONS(v.Path, v.F)
		case "DELETE":
			ts.engine.DELETE(v.Path, v.F)
		default:
			kinit.LogError.Println("not method :", v.Method)
		}
	}
}
func (ts *RouteStruct) SetMode(mode string) {
	gin.SetMode(mode)
}

func (ts *RouteStruct) SetMiddleware(middleware ...gin.HandlerFunc) {
	ts.engine.Use(middleware...)
}

func (ts *RouteStruct) Run() {
	addr := fmt.Sprintf(":%d", ts.port)
	ts.engine.Run(addr)
}

//-------------------------------------------------------------------------
func SetCommonHeader(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET,POST")
}
