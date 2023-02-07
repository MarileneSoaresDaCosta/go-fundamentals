package main

import (
	"fmt"
	"strings"
)

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

	type rectangle struct {
		length float64
		width  float64
	}

	fmt.Println(rectangle{10.5, 21.3})

	var rec rectangle
	rec.length = 1.3
	rec.width = 6.9
	fmt.Println(rec)

	type menuItem struct {
		name   string
		prices map[string]float64
	}
	// menu is a slice of menuItems - a slice of structs
	menu := []menuItem{
		{name: "Coffee", prices: map[string]float64{"small": 1.65, "medium": 1.80, "large": 1.95}},
		{name: "Espresso", prices: map[string]float64{"single": 1.90, "double": 2.25, "triple": 2.55}},
	}
	for _, item := range menu {
		fmt.Println(item.name)
		fmt.Println(strings.Repeat("-", 10))
		for size, price := range item.prices {
			fmt.Printf("\t%10s%10.2f\n", size, price)
		}
	}

}

// part of the 'web service'
// func Handler(w http.ResponseWriter, r *http.Request) {
// 	f, _ := os.Open("./menu.txt")
// 	io.Copy(w, f)
// }
