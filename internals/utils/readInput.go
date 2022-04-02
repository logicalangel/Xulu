package utils

import (
	"bufio"
	"os"
	"xulu/internals/consts"
)

func ReadLine() string {
	reader := bufio.NewReader(os.Stdin)
	lineB, _, err := reader.ReadLine()
	if err != nil {
		panic(consts.Internal_error)
	}

	return string(lineB)
}
