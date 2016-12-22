package pullword

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

const (
	HTTP   = "http://"
	HTTPS  = "https://"
	APIURL = "api.pullword.com/get.php"
)

type Request struct {
	source string
	param1 int
	param2 int
}

func (request *Request) Get() string {
	addr := fmt.Sprintf(HTTP+APIURL+"?"+"source=%s&param1=%d&param2=%d", request.source, request.param1, request.param2)
	resp, err := http.Get(addr)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return string(body)
}

func (request *Request) Post() {

}

// debug模式下返回map
// 非debug返回slice
func (request *Request) result() (interface{}, error) {
	res := request.Get()
	resElements := strings.Split(res, "\r\n")
	if request.param2 == 0 {
		return resElements, nil
        
	} else if request.param2 == 1 {
		m := make(map[string]float64)
		for _, resEle := range resElements {
			eles := strings.Split(resEle, ":")
			if len(eles) == 2 {
				word, probabilityStr := eles[0], eles[1]
				probabilityFloat, err := strconv.ParseFloat(probabilityStr, 32)
				if err != nil {
					return nil, errors.New(fmt.Sprint(err))
				}
				m[word] = probabilityFloat
			}
		}
		return m, nil
	}
	return nil, errors.New("request.param2 error")
}

func (request *Request) GetM() (map[string]float64, error) {
	res, err := request.result()
	if m, ok := res.(map[string]float64); ok {
		return m, err
	} else {
		return nil, err
	}
}

func (request *Request) GetS() ([]string, error) {
	res, err := request.result()
	if s, ok := res.([]string); ok {
		return s, err
	} else {
		return nil, err
	}
}

func NewRequest(source string, strict bool, debug bool) *Request {
	request := Request{source: source}
	if strict {
		request.param1 = 1
	} else {
		request.param1 = 0
	}

	if debug {
		request.param2 = 1
	} else {
		request.param2 = 0
	}
	return &request
}
