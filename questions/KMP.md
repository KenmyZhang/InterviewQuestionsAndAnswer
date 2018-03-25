### Description
	实现KMP算法

### Example
    Given s="anmincd" t="mincd", return 2
    Given s="anmincd" t="abcde", return -1

[answer](https://github.com/KenmyZhang/InterviewQuestionsAndAnswer/blob/master/answers/KMP.go)



### KMP

* KMP算法由Donald Knuth、Vaughan Pratt、James H. Morris三人于1977年联合发表，故取这3人的姓氏命名此算法。

* KMP算法是一种模式匹配算法的改进版，其通过减少匹配的次数以及使主串不回朔来减少字符串匹配的次数，从而较少算法的相应代价

* 失配时，模式串向右移动的位数为：已匹配字符数 - 失配字符的上一位字符所对应的最大长度值，即 6 - 2 等于4，所以将搜索词向后移动4位。

* 问题的关键就是寻找模式串中最大长度的相同前缀和后缀，找到了模式串中每个字符之前的前缀和后缀公共部分的最大长度后，便可基于此匹配。

* next 数组各值的含义：代表当前字符之前的字符串中，有多大长度的相同前缀后缀


### KMP的算法流程
	假设现在文本串S匹配到 i 位置，模式串P匹配到 j 位置
	如果j = -1，或者当前字符匹配成功（即S[i] == P[j]），都令i++，j++，继续匹配下一个字符；
	如果j != -1，且当前字符匹配失败（即S[i] != P[j]），则令 i 不变，j = next[j]。此举意味着失配时，模式串P相对于文本串S向右移动了j - next [j] 位。
	换言之，当匹配失败时，模式串向右移动的位数为：失配字符所在位置 - 失配字符对应的next 值，即移动的实际位数为：j - next[j]，且此值大于等于1。    
	next 数组各值的含义：代表当前字符之前的字符串中，有多大长度的相同前缀后缀。例如如果next [j] = k，代表j 之前的字符串中有最大长度为k 的相同前缀后缀。

	在某个字符失配时，该字符对应的next 值会告诉你下一步匹配中，模式串应该跳到哪个位置（跳到next [j] 的位置）。如果next [j] 等于0或-1，则跳到模式串的开头字符，若next [j] = k 且 k > 0，代表下次匹配跳到j 之前的某个字符，而不是跳到开头，且具体跳过了k 个字符。

### a.寻找前缀后缀最长公共元素长度  最大长度表
![](http://i2.51cto.com/images/blog/201803/24/7c85ac440d7078bd1954ac79d3f79354.png?x-oss-process=image/watermark,size_16,text_QDUxQ1RP5Y2a5a6i,color_FFFFFF,t_100,g_se,x_10,y_10,shadow_90,type_ZmFuZ3poZW5naGVpdGk=)

### b.求next数组
next 数组考虑的是除当前字符外的最长相同前缀后缀，所以通过第①步骤求得各个前缀后缀的公共元素的最大长度后，只要稍作变形即可：将第①步骤中求得的值整体右移一位，然后初值赋为-1
![](http://i2.51cto.com/images/blog/201803/24/9627c21e5f6a18c109c30f6bcc85091d.png?x-oss-process=image/watermark,size_16,text_QDUxQ1RP5Y2a5a6i,color_FFFFFF,t_100,g_se,x_10,y_10,shadow_90,type_ZmFuZ3poZW5naGVpdGk=)

### c.根据next数组进行匹配
匹配失配，j = next [j]，模式串向右移动的位数为：已匹配字符数 - 失配字符的上一位字符所对应的最大长度值（即：j - next[j]）

当模式串中的某个字符跟文本串中的某个字符匹配失配时，模式串下一步应该跳到哪个位置。如模式串中在j 处的字符跟文本串在i 处的字符匹配失配时，下一步用next [j] 处的字符继续跟文本串i 处的字符匹配，相当于模式串向右移动 j - next[j] 位。

失配时，模式串向右移动的位数为：失配字符所在位置 - 失配字符对应的next 值
![](http://i2.51cto.com/images/blog/201803/24/0e958b180efbcebc955d358800591762.png?x-oss-process=image/watermark,size_16,text_QDUxQ1RP5Y2a5a6i,color_FFFFFF,t_100,g_se,x_10,y_10,shadow_90,type_ZmFuZ3poZW5naGVpdGk=)

next[j] = k 代表p[j] 之前的模式串子串中，有长度为k 的相同前缀和后缀
	
### Reference

    https://kb.cnblogs.com/page/176818/
    https://blog.csdn.net/v_july_v/article/details/7041827