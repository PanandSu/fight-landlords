package main

import "fmt"

func main() {
	//foo()
	//bar()
	fmt.Println(baz(2, 1))
	fmt.Println(baz(2.2, 1.2))
}

func foo() {
	fmt.Println("Hello World")
	var p *string
	s := "ab"
	fmt.Println(&s) //0xc00008a030
	p = &s
	*p += "cd"
	for range 1000 {
		*p += "0"
	}
	fmt.Println(s, &s, p) //0xc00008a030 0xc00008a030

	str := "AB"
	fmt.Println(&str) //0xc00008a050
	str += "CD"
	fmt.Println(&str) //0xc00008a050
	for range 100 {
		str += "F"
	}
	fmt.Println(&str) //0xc00008a050

	//	go中字符串为可变对象,+=后会在原地操作,不会创建新的string
}

func bar() {
	s := ""
	// 想要修改一个变量的值,就不要把他放在等号右边的表达式,右边编译器只会计算出值
	// 只有在左边,才有内存地址的含义
	// 如果非要放在右边,就传指针
	s = s + "A"
	fmt.Println(s)

	s += "B"
	fmt.Println(s)

	tmp := s
	tmp += "C"
	fmt.Println(s, tmp) //修改的是tmp

	p := &s
	*p += "D"
	fmt.Println(s)

	addStr1(s)
	//相当于
	//param := s
	//addStr1(param)
	fmt.Println(s)

	addStr2(s)
	fmt.Println(s)

	addStr3(&s)
	fmt.Println(s)

	addStr4(&s)
	fmt.Println(s)
}

func addStr1(s string) {
	s += "E"
}

func addStr2(s string) {
	tmp := s
	tmp += "F"
}

func addStr3(s *string) {
	*s += "G"
}

func addStr4(s *string) {
	tmp := s
	*tmp += "H"
}

// baz(2,1)=>1
// baz(2.1,1.1)=>1
func baz(a, b any) int {
	switch a := a.(type) {
	case int:
		fmt.Println(a)
		switch b := b.(type) {
		case int:
			fmt.Println(b)
			return a - b
		}
	case float64:
		fmt.Println(a)
		switch b := b.(type) {
		case float64:
			fmt.Println(b)
			return int(a - b)
		}
	default:
		return 0
	}
	return 0
}
