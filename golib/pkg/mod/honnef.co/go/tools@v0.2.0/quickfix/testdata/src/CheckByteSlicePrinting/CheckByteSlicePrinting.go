package pkg

import "fmt"

type Stringable []byte

func (*Stringable) String() string { return "" }

func fn() {
	type ByteSlice []byte
	var b1 []byte
	var b2 ByteSlice
	var b3 *Stringable
	var b4 Stringable

	var s string
	fmt.Print(1, b1, 2, []byte(""), b2, s)       // want `could convert argument to string` `could convert argument to string` `could convert argument to string`
	fmt.Fprint(nil, 1, b1, 2, []byte(""), b2, s) // want `could convert argument to string` `could convert argument to string` `could convert argument to string`
	fmt.Print()
	fmt.Fprint(nil)

	fmt.Println(b3)
	fmt.Println(b4) // want `could convert argument to string`
}
