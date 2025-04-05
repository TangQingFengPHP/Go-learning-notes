package main

func main() {
	// for 数字索引的方式遍历
	echo()

	// for range 方式遍历
	echo2()

	// 使用strings.Join()函数拼接字符串
	echo3()

	// 打印所有参数，包括函数名和参数
	echoAll1()

	// 打印每个参数的索引和值，每个一行
	echoKV1()
}
