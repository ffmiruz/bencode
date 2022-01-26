package bencode

import (
	"bytes"
	"testing"
)

func TestEncode(t *testing.T) {
	want := "d8:announce38:udp://tracker.publicbt.com:80/announce13:announce-listll38:udp://tracker.publicbt.com:80/announceel44:udp://tracker.openbittorrent.com:80/announceee7:comment33:Debian CD from cdimage.debian.org4:infod6:lengthi170917888e4:name30:debian-8.8.0-arm64-netinst.iso12:piece lenghti262144eee"

	dict := make(map[string]interface{})
	info := make(map[string]interface{})
	info["length"] = 170917888
	info["piece lenght"] = 262144
	info["name"] = "debian-8.8.0-arm64-netinst.iso"
	dict["info"] = info
	dict["announce"] = "udp://tracker.publicbt.com:80/announce"
	list := make([]interface{}, 2)
	innerList1 := make([]interface{}, 1)
	innerList2 := make([]interface{}, 1)
	list[0], list[1] = innerList1, innerList2
	innerList1[0] = "udp://tracker.publicbt.com:80/announce"
	innerList2[0] = "udp://tracker.openbittorrent.com:80/announce"
	dict["announce-list"] = list
	dict["comment"] = "Debian CD from cdimage.debian.org"

	w := new(bytes.Buffer)
	e := NewEncoder(w)
	err := e.Encode(dict)
	if err != nil {
		t.Fatal(err)
	}
	if len(w.String()) != len(want) {
		t.Errorf("len is not equal. want:%v, got:%v", len(want), len(w.String()))
	}
	if w.String() != want {
		t.Errorf("\nwant: \"%s\", \ngot:  \"%s\"", want, w.String())
	}
}

func TestBasicEncode(t *testing.T) {
	want := "l4:spam4:eggsi3ee"
	slice := []interface{}{"spam", "eggs", 3}

	wantDict := "d9:publisher3:bob17:publisher-webpage15:www.example.com18:publisher.location4:homee"
	dict := make(map[string]interface{})
	dict["publisher"] = "bob"
	dict["publisher-webpage"] = "www.example.com"
	dict["publisher.location"] = "home"

	w := new(bytes.Buffer)
	e := NewEncoder(w)
	err := e.Encode(slice)
	if err != nil {
		t.Fatal(err)
	}
	if w.String() != want {
		t.Errorf("want: \"%s\", got \"%s\"", want, w.String())
	}

	w.Reset()
	err = e.Encode(dict)
	if err != nil {
		t.Fatal(err)
	}
	if w.String() != wantDict {
		t.Errorf("want: \"%s\", got \"%s\"", wantDict, w.String())
	}
}

func TestEncodeListOfList(t *testing.T) {
	// [["spam"],["eggs"]]
	want := "ll4:spamel4:eggsee"
	list := make([]interface{}, 2)
	innerList1 := make([]interface{}, 1)
	innerList2 := make([]interface{}, 1)
	list[0], list[1] = innerList1, innerList2
	innerList1[0] = "spam"
	innerList2[0] = "eggs"

	w := new(bytes.Buffer)
	e := NewEncoder(w)
	err := e.Encode(list)
	if err != nil {
		t.Fatal(err)
	}
	if len(w.String()) != len(want) {
		t.Errorf("len is not equal. want:%v, got:%v", len(want), len(w.String()))
	}
	if w.String() != want {
		t.Errorf("\nwant: \"%s\", \ngot:  \"%s\"", want, w.String())
	}
}
