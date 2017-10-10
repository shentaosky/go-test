package main

func main() {
	var a int
	a = -1
	if a > 0 {
		println(a)
		defer println(a)

	}
}
