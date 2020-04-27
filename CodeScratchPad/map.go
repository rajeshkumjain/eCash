package main

import "fmt"

//const C001 = "ERROR: Invalid header format : Expecting application/jon in the header ... "
//const C002 = "ERROR: Unable to decode the content from request body ... "
//const C003 = "ERROR: Unable to decode the JSON packet ... "
//const C004 = "ERROR: Unable to complete the registration process ... "
//const S001 = "SUCCESS: Thanks"

var ErrorMap = map[string]string{
	"C001": "ERROR: Invalid header format : Expecting application/jon in the header ... ",
	"C002": "ERROR: Unable to decode the content from the http request body ... ",
	"C003": "ERROR: Unable to decode the JSON packet ... ",
	"C004": "ERROR: Unable to complete the registration process ... ",
	"S001": "SUCCESS: Thank you - your registration completed successfully ...",
}

func main() {
	fmt.Println(ErrorMap["C001"])
	print(ErrorMap["C001"])

}
