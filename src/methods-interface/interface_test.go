package methods_interface




import (
	"fmt"
	"testing"
)

type USB interface {
	start()
}
type Computer interface {
	start()
	end()
}
// ----------------------------Usb 类-----------------------------------------------------------------------
type U1 struct {
	name string
	price int
}
func (u U1) start()  {
	fmt.Println("我是Usb的 start : ",u.name)
}
func (u U1) end()  {
	fmt.Println("我是usb的 end : ", u.name)
}

// ---------------------------Computer 类------------------------------------------------------------------------
type C1 struct {
	name string
}
func (c C1) start()  {
	fmt.Println("我是电脑的 start : ",c.name)
}
func (c C1) end()  {
	fmt.Println("我是电脑的 end : ", c.name)
}
// -----------------------------------------------------------------------------------------------------

// go 语言中，接口不像 java 语言，
// 鸭子模型   那什么鸭子模型？
// 鸭子模型的解释，通常会用了一个非常有趣的例子，一个东西究竟是不是鸭子，取决于它的能力。游泳起来像鸭子、叫起来也像鸭子，那么就可以是鸭子。
// 通俗理解， 能像鸭子一样跑、或者像鸭子一样叫 就都认定为是鸭子

// 以上的栗子 是  usb  和 computer 都有 start 功能 ， 而 computer 还有 end 功能
// 当 U1 和 C1 两个类都实现了 start 接口时候， 那么就可以认定他们即是  usb 也可以是 computer ，


func TestInter( t *testing.T) {
	c := C1{
		name: "电脑小明",
	}
	u := U1{
		name: "usb 小红",
	}
	
	var u1 USB
	u1 = u
	u1.start()
	
	fmt.Println("-------------------------------------------")
	
	var c1 Computer
	c1 = c
	c1.start()
	c1.end()
	
	
}
