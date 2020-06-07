package dev

import (
	"fmt"
	"strings"
	"testing"
)

func TestCPUTempParser(t *testing.T) {
	data := "45277\n"
	fmt.Println(data)
	ret := strings.ReplaceAll(string(data), "\n", "")
	fmt.Println(ret)
}
