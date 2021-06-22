package movie

import (
	"fmt"
	"time"
)

func main() {
	//创建对象
	movie := Movie{"century love", []string{"zhangsan", "lisi"}, time.Now()}
	//对象调用方法
	movie.summary()
	//函数调用
	Summary(&movie)
}

//Movie 电影结构
type Movie struct {
	name        string
	actors      []string
	releaseTime time.Time
}

//Summary 电影介绍函数
func Summary(movie *Movie) {
	filmName := movie.name
	filmActor := ""
	for _, actor := range movie.actors {
		filmActor += actor + ","
	}
	filmActor = filmName[:len(filmName)-1]

	releaseTime := movie.releaseTime
	rt := releaseTime.Format("2006-01-02 15:04:05")

	fmt.Printf("function summary[film name:%s ,film actor:%s, releaseTime:%s]:", filmName, filmActor, rt)
	println()
}

//Movie 电影对象介绍
func (movie *Movie) summary() {
	filmName := movie.name
	filmActor := ""
	for _, actor := range movie.actors {
		filmActor += actor + ","
	}
	filmActor = filmName[:len(filmName)-1]

	releaseTime := movie.releaseTime
	rt := releaseTime.Format("2006-01-02 15:04:05")

	fmt.Printf("main summary[film name:%s ,film actor:%s, releaseTime:%s]:", filmName, filmActor, rt)
	println()
}
