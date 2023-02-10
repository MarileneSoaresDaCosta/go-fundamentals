package main

import (
	"bufio"
	"demo/menu"
	"fmt"
	"os"
	"strings"
)

var in = bufio.NewReader(os.Stdin)

func main() {
	// fmt.Println("Hello, Gophers!")
	// fmt.Println("Here's a escape char \n in an interpreted string")
	// fmt.Println(`Here's a escape char \n in a raw string`)
	// fmt.Println(`raw strings
	// ignore new lines`)
	// d := 3.14
	// fmt.Println(d)
	// var e int = int(d)
	// fmt.Println(e)

	// a := 42

	// b := &a

	// fmt.Println(*b)
	// // we can change the content of 'a' by dereferencing b!
	// *b = 27
	// fmt.Println(a) // 27
	// // another way to generate a pointer using 'new' - anonymous memory
	// p = new(string)

	/*
		*  Building a cli application
		/*
		// fmt.Println("What would you like me to scream?")
		// in := bufio.NewReader(os.Stdin)
		// s, _ := in.ReadString('\n')
		// s = strings.TrimSpace(s)
		// s = strings.ToUpper(s)
		// fmt.Println(s + "!!")

		/*
		*  Building a web service
	*/
	// http.HandleFunc("/", Handler)
	// http.ListenAndServe(":3000", nil) // Go knows we mean 'localhost:...'

	/*
		slices
	*/
	// primes := [6]int{2, 3, 5, 7, 11, 13}
	// primes2 := primes
	// fmt.Println(primes == primes2)

	// var s []int = primes[1:4]
	// fmt.Println(s)

	// sc := s
	// fmt.Println(sc)

	// s = append(s, 23)
	// fmt.Println(s)
	// fmt.Println(sc)

	// s = slices.Delete(s, 3, 4)
	// fmt.Println(s)
	// fmt.Println(sc)
	// // fmt.Println(s == sc) // error - slices cannot be compared

	// x := []int{1, 2, 3}
	// x = append(x, 3)
	// fmt.Println(x)

	/*
	* Map
	 */
	// var m map[string]int
	// fmt.Println(m)

	// m2 := map[string]int{"a": 1, "b": 2}
	// fmt.Println(m2)
	// fmt.Println(m2["foo"]) // expect no results, but get zero
	// v, ok := m2["foo"]     // comma ok syntax verifies presence
	// fmt.Println(v, ok)

	// m3 := map[string]int{
	// 	"foo": 1,
	// 	"bar": 2,
	// 	"baz": 3,
	// }
	// // maps are copied by reference
	// m4 := m3
	// m3["foo"], m4["bar"] = 99, 42
	// fmt.Println(m3)
	// fmt.Println(m4)

	/*
	* Struct
	 */

	// type rectangle struct {
	// 	length float64
	// 	width  float64
	// }

	// fmt.Println(rectangle{10.5, 21.3})

	// var rec rectangle
	// rec.length = 1.3
	// rec.width = 6.9
	// fmt.Println(rec)

	/* if statements */
	// i := 4
	// if i < 5 {
	// 	fmt.Println("i is less than 5")
	// } else if i < 10 {
	// 	fmt.Println("i is NOT less than 5 but less than 10")
	// } else {
	// 	fmt.Println("i is at least 10")
	// }
	// fmt.Println("after the if statement")

	/* switch */

	/* logical switch */
	// switch i := 8; {
	// case i < 5:
	// 	fmt.Println("i is less than 5")
	// case i < 10:
	// 	fmt.Println("i is less than 10")
	// default:
	// 	fmt.Println("i is greater than 10")
	// }

	/* Deferred Functions
	 * we expect to see main 1, main 2, defer 2 and defer 1 in the output
	 * because the deferred functions are executed in FIRST IN LAST OUT fashion
	 * this is important when we need to 'release resources', such as closing a db
	 * connection.
	 */
	// fmt.Println("main 1")
	// defer fmt.Println("defer 1")
	// fmt.Println("main 2")
	// defer fmt.Println("defer 2")

	/* Panic and Recover */
	// dividend, divisor := 10, 5

	// fmt.Printf("%v divided by %v is %v\n", dividend, divisor, divide(dividend, divisor))
	// // inducing panic >> panic: runtime error: integer divide by zero
	// dividend, divisor = 10, 0
	// fmt.Printf("%v divided by %v is %v\n", dividend, divisor, divide(dividend, divisor))

	// 	result, ok := divide2(1, 2)
	// 	if ok {
	// 		fmt.Println(result)
	// 	}

	// 	result, ok = divide2(1, 0)
	// 	if ok {
	// 		fmt.Println(result)
	// 	}

	/* our menu app */

	// a label loop
loop:
	for {
		fmt.Println("Please select an option")
		fmt.Println(" 1) Print menu")
		fmt.Println(" 2) Add item")
		fmt.Println(" q) Quit")
		choice, _ := in.ReadString('\n')

		switch strings.TrimSpace(choice) {
		case "1":
			menu.Print()
		case "2":
			menu.AddItem()
		case "q":
			break loop
		default:
			fmt.Println("Unknown option")
		}
	}
}

// /* function used in the panic and recover example */
// func divide(dividend, divisor int) int {
// 	defer func() {
// 		// condition avoids nil being printed when there is no panic
// 		// but, at this point, we get runtime error: integer divide by zero
// 		// and "10 divided by 0 is 0", which is obviously wrong.
// 		// we'll find out how to fix this in the next module
// 		if msg := recover(); msg != nil {
// 			fmt.Println(msg)
// 		}
// 	}()
// 	return dividend / divisor
// }

// func divide2(l, r int) (int, bool) {
// 	if r == 0 {
// 		return 0, false
// 	}
// 	return l / r, true
// }

/*
*  part of the 'web service'
 */
// func Handler(w http.ResponseWriter, r *http.Request) {
// 	f, _ := os.Open("./menu.txt")
// 	io.Copy(w, f)
// }
