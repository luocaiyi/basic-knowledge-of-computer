# Big O Notation

## 高效解

### O(1)

- 常数（constant function）
- 这类算法效率很高
- 不含转向（循环、调用、递归），必须顺序执行，即是O(1)

### O(logn)

- 对数 O(logn)
- 常底数无所谓：log a n = log a b * log b n = c * log b n
- 常数次幂无所谓：(log n)^c = c * log n
- 对数多项式(ploy-log function) 

## 有效解

### O(n^c)

- 多项式(polynomial function)
- 线性(linear function): O(n)
- O(n)~O(n^2):编程习题主要覆盖范围

## 难解

### O(2^n)

- 指数(exponential function)
- 这类算法计算成本增长极快，通常被认为不可接受
- O(n^c)~O(2^n)，是从有效算法到无效算法的分水岭

## 2-Subset

问题描述： s 包含 n 个正整数，ΣS = 2m, 是否有子集T，满足 ΣT = m

直觉算法：逐一枚举S的每一子集，并统计其中元素总和

复杂度：O(2^n)

定理：2-Subset is NP-complete

就目前的计算模型而言，不存在可在多项式时间内回答此问题的算法

就此意义而言上述的直觉算法已属最优
