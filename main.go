package main

import (
	"github.com/Likhon22/ecom/cmd"
	_ "github.com/Likhon22/ecom/database"
)

type User interface {
	PrintDetails()
	ReceiveMoney(amount float64) float64
}

func main() {

	cmd.Serve()
}
