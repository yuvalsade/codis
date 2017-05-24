// Copyright 2016 CodisLabs. All Rights Reserved.
// Licensed under the MIT (MIT-LICENSE.txt) license.

package models

import (
	"bytes"
	"compress/gzip"
	"encoding/gob"
	"encoding/json"

	"github.com/CodisLabs/codis/pkg/utils/errors"
	"github.com/CodisLabs/codis/pkg/utils/log"
)

func JsonEncode(v interface{}) []byte {
	b, err := json.MarshalIndent(v, "", "    ")
	if err != nil {
		log.PanicErrorf(err, "encode to json failed")
	}
	return b
}

func JsonString(v interface{}) string {
	return string(JsonEncode(v))
}

func JsonDecode(v interface{}, b []byte) error {
	if err := json.Unmarshal(b, v); err != nil {
		return errors.Trace(err)
	}
	return nil
}

func GzipGobEncode(v interface{}) []byte {
	var buf bytes.Buffer
	w, err := gzip.NewWriterLevel(&buf, gzip.BestCompression)
	if err != nil {
		log.PanicErrorf(err, "encode to gob+gzip failed")
	}
	if err := gob.NewEncoder(w).Encode(v); err != nil {
		log.PanicErrorf(err, "encode to gob+gzip failed")
	}
	if err := w.Close(); err != nil {
		log.PanicErrorf(err, "encode to gob+gzip failed")
	}
	return buf.Bytes()
}

func GzipGobDecode(v interface{}, b []byte) error {
	r, err := gzip.NewReader(bytes.NewReader(b))
	if err != nil {
		return errors.Trace(err)
	}
	defer r.Close()
	if err := gob.NewDecoder(r).Decode(v); err != nil {
		return errors.Trace(err)
	}
	return nil
}
