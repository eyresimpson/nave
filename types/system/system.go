package system

// 数据库配置
type DbConf struct {
	// 配置ID
	id int64
	// 配置名
	name string
	// 数据库URL
	url string
	// 数据库类型
	dbType string
	// 数据库编码
	encoding string
	// 数据库端口
	port int64
	// 指定表（可选，仅在exec为true的情况下此项生效）
	table string
	// 用户名
	user string
	// 密码
	passwd string
	// 密钥
	secret string
	// 是否为可执行数据库配置，如果为true，如果指定了table，将以此库为基础表执行单表执行操作，如果未指定，将自行创建默认数据库和数据表
	exec bool
}
