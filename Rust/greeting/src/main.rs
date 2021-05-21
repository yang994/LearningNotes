fn main() {
    test4();
}

const A:u32 = 1000; //常量

//变量定义(可变变量与不可变变量)
fn test1(){
    let a = 1;  //不可变
    let mut b = 2;  //可变
    println!("a={},b={}",a,b);
    b = 3;
    println!("a={},b={}",a,b);
    let b= 3.3; //重影
    println!("a={},b={}",a,b);
    println!("A={}",A);
}

fn test2(){
    let is_true:bool = true;    //bool类型
    println!("is_true = {}",is_true);
    let c:char = 'a';//字符类型 32位，包含汉字
    println!("c = {}",c);
    let max:usize = usize::max_value();
    println!("无符号变量最大数max={}",max);//与系统位数有关

    let arr : [u32; 5] =[1,2,3,4,5];        //在Rust中 size也是数组的一部分
    println!("arr[0] = {}",arr[0]);

    let tup = (1,1.2,'a',"adf");
    let (x,y,z,w) = tup;
    println!("tup.0 = {},x = {}",tup.0,x);  //元祖以及拆解
}

fn test3(a:i32,b:u32 ) -> u32{ //函数定义中 参数类型无法推导
    println!("a={},b={}",a,b);
    return 1+2;
}

//if else的使用
fn test4(){
    
    let mut y = if true{5} else{6}; 
    if y==1{
        println!("y==1");
    } else if y==0{
        println!("y!=1,y==0");
    } else{
        println!("over");
    }
    //loop  以及break的赋值操作
    let result = loop{
        println!("in loop");
        if y==10 {
            break y*2;
        }
        y=y+1;
    };
    println!("result = {}",result);
    //while 操作
    while y==10{
        y+=1;
        println!("in while,y={}",y);
    }

    //for 与数组迭代器
    let arr = [1,2,3,4,5];
    for x in arr.iter(){
        println!("in for: x={}",x);
    }

} 

 //堆上数据类型
fn test5(){ 
    {
        let  mut s1 = String::from("hello");    //定义在堆上的可变字符串 数据类型大小不固定。固定数据类型大小分配在栈上
        s1.push_str(" world");
        println!("s1 = {}",s1); //离开作用域  调用drop方法 释放内存
        
        let s2 = s1;    //字符串移动操作，传指。 s1变为无效 不可再使用
        println!("s2 = {}",s2);
        //println!("s1 = {}",s1); //离开作用域时，s1 s2释放同一块内存，所以设置s1无效
        
        let s3 = s2.clone();    //s3深拷贝s2 传值
        println!("s3 = {}",s3);
    }

    //栈上数据类型
    {
        let a = 1;
        let b = a;  //栈上数据类型传值
        println!("a={}",a);
        println!("b={}",b);
    }

}
//引用  & 创建一个指向变量值的引用，不会影响变量的所有权
    //在有可变引用（借用）之后，不可再有不可变引用
fn test6(s:&String)->usize{
    
    s.len()
}

fn test7(){
    //切片操作，给出一个引用
    let s = String::from("hello world");

    let h = &s[0..5];
    println!("h={}",h);

    let s3 = "hh"; //指向字符串“hh”的引用 不可变
}

//结构体
struct User{
    name:String,
    count:String,
    nonce:u64,
}
//元祖结构体，字段没有名字
struct Point(i32,i32);
//无字段结构体，用来携带结构体方法
struct A{}
//定义结构体中的方法
impl User{
    //获取结构体中的元素值，使用引用的方式
    fn get_name(&self)->&str{
        &(self.name[..])
    }
    fn get_count(&self)->&str{
        &(self.count[..])
    }
    fn show(){  //这个函数没有传入self
        println!("hello")
    }
}

fn test8(){
    //创建结构体
    let mut u = User{
        name:String::from("name"),
        count:String::from("count"),
        nonce:5,
    };
    let a = Point(1,2);
    println!("{}",a.0) ;  //访问元祖结构体元素

    //修改结构体
    u.nonce = 20000;

    //从其他结构体创建新结构体，但原结构体变为不可用
    let u1 = User{
        ..u
    };

    User::show();//结构体中的方法

}


//枚举 定义以及使用
fn test9(){
    //第一种方式
    enum kind1{
        v1,
        v2,
    };
    struct IpAddr{
        ip_kind:kind1,
        address:String,
    };
    let a = IpAddr{
        ip_kind:kind1::v1,
        address:String::from("123123123"),
    };
}
    //第二种方式
    enum Kind2{
        V1(String),
        V2(u8,u8,u8,u8),
    }
    //枚举类型匹配
    impl Kind2{
        fn print(&self){
            match self {
                Kind2::V1(x) => println!("x={}",x),
                Kind2::V2(a,b,c,d) => println!("{}{}{}{}",a,b,c,d),
                _=>println!("void"),
            }
        }
    }