package main

import (
	"fmt"
	"github.com/utopia-go/http/http"
)

func main() {
	fmt.Println(http.Run(3))
	fmt.Println(http.PreparePath("/v1/hea/:document/collection"))
}
