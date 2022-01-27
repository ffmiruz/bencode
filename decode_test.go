package bencode

import (
	"strings"
	"testing"
)

func TestDecodeBasic(t *testing.T) {
	// ["spam", 23, ["eggs", -9999], {"a":"xyz", "b": 222}]
	want := "l4:spami23el4:eggsi-9999eed1:a3:xyz1:bi222eee"
	r := strings.NewReader(want)

	var str, str2, str3 string
	var num, num2, num3 int
	innerList := []interface{}{str2, num2}
	dict := make(map[string]interface{})
	dict["a"] = str3
	dict["b"] = num3
	list := []interface{}{str, num, innerList, dict}
	dec := NewDecoder(r)
	err := dec.Decode(&list)
	if err != nil {
		t.Errorf("decode error: %s", err.Error())
	}
	got, _ := Marshal(list)
	if string(got) != want {
		t.Errorf("\nwant: %s, \ngot : %s", want, string(got))
	}
}
