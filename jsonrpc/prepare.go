package jsonrpc

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"strings"
)

func prepareBody(r *http.Request) (*bytes.Buffer, bool, error) {
	if !strings.HasPrefix(r.Header.Get(contentType), contentTypeJSON) {
		return nil, false, errors.New("invalid content-type")
	}

	body := bytes.NewBuffer(make([]byte, 0, r.ContentLength))
	if _, err := body.ReadFrom(r.Body); err != nil {
		return nil, false, errors.New("invalid body")
	}

	batch, err := isBatch(body)
	if err != nil {
		return nil, false, errors.New("invalid body")
	}

	return body, batch, nil
}

func prepareRequests(body *bytes.Buffer, batch bool) ([]*request, error) {
	requests := make([]*request, 0)

	if batch {
		if err := json.Unmarshal(body.Bytes(), &requests); err != nil {
			return nil, err
		}
	} else {
		req := new(request)
		if err := json.Unmarshal(body.Bytes(), req); err != nil {
			return nil, err
		}
		requests = append(requests, req)
	}

	return requests, nil
}
