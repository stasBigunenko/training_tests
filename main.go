package main

import "fmt"

type GetArgs interface {
	GetA() int
	GetB() int
	//GetArgs() (int, int)
}

type Args struct {
	args1 int
	args2 int
}

func Sum(args GetArgs) int {
	//num1, num2 := args.GetArgs()
	//return num1 + num2
	return args.GetA() + args.GetB()
}

//func (n Args) GetArgs() (int, int) {
//	return n.GetA(), n.GetB()
//}

func (n Args) GetA() int {
	return n.args1
}

func (n Args) GetB() int {
	return n.args2
}

func main() {
	args := Args{
		args1: 2,
		args2: 4,
	}
	fmt.Println(Sum(args))
}