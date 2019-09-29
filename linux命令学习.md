### `cd`
  Change directory　修改目录

### `pwd`
  Print Working Directory 打印当前目录

### `touch`
  这个是全名，创建文件 后面根文件名 如```touch test.txt``` 在当前目录创建test.txt文件

### `mkdir`
  Make Directory 创建目录，就是创建一个文件夹  
  ```mkdir test``` 在当前目录创建一个test文件夹  
  ```-p``` 当想要创建多级文件夹时，直接创建会失败，因为上级目录没有创建，```mkdir -p test/test/test```

### `cp`
  Copy 复制文件操作 将a复制到b a和b都是带路径的文件  
  ```cp a.txt ..``` 将本目录的a.txt复制到上级目录.  
  ```cp a.txt ../b.txt```将本目录的a.txt复制到上级目录 并且重命名b.txt

### `mv`
  Move 移动文件操作，理解起来应该是修改带路径的文件名。  
  ```mv path1/a.txt path2/b.txt```将path1下的a.txt移动到path2并且改名b.txt，其实就是修改了这个路径索引

### `rm`
  Remove 删除 目录或者文件  
  ```rm a.txt``` 删除当前目录下的a.txt
  ```-r```递归删除整个目录，包括其中的文件
  ```-f``` force 强制删除操作

### `rmdir`
  Remove Directory 删除空目录，不怎么用

### `cat`
  concatenate 显示文件内容  
  ```cat a.txt``` 将本目录的a.txt内容打印出来
### `more`
  More 类比cat 分页显示文件内容

### `find`
  find 查找目标文件夹或者文件
  ```find path -maxdepth 1 -name tes``` 查找path文件夹下的名字带有tes的文件，深度为1（不查找子文件夹）

### `grep`
  在指定文件中查找对应字符串的行  
  ```grep aaa test.txt```在test.txt文件中查找包含aaa的行
  ```-i```忽略大小写  
  ```-v```反向查找，找没有这个字符的行  
  更多选项之后补充

###