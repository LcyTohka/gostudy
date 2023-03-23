package main

type E1 interface {
	M1()
	M2()
	M3()
}

type E2 interface {
	M1()
	M2()
	M4()
}

type K struct {
	E1
	E2
}

func main() {
	//t := K{}
	//t.M1()
	//t.M2()
}
