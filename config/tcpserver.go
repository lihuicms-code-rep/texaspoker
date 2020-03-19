package config

//服务器框架层所需配置字段
//约定,json字段部分也按驼峰
type TCPServer struct {
	Host             string `json:"host"`             //主机
	IPVersion        string `json:"ipVersion"`        //ip版本
	Port             uint32 `json:"port"`             //监听端口
	Name             string `json:"name"`             //服务器名
	Version          string `json:"version"`          //服务器版本
	MaxConn          uint32 `json:"maxConn"`          //单服最大连接数
	MaxPackageSize   uint32 `json:"maxPackageSize"`   //单包最大字节数
	WorkerPoolSize   uint32 `json:"workerPoolSize"`   //工作池数
	MaxWorkerTaskLen uint32 `json:"maxWorkerTaskLen"` //工作goroutine最大任务队列数
}

func NewTCPServerConfig() *TCPServer {
	return &TCPServer{}
}
