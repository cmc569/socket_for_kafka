package handler

import (
	"api/util/time_date"
	"encoding/json"
	"sync"
	"time"
)

var CreateTransactionMappingMap = make(map[string][]byte, 0)

type DeleteTimeOutData struct {
	CreatedTime time.Time `json:"createdTime"`
}

func RspListen(mappingMap map[string][]byte, uuid string) map[string]interface{} {
	dataList := map[string]interface{}{
		"status":  false,
		"message": "time out",
	}
	var SendTime = time_date.TimeNow()
	var mutex sync.RWMutex
	for {
		isLoop := true
		if time_date.TimeNow().Sub(SendTime).Seconds() >= 30 {
			break
		}
		mutex.RLock()
		if val, ok := mappingMap[uuid]; ok {
			type receivedData struct {
				Uuid    string `json:"uuid"`
				Status  bool   `json:"status"`
				Message string `json:"message"`
			}
			var data receivedData
			if err := json.Unmarshal(val, &data); err == nil {
				dataList["status"] = data.Status
				dataList["message"] = data.Message
			}
			isLoop = false
		}
		mutex.RUnlock()
		if !isLoop {
			break
		}
		time.Sleep(500 * time.Millisecond)
	}
	return dataList
}

func ListenCreateTransaction(dataByte []byte) (isBreak bool) {
	type receivedData struct {
		Uuid    string `json:"uuid"`
		Status  bool   `json:"status"`
		Message string `json:"message"`
	}
	var data receivedData
	var mutex sync.RWMutex
	if err := json.Unmarshal(dataByte, &data); err == nil {
		mutex.Lock()
		CreateTransactionMappingMap[data.Uuid] = dataByte
		if len(CreateTransactionMappingMap) > 1000 {
			var data DeleteTimeOutData
			for key, value := range CreateTransactionMappingMap {
				if err := json.Unmarshal(value, &data); err == nil {
					if time_date.TimeNow().Sub(data.CreatedTime).Minutes() >= 5 {
						delete(CreateTransactionMappingMap, key)
					}
				}
			}
		}
		mutex.Unlock()
	}
	return false
}