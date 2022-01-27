package bencode

import (
	"strings"
	"testing"
)

func TestDecodeBasic(t *testing.T) {
	// {
	//	"content": ["spam", 23, ["eggs", -9999], {"a":"xyz", "b": 222}]
	// }
	want := "d7:contentl4:spami23el4:eggsi-9999eed1:a3:xyz1:bi222eeee"
	r := strings.NewReader(want)

	var str, str2, str3 string
	var num, num2, num3 int
	innerList := []interface{}{str2, num2}
	dict := make(map[string]interface{})
	dict["a"] = str3
	dict["b"] = num3
	list := []interface{}{str, num, innerList, dict}
	data := make(map[string]interface{})
	data["content"] = list

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
