package main

func main() {

	users := []Emp{
		{ename: "Sai", age: 24, role: "Developer"},
		{ename: "Aravindh", age: 23, role: "Manager"},
		{ename: "SivaKumar", age: 23, role: "Developer"},
		{ename: "Kalaiselvi", age: 22, role: "Manager"},
		{ename: "keerthika", age: 23, role: "Developer"},
		{ename: "Dinesh", age: 22, role: "Manager"},
	}

}

type Employee interface {
	empSort()
}

type Emp struct {
	ename string
	age   uint
	role  string
}
