package main

func readByExt(ext string) {
	if ext == "json" {
		println("read json file")
	} else if ext == "jpg" || ext == "jpeg" || ext == "png" || ext == "gif" {
		println("read image file")
	} else if ext == "txt" || ext == "md" {
		println("read text file")
	} else if ext == "yml" || ext == "yaml" {
		println("read yaml file")
	} else if ext == "ini" {
		println("read ini file")
	} else {
		println("unsupported file extension:", ext)
	}
}

func readByExtBySwitch(ext string) {
	switch ext {
	case "json":
		println("read json file")
	case "jpg", "jpeg", "png", "gif":
		println("read image file")
	case "txt", "md":
		println("read text file")
	case "yml", "yaml":
		println("read yaml file")
	case "ini":
		println("read ini file")
	default:
		println("unsupported file extension:", ext)
	}
}

func case1() int {
	println("eval case1 expr")
	return 1
}

func case2_1() int {
	println("eval case2_1 expr")
	return 0
}
func case2_2() int {
	println("eval case2_2 expr")
	return 2
}

func case3() int {
	println("eval case3 expr")
	return 3
}

func switchexpr() int {
	println("eval switch expr")
	return 2
}

func main() {
	switch 2 {
	case case3():
		println("exec case3")
	case case2_1(), case2_2():
		println("exec case2")
		fallthrough
	case case1():
		println("exec case1")

	default:
		println("exec default")
	}
}
