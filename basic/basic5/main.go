package main

import "fmt"

/*
	数值类型：
		1. 整型
		2. 浮点型
		3. 复数
		4. 指针

	类型名称							字节数				二进制存储位数					取值范围
	整型
	int 								4/8					32/64								-2^31~2^31-1/-2^63-2^63-1
	int8(java byte)			1						8										-2^7-2^7-1(-128~127)
	int16(java short)   2						16									-2^15~2^15-1(-32768~32768)
	int32(int)					4						32									-2^31-2^31-1(-2147483648~2147483648)
	int64(long)					8						64									-2^63-2^63-1()

																											负数取值范围							正数的取值范围
	浮点型
	float32(float)			4						32									-3.4E+38~-1.4E-45				1.4E-45~3.4E+38
	float64(double)			8 					64									-1.7E+308~-4.9E-324			4.9E-324~1.7E+308

	复数
	complex64 	两个 float32 	=> 一个表示 32 位实部, 一个表示 32 位虚部
	complex128	俩个 float64	  => 一个表示 64 位实部, 一个表示 64 位虚部

	有符号 (可表示正负数)的类型: int, int8,  int16, int32, int64
	无符号 (从 0 开始的取值范围)的类型: uint, uint8,  uint16, uint32, uint64,
			uintptr 4/8 个字节 存储指针 无符号整型

	区别：
		1. 取值范围不同
		2. int 内存占用略大于 uint
		3. 使用场景不同

	complex: 一般用于科学计算

	整型默认类型：int(32/64)
	浮点型默认类型：float(64)


	type byte = uint8 => ASCII 码的一个字符  97 a
	type rune	= int32 => 可以表示 unicode 的一个字符
*/

func main() {
	var (
		a int     = 1
		b int8    = 127
		c int16   = 32767
		d int32   = 2147483647
		e float32 = 3.14
		f float64 = 3.1415926
		g byte    = 'a'
		h byte    = 97
		i uint8   = 'a'
		k rune    = '中'
	)

	fmt.Println(a, b, c, d, e, f, g, h, i, k)
	fmt.Printf("%c", g)

	fmt.Printf("%c", k)
}
