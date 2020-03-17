package config

//数据存储层所需配置字段
type DB struct {
	MySQL Mysql `json:"mysql"`
	Redis Redis `json:"redis"`
}

//MySQL连接参数
type Mysql struct {
	UserName string            `json:"username"` //用户名
	Password string            `json:"password"` //密码
	Protocol string            `json:"protocol"` //协议
	Address  string            `json:"address"`  //地址
	DBName   string            `json:"dbname"`   //数据库名
	Params   map[string]string `json:"params"`   //其他参数
}

//Redis连接参数
type Redis struct {
	Address  string `json:"address"`  //地址
	Password string `json:"password"` //密码
	DBNumber int    `json:"dbnumber"` //数据库编号
}

func NewDBConfig() *DB {
	return &DB{}
}
