package general

import (
	"math/rand"
	"os/exec"
	"strconv"
	"strings"
)

// BuildInvitationCode 產生英文隨機碼
func BuildInvitationCode() (invitationCode string) {
	var letters = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, 6)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func CreateUUid() string {
	out, _ := exec.Command("uuidgen").Output()
	uuid := string(out)
	return strings.ReplaceAll(uuid, "\n", "")
}

func StringSplitToIntAry(splitValue, splitSep string) []int64 {
	var filterStatusAry []int64
	if splitValue == "" {
		return filterStatusAry
	}
	strAry := strings.Split(splitValue, splitSep)
	for _, value := range strAry {
		status, _ := strconv.ParseInt(value, 10, 64)
		filterStatusAry = append(filterStatusAry, status)
	}

	return filterStatusAry
}



func MappingEnumToList(modelMap map[int64]string) []map[string]interface{} {
	var dataList = make([]map[string]interface{}, 0)
	for k, v := range modelMap {
		dataList = append(dataList, map[string]interface{}{
			"name":  v,
			"value": k,
		})
	}

	return dataList
}