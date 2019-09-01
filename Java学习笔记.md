# JAVA 面向对象编程专精

### 字符串处理 string，StringBuffer
#### StringBuffer
  + StringBuffer.charAt（int index）  
    返回这个StringBuffer 索引index位置的字符。 理解为StringBuffer[index]
  + StringBuffer.toString()
    将StringBuffer转化为String
  + StringBuffer.append(String s)
    将字符串s添加到后面,类比String中 s1+s2

#### String
  + String.replace(String s1,String s2)
    返回一个字符串，当前String中将s1替换成s2

#### 整数进制转化
  + Integer包，用函数将10进制转化成其他进制
  + Integer.toBinaryString(int n)
  + Integer.toString(int n,int r)  
    这个为通用函数，将十进制转化成r进制，返回的是字符串String
  + Integer.parseInt(String s,int r)
    这个用于将r进制的字符串s 转化为10进制