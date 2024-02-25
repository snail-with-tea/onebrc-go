package generator

import (
	"flag"
	"fmt"
)

var GenSize string

func init() {
	const (
		deflt = "50k"
		usage = "set generation size;"
	)
	flag.StringVar(&GenSize, "genSize", deflt, usage)
	flag.StringVar(&GenSize, "S", deflt, usage+" (shorthand)")
}

func main() {
	fmt.Println("Let's generate some files")
	flag.Parse()
	fmt.Println(GenSize)
}

func Start() {
	fmt.Println("Let's generate some files")
}
