package main

/*
*

		++ 自加 1 自增1
	  -- 自减 1

		1. Go 语言中 ++ -- 不是表达式 => Expression 不能参与运算或者赋值
		2. ++ -- 是语句 statement，只能单独出现
		3. ++ -- 在变量后，不能出现在变量前，推荐使用在变量后 ++ --
*/
func main() {
	a := 1
	// b := a ++ // ❌ ++ -- 是不能参与其他运算或者赋值的
	// b := ++ a
	// a++ // => a = a + 1
	a--
	b := a
	println(b)
}
