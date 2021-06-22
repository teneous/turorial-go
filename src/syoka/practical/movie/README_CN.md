## 如何区分函数和方法

## 快速开始

首先定义了一个结构Movie，内部含有

函数

```go
func summary(movie *Movie)
```

方法

```go
func (movie *Movie) summary() 
```

调用方式的不同

```go
func main() {
    //创建对象
    movie := Movie{"century love", []string{"zhangsan", "lisi"}, time.Now()}
    //对象调用方法
    movie.summary()
    //函数调用
    summary(&movie)
}
```
