//test memory pointer
//* 解析饮用 & 获取变量地址
package basic

import (
	"testing"
)

//测试函数传递指针的情况
func TestPointer(t *testing.T) {
	cmd := "gh clone repo"
	println("origin address:", &cmd)
	//解析指针
	println("origin   value:", *(&cmd))
	memoryAddr(cmd)
	println("origin  return:", cmd)
}

func memoryAddr(cmd string) {
	//可以看到入参的指针地址和原地址不一样，因此推测函数入参值传递，传递了指针地址的副本
	println("func   address:", &cmd)
	println("func     value:", *(&cmd))
	//为了进一步验证结果,我们在函数内部修改了值，在方法结束后，变量的作用就结束了，由于是副本，不会影响外部调用方的值
	cmd = "gh status"
}

//测试函数传递指针的情况
func TestPointer2(t *testing.T) {
	cmd := "gh clone repo"
	println("origin address:", &cmd)
	println("origin   value:", *(&cmd))
	memoryAddr2(&cmd)
	println("origin  return:", cmd)
}

//传递了指针
func memoryAddr2(cmd *string) {
	//打印指针，可以看到如果传递的是指针，那么方法内外是一致的
	println("func   address:", cmd)
	//解析指针引用的值
	println("func * address:", *cmd)
	//为了进一步验证结果,我们在函数内部修改了值.在方法结束后，变量的作用就结束了，由于是指针，那么会影响外部调用值
	*cmd = "git status"
}
