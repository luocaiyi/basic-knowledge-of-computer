# Computing Model

## 图灵机

TM: Turing Machine

- Tape: 依次均匀地划分为单元格，各注有某一字符，默认为'#'
- Alphabet: 字符的种类有限
- Head: 总是对准某一单元格，并可读取和改写其中的字符，每经过一个节拍，可转向左侧或右侧的邻格
- State: TM总是处于有限状态中的某一种，每经过一个节拍，可(按照规则)转向另一种状态。
- Transition Function: (q, c; d, L/R, p) 若当前状态为q且当前字符为c，则将当前字符改写为d；转向左侧/右侧的邻格；转入p状态。一旦转入特定状态'h'，则停机

## RAM模型

RAM: Ramdom Access Machine

- 寄存器顺序编号，总数没有限制
- 每一种基本操作仅需要常熟时间

    ```text
    R[i] <- c
    R[i] <- R[j]
    R[i] <- R[R[j]]
    R[R[i]] <- R[j]
    // 循环及子程序本身非基本操作
    R[i] <- R[j] + R[k]
    R[i] <- R[j] - R[k]
    IF R[i] = 0 GOTO 1
    IF R[i] > 0 GOTO 1
    GOTO 1
    STOP
    ```
  
- 与TM模型一样，RAM模型也是对一般计算工具的简化与抽象，是我们可以独立于具体的平台，对算法效率做出可信的比较与评测

