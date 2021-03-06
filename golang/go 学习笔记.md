# go学习笔记



## fmt包
* 格式化字符串输入
  fmt.Springf("%d%s%c%f`%T`",整形，字符串，字符，浮点,`变量类型`)
## time包
* basetime
  time包里的示例时间点，用作格式化`2006-01-02 15:04:05 -0700 MST`
  具体参见time包里的const部分
* time.Time结构
  包含从`初始时间`（1979年1月1日00:00:00UTC）到表示时间的秒数和纳秒数（绝对时间），和时区信息（相对时间)，该结构大小：24B
  `func (t Time) In(loc *Location) Time`
  //返回一个规定时区的time,原来的绝对时间不变,类似的操作有UTC()、local（）
  `func (t Time) IsZero() bool`
  //判断这个时间是否为0时间，一般用于判断是否初始化
  `func Date(year int, month Month, day, hour, min, sec, nsec int, loc *Location) Time`
  //根据要求构造时间，参数不可少
  `func (t Time) Format(layout string) string`
  //传入需要的格式，把时间格式转化成字符串，格式时间为示例时间
  `func Parse(layout, value string) (Time, error)`
  //说明字符串的格式，将字符串转化成time格式
  `func (t Time) Unix() int64`
  //获取时间戳，距离初始时间的秒数
* time.Time计算
  `func (t Time) After(u Time) bool`
  `func (t Time) Before(u Time) bool`
  `func (t Time) Before(u Time) bool`
  //时间结构的比较函数
* time.now()
获取当前的时间，格式为time.time
## unsafe包
* 获取变量的大小
  `unsafe.Sizeof()` //返回变量的大小 单位：Byte
## sync.map 线程安全的数据结构
* 定义：
  ```
    var test sync.map
  ```
  线程安全的kv对类型的数据结构，key和value可以是任意结构,不需要相同
* 方法（增删改查）：
  ```
  testvalue,ok := test.Load(testkey)
  ```
  读出对应key的value，读失败时ok为false,`注意`,此时返回的testvalue为结构格式，需要进行强制格式转换，变为原来的格式
  ```
  test.Store(testkey,testvalue)
  ```
  将testkey/testvalue作为kv对放入结构中
  ```
  test.Delete(testkey)
  ```
  传入key值删除结构体中对应的kv对
  ```
  test.Range( func(k, v interface{}) bool {
					return true
				} )
  ```
  迭代函数，结构体中的每一个变量进行循环操作，做法是传入一个回调函数，这个函数的参数为结构体中的kv对，每一次循环，就运行一次传入的函数。函数的返回值用于控制循环是否继续，继续循环返回true，否则返回false