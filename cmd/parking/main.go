package main

import (
	"os"
	"parkee/internal/cli"
	"parkee/internal/repository/memory"
	"parkee/internal/service"
)

func main() {
	repo := memory.NewMemoryParkingRepository()

	svc := service.NewParkingService(repo)

	c := cli.NewCLI(svc)

	if len(os.Args) < 2 {
		panic("Please provide input file")
	}

	filename := os.Args[1]

	c.RunFromFile(filename)
}
