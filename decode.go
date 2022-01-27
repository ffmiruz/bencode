package bencode

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"sort"
	"strconv"
)

type Decoder struct {
	rd *bufio.Reader
	// Temporary store buffer
	buf *bytes.Buffer
}

// NewDecoder returns a new decoder that reads from r.
func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{rd: bufio.NewReader(r), buf: &bytes.Buffer{}}
}

// Decode reads bencode data from its input and store to value pointed by v.
func (d *Decoder) Decode(v interface{}) error {
	return d.decode(v)
}

func (d *Decoder) decode(v interface{}) error {
	var err error
	switch value := v.(type) {
	case *int:
		err = d.decodeInt(value)
	case *string:
		err = d.decodeString(value)
	case *[]interface{}:
		err = d.decodeList(value)
	case *map[string]interface{}:
		err = d.decodeDict(value)
	default:
		return errors.New("v points to invalid type to decode to")
	}

	return err
}

func (d *Decoder) decodeDict(dict *map[string]interface{}) error {
	// consume "d"
	_, err := d.rd.Discard(1)
	if err != nil {
		return err
	}
	obj := *dict

	// grab all keys into a slice and sort it.
	keys := make(sort.StringSlice, len(obj))
	i := 0
	for key := range obj {
		keys[i] = key
		i++
	}
	keys.Sort()

	for _, key := range keys {
		// consume bytes of key from input reader, do nothing with it.
		var _dump string
		err = d.decodeString(&_dump)
		if err != nil {
			return err
		}

		switch v := obj[key].(type) {
		case string:
			var str string
			d.decode(&str)
			obj[key] = str
		case int:
			var num int
			d.decode(&num)
			obj[key] = num
		case []interface{}:
			d.decode(&v)
			obj[key] = v
		case map[string]interface{}:
			d.decode(&v)
			obj[key] = v
		default:
			return errors.New("v points to invalid type to decode to")
		}
	}
	// consume "e"
	_, err = d.rd.Discard(1)
	if err != nil {
		return err
	}

	return nil
}

// TODO: check list format.
func (d *Decoder) decodeList(list *[]interface{}) error {
	// consume "l"
	_, err := d.rd.Discard(1)
	if err != nil {
		return err
	}
	l := *list
	// fill in the slice
	for i := range l {
		switch value := l[i].(type) {
		case int:
			var num int
			err = d.decode(&num)
			if err != nil {
				return err
			}
			l[i] = num
		case string:
			var str string
			err = d.decode(&str)
			if err != nil {
				return err
			}
			l[i] = str
		case []interface{}:
			err = d.decodeList(&value)
			if err != nil {
				return err
			}
		case map[string]interface{}:
			err = d.decodeDict(&value)
			if err != nil {
				return err
			}
		}
	}
	// consume "e"
	_, err = d.rd.Discard(1)
	if err != nil {
		return err
	}
	return nil
}

// TODO: check illegal int format
func (d *Decoder) decodeInt(num *int) error {
	numstr, err := d.rd.ReadBytes('e')
	if err != nil {
		return err
	}
	// trim "i" and "e"
	numstr = numstr[1 : len(numstr)-1]
	*num, err = strconv.Atoi(string(numstr))
	if err != nil {
		return err
	}

	return nil
}

func (d *Decoder) decodeString(str *string) error {
	lenstr, err := d.rd.ReadBytes(':')
	if err != nil {
		return err
	}
	// get the string length
	n, err := strconv.Atoi(string(lenstr[:len(lenstr)-1]))
	if err != nil {
		return err
	}
	// copy n bytes into temporary buf.
	_, err = io.CopyN(d.buf, d.rd, int64(n))
	if err != nil {
		return err
	}
	*str = d.buf.String()
	d.buf.Reset()

	return nil
}
