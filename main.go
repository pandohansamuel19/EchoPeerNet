
// Server Instance ke Object dan gunakan sebagai handler

package main

import (
	// User Defined Function
	// "Student-Management-Golang/controller"
	// "Student-Management-Golang/app"
	// "Student-Management-Golang/exception"
	// "Student-Management-Golang/helper"
	// "Student-Management-Golang/middleware"
	// "Student-Management-Golang/model"
	// "Student-Management-Golang/model/web"
	// "Student-Management-Golang/repository"
	// "Student-Management-Golang/service"
	// "Student-Management-Golang/test"
	
	// cv""
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	
	"fmt"
	// "strings"
	// "io"
	// "os"
	// "flag"
	// "math"
	// "sort"
	// "time"
	// "reflect"
	"net/http"
	// "encoding"
	// "errors"
)

func main() {
	e := echo.New()
    e.GET("/", func(c echo.Context) error {
        return c.String(http.StatusOK, "Hello, World!")
    })
    e.Logger.Fatal(e.Start(":1323"))
}
