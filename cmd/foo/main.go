package main

import (
	"math/rand"
	"time"

	_ "go.uber.org/automaxprocs"

	"iam-based-app/internal/apiserver"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	apiserver.NewApp("iam-apiserver").Run()
}
