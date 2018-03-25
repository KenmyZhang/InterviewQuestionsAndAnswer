### Description
    子串定位运算又称为模式匹配（Pattern Matching）或串匹配（String Matching）;在串匹配中，一般将主串称为目标串，将子串称为模式串。

    从目标串S的第一个字符开始和模式串T的第一个字符进行比较，如果相等则进一步比较二者的后继字符，否则从目标串的第二个字符开始再重新与模式串T的第一个字符进行比较，以此类推，直到模式串T与目标串S中的一个子串相等，称为匹配成功，返回T在S中的位置；或者S中不存在值与T相等的子串，称匹配失败，返回-1.此算法也称为BF（Brute-Force）算法。

    暴力匹配的思路，并假设现在文本串S匹配到 i 位置，模式串P匹配到 j 位置，则有：
		如果当前字符匹配成功（即S[i] == P[j]），则i++，j++，继续匹配下一个字符；
		如果失配（即S[i]! = P[j]），令i = i - (j - 1)，j = 0。相当于每次匹配失败时，i 回溯，j 被置为0。


### Example
    Given s="anmincd" t="mincd", return 2
    Given s="anmincd" t="abcde", return -1


[answer](https://github.com/KenmyZhang/InterviewQuestionsAndAnswer/blob/master/answers/BruteForce.go)