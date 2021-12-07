package main

var a, b, c int

func Check() string {
	if a == b && b == c {
		return "Segitiga sama sisi"
	} else if a == b && b != c {
		return "Segitiga sama kaki"
	} else if b == c && a != b {
		return "Segitiga sama kaki"
	} else if a == c && a != b {
		return "Segitiga sama kaki"
	} else {
		return "Segitiga sembarang"
	}
}
func main() {}
