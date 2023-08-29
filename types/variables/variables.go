package variables

type Variables struct {
	// 存储所有的变量
	maps map[string]any
}

type String struct {
	key   string
	value string
}

type Int struct {
	key   string
	value int
}

type Bool struct {
	key   string
	value bool
}

// 尝试初始化流中的所有变量
func init() {

}
