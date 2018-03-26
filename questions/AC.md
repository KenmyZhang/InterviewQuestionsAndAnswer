



Step1: 将由patterns组成的集合（要同时匹配多个patterns嘛）构成一个有限状态自动机。
Step2: 将要匹配的text作为自动机输入，输出含有哪些patterns及patterns在全文中位置。

自动机的执行动作由三个部分组成：
    一个goto function
    一个failure function
    一个output function
