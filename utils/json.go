package utils

import (
	"bytes"
	"encoding/json"
)

func JsonEncode(v interface{}, escape bool) string {
	bf := bytes.NewBuffer([]byte{})
	jsonEncoder := json.NewEncoder(bf)
	jsonEncoder.SetEscapeHTML(escape)
	jsonEncoder.Encode(v)

	return bf.String()
}
