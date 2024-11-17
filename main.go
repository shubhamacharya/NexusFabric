package main

import (
	"fmt"
	"os"

	"github.com/shubhamacharya/NexusFabric/cmd"
	"github.com/shubhamacharya/NexusFabric/internal/controllers"
)

func main() {
	controllers.ReadSpecsFile("")
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
