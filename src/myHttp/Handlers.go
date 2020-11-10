package myHttp

import (
	"andy-golang-test/src/utils"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// 自定义个 Handler 结构体
type MyHandler struct{}

// 注册所有的 请求地址
var mux = map[string]func(*Context, http.ResponseWriter, *http.Request){
	"/test": testRequest,
}

// 实现 serve 接口， 用于监控 和封装一些基本信息
func (h MyHandler) ServeHTTP(responseWriter http.ResponseWriter, req *http.Request) {

	context := &Context{}
	context.IpAddress = utils.GetIpAddress(req, []string{})
	context.UserAgent = req.UserAgent()
	context.Refer = req.Referer()
	context.Path = req.URL.Path

	// 如果不包含则直接返回404
	if handlerFunc, ok := mux[context.Path]; !ok {
		responseWriter.WriteHeader(404)
		responseWriter.Write([]byte(`{"status":"fail"}`))
		return
	} else {
		handlerFunc(context, responseWriter, req)
	}

	fmt.Printf("当前访问路径是: %v \n", context.Path)
	fmt.Printf("当前UserAgent是: %v \n", context.UserAgent)

	// 如果包含则执行函数
	//h.HandleFunc(c, &responseWriter, req)

}

func testRequest(c *Context, w http.ResponseWriter, r *http.Request) {
	r.ParseForm()                      //解析url传递的参数，对于POST则解析响应包的主体（request body）
	fmt.Fprintf(w, "service start...") //这个写入到w的是输出到客户端的 也可以用下面的 io.WriteString对象
	fmt.Printf("当前访问的 ip : %v", c.IpAddress)

	r.ParseForm()       //解析url传递的参数，对于POST则解析响应包的主体（request body）
	fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	// 打印header 信息
	for i, ele := range r.Header {
		fmt.Println(i, ":", ele)
	}
	// 打印 请求参数信息
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("value:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "service start...") //这个写入到w的是输出到客户端的 也可以用下面的 io.WriteString对象

	//注意:如果没有调用ParseForm方法，下面无法获取表单的数据
	query := r.URL.Query()
	fmt.Println(query)
	var uid string // 初始化定义变量
	if r.Method == "GET" {
		uid = r.FormValue("uid")
	} else if r.Method == "POST" {
		uid = r.PostFormValue("uid")
	}
	io.WriteString(w, "uid = "+uid)
}
