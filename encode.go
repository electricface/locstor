// Copyright 2015 Alex Browne.  All rights reserved.
// Use of this source code is governed by the MIT
// license, which can be found in the LICENSE file.

package locstor

import (
	"github.com/cathalgarvey/fmtless/encoding/json"
)

var (
	// JSONEncoding is a ready-to-use implementation of EncoderDecoder which
	// encodes data structures as json.
	JSONEncoding = &jsonEncoderDecoder{}
)

// Encoder is an interface implemented by objects which can encode an arbitrary
// go object into a slice of bytes.
type Encoder interface {
	Encode(interface{}) ([]byte, error)
}

// Decoder is an interface implemented by objects which can decode a slice
// of bytes into an arbitrary go object.
type Decoder interface {
	Decode([]byte, interface{}) error
}

// EncoderDecoder is an interface implemented by objects which can both encode
// an arbitrary go object into a slice of bytes and decode that slice of bytes
// into an arbitrary go object. EncoderDecoders should have the property that
// Encode(Decode(x)) == x for all objects x which are encodable.
type EncoderDecoder interface {
	Encoder
	Decoder
}

// jsonEncoderDecoder is an implementation of EncoderDecoder which uses json
// encoding.
type jsonEncoderDecoder struct{}

// Encode implements the Encode method of Encoder
func (jsonEncoderDecoder) Encode(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

// Decode implements the Decode method of Decoder
func (jsonEncoderDecoder) Decode(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

