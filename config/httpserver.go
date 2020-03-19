package config

type HTTPServer struct {
	Host      string `json:"host"`      //主机
	IPVersion string `json:"ipVersion"` //ip版本
	Port      uint32 `json:"port"`      //监听端口
	Name      string `json:"name"`      //服务器名
	Version   string `json:"version"`   //服务器版本
	Mode      string `json:"mode"`      //运行模式debug/release
}

func NewHTTPServerConfig() *HTTPServer {
	return &HTTPServer{}
}
