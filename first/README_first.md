## ways to the first work 

- [解方程](https://www.luogu.com.cn/problem/P1022)
  输入的同时记录系数，注意未知数没有系数的情况，如`+a`或`a`，同时记录下系数的符号以及标记未知数对应的系数，
  通过标记计算出未知数的系数以及常数项的和，然后移项相除
- [The Tamworth Two](https://www.luogu.com.cn/problem/P1518)
  一直循环直到相遇或者达成没有解的情况break并输出步数，  
  边界条件在于直到回到以前走过的某个条件或者走的步数超过一定阙值也算进入死循环。
  emmm，有点bug，没想出来，先提交了吧。
- [凯撒密码](https://www.luogu.com.cn/problem/P1914)
  直接对字符的ascii值进行操作。如果移动位数，没有超过小写字符`a`-`z`的范维直接ascii值加1既可；如果超过直接取模或者差值

