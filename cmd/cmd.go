package main

import (
	"back-end-2020-1/app/account"
	"back-end-2020-1/bootstrap"
)

func main() {
	account.G_username = "abc"
	bootstrap.Init()
}
