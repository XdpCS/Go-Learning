package main

import "fmt"

// Slice 是 Go 中一个重要的数据类型，它提供了比数组更强大的序列交互方式
func main() {
	// slice 的类型仅由它所包含的元素决定（与元素个数无关)
	// 要创建一个长度非零的空slice，需要使用内建的方法 make
	// 这里我们创建了一个长度为3的 string 类型 slice（初始化为零值）
	s := make([]string, 3)
	fmt.Println("emp:", s)

	// 可以和数组一样设置和得到值
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	// len 返回 slice 的长度
	fmt.Println("len:", len(s))

	// slice 支持比数组更多的操作
	// 其中一个是内建的 append，会返回一个包含了一个或者多个新值的 slice
	// 注意由于 append 可能返回一个新的 slice
	// 所以需要接收其返回值
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	// Slice 也可以被 copy
	// 这里创建一个空的和 s 有相同长度的 slice c，并且将 s 复制给 c。
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	// Slice 支持通过 slice[low:high] 语法进行“切片”操作
	// 例如，得到一个包含元素 s[2], s[3],s[4] 的 slice。
	l := s[2:5]
	fmt.Println("sl1", l)

	// 这个 slice 从 s[0] 到（但是不包含）s[5]
	l = s[:5]
	fmt.Println("sl2:", l)

	// 这个 slice 从（包含）s[2] 到 slice 的后一个值
	l = s[2:]
	fmt.Println("sl3:", l)

	// 可以在一行代码中声明并初始化一个 slice 变量
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	// Slice 可以组成多维数据结构
	// 内部的 slice 长度可以不同，这和多位数组不同。
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	// 注意，slice 和数组不同，虽然它们通过 fmt.Println 输出差不多
	fmt.Println("2d: ", twoD)
}
