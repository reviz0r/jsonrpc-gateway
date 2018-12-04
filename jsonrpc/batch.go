package jsonrpc

import (
	"bytes"
	"errors"
)

func isBatch(buf *bytes.Buffer) (bool, error) {
	if buf == nil {
		return false, errors.New("buffer is nil")
	}

	f, _, err := buf.ReadRune()
	if err != nil {
		return false, err
	}

	if err := buf.UnreadRune(); err != nil {
		return false, err
	}

	return f == batchKey, nil
}
