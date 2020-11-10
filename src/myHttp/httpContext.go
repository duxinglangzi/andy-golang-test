package myHttp

type Context struct {
	siteURLHeader string
	RequestId     string
	// ip 地址
	IpAddress     string
	// 路径
	Path          string
	// user agent
	UserAgent     string
	Refer         string
}


