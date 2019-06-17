package main

import (
	"fmt"
	"github.com/mogbeyi-david/go_crash_course/03_packages/strutil"
	"math"
)

func main() {
	fmt.Println(math.Floor(2.7))
	fmt.Println(math.Ceil(2.7))
	fmt.Println(math.Sqrt(54))
	fmt.Println(strutil.Reverse("Hello world"))
}
