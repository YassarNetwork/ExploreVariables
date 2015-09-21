// ExploreVariable project main.go
package main

import (
	"fmt"
)

func main() {
	
		// Changed to test Git Hub
	//TODO : Naming covention 4variables
	
	//////////////////////// Int //////////////////////////
	var intSize int=3
	fmt.Println(intSize)
	
	//////////////////////// Float //////////////////////////
	// Its good practise to explicitly put ".", as Go will not do implicit conversion
	var fltWeight float32=42.12
	
	// Float gets printed differently than example on plural
	fmt.Println(fltWeight)
	
	//////////////////////// Implicit Declaration //////////////////////////
	// In this case, we are not using var keyword and also declaring variable type, Go is going to automatically determine type
	// from value being assigned
	myString:="Helloooooooooo"	
	fmt.Println(myString)
	intA:=6
	fmt.Println(intA)
	
	//////////////////////// Copmplex Variable //////////////////////////
	myComplex:=complex(2,3)
	
	
	
	
	fmt.Println(imag(myComplex))
	
	fmt.Println("Explore Constants")
	
	// Any primitive data type can be assigned as constant values.
	// No data Type defined for const
	const 
	(
		First="First"
		Second=2
		Third="Third"
		
	)
	
	fmt.Println(First)
	fmt.Println(Second)
	const 
	(
		Red=iota
		// Note : No values are assigned to these constants
		Black
		Orange
		
	)
	// Iota is a keywork which auto incerement values in a constant block. We can use it
	// for enum
	fmt.Println(Red)
	fmt.Println(Black)
	fmt.Println("Explore Constant Expression")
	
	// Constant expression have to be resolveab;e at compile time so they can not have function calls
	const 
	(
		fi=1<<(10*iota) 
		// Note : No values are assigned to these constants
		sec
		thi
		
	)
	fmt.Println(fi)
	fmt.Println(sec)
	fmt.Println(thi)
}
