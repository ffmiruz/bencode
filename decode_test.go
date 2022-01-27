package bencode

import (
	"strings"
	"testing"
)

func TestDecodeBasic(t *testing.T) {
	// {
	//	"content": ["spam", 777, ["eggs", -9999], {"a":"xyz", "b": 222}],
	//	"info"	 : {"day": "today", "xdate": 13}
	// }
	want := "d7:contentl4:spami777el4:eggsi-9999eed1:a3:xyz1:bi222eee4:infod3:day5:today5:xdatei13eee"
	r := strings.NewReader(want)

	var str, str2, str3, str4 string
	var num, num2, num3, num4 int
	innerList := []interface{}{str2, num2}
	dict := make(map[string]interface{})
	dict["a"] = str3
	dict["b"] = num3

	content := []interface{}{str, num, innerList, dict}
	info := make(map[string]interface{})
	info["day"], info["xdate"] = str4, num4

	data := make(map[string]interface{})
	data["content"] = content
	data["info"] = info

	dec := NewDecoder(r)
	err := dec.Decode(&data)
	if err != nil {
		t.Errorf("decode error: %s", err.Error())
	}
	got, _ := Marshal(data)
	if string(got) != want {
		t.Errorf("\nwant: %s, \ngot : %s", want, string(got))
	}
}
