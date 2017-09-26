package app

import (
	"fmt"
)

func print(str string) string {
	return fmt.Sprintf("%v%v%v", str, str, str)
}
