// GoTest project GoTest.go
package GoTest

import (
	"fmt"
	"os"
)

type T struct {
	name  string
	value int
}

func main() {
	fmt.Println("Hello World!")
	//	os.Create("c:/TEMP/gofile.txt")
	fp, e := os.OpenFile("C:/TEMP/gofile.txt", os.O_APPEND, 0777)

	if e == nil {
		fp.WriteString("Hell\r\no\r\n")
	}

	fp.Close()
}
