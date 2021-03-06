# Python 格式化输出

## 一、print函数的格式控制的完整格式：

```sh
%     -     0     m.n     l或h     格式字符

下面对组成格式说明的各项加以说明：

①%：表示格式说明的起始符号，不可缺少。

②-：有-表示左对齐输出，如省略表示右对齐输出。

③0：有0表示指定空位填0,如省略表示指定空位不填。

④m.n：m指域宽，即对应的输出项在输出设备上所占的字符数。N指精度。用于说明输出的实型数的小数位数。为指定n时，隐含的精度为n=6位。

⑤l或h:l对整型指long型，对实型指double型。h用于将整型的格式字符修正为short型。
```

---

### 格式字符

**格式字符用以指定输出项的数据类型和输出格式。**

> ①d格式：用来输出十进制整数。有以下几种用法：
>
> ​	%d：按整型数据的实际长度输出。
>
> ​	%md：m为指定的输出字段的宽度。如果数据的位数小于m，则左端补以空格，若大于m，则		     按实际位数输出。
>
> ​	%ld：输出长整型数据。
>
> ②o格式：以无符号八进制形式输出整数。对长整型可以用"%lo"格式输出。同样也可以指定字段宽        度用“%mo”格式输出。

例：

```python
# 进制输出
>>> print('%o' %20)		# 以octonary 形式输出
24
>>> print('%x' %20)		# 以hexadecimal 形式输出
14
>>> print('%d' %20)		# 以decimalism 形式输出
20



# 浮点输出，当n > m 时默认可输出全部内容，当域宽大于位数时，补0或者补空格
>>> print('%f' %1.11)		# 输出浮点数，默认保留6位小数
1.110000
>>> print('%lf' %1.11)		# 双精度浮点数，默认保留6位小数
1.110000
>>> print('%20.5f' %3223623.11111656)		# 域宽20，保留5位小数的浮点数输出
       3223623.11112
>>> print('%20.5lf' %3223623.11111656)		# 同上为双浮点数
       3223623.11112
>>> print('%020.5lf' %3223623.11111656)		# 域宽20空位补0,5为小数的浮点输出
00000003223623.11112
>>> print('%-20.5lf' %3223623.11111656)		# 左对齐，域宽20空位补0,5为小数的浮点输出
3223623.11112       




# 指数以及保留相应有效数字
>>> print('%g' %1.111)		# 有效数字少于6位则直接输出
1.111
>>> print('%g' %1.111315132213)		# 保留6位有效数字
1.11132
>>> print('%10.7g' %1.111)      
     1.111
>>> print('%10.7g' %11111111.111)
1.111111e+07
>>> print('%e' %1.11)
1.110000e+00
>>> print('%0.2g' %1.111)
1.1
>>> print('%0.2g' %1111.111)
1.1e+03
>>> print('%2.3e' %1.11111)		# 保留3位小数的科学计数法表示
1.111e+00



 
```
> ③x格式：以无符号十六进制形式输出整数。对长整型可以用"%lx"格式输出。同样也可以指定字段宽度用"%mx"格式输出。
>
> ④u格式：以无符号十进制形式输出整数。对长整型可以用"%lu"格式输出。同样也可以指定字段宽度用“%mu”格式输出。
> ⑤c格式：输出一个字符。
> ⑥s格式：用来输出一个串。有几中用法
> %s：例如:printf("%s", "CHINA")输出"CHINA"字符串（不包括双引号）。
> %ms：输出的字符串占m列，如字符串本身长度大于m，则突破获m的限制,将字符串全部输出。若串长小于m，则左补空格。
> %-ms：如果串长小于m，则在m列范围内，字符串向左靠，右补空格。
> %m.ns：输出占m列，但只取字符串中左端n个字符。这n个字符输出在m列的右侧，左补空格。
> %-m.ns：其中m、n含义同上，n个字符输出在m列范围的左侧，右补空格。如果n>m，则自动取n值，即保证n个字符正常输出。
> ⑦f格式：用来输出实数（包括单、双精度），以小数形式输出。有以下几种用法：
> %f：不指定宽度，整数部分全部输出并输出6位小数。
> %m.nf：输出共占m列，其中有n位小数，如数值宽度小于m左端补空格。 
> %-m.nf：输出共占n列，其中有n位小数，如数值宽度小于m右端补空格。
> ⑧e格式：以指数形式输出实数。可用以下形式：
> %e：数字部分（又称尾数）输出6位小数，指数部分占5位或4位。
> %m.ne和%-m.ne：m、n和”-”字符含义与前相同。此处n指数据的数字部分的小数位数，m表示整个输出数据所占的宽度。
> ⑨g格式：自动选f格式或e格式中较短的一种输出，且不输出无意义的零。



---

```c
关于c语言printf函数的进一步说明：

如果想输出字符"%",则应该在“格式控制”字符串中用连续两个%表示，如:

printf("%f%%", 1.0/3);

输出0.333333%。
对于单精度数，使用%f格式符输出时，仅前7位是有效数字，小数6位．
对于双精度数，使用%lf格式符输出时，前16位是有效数字，小数6位．

对于m.n的格式还可以用如下方法表示（例）
char ch[20];
printf("%.s\n",m,n,ch);
前边的*定义的是总的宽度，后边的定义的是输出的个数。分别对应外面的参数m和n 。我想这种方法的好处是可以在语句之外对参数m和n赋值，从而控制输出格式


```

---

## 二、round函数

```python
round(x[, n])
对于x的值采用“四舍六入五看齐，奇进偶不进”的近似策略，如果第二个参数不存在时，默认保留到整数位例如：
>>> round(10.50)
10
>>> round(10.51)
11
>>> round(11.50)
12

第二个参数表示保留的小数位数，但是由于浮点数无法在计算机中精确表示，所以不能完全满足例如四舍六入五看齐“，所以不可用于对精度要求过高的场合
>>> round(5.215, 2)
5.21
>>> round(5.225, 2)
5.22
>>> round(5.235, 2)
5.24
>>> round(5.255, 2)
5.25
>>> round(5.245, 2)
5.25
ps：如果保留位数的后一位如果是5，且该位数后有数字。则进上去，例如5.2152保留两位小数为5.22，5.2252保留两位小数为5.23，5.22500001保留两位小数为5.23。
>>> round(5.2152, 2)		# 满足精度要求，5.2152 离5.22更近
5.22
>>> round(5.2252, 2)
5.23
```

doc原文解释：

```
Note

The behavior of round() for floats can be surprising: for example, round(2.675, 2) gives 2.67 instead of the expected 2.68. This is not a bug: it’s a result of the fact that most decimal fractions can’t be represented exactly as a float. See Floating Point Arithmetic: Issues and Limitations for more information.
```

# 三、format函数

> Python2.6 开始，新增了一种格式化字符串的函数 str.format()，它增强了字符串格式化的功能。
>
> 基本语法是通过 {} 和 : 来代替以前的 % 。
>
> format 函数可以接受不限个参数，位置可以不按顺序。

```python
>>>"{} {}".format("hello", "world")    # 不设置指定位置，按默认顺序
'hello world'

>>> "{0} {1}".format("hello", "world")  # 设置指定位置
'hello world'
 
>>> "{1} {0} {1}".format("hello", "world")  # 设置指定位置
'world hello world'
```

**设置参数：**

```python
>>> print('Name:{n}   Addr:{a}'.format(n='Jack Ma', a='www.inno.com'))		
Name:Jack Ma   Addr:www.inno.com
>>> list = ['Jack Ma', 'www.inno.com']		# 利用列表设置
>>> print('Name:{0[0]}   Addr:{0[1]}'.format(list))
Name:Jack Ma   Addr:www.inno.com

>>> site = {'Name':'Jack Ma', 'Url':'www.mayun.com'}# 利用字典设置，注意format参数
>>> print('Name:{Name}   Addr:{Url}'.format(**site))
Name:Jack Ma   Addr:www.mayun.com


```

**数字格式化：**

| 数字       | 格式                                                         | 输出                                             | 描述                         |
| ---------- | ------------------------------------------------------------ | ------------------------------------------------ | ---------------------------- |
| 3.1415926  | {:.2f}                                                       | 3.14                                             | 保留小数点后两位             |
| 3.1415926  | {:+.2f}                                                      | +3.14                                            | 带符号保留小数点后两位       |
| -1         | {:+.2f}                                                      | -1.00                                            | 带符号保留小数点后两位       |
| 2.71828    | {:.0f}                                                       | 3                                                | 不带小数                     |
| 5          | {:0>2d}                                                      | 05                                               | 数字补零 (填充左边, 宽度为2) |
| 5          | {:x<4d}                                                      | 5xxx                                             | 数字补x (填充右边, 宽度为4)  |
| 10         | {:x<4d}                                                      | 10xx                                             | 数字补x (填充右边, 宽度为4)  |
| 1000000    | {:,}                                                         | 1,000,000                                        | 以逗号分隔的数字格式         |
| 0.25       | {:.2%}                                                       | 25.00%                                           | 百分比格式                   |
| 1000000000 | {:.2e}                                                       | 1.00e+09                                         | 指数记法                     |
| 13         | {:10d}                                                       | 13                                               | 右对齐 (默认, 宽度为10)      |
| 13         | {:<10d}                                                      | 13                                               | 左对齐 (宽度为10)            |
| 13         | {:^10d}                                                      | 13                                               | 中间对齐 (宽度为10)          |
| 11         | '{:b}'.format(11)<br />'{:d}'.format(11)<br />'{:o}'.format(11)<br />'{:x}'.format(11)<br />'{:#x}'.format(11)<br />'{:#X}'.format(11) | 1011<br />11<br />13 <br />b <br />0xb <br />0XB | 进制                         |

* ^ <>分别为居中、左右对齐后面为宽度

* ：后面带填充字符，只能是一个，不指定的空格填充

* +表示为数字添加正负号

* 可以用{}转义{}

  ```python
  #!/usr/bin/python
  # -*- coding: UTF-8 -*-
   
  print ("{} 对应的位置是 {{0}}".format("runoob"))

  runoob 对应的位置是 {0}
  ```

  ​

