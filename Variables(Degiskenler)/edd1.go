package main

import "fmt"

func main() {
	var sinama bool
	var kavram1 string = "Derece"
	var kavram2 string = "Fahrenhiet"
	sinama = kavram1 == kavram2 //önce bu

	fmt.Println(sinama)
	fmt.Printf("typedef: % T\n", sinama) //tip belirleme

	sinama = kavram1 != kavram2
	fmt.Println(sinama)  //üsttekini sına
	fmt.Println(!sinama) //ters
}
