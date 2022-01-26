// This package implement encoding and decoding of bencode data.
package bencode

import (
	"bytes"
	"errors"
	"io"
	"sort"
	"strconv"
)

type Encoder struct {
	w io.Writer
}

// NewEncoder returns a new Encoder which perform bencode encoding. Encoded
// data is written to w.
func NewEncoder(w io.Writer) *Encoder {
	return &Encoder{w: w}
}

func (e *Encoder) Encode(v interface{}) error {
	return e.encode(v)
}

func (e *Encoder) encode(v interface{}) error {
	switch data := v.(type) {
	case map[string]interface{}:
		e.encodeDict(data)
	case []interface{}:
		e.encodeList(data)
	case int:
		e.encodeInt(data)
	case string:
		e.encodeString(data)
	default:
		return errors.New("error type switching")
	}
	return nil
}

func (e *Encoder) encodeString(v string) error {
	strlen := strconv.Itoa(len(v))
	e.w.Write([]byte(strlen))
	e.w.Write([]byte(":"))
	e.w.Write([]byte(v))

	return nil
}

func (e *Encoder) encodeInt(v int) error {
	e.w.Write([]byte("i"))
	num := strconv.Itoa(v)
	e.w.Write([]byte(num))
	e.w.Write([]byte("e"))

	return nil
}

func (e *Encoder) encodeList(v []interface{}) error {
	e.w.Write([]byte("l"))
	for i := range v {
		e.encode(v[i])
	}
	e.w.Write([]byte("e"))

	return nil
}

func (e *Encoder) encodeDict(v map[string]interface{}) error {
	e.w.Write([]byte("d"))

	// grab all keys into a slice and sort it.
	keys := make(sort.StringSlice, len(v))
	i := 0
	for key := range v {
		keys[i] = key
		i++
	}
	keys.Sort()

	for _, k := range keys {
		e.encodeString(k)
		value := v[k]
		e.encode(value)
	}
	e.w.Write([]byte("e"))

	return nil
}

// Marshal returns bencode encoding of v in bytes.
func Marshal(v interface{}) ([]byte, error) {
	w := new(bytes.Buffer)
	e := NewEncoder(w)
	err := e.Encode(v)
	if err != nil {
		return nil, err
	}
	return w.Bytes(), nil
}
