
// Layer yang berisi business logic dari aplikasi. Dimana service akan menggunakan repository untuk melakukan interaksi ke database. Buat dalam interface


package service

import (
	"fmt"
	"strings"
)

func printMessage(message string, arr []string) {
    var nameString = strings.Join(arr, " ")
    fmt.Println(message, nameString)
}