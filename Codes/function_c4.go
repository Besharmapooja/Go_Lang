package main

import "fmt"

func main() {
	simplefunnction()
	functionwPara("age", 25)
	//var i, j int = 5, 10
	i, j := 5, 10
	functionwParas(i, j)
	fmt.Println(functionWithReturnValue())

	x, y := retsameValue(2, 4)
	fmt.Println(x, y)

	fmt.Println(incrementnum(2, 3))

}

func simplefunnction() {
	fmt.Println("print-something")
}

//function with diff parameters
func functionwPara(param1 string, param2 int) {
	fmt.Println(param1, param2)
}

//functionwith same parametes
func functionwParas(param1, param2 int) {
	fmt.Println(param1, param2)
}

func functionWithReturnValue() int {
	return 2325
}

func fucwithmultiRet() (string, int) {
	return "age", 12
}
func retsameValue(param1, param2 int) (int, int) {
	return param1, param2
}
func incrementnum(i, j int) (int, int) {
	i++
	j++
	return i, j
}
