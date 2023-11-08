package main

func main() {

	/**
	  1. 什么是哈希（Hash）?
			{ myName: 'xiaoyesensen' }
			input(key)                                    output(摘要)
			‘myName’ ===> hash 算法 (createHash(key)) ===> hash 值
				1. 相同的 key 产生相同的 hash 值（hash的单一性）
				2. 理论上无穷大的 key 都可以通过哈希算法
						得到一个一定大的 hash 值 (hash 的压缩性)
				3. 映射就是 key 值与 hash 值之间的对应关系(hash 的映射性)
		    4. 不同的 hash 值对应的数据存储在不同的空间内
				5. hash 算法是不能从 hash 值返推回 key 值
				6.key 的可能性是无限的，hash 值是有限的
					两个不同的 key 通过 hash 算法后可能得到相同的 hash 值（hash 的冲突性）
	*/
}
