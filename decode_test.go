package bencode

import (
	"strings"
	"testing"
)

func TestDecodeBasic(t *testing.T) {
	ctl := "l4:spami23ee"
	r := strings.NewReader(ctl)

	var str string
	var num int
	list := []interface{}{str, num}
	dec := NewDecoder(r)
	err := dec.Decode(&list)
	if err != nil {
		t.Errorf("decode error: %s", err.Error())
	}
	got, _ := Marshal(list)
	if string(got) != ctl {
		t.Errorf("want: %s, got: %s", ctl, string(got))
	}
}
