package bencode

import (
	"strings"
	"testing"
)

func TestDecodeBasic(t *testing.T) {
	// ["spam", 23, ["eggs", -9999]]
	ctl := "l4:spami23el4:eggsi-9999eee"
	r := strings.NewReader(ctl)

	var str, str2 string
	var num, num2 int
	innerList := []interface{}{str2, num2}
	list := []interface{}{str, num, innerList}
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
