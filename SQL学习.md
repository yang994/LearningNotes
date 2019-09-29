## SQL常识
  + `RDBMS` 数据库管理系统（MS Access, SQL Server, MySQL）  
  + `ANSI` 没过国家标准化组织，SQL是ANSI标准化语言，不同版本的数据库系统都兼容ANSI
  + 大小写不敏感
  + `DML` 数据库操作语言（对于数据增删改查）
  + `DDL` 数据库定义语言（对于库、表增删改查）
  + `‘’`单引号 为字符串，否则为数值或变量

## 查询操作 SELECT
  + 格式： `SELECT a FROM b` 从b表中查询a列，当a为`*` ，表示所有列
  + 操作返回结果集（result-set），可使用其他语言调用
  + `DISTINCT`关键字  `SELECT DISTINCT a FROM b` 和查询语句一样，此时查出的a列会丢掉重复项（做set操作）
  + `WHERE`子句 添加选取数据的条件
    + 都知道的逻辑符 `=`（不是`==`）、`> >=`、`< <=`、`<>`(有些地方是`！=`)
    + 不知道的逻辑符 `BETWEEN` 在某范围内
    + 不知道的逻辑符 `LIKE`搜索某种模式
    + 使用：`SELECT a FROM b WHERE c=‘1’` 加上了b表中c列元素是1的限制
    + `AND``OR`操作，就是判断语句中的`&&`和`||`
  + `ORDER BY` 放在最后，按照某种规则进行排序，默认升序
    + `DESC` 降序排列标志
    + `ASC` 升序排列标志
    + `SELECT a FROM b WHERE c=‘1’ ORDER BY c ASC，d DESC` 按照c列元素升序排列，当c一样时，按照d列元素降序排列

## 插入操作INSERT INTO
  + 格式：
    + 完全形式：`INSERT INTO a VALUES (1,2,3,4)` 向a表中插入一行，数据是1,2,3,4（这个表有4列）
    + 缺省形式 插入部分数据：`INSERT INTO a（b，c） VALUES (1,2)`向a表中插入一行，b列为1，c列为2，其他数据缺省

## 修改操作 Update
  + 格式： `UPDATE a SET b=1 WHERE....` 修改a表中的b列数据，修改数据条目满足WHERE语句

## 删除操作 DELETE
  + 格式：`DELETE FROM a WHERE....`删除表中的行，行满足WHERE语句