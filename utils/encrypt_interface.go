package utils

import (
	"encoding/gob"
	"github.com/huyinghuan/cfb"
	"bytes"
)

func EncryptInterface(key string, value interface{}) (body []byte, err error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	if err = enc.Encode(value); err != nil {
		return
	}
	body, err = cfb.Encrypt([]byte(key), buf.Bytes())
	return
}
