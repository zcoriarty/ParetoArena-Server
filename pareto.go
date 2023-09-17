package pareto

import (
	"github.com/zcoriarty/Backend/cmd"
	"github.com/zcoriarty/Backend/route"
)

// New creates a new Pareto instance
func New() *Pareto {
	return &Pareto{}
}

// Pareto allows us to specify customizations, such as custom route services
type Pareto struct {
	RouteServices []route.ServicesI
}

// WithRoutes is the builder method for us to add in custom route services
func (g *Pareto) WithRoutes(RouteServices ...route.ServicesI) *Pareto {
	return &Pareto{RouteServices}
}

// Run executes our pareto functions or servers
func (g *Pareto) Run() {
	cmd.Execute(g.RouteServices)
}
