package main

import "flag"

func options() {
	var optVer bool
	flag.BoolVar(&optVer, "v", false, "print version infomation")
	flag.Func("t", "test flag func", func(s string) error {
		print(flag.Func, s)
		return nil
	})

	flag.Parse()
	if optVer {
		print("v0.0.0")
	}

}
func main() {
	options()
}
