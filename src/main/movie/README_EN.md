## How to distinguish function and method in go

## Quick Start

Define a struct named Movie Firstly,which involve

function definition

```go
func summary(movie *Movie)
```

method definition

```go
func (movie *Movie) summary() 
```

The way of calling is also different

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
