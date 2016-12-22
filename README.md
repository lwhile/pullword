# pullword

pullword中文分词Go语言客户端

## Install
> go get github.com/lwhhhh/pullword

## example

    func main() {
	    request := pullword.NewRequest("感谢pullword的中文分词服务", false, true)
	    resM, err := request.GetM()
	    if err != nil {
		    log.Fatal(err)
	    }
	    for k, v := range resM {
		    fmt.Println(k, v)
	    }
    }    

## API document

- > func NewRequest(source string, strict bool, debug bool) *Request

    - source: 待带查询的文本,
    - strict: 是否只返回概率为100%的分词结果
    - debug: 是否返回每个分词结果的概率


- > func (request *Request) GetM() (map[string]float64, error) 

    在debug模式下,返回一个map,其中key为分词结果,value为其概率

- > func (request *Request) GetS() ([]string, error)

    在非debug模式下,返回一个分词结果的切片