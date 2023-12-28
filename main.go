package main

import (
	apps "github.com/chanaseptiari/learn-auth-golang/apps"
	"github.com/chanaseptiari/learn-auth-golang/apps/initializer"
)

func init() {
	// Load ENV
	initializer.LoadEnv()
	// Connect Database
	initializer.ConnectDatabase()
}

func main() {
	// Running Apps
	apps.SetupRouter().Run(":8000")
}
