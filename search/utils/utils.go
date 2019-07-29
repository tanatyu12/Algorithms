package utils

import (
	"flag"
	"strconv"
	"errors"
)

func GetCommandArg() (int, int, error) {
	flag.Parse()
	n, err1 := strconv.Atoi(flag.Arg(0))
	target, err2 := strconv.Atoi(flag.Arg(1))
	if err1 != nil || err2 != nil {
		err := errors.New("arg is not Int")
		return 0, 0, err
	}
	return n, target, nil
}

func MakeCollection(len int) []int {
	var arr []int
	for i := 1; i <= len; i++ {
		arr = append(arr, i)
	}
	return arr
}