package base

import (
	"github.com/gin-gonic/gin"
)

const (
	ZERO           = 0
	CHECKTOKEN     = 1 << 0
	ORDERFREQUENCY = 1 << 2
)

type RouteWrapStruct struct {
	Method string
	Path   string
	F      func(*gin.Context)
}

func Wrap(Method string, Path string, f func(*gin.Context), types int) RouteWrapStruct {
	wp := RouteWrapStruct{
		Method: Method,
		Path:   Path,
	}

	wp.F = func(c *gin.Context) {
		if (types & CHECKTOKEN) > 0 {
			err := CheckToken(c)
			if err != nil {
				return
			}
		}

		f(c)
	}
	return wp
}

func CheckToken(c *gin.Context) error {

	return nil
}
