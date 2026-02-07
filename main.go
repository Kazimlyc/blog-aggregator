package main

import (
	"fmt"

	"github.com/Kazimlyc/blog-aggregator/internal/config"
)

func main() {

	cfg, err := config.Read()
	if err != nil {
		panic(err)
	}

	if err := cfg.SetUser("Kazim"); err != nil {
		panic(err)
	}

	cfg2, err := config.Read()
	if err != nil {
		panic(err)
	}

	fmt.Println(cfg2)

}
