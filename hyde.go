package main

import (
	"fmt"
	"github.com/gr4y/hyde/lib"
)

func main() {
	conf := lib.Configuration{}
	if err := conf.Read(); err != nil {
		panic(err)
	}

	buildContext := lib.BuildContext{Configuration: conf}
	if err := buildContext.Build(); err != nil {
		panic(err)
	}

	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	// TODO: write to outputPath
	// fmt.Println(fmt.Sprintf("Reading File %s", f))

}
