package Utils

import (
	"strconv"
	"strings"
)

func GetMessageAndStatus(err error) (status int, message string) {
	arr := strings.Split(err.Error(), ":")
	result, _ := strconv.Atoi(arr[0])
	return result, arr[1]
}
