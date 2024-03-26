package main

import "fmt"

type Animal interface {
	Sound() string
}

func Dongwu() {
	var animal Animal

	// createa dog object and assign it to the interface variable
	animal = Dog{}
	dog, ok := animal.(Dog)
	if ok {
		fmt.Println(dog.Sound()) // output: Woof!
	}

	// create a cat object and assign it to the interface variable
	animal = Cat{}
	cat, ok := animal.(Cat)
	if ok {
		fmt.Println(cat.Sound()) // output: Meow!
	}
}

// WashingMachine 洗衣机
type WashingMachine interface {
	wash()
	dry()
}

// 甩干器
type dryer struct {

}

// 实现WashingMaching 接口的 dry() 方法
func (d dryer) dry() {
	fmt.Println("甩一甩")
}

// 海尔洗衣机
type haier struct {
	dryer // 嵌入甩干器
}

// 实现WashingMachine 接口的 wash() 方法
func (h haier) wash() {
	fmt.Println("洗刷刷")
}

type person interface {
	Say()
}

func Say(p person) {
	p.Say()
}

type man struct {

}

func (m *man) Say() {
	fmt.Println("man")
}

type woman struct {

}

func (w *woman) Say() {
	fmt.Println("woman")
}

func main() {
	m := new(man)
	w := new(woman)

	person.Say(m)
	person.Say(w)
	Say(m)
}