package main

import "fmt"

type Details struct {
	name   string
	age    uint
	salary uint
}

func main() {

	var array2 = [...]int{10, 30, 90, 20, 60, 80, 100, 40, 70}
	// // array1 = append(array1, array2)
	// fmt.Println("array2", array2)

	// var arr = [5]int{}
	// fmt.Println("arr", arr)

	// // arr = append(arr, array2)
	// // fmt.Println("arr", arr)

	// var my_array_1 = [3]string{"Geeks", "for", "Geeks"}
	// c := my_array_1
	// c[0] = "geek"
	// fmt.Println(c)
	// fmt.Println(my_array_1)

	// d := []int{}
	// fmt.Println(d)

	// d = append(d, 10, 20, 30)
	// e := d

	// fmt.Println(d)
	// fmt.Println(e)
	// fmt.Printf("d address :%p \n", &d)
	// fmt.Printf("e address :%p \n", &e)

	// e[0] = 100 // *
	// fmt.Println(d)
	// fmt.Println(e)
	// fmt.Printf("d address :%p \n", &d)
	// fmt.Printf("e address :%p \n", &e)

	// e[1] = 2000 // *
	// fmt.Println(d)
	// fmt.Println(e)
	// fmt.Printf("d address :%p \n", &d)
	// fmt.Printf("e address :%p \n", &e)

	// fmt.Println("After Appending ")
	// e = append(e, 5000) //
	// fmt.Println(d)
	// fmt.Println(e)
	// fmt.Printf("d address :%p \n", &d)
	// fmt.Printf("e address :%p \n", &e)

	// e[1] = 100 // x
	// fmt.Println(d)
	// fmt.Println(e)
	// fmt.Printf("d address :%p \n", &d)
	// fmt.Printf("e address :%p \n", &e)
	// fmt.Println("********************************************************************************")

	// e = append(e, 1000)
	// d = append(d, 2000)

	// fmt.Println(d)
	// fmt.Println(e)

	// e[1] = 50
	// fmt.Println(d)
	// fmt.Println(e)

	// var number float32 = 45.85
	// var name string = "jeevan"
	// fmt.Printf("The value is %T and the name is %T \n", number, name)

	// var p int = 10
	// q := p
	// fmt.Println(p)
	// fmt.Println(q)

	// q = 20
	// fmt.Println(p)
	// fmt.Println(q)

	// g := []int{10, 20, 30, 40, 50}
	// z := []int{60, 70, 80, 90, 100}

	// z = append(z, g...)
	// fmt.Println(z)

	// z = append(z, 11)
	// fmt.Println(g)

	// a := make(map[int]string)
	// fmt.Println("a : ", a)

	// a[1] = "a"
	// a[2] = "b"
	// a[3] = "c"
	// a[4] = "d"
	// a[5] = "e"

	// fmt.Println(z)
	// delete(a, 1)

	// fmt.Println(a)

	// fmt.Println(z[1])

	// var j = len(z)

	// var k []int

	// for i := 0; i < len(g); i++ {
	// 	k[i+j] = z[i]
	// 	k[i] = g[i]
	// }

	// h := make([]int, 5)
	// i := [5]int{}

	// fmt.Println(g)
	// fmt.Println(h)
	// fmt.Println("***************************")

	// copy(h, g)
	// for index, value := range g {
	// 	i[index] = value
	// }

	// fmt.Println("g : ", g)
	// fmt.Println("h : ", h)
	// fmt.Println("i : ", i)

	// h[2] = 100
	// fmt.Println(g)
	// fmt.Println(h)

	// fmt.Println(cap(h))

	// h = append(h, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15)
	// fmt.Println(h)
	// fmt.Println(cap(h))

	// k := []int{5, 10, 15}

	// g = append(g, k)

	// for j := 0; j < len(h); j++ {
	// 	k[j] = h[j]
	// }
	// fmt.Println(k)

	// fmt.Printf("type of k is %T and type of h is %T \n", k, g)

	// a := make(map[int]byte)
	// a[1] = 65
	// a[2] = 66
	// a[3] = 67
	// a[4] = 68
	// a[5] = 69
	// a[6] = 70
	// fmt.Println(a)

	// b := make(map[int]byte)

	// // delete(a, 1)
	// // fmt.Println(a)
	// for i, value := range a {
	// 	// fmt.Println("the key is ", i, " and the value is ", value)
	// 	if i%2 != 0 {
	// 		continue
	// 	} else {
	// 		b[i] = value
	// 	}
	// }

	// fmt.Println(b)

	// var c float32 = 45.58

	// println(c)
	// fmt.Println(c)

	// var user1 Details
	// user1.name = "Jeeva"
	// user1.age = 22
	// user1.salary = 854500

	// var user2 = Details{"Jee", 21, 755555}
	// // user1.health = "good"

	// fmt.Println(user1.name)
	// fmt.Println(user1.age)
	// fmt.Println(user1.salary)
	// // fmt.Println(user1.health)
	// fmt.Println("User 1 : ", user1)

	// fmt.Println(user2.name)
	// fmt.Println(user2.age)
	// fmt.Println(user2.salary)
	// // fmt.Println(user1.health)
	// fmt.Println("User 2 : ", user2)

	// user3 := &Details{"sai", 23, 555555}
	// fmt.Println("User 3 : ", *user3)

	// var num int

	// fmt.Println("Enter a number : ")
	// fmt.Scan(&num)

	// fmt.Println("Number is ", num)

	var l = 10

	fmt.Println("l = ", l)
	fmt.Printf("l address : %p \n", &l)

	var m = l
	fmt.Println("m = ", m)
	fmt.Printf("m address : %p \n", &m)

	n := &l
	fmt.Println("n = ", n)
	fmt.Printf("n address : %p \n and type is %T \n", &n, n)

}
