package gin_extend

import (
	"github.com/gin-gonic/gin"
	"github.com/iancoleman/strcase"
	"net/http"
	"strings"
	"time"
)

type Gin struct {
	Context *gin.Context
}

type Response struct {
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func BeforeAction(content *gin.Context) *Gin {
	c := Gin{Context: content}
	return &c
}

func (gin *Gin) SuccessResponse(msg string, data ...interface{}) {
	var resData interface{}
	if data == nil {
		resData = ""
	} else {
		resData = data[0]
	}
	gin.Context.JSON(http.StatusOK, Response{
		Msg:  msg,
		Data: lowerCamelCase(&resData),
	})
	return
}

func (gin *Gin) ErrorResponse(msg string, data ...interface{}) {
	var resData interface{}
	if data == nil {
		resData = ""
	} else {
		resData = data[0]
	}

	gin.Context.JSON(http.StatusBadRequest, Response{
		Msg:  msg,
		Data: lowerCamelCase(&resData),
	})
	return
}

func lowerCamelCase(data *interface{}) interface{} {
	mapData, ok := (*data).(map[string]interface{})
	if ok {
		for key, value := range mapData {
			processingData(key, &mapData)
			lowerCamelCase(&value)
		}
		return mapData
	}

	mapAryData, ok := (*data).([]map[string]interface{})
	if ok {
		for _, mapData := range mapAryData {
			for key, value := range mapData {
				processingData(key, &mapData)
				lowerCamelCase(&value)
			}
		}
		return mapAryData
	}
	return data
}

func processingData(key string, mapData *map[string]interface{}) {
	//format time
	if _, ok := (*mapData)[key].(time.Time); ok {
		(*mapData)[key] = (*mapData)[key].(time.Time).Format("2006-01-02 15:04:05")
	}

	if (*mapData)[key] == nil {
		(*mapData)[key] = ""
	}

	// snake case to lower camel
	if i := strings.Index(key, "_"); i != -1 {
		(*mapData)[strcase.ToLowerCamel(key)] = (*mapData)[key]
		delete(*mapData, key)
	}
}
