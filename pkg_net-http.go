package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// 使用net/http可以对http和https接口进行请求操作，也可以用来创建web服务器
// 这里主要用来进行接口请求处理，不做服务器配置

// 基本的GET请求

func TestGet(url string) {
	res, err := http.Get(url)
	defer res.Body.Close()
	if err != nil {
		println("请求失败")
		return
	} else {
		body, _ := ioutil.ReadAll(res.Body)
		println(res.StatusCode, string(body))
	}
}

// 将参数不放入url中的GET请求，其实也是进行拼接

func GetWithParam() {
	params := url.Values{}
	Url, err := url.Parse("http://39.98.138.157:5000/api/demo?")
	if err != nil {
		return
	}
	params.Set("limit", "1")
	Url.RawQuery = params.Encode() // 如果参数中有中文，使用此方法进行UrlEncode
	urlpath := Url.String()
	fmt.Println(urlpath) // 打印拼接的地址和参数
	res, err := http.Get(urlpath)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
}

// 解析json的返回结果

type result struct {
	Args    string            `json:"args"`
	Headers map[string]string `json:"headers"`
	Origin  string            `json:"origin"`
	Url     string            `json:"url" volid:"123"` // 标签，解析时使用url作为名称，volid可以做验证
}

func JsonParse() {
	res, err := http.Get("http://39.98.138.157:5000/api/demo?limit=1")
	if err != nil {
		return
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
	var rest result
	_ = json.Unmarshal(body, &rest)
	fmt.Printf("%#v", rest)
}

// GET添加请求头

func GetWithHeader() {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://39.98.138.157:5000/api/demo?limit=1", nil)
	req.Header.Add("content-type", "application/json")
	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	fmt.Println(string(body))
}

// 普通的POST请求

func TestPost1() {
	urlval := url.Values{}
	urlval.Add("postKey", "postValue")
	res, err := http.PostForm("http://xxx.com/post", urlval)
	defer res.Body.Close()
	if err != nil {
		return
	}
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
}

func TestPost2() {
	urlValues := url.Values{
		"name": {"Paul_Chan"},
		"age":  {"26"},
	}
	reqBody := urlValues.Encode()
	resp, _ := http.Post("http://xxx.com/post", "text/html", strings.NewReader(reqBody))
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

// 发送json数据的POST请求

func PostWithParam() {
	client := &http.Client{}
	data := make(map[string]interface{})
	data["name"] = "zhaofan"
	data["age"] = "23"
	bytesData, _ := json.Marshal(data) // 序列化成json格式
	req, _ := http.NewRequest("POST", "http://httpbin.org/post", bytes.NewReader(bytesData))
	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

// 不用client进行post数据传递

func PostWithOutClient() {
	data := make(map[string]interface{})
	data["name"] = "zhaofan"
	data["age"] = "23"
	bytesData, _ := json.Marshal(data)
	resp, _ := http.Post("http://httpbin.org/post", "application/json", bytes.NewReader(bytesData))
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

func main() {
	// 普通的GET请求
	//TestGet("http://39.98.138.157:5000/api/demo?limit=1")

	// 分离参数的GET请求
	//GetWithParam()

	// 解析json数据
	//JsonParse()

	// GET添加请求头
	//GetWithHeader()

	// 普通POST（不能运行）
	//TestPost1()

	// post添加json请求数据(不能运行)
	//PostWithParam()

	// Post请求不使用client传递json数据(不能运行)
	//PostWithOutClient()
}
