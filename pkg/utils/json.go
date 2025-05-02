package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func EncodeJson[T any](w io.Writer, v T, indent bool) error {
	if rw, ok := w.(http.ResponseWriter); ok {
		rw.Header().Set("Content-Type", "application/json")
	}
	enc := json.NewEncoder(w)
	if indent {
		enc.SetIndent("", "\t")
	}
	return enc.Encode(v)
}

func MustMarshall[T any](v T, indent bool) string {
	var b []byte
	var err error
	if indent {
		b, err = json.MarshalIndent(v, "", "\t")
	} else {
		b, err = json.Marshal(v)
	}
	if err != nil {
		panic(err)
	}
	return string(b)
}

func DecodeJson[T any](r io.Reader) (*T, error) {
	v := new(T)
	if err := json.NewDecoder(r).Decode(&v); err != nil {
		return v, fmt.Errorf("encode error: %w", err)
	}
	return v, nil
}
