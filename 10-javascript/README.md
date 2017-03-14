# *javascript*

### 数据类型

- typeof 操作符

```js 
    typeof 是操作符不是函数
    typeof(xxx) / typeof xxx

    "undefined": 这个值未定义(var声明变量但未对其加以初始化)
    "boolean"  : 布尔值
    "string"   : 字符串
    "number"   : 数值
    "object"   : 对象或null(null是一个空的对象引用)
    "function" : 函数

    alert(undefined == undefined); // true
    //null 派生 undefined 值
    alert(null == undefined); // true

    // NaN 是一个特殊的数值类型，任何数值除以0返回NaN。NaN不与任何值相等，包括NaN 本身
    alert(NaN == NaN); // false
    // 任何可以转换为数值的，isNaN(xxx) 都是false
    alert(isNaN("blue")); // false
    alert(isNaN("true"));

```

Booleand() 函数 转换(if语句自动执行 Boolean()) 

| 数据类型        | 转为true的值  | 转为false的值  |
| ------------- |:------------:| ------------:|
| Boolean       | true         | false        |
| String        | 任何非空字符串  | ""(空字符串) |
| Number | 任何非0数字值(包括无穷大)  |    0和NaN |
| Object | 任何对象 | null |
| Undefined | n/a | undefined |






































END


