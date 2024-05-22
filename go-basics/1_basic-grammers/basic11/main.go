package main

/**
赋值运算
	= 赋值符号
	a = 1
	二元运算
	a = a + 1 => a ++

	求和赋值
	x = 10
	二元运算			一元运算
	a = a + x => a += x

	求差赋值
	a = a - x => a -= x

	求积赋值
	a = a * x => a *= x

	求商赋值
	a = a / x => a /= x

	求余赋值
	a = a % x => a %= x
*/

/**
关系运算
	所有的关系运算都会返回出布尔值
		true 成立， false 不成立

	> 大于
	< 小于
	>= 大于等于
	<= 小于等于
	== 相等
	!= 不想等
*/

/*
*
逻辑运算

		所有逻辑运算都会返回布尔值，真假条件

		&& 与运算

			a && b => true => a: true && b: true  => 两个同时为 true
	 					 => false => 如果有一个为 false 或者 两个都为 false

			总结：a 和 b 条件的返回值都为真，那么与运算返回值返回真
					 有一个或者多个为假，与运算返回假

			a && b: 如果 a 为 false 停止 => 返回假


		|| 或运算
			a || b => true => 有一个或两个为 true
	  				 => false => 两个都为false

			总结：a 和 b 的条件返回值有一个或者多个为真，或运算返回为真
	  			 全部都为假，或运算返回为假
			a || b: a => true => 停止 => 返回真

		总结：
			与运算中，有一个为假，就返回假
							全部为真，就返回真
			或运算中，有一个为真，就返回真
							全部为假，就返回假

	 	!运算
			a => true
			!a => false
*/
func main() {
	// a := 1
	// b := 2

	// a += b // a = a + b
	// fmt.Println(a)
	//
	// a -= b // a = a -b
	// fmt.Println(a)
	//
	// a *= b // a = a * b
	// fmt.Println(a)
	//
	// a /= b // a = a / b
	// fmt.Println(a)
	//
	// a %= b // a = a % b
	// fmt.Println(a)

	// fmt.Println(a > b) // => false
	// fmt.Println(a < b) // => true

	// a := 1
	// b := 1
	//
	// // > < 或者 = 满足其一返回 true
	// fmt.Println(a >= b) // => true
	// fmt.Println(a <= b) // => true
	//
	// fmt.Println(a == b) // => true
	// fmt.Println(a != b) // => false

	// -------------------------------------
	// a := 1
	// b := 2

	// temp := a
	// a = b
	// b = temp
	//
	// fmt.Println(a, b)

	// a = a + b
	// b = a - b
	// a = a - b
	//
	// fmt.Println(a, b)

	// a := "abc"
	// b := "abc"
	//
	// fmt.Println(a == b)

	// a := "abc"
	// b := 1
	// // 无效运算: a == b(类型 string 和 int 不匹配)
	// fmt.Println(a == b)

	// a := 1
	// b := 1.0
	// 无效运算: a == b(类型 int 和 float64 不匹配)
	// fmt.Println(a == b)
	// fmt.Println(a != b)

	// ------------------------------------------

	// a := 1
	// b := 2
	// c := 3
	//
	// fmt.Println(a > 1 && b > 1)  // => false
	// fmt.Println(a >= 1 && b > 1) // => true
	//
	// fmt.Println(a >= 1 && b > 2 && c >= 3) // => false
	//
	// fmt.Println(a >= 1 && b >= 2 && c > 3) // => false
	//
	// fmt.Println(a >= 1 || b > 2 || c > 3) // => true
	//
	// fmt.Println(a > 1 || b >= 2 || c > 3) // => true

	/**
	a > 1 => false
	b > 2 => false
	c > 3 => false
	=> false
	*/
	// fmt.Println(a > 1 || b > 2 || c > 3) // => false

	/**
	a >= 1 => true
	true
	c >= 3 => true
	=> true
	*/
	// fmt.Println(a >= 1 && c >= 3 || b > 2) // => true
	//
	// fmt.Println(b > 2 && c >= 3 || a >= 1) // => false
	//
	// fmt.Println(b >= 2 && (a > 1 || c > 3)) // => false
	//
	// flag := false
	// flag = true
	//
	// fmt.Println(!flag) // => false

}
