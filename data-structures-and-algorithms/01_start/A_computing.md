# Computing

## Hailstone

```text
                 {1}                    n <= 1
Hailstone(n) =   {n} U Hailstone(n/2)   n 偶
                 {n} U Hailstone(3n+1)  n 奇

Hailstone(42) = { 42, 21, 64, 32, ..., 1 }

目前还没有办法证明是否所有的数都满足这个算法是有穷的。
也就是说既没有找到无穷的那个数，也无法证明那个数不存在。
所以这并不能算一个算法，即程序不等于算法。
```

```go
package main

// hailstone 计算Hailstone()的长度
func hailstone(n int) int {
	length := 1
	for ; n > 1; length++ {
		if n%2 == 1 {
			n = 3 * n + 1
		} else {
			n = n / 2
		}
	}
	return length
}
```

## 好算法

- 正确
- 健壮
- 可读： 结构化 + 准确命名 + 注释
- 效率
