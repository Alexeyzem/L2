package unpacking_test

import (
	"errors"
	"testing"
	"unpacking_task/pkg/unpacking"
)

type testCase struct {
	data   string
	result string
	err    error
}

func TestUnpacking(t *testing.T) {
	cases := []testCase{
		{data: "a4b5", result: "aaaabbbbb", err: nil},
		{data: "45", result: "", err: unpacking.WrongInputErr},
		{data: "", result: "", err: nil},
		{data: "a45", result: "", err: unpacking.WrongInputErr},
	}
	for i, v := range cases {
		result, err := unpacking.Unpacking(v.data)
		if result != v.result {
			t.Errorf("Error in case[%v] wrong result. Expected:%v, got:%v", i, v.result, result)
		}
		if err != nil && v.err == nil {
			t.Errorf("Case[%v]. Unexpected error: %v", i, err)
		}
		if err == nil && v.err != nil {
			t.Errorf("Case[%v]. Expected error: %v, got nil", i, v.err)
		}
		if !errors.Is(err, v.err) {
			t.Errorf("Case[%v]. Expexted error:%v, got:%v", i, v.err, err)
		}
	}
}
