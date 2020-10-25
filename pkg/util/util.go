package util

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os/user"
	"path"

	jsoniter "github.com/json-iterator/go"
	http "github.com/valyala/fasthttp"
)

func getTokenPath() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	return path.Join(usr.HomeDir, "telegraph.token")
}

var tokenPath string = getTokenPath()

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

func StoreAccessToken(token string) (int, error) {
	err := ioutil.WriteFile(tokenPath, []byte(token), 0644)
	if err != nil {
		fmt.Println(err)
		return 1, errors.New("Couldn't save token :(")
	}
	return 0, nil
}

func FetchAccessToken() (string, error) {
	data, err := ioutil.ReadFile(tokenPath)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
