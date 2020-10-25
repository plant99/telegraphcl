package util

import (
	"encoding/json"
	"errors"

	jsoniter "github.com/json-iterator/go"
	http "github.com/valyala/fasthttp"
)

type Response struct {
	Ok     bool            `json:"ok"`
	Error  string          `json:"error,omitempty"`
	Result json.RawMessage `json:"result,omitempty"`
}

func MakeRequest(path string, payload interface{}) ([]byte, error) {
	parser := jsoniter.ConfigFastest

	src, err := parser.Marshal(payload)
	emptyBytes := []byte{}
	if err != nil {
		return emptyBytes, errors.New("Failed to marshal payload")
	}
	u := http.AcquireURI()
	defer http.ReleaseURI(u)
	u.SetScheme("https")
	u.SetHost("api.telegra.ph")
	u.SetPath(path)

	req := http.AcquireRequest()
	defer http.ReleaseRequest(req)
	req.SetRequestURIBytes(u.FullURI())
	req.Header.SetMethod(http.MethodPost)
	req.Header.SetUserAgent("toby3d/telegraph")
	req.Header.SetContentType("application/json")
	req.SetBody(src)

	resp := http.AcquireResponse()
	defer http.ReleaseResponse(resp)

	if err := http.Do(req, resp); err != nil {
		return nil, err
	}

	r := new(Response)
	if err := parser.Unmarshal(resp.Body(), r); err != nil {
		return nil, err
	}

	if !r.Ok {
		return nil, errors.New(r.Error)
	}

	return r.Result, nil
}
