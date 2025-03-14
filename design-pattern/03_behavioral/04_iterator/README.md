# 迭代器模式 (Iterator Pattern)

## 特点
- 提供一种方法顺序访问一个聚合对象中的各个元素，而又不暴露其内部的表示
- 将遍历集合的责任从集合中分离出来，放到迭代器中
- 支持以不同的方式遍历一个集合

## 实现特点
1. 迭代器（Iterator）：定义访问和遍历元素的接口
2. 具体迭代器（ConcreteIterator）：实现迭代器接口
3. 集合（Aggregate）：定义创建迭代器的接口
4. 具体集合（ConcreteAggregate）：实现创建迭代器的接口

## 常见使用场景
1. 访问一个集合对象的内容而无需暴露它的内部表示
2. 支持对集合对象的多种遍历方式
3. 为遍历不同的集合结构提供一个统一的接口
4. 在不同的集合类型中使用相同的遍历接口

## 优点
1. 支持以不同的方式遍历一个集合
2. 简化了集合的接口
3. 在同一个集合上可以有多个遍历
4. 迭代器模式将集合的遍历行为分离出来

## 缺点
1. 对于简单的遍历可能会过度设计
2. 增加了系统的复杂性
3. 迭代器模式在某些情况下可能比直接遍历集合要慢 