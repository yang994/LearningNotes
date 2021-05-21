# Rust

## Cargo
  Rust的构建系统与包管理器，可以理解为Rust语言的专属构建工具，类比golang的go命令  
  ``` sh
  cargo check  //检查
  cargo build  //编译
  cargo run     //运行
  ```
  编译以及运行操作
  + 重影
     对变量名称的重新利用
     ```rust
     let x=5;
     let x=x+1;
     ```
     Rust中不允许直接对“不可变变量”进行修改，只能使用重影机制