package unpacking

import (
	"errors"
	"strconv"
)

var WrongInputErr error = errors.New("wrong string")

func Unpacking(str string) (string, error) {
	out := []rune{}

	sl := []rune(str)
	if len(sl) == 0 {
		return "", nil
	}
	if sl[0] >= 48 && sl[0] <= 57 {
		return "", WrongInputErr
	}

	flag := true
	last := sl[0]

	for _, v := range sl {
		if v >= '0' && v <= '9' && flag {
			flag = false
			c, _ := strconv.Atoi(string(v)) //так как выше в if проверка на то, что v - строго число.
			for i := 1; i < c; i++ {
				out = append(out, last)
			}
		} else if v >= 48 && v <= 57 {
			return "", WrongInputErr
		} else {
			last = v
			out = append(out, v)
			flag = true
		}
	}
	return string(out), nil
}
