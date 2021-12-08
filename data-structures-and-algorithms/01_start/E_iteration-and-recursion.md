# Iteration and Recursion

## 迭代与递归

### 数组求和：迭代

问题：计算任意n个整数之和

见`sumi/SumI`

## 减而治之

Decrease-and-conquer: 为求解一个大规模的问题，可以将其划分为两个子问题，其一平凡，另一规模缩减，分别求解子问题，由子问题的解，得到原问题的解

见`sumi/SumDecrease`

递归跟踪(recursion trace) 分析：检测每个示例，累计所需时间即算法执行时间。

- 递归跟踪：直观形象，仅适用于简明的递归模式
- 递推方程：间接抽象，更适用于复杂递归模式

## 分而治之

Divide-and-conquer: 为求解一个大规模的问题，可以将其划分为若干（通常两个）子问题，规模大体相当，分别求解子问题，由子问题的解，得到原问题的解

见`sumi/SumDivide`

## 大师定理

[Master Theorem](https://en.wikipedia.org/wiki/Master_theorem_(analysis_of_algorithms))

[主定理](https://zh.wikipedia.org/wiki/%E4%B8%BB%E5%AE%9A%E7%90%86)

- 分治策略对应的递推式，通常（尽管不总是）形如：T(n) = a*T(n/b) + O(f(n))
  (原问题被分为 a 个规模为 n/b 的子任务；任务的划分、解的合并耗时f(n))
- 若 f(n) = O(n^(log b a - ε))，则 T(n) = Θ(n^(log b a))
  - kd-search: T(n) = 2*T(n/4) + O(1) = O(n^(1/2))
- 若 f(n) = Θ(n^(log b a) * (log n)^k)，则T(n) = Θ(n^(log b a) * (log n)^(k+1))
  - binary search: T(n) = 1*T(n/2) + O(1) = O(log n)
  - mergesort: T(n) = 2*T(n/2) + O(n) = O(n * log n)
  - SLT mergesort: T(n) = 2*T(n/2) + O(n * log n) = O(n * (log n)^2)
- 若 f(n) = Ω(n^(log b a + ε))，则 T(n) = Θ(f(n))
  - quickSelect (average case): T(n) = 1*T(n/2) + O(n) = O(n)
