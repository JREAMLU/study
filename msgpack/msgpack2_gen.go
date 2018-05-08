package main

// NOTE: THIS FILE WAS PRODUCED BY THE
// MSGP CODE GENERATION TOOL (github.com/tinylib/msgp)
// DO NOT EDIT

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *Foo) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "bar":
			z.Bar, err = dc.ReadString()
			if err != nil {
				return
			}
		case "baz":
			z.Baz, err = dc.ReadFloat64()
			if err != nil {
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z Foo) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "bar"
	err = en.Append(0x82, 0xa3, 0x62, 0x61, 0x72)
	if err != nil {
		return
	}
	err = en.WriteString(z.Bar)
	if err != nil {
		return
	}
	// write "baz"
	err = en.Append(0xa3, 0x62, 0x61, 0x7a)
	if err != nil {
		return
	}
	err = en.WriteFloat64(z.Baz)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z Foo) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "bar"
	o = append(o, 0x82, 0xa3, 0x62, 0x61, 0x72)
	o = msgp.AppendString(o, z.Bar)
	// string "baz"
	o = append(o, 0xa3, 0x62, 0x61, 0x7a)
	o = msgp.AppendFloat64(o, z.Baz)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Foo) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "bar":
			z.Bar, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "baz":
			z.Baz, bts, err = msgp.ReadFloat64Bytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z Foo) Msgsize() (s int) {
	s = 1 + 4 + msgp.StringPrefixSize + len(z.Bar) + 4 + msgp.Float64Size
	return
}
