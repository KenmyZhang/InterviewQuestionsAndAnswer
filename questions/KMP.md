### Description
    KMP算法的关键是利用匹配失败后的信息，尽量减少模式串与主串的匹配次数以达到快速匹配的目的。具体实现就是实现一个next()函数，函数本身包含了模式串的局部匹配信息。时间复杂度O(m+n)。

### Example
    Given s="anmincd" t="mincd", return 2
    Given s="anmincd" t="abcde", return -1

[answer](https://github.com/KenmyZhang/InterviewQuestionsAndAnswer/blob/master/answers/KMP.go)