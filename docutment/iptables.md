Iptables防火墙的四表五链概念以及使用技巧
========================

【摘要】 Iptables防火墙的四表五链概念以及使用技巧 文章目录 Iptables防火墙的四表五链概念以及使用技巧1.链的概念2.Iptables五种链的概念3.Iptables数据流向经过的表4.I...

[](https://bbs.huaweicloud.com/blogs/364191)[](https://bbs.huaweicloud.com/blogs/364191)Iptables防火墙的四表五链概念以及使用技巧
================================================================================================================

### [](https://bbs.huaweicloud.com/blogs/364191)文章目录

*   [Iptables防火墙的四表五链概念以及使用技巧](https://bbs.huaweicloud.com/blogs/364191)
*   *   [1.链的概念](https://bbs.huaweicloud.com/blogs/364191)
*   [2.Iptables五种链的概念](https://bbs.huaweicloud.com/blogs/364191)
*   [3.Iptables数据流向经过的表](https://bbs.huaweicloud.com/blogs/364191)
*   [4.Iptables防火墙四种表的概念](https://bbs.huaweicloud.com/blogs/364191)
*   [5.Iptables防火墙表与链之间的优先级概念](https://bbs.huaweicloud.com/blogs/364191)
*   [6.Iptables防火墙表和链之间的使用技巧](https://bbs.huaweicloud.com/blogs/364191)
*   [7.Iptables防火墙几种动作](https://bbs.huaweicloud.com/blogs/364191)

[](https://bbs.huaweicloud.com/blogs/364191)[](https://bbs.huaweicloud.com/blogs/364191)1.链的概念
----------------------------------------------------------------------------------------------

在防火墙中，用户想要成功进入内网环境，就需要发送请求报文，请求报文要和防火墙设置的各种规则进行匹配和判断，最后执行相应的动作（放行或者拒绝），一个防火墙中通常针对不同的来源设置很多种策略，多个策略形成一个链，其实也可以理解成是分组的概念，在Iptables防火墙中针对不同的链路共分为五种不同的链。

如下图所示，当数据报文进入链之后，首先匹配第一条规则，如果第一条规则通过则访问，如果不匹配，则接着向下匹配，如果链中的所有规则都不匹配，那么就按照链的默认规则处理数据报文的动作。  
![在这里插入图片描述](https://img-blog.csdnimg.cn/288eaac5c0fd49da9f24d7da6b92e2d2.png)

[](https://bbs.huaweicloud.com/blogs/364191)[](https://bbs.huaweicloud.com/blogs/364191)2.Iptables五种链的概念
--------------------------------------------------------------------------------------------------------

Iptables有五种不同的链，分别是INPUT、OUTPUT、FORWARD、PREROUTING、POSTROUTING。

*   INPUT：从外界进入防火墙的数据包会应用此规则链中的策略。
*   OUTPUT：当前服务器从防火墙外出的数据表会应用此规则链中的策略。
*   FORWARD：转发数据包时会应用此规则链中的策略。
*   PREROUTING：主机外的报文要进入防火墙，所有的数据包进来的时候都会由PREROUTING链进行处理。
*   POSTROUTING：主机内的报文要从防火墙出去，需要经过POSTROUTING链进行处理。

[](https://bbs.huaweicloud.com/blogs/364191)[](https://bbs.huaweicloud.com/blogs/364191)3.Iptables数据流向经过的表
----------------------------------------------------------------------------------------------------------

**请求报文流入本地要经过的链：**

请求报文要进入本机的某个应用程序，首先会到达Iptables防火墙的PREROUTING链，然后又PREROUTING链转发到INPUT链，最后转发到所在的应用程序上。

> PREROUTING—>INPUT—>PROCESS

**请求报文从本机流出要经过的链：**

请求报文读取完应用程序要从本机流出，首先要经过Iptables的OUTPUT链，然后转发到POSTROUTING链，最后从本机成功流出。

> PROCESS—>OUTPUT—>POSTROUTING

**请求报文经过本机向其他主机转发时要经过的链：**

请求报文要经过本机向其他的主机进行换发时，首先进入A主机的PREROUTING链，此时不会被转发到INPUT链，因为不是发给本机的请求报文，此时会通过FORWARD链进行转发，然后从A主机的POSTROUTING链流出，最后到达B主机的PREROUTING链。

> PREROUTING—>FORWARD—>POSTROUTING

[](https://bbs.huaweicloud.com/blogs/364191)[](https://bbs.huaweicloud.com/blogs/364191)4.Iptables防火墙四种表的概念
-----------------------------------------------------------------------------------------------------------

\*\*Iptables防火墙中表的概念：\*\*在一个链中会有很多的防火墙规则，我们将具有同一种类型的规则组成一个集合，这个集合就叫做表，表可以简单的列成是一些具有同样类型的规则的分组，例如关于IP地址转换的策略都放在一个表中、修改数据保报文的策略都放在一个表中。

在Iptables防火墙中包含四种常见的表，分别是filter、nat、mangle、raw。

*   filter：负责过滤数据包。
    *   filter表可以管理INPUT、OUTPUT、FORWARD链。
*   nat：用于网络地址转换。
    *   nat表可以管理PREROUTING、INPUT、OUTPUT、POSTROUTING链。
*   mangle：修改数据包中的内容，例如服务类型、TTL、QOS等等。
    *   mangle表可以管理PREROUTING、INPUT、OUTPUT、POSTROUTING、FORWARD链。
*   raw：决定数据包是否被状态跟踪机制处理。
    *   raw表可以管理PREROUTING、OUTPUT链。

[](https://bbs.huaweicloud.com/blogs/364191)[](https://bbs.huaweicloud.com/blogs/364191)5.Iptables防火墙表与链之间的优先级概念
----------------------------------------------------------------------------------------------------------------

在Iptables防火墙中，表与链之间是存在优先级的关系的，因为每张表的作用都是不同的，一张表会同时存放在多个链中，当一条数据报文进入一个链后，会按照表的优先级顺序依次匹配对应的规则。

Iptables防火墙表的优先级顺序：raw—>mangle—>nat—>filter。

如下图所示，当数据报文进入PREROUTING链时，首先规匹配raw表中的规则，然后在匹配mangle表中的规则，最后在匹配nat表的规则，按照优先级顺序依次匹配。

\[外链图片转存失败,源站可能有防盗链机制,建议将图片保存下来直接上传(img-4e31sEnS-1657242939282)(G:\\01-运维技术文档整理\\Iptables\\01-Iptables防火墙核心理论概念.assets\\image-20220428235007852.png)\]

[](https://bbs.huaweicloud.com/blogs/364191)[](https://bbs.huaweicloud.com/blogs/364191)6.Iptables防火墙表和链之间的使用技巧
---------------------------------------------------------------------------------------------------------------

从以下三个问题中掌握防火墙表和链之间的使用技巧。

**首先牢记各个链对应的表有那些：**

| 链名 | 对应的表名 |
| --- | --- |
| INPUT | mangle、nat、filter |
| OUTPUT | raw、mangle、Nat、filter |
| FORWARD | mangle、filter |
| PREROUTING | raw、mangle、nat |
| POSTROUTING | mangle、nat |

记忆技巧：进出第一关的链都没有fileter表，第一个进链除fileter都包含，input除raw都有、output全有、出链只有mangle和nat、forward只有mongle和filter。

\*\*问题1：\*\*来自于10.0.0.1的地址，访问本机的web服务请求都不允许，应该在哪个表的哪个链上设定规则？

> 源地址为10.0.0.1的IP访问本机的WEB请求时不允许，属于数据流入的操作，首先要分析会经过那些Iptables链？
>
> 源地址到本地服务会经过Iptables的PREROUTING和INPUT链，做这种规则时，都会坚定最近位置处做策略，但是也要结合实际的功能，PREROUTING链离源地址最近，但是没有过滤的表，在PREROUTING中的表有mangle、nat，没有负责做过滤的表，因此就要判断第二个链，也就是INPUT链，在INPUT链中包含mangle、nat、filter表，在INPUT链中添加策略是最合适的。
>
> 最终的结果就是在INPUT链的filter表中添加具体的策略。

\*\*问题2：\*\*所有由本机发往10.0.0.0/24网段的TCP服务都不允许？

> 从本机发往其他主机的TCP服务请求，属于数据流出的操作，会经过PREROUTING链和OUTPUT链，到达目标地址的数据保温都拒绝，这种类似的需求，一般都是由过滤表filter来实现，PREROUTING链包含的表有mangle、nat这两张，没有过滤表，OUTPUT链有raw、mangle、Nat、filter四张表，最终的结果就是在OUTPUT链的filter表添加具体的策略。

\*\*问题3：\*\*所有来自己本地内部网络的主机，向互联网发送web服务器请求都允许？

> 到达本机的请求报文向互联网发送请求，属于数据转发的操作，会经过PREROUTING、FORWARD和POSTROUTING三个连，这种允许和拒绝都是在filter表中操作的，因此找到含有filter表并且距离目标端最近的链中添加合适的规则，最终的结果就是在filter表中添加具体的操作。

**结论：**

1）首先要知道要实现的需求含义，然后根据需求判断出要在哪一个表中实现该策略。

2）然后摸清楚报文要经过的链，坚定在距离源/目的最近的链做策略。

3）最后根据链包含的表，判断出要将规则添加到哪一个链的表中。

[](https://bbs.huaweicloud.com/blogs/364191)[](https://bbs.huaweicloud.com/blogs/364191)7.Iptables防火墙几种动作
---------------------------------------------------------------------------------------------------------

ACCEPT：将数据包放行。

REJECT：拒绝该数据包通行，阻拦数据包。

DROP：丢弃数据包，不给予任何处理。

REDIRECT：重定向。

文章来源: jiangxl.blog.csdn.net，作者：Jiangxl~，版权归原作者所有，如需转载，请联系作者。

原文链接：jiangxl.blog.csdn.net/article/details/125671469