package log

import (
	"api/config/setting"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"runtime"
	"strings"
)

func Verbose(text string){
	_, fn, line, _ := runtime.Caller(1)
	fmt.Printf("[Verbose] %s ,line %d: %s", fn, line, text)

	if gin.Mode() == gin.ReleaseMode {
		var builder strings.Builder
		f, _ := os.OpenFile(setting.PathConfig.VerboseLogPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		defer closeFile(f)
		log.SetOutput(f)
		_, _ = fmt.Fprintf(&builder, "[Verbose] %s ,line %d: %s \n", fn, line, text)
		log.Println(builder.String())
	}
}

func Error(err error){
	if err != nil {
		_, fn, line, _ := runtime.Caller(1)
		fmt.Printf("[ERROR] %s ,line %d: %s \n", fn, line, err.Error())

		if gin.Mode() == gin.ReleaseMode {
			var builder strings.Builder
			f, _ := os.OpenFile(setting.PathConfig.ErrorLogPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
			defer closeFile(f)
			log.SetOutput(f)
			_,_ = fmt.Fprintf(&builder, "[ERROR] %s ,line %d: %s \n", fn, line, err.Error())
			log.Println(builder.String())
		}
	}
}

func closeFile(f *os.File){
	if f != nil {
		_ = f.Close()
	}
}

func Sql(text interface{}){
	if gin.Mode() == gin.DebugMode {
		_, fn, line, _ := runtime.Caller(1)
		fmt.Printf("[SQL] %s ,line %d: %v \n", fn, line, text)
	}
}