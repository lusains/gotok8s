# 原型模式 (Prototype Pattern)

## 特点
- 用原型实例指定创建对象的种类，并且通过拷贝这些原型创建新的对象
- 不需要知道任何创建的细节，不调用构造函数
- 一般用于创建复杂的或者耗时的实例，因为这种情况下，复制一个已经存在的实例更加高效

## 实现特点
1. 定义克隆接口
2. 实现克隆方法
3. 提供浅拷贝和深拷贝两种实现方式
4. 管理原型注册表（可选）

## 常见使用场景
1. 对象的创建成本比较大（如需要经过复杂计算、数据库查询等）
2. 系统需要优化创建对象的性能
3. 创建一个对象需要繁琐的数据准备或访问权限
4. 系统要保存对象的状态，而要避免多次创建相同的对象

## 优点
1. 性能提高，减少创建对象的开销
2. 简化对象创建过程
3. 逃避构造函数的约束
4. 保护性拷贝

## 缺点
1. 必须实现克隆方法，可能很困难
2. 深拷贝和浅拷贝的问题需要考虑
3. 克隆复杂对象或有循环引用的对象可能很困难 