package Tools

import (
	"gopkg.in/h2non/gentleman.v2"
	"gopkg.in/h2non/gentleman.v2/plugins/body"
	"gopkg.in/h2non/gentleman.v2/plugins/bodytype"
	"gopkg.in/h2non/gentleman.v2/plugins/query"
	"gopkg.in/h2non/gentleman.v2/plugins/timeout"
	"strings"
	"time"
)

type HttpClient struct {
	Host         string            //域名或者IP
	Url          string            //访问路径
	Header       map[string]string //请求头参数
	Data         interface{}       //请求数据
	Code         int               //响应编码
	Msg          string            //响应信息提示
	ResponseData string            //响应数据
}

func NewHttpClient(host, url string, data interface{}, header map[string]string) *HttpClient {
	return &HttpClient{
		Host:   host,
		Url:    url,
		Data:   data,
		Header: header,
	}
}

func (h *HttpClient) getClient() *gentleman.Request {
	//创建客户端
	cli := gentleman.New()
	//请求超时时间
	cli.Use(timeout.Request(10 * time.Second))
	//请求拨号时间与活跃时间
	cli.Use(timeout.Dial(5*time.Second, 30*time.Second))
	cli.URL(h.Host)
	req := cli.Request()
	req.Path(h.Url)
	if len(h.Header) > 0 {
		req.SetHeaders(h.Header)
	}
	return req
}

// SendGet 发送Get方法 支持map[string]string
func (h *HttpClient) SendGet() *HttpClient {
	req := h.getClient()
	req.Method("GET")
	if h.Data != nil {
		data, ok := h.Data.(map[string]string)
		if ok {
			req.Use(query.SetMap(data))
		} else {
			h.Code = REQUEST_ERROR
			h.Msg = "请求参数类型转换错误，Get请求参数类型必须为：map[string]string"
		}
	}

	res, err := req.Send()
	if err != nil {
		h.Code = REQUEST_ERROR
		h.Msg = err.Error()
		return h
	}
	if !res.Ok {
		h.Code = res.StatusCode
		h.Msg = res.Error.Error()
		return h
	}
	h.Code = REQUEST_SUCCESS
	h.ResponseData = res.String()

	return h
}

// SendPost 发送数据application/x-www-form-urlencoded类型的Post方法 支持a=b&c=d字符串、map[string]string
func (h *HttpClient) SendPost() *HttpClient {
	req := h.getClient()
	req.Method("POST")
	var data string
	if h.Data != nil {
		switch h.Data.(type) {
		case string:
			data = h.Data.(string)
		case map[string]string:
			for key, value := range h.Data.(map[string]string) {
				data = data + key + "=" + value + "&"
			}
			data = strings.TrimRight(data, "&")
		}
	}
	req.Use(bodytype.Type("form-data"))
	res, err := req.Send()
	if err != nil {
		h.Code = REQUEST_ERROR
		h.Msg = err.Error()
		return h
	}
	if !res.Ok {
		h.Code = res.StatusCode
		h.Msg = res.Error.Error()
		return h
	}

	h.Code = REQUEST_SUCCESS
	h.ResponseData = res.String()

	return h
}

// SendJsonPost 发送json数据类型的Post方法 支持结构体、`{"a":1,"v":"d"}`、map[string]string
func (h *HttpClient) SendJsonPost() *HttpClient {
	req := h.getClient()
	req.Method("POST")

	if h.Data != nil {
		req.Use(body.JSON(h.Data))
	}

	res, err := req.Send()
	if err != nil {
		h.Code = REQUEST_ERROR
		h.Msg = err.Error()
		return h
	}
	if !res.Ok {
		h.Code = res.StatusCode
		h.Msg = res.Error.Error()
		return h
	}

	h.Code = REQUEST_SUCCESS
	h.ResponseData = res.String()
	return h
}
