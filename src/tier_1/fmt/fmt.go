package fmt_tier_1

import (
	"fmt"
)

type LocalTier1 struct{}

func (t *LocalTier1) Fmt_all_packages() {

	fmt.Printf("\n================================== Printing all tier1 standard packages ==============================\n\n")

	fmt.Println("1. Hello world this is shubh printing this statement using fmt.Println from fmt package")
	fmt.Print("2.", "Printing", "this", "line", "line", "using", "fmt.Print")
	name := "shubham"
	fmt.Printf("\n3. Printing my name using fmt.Printf : %s\n", name)
	age := 27
	fmt.Printf("4. Printing my age using fmt.Printf : %d\n", age)

}
