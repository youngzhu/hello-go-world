# 学习过程中的疑问

_注：答案不是100%准确，是基于仅有的经验的总结。_

## Q1: build/run命令和gopath及项目目录之间的关系

## Q2: *和&的区别和用途

## Q3: package和目录名必须一致吗？
必须一致。可参见ch03 SMP，根据原文应该是path：`mlib`，package：`libaray`，但怎么都编译不通过。将path改成`libaray`就好了。  
main则不必。

~~其实想多了，main包就不在main目录下。~~  以偏概全了
