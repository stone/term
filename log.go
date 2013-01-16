package term

import (
	"fmt"
	"log"
	"strings"
)

func Sprintf(c Color, format string, args ...interface{}) string {
	return strings.Join([]string{
		string(c),
		fmt.Sprintf(format, args...),
		string(Reset),
	}, "")
}

func Logf(c Color, format string, args ...interface{}) {
	log.Print(Sprintf(c, format, args...))
}

func Loglnf(c Color, format string, args ...interface{}) {
	log.Println(Sprintf(c, format, args...))
}
