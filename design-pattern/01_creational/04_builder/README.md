# 建造者模式 (Builder Pattern)

## 特点
- 将一个复杂对象的构建与它的表示分离，使得同样的构建过程可以创建不同的表示
- 适用于构建过程稳定，但构建的对象具有不同表示的场景
- 常用于创建复杂的组合对象

## 实现特点
1. 定义建造者接口，声明构建步骤
2. 实现具体建造者，完成具体构建步骤
3. 定义指挥者，控制构建过程
4. 定义产品，表示被构建的复杂对象

## 常见使用场景
1. 构建复杂的对象，如计算机（CPU、内存、硬盘等）
2. 生成不同格式的文档（HTML、PDF、Word等）
3. 创建复杂的数据库查询构建器
4. 构建复杂的GUI表单

## 优点
1. 可以精细地控制构建过程
2. 将复杂对象的创建步骤分解
3. 隔离复杂对象的构建和表示
4. 同样的构建过程可以创建不同的表示

## 缺点
1. 产品必须有足够的复杂性才值得使用建造者模式
2. 产品的组成部分必须相对稳定
3. 需要额外定义建造者类，增加代码量 