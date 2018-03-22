### Description
    求两个节点的最近公共祖先

### Example
    Given 1, return 0
    Given 2, return 1
    Given 10, return 34

[answer](https://github.com/KenmyZhang/InterviewQuestionsAndAnswer/blob/master/answers/GetAncestorForSearchBTree.go)



如果两个节点都比根节点小，则递归左子树 ；
如果两个节点都比跟节点大，则递归右子树 ；
否则，两个节点一个在左子树，一个在右子树，则当前节点就是最近公共祖先节点。