package main

import "fmt"

// map 是 Go 内置关联数据类型（在一些其他的语言中称为哈希或者字典）
func main() {
	// 创建一个空 map，需要使用内建的 make(map[key-type]val-type).
	m := make(map[string]int)

	// 使用典型的 make[key] = val 语法来设置键值对
	m["k1"] = 7
	m["k2"] = 13

	// 使用例如 Println 来打印一个 map 将会输出所有的键值对
	fmt.Println("map:", m)

	// 使用 name[key] 来获取一个键的值
	v1 := m["k1"]
	fmt.Println("v1: ", v1)

	// 当对一个 map 调用内建函数 len 时，返回的是键值对数目
	fmt.Println("len:", len(m))

	// 内建函数 delete 可以从一个 map 中移除键值对
	delete(m, "k2")
	fmt.Println("map:", m)

	// 当从一个 map 中取值时，可以选择是否接收的第二个返回值，该值表明了 map 中是否存在这个键
	// 这可以用来消除键不存在和键有零值产生的歧义，像 0 或者 ""
	// 这里不需要值，所以用 空白标识符(blank identifier) _ 将其忽略
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	// 可以通过这个语法在同一行声明并初始化一个新的map
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
	// 注意一个 map 在使用 fmt.Println 打印的时候，是以 map[k:v k:v]的格式输出的
}
