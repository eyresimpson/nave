package blueprint

type Worker struct {
	// 执行器主键
	Sid string `json:"sid"`
	// 下一个需要执行的节点
	Next string `json:"next"`
	// 执行路径，注意逻辑操作的这个值固定为空，被编译器忽略
	Path string `json:"path"`
	// 类型
	// 0：逻辑操作，流重定向
	// 1：Restful操作，通过原生rw返还数据到请求
	// 2：无返回操作，常见于CRON流
	Type int8 `json:"type"`
	// 名称
	Name string `json:"name"`
	// 子节点块，仅在逻辑语句或组合语句时会被解析
	Children []Worker `json:"children"`
	// 条件（一般opt满足此条件，将执行，对于逻辑opt，会执行分流操作）
	Condition []string `json:"condition"`
	// 执行参数
	OptParams []string `json:"optParams,omitempty"`
	// 执行返回，逻辑操作返回固定为空
	OptReturn []any `json:"optReturn,omitempty"`
}
