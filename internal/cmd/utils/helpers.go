package utils

import (
	"fmt"
	"os"
)

func ExitIfErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func StringPointer(s string) *string {
	return &s
}

func StringInSlice(s []string, v string) bool {
	for _, a := range s {
		if a == v {
			return true
		}
	}
	return false
}
