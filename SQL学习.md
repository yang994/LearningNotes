# SQL初学笔记

## SQL常识

+ `RDBMS` 数据库管理系统（MS Access, SQL Server, MySQL）  
+ `ANSI` 美国国家标准化组织，SQL是ANSI标准化语言，不同版本的数据库系统都兼容ANSI
+ 大小写不敏感
+ `DML` 数据库操作语言（对于数据增删改查）
+ `DDL` 数据库定义语言（对于库、表增删改查）
+ `‘’`单引号 为字符串，否则为数值或变量

## 基本操作-增删改查

### 查询操作 `SELECT`

+ 格式： `SELECT a FROM b` 从b表中查询a列，当a为`*` ，表示所有列
+ 操作返回结果集（result-set），可使用其他语言调用
+ `DISTINCT`关键字  `SELECT DISTINCT a FROM b` 和查询语句一样，此时查出的a列会丢掉重复项（做set操作）
+ `WHERE`子句 添加选取数据的条件
  + 都知道的逻辑符 `=`（不是`==`）、`> >=`、`< <=`、`<>`(有些地方是`！=`)
  + 不知道的逻辑符 `BETWEEN...AND...` 在某范围内
  + 不知道的逻辑符 `LIKE`搜索某种模式 后面接通配符的语句
  + 使用：`SELECT a FROM b WHERE c=‘1’` 加上了b表中c列元素是1的限制
  + `AND``OR`操作，就是判断语句中的`&&`和`||`
+ `ORDER BY` 放在最后，按照某种规则进行排序，默认升序
  + `DESC` 降序排列标志
  + `ASC` 升序排列标志
  + `SELECT a FROM b WHERE c=‘1’ ORDER BY c ASC，d DESC` 按照c列元素升序排列，当c一样时，按照d列元素降序排列

### 插入操作`INSERT INTO`

+ 格式：
  + 完全形式：`INSERT INTO a VALUES (1,2,3,4)` 向a表中插入一行，数据是1,2,3,4（这个表有4列）
  + 缺省形式 插入部分数据：`INSERT INTO a（b，c） VALUES (1,2)`向a表中插入一行，b列为1，c列为2，其他数据缺省

### 修改操作 `Update`

+ 格式： `UPDATE a SET b=1 WHERE....` 修改a表中的b列数据，修改数据条目满足WHERE语句

### 删除操作 `DELETE`

+ 格式：`DELETE FROM a WHERE....`删除表中的行，行满足WHERE语句

## 进阶语句

### 选取行数 `TOP`

+ `SELECT TOP 2 a FROM b` 从表b中选a列的头两行
+ `SELECT TOP 2 PERCENT a FROM b` 选头2%的数据

### 通配符 用在LIKE语句之后

+ `%`替代任何字符串
+ `_`替代任何一个字符
+ 字符列表`[]`:`[ABC]`ABC中任何一个字符，`[！ABC]`除了ABC之外的字符

### 匹配多个候选项 `IN`

+ `WHERE a IN (v1,v2....)`a列的值在列表中存在
+ 应该等价于 `a=v1 OR a=v2 OR .......`

### 在两数之间 `BETWEEN AND`

+ 数字好说，对于字符串，会对字母顺序排序比较。
+ `WHERE a BETWEEN ‘abc’ AND ‘def’`

### 别名操作 `AS`

+ 在最终返回结果的时候，会按照别名返回，在脚本编写的时候，可以简化语言
+ `SELECT a AS b FROM c`从c中取a行，返回时名字为b

### !连接操作！ `JOIN`

+ 从多个表中获取结果
+ 基本操作：`SELECT A.a,B.b FROM A,B WHERE A.p1=B.p2` A和B通过A中p1列和B中p2列连接起来
+ JOIN使用 `A JOIN B ON 条件`按某种条件进行A表和B表的链接，有四种模式 `INNER JOIN`、`LEFT JOIN`、`RIGHT JOIN`、`FULL JOIN`
  + `INNER JOIN`只返回满足条件的行
  + `LEFT JOIN`在INNER JOIN的基础上，A不满足条件的行也会返回，其中B链接的位置为空
  + `RIGHT JOIN`同LEFT JOIN，A变为B
  + `FULL JOIN` LEFT JOIN + RIGHT JOIN

### 合并操作 UNION

+ 合并两个同类型的表，类比编程语言中的append操作。
+ `A UNION B`B中与A不同的行添加到A后面
+ `A UNION ALL B`B中所有行都添加到后面，不管是否重复

## 下面应该是DDL部分

### 创建备份 SELECT INTO

+ `SELECT a 【】 FROM .....`SELECT语句中插入INTO操作 在【】位置
+ 【】-> `INTO backup` 做一个备份表backup
+ 【】-> `INTO b IN x.mdb` 在其他的数据库中创建表b

### 新建操作 CREATE

+ `CREATE DATABASE a`创建新数据库a
+ 创建表：
  
  ``` SQL
    CREATE TABLE A
    (
      a int,           //整形
      b varchar(255), //变长字符串
      c char(255),    //定长字符串
      d date,         //日期
      e decimal,      //小数
    )
  ```
