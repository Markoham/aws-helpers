package main

import (
	"fmt"
	"os"

	"github.com/logrusorgru/aurora"
)

func main() {
	region := "eu-west-1"
	if len(os.Args) > 1 {
		region = os.Args[1]
	}

	fmt.Print(aurora.Magenta("Please wait... "))
	result, err := getInstances(region)
	if err != nil {
		panic(err)
	}
	fmt.Println(aurora.Green("Done!").Bold())
	instances := parseResult(result)

	showInstances(instances)
}
