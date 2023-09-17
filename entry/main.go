package main

import (
	"fmt"

	pareto "github.com/zcoriarty/Backend"
)

func main() {
	pareto.New().
		WithRoutes(&MyServices{}).
		Run()
}

// MyServices implements github.com/zcoriarty/Backend/route.ServicesI
type MyServices struct{}

// SetupRoutes is our implementation of custom routes
func (s *MyServices) SetupRoutes() {
	fmt.Println("set up our custom routes!")
}
