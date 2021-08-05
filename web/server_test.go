package web

import (
	"strconv"
	"testing"
)

type args struct {
	data string
}

type ts struct {
	in string
	want string
}

func TestRun(t *testing.T) {
	go Run()
	list := []ts{
		{
			"a", "",
		},
		{
			"a", "exist",
		},
		{
			"b", "",
		},
		{
			"a", "exist",
		},
		{
			"c", "",
		},
		{
			"b", "exist",
		},
	}
	for i, l := range list {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			r := DoClientSend(l.in)
			if r != l.want {
				t.Errorf("want (%s), got (%s)", l.want, r)
			}
		})
	}
}
