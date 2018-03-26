### Description
    BF算法(Brute Force)是普通的模式匹配算法，BF算法的思想就是将目标串S的第一个字符与模式串T的第一个字符进行匹配，若相等，则继续比较S的第二个字符和 T的第二个字符；若不相等，则比较S的第二个字符和T的第一个字符，依次比较下去，直到得出最后的匹配结果。BF算法是一种蛮力算法。时间复杂度为O(M*N)。

    暴力匹配的思路，并假设现在文本串S匹配到 i 位置，模式串P匹配到 j 位置，则有：
		如果当前字符匹配成功（即S[i] == P[j]），则i++，j++，继续匹配下一个字符；
		如果失配（即S[i]! = P[j]），令i = i - (j - 1)，j = 0。相当于每次匹配失败时，i 回溯，j 被置为0。


### Example
    Given s="anmincd" t="mincd", return 2
    Given s="anmincd" t="abcde", return -1


[代码链接](https://github.com/KenmyZhang/InterviewQuestionsAndAnswer/blob/master/answers/BruteForce.go)