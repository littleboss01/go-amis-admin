package amis

import "encoding/json"

type AmisPage interface {
	AddBody(i interface{})
}

func NewPage(title string) *Page {
	return &Page{
		Id:    "",
		Type:  "page",
		Title: title,
		Body:  nil,
		Data:  nil,
		opt:   map[string]interface{}{},
	}
}

type Opt interface {
	SetOptions(k string, v interface{})
}

type Page struct {
	Id      string      `json:"id,omitempty"`
	Type    string      `json:"type,omitempty"`
	Title   string      `json:"title,omitempty"`
	Body    interface{} `json:"body,omitempty"`
	OnEvent *OnEvent    `json:"onEvent,omitempty"`
	InitApi *InitApi    `json:"initApi,omitempty"`
	Data    interface{} `json:"data,omitempty"`

	opt map[string]interface{}
}

func (p *Page) SetOptions(k string, v interface{}) {
	p.opt[k] = v
}

// MarshalJSON opt 会合并到整个ColumnConfig上再输出到前端
func (p *Page) MarshalJSON() ([]byte, error) {
	newStruct := *p
	if len(p.opt) == 0 {
		return json.Marshal(newStruct)
	}
	mm := map[string]interface{}{}
	by, _ := json.Marshal(newStruct)
	_ = json.Unmarshal(by, &mm)
	for k, v := range p.opt {
		mm[k] = v
	}
	return json.Marshal(mm)
}

type InitApi struct {
	Url    string `json:"url"`
	Method string `json:"method"`
}

type OnEvent struct {
	Init struct {
		Weight  int           `json:"weight"`
		Actions []interface{} `json:"actions"`
	} `json:"init"`
	Inited struct {
		Weight  int           `json:"weight"`
		Actions []interface{} `json:"actions"`
	} `json:"inited"`
}

func (p *Page) AddBody(i interface{}) {
	if _, ok := p.Body.([]interface{}); !ok {
		p.Body = []interface{}{}
	}
	p.Body = append(p.Body.([]interface{}), i)
}

// curd是单独的,这里设置了也无效 //TODO 增加一个读取加密json的函数
func (p *Page) AddBodyByJsonStr(s string) {
	var v interface{}
	err := json.Unmarshal([]byte(s), &v)
	if err != nil {
		println(err.Error())
		return
	}

	switch v := v.(type) {
	case map[string]interface{}:
		// JSON 是对象
		p.AddBody(v)
	case []interface{}:
		// JSON 是数组
		for _, item := range v {
			p.AddBody(item)
		}
	default:
		// 其他类型，可以按需处理或忽略
		println("Unexpected JSON type")
	}
}
