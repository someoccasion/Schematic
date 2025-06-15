package constraints

import "github.com/hydridity/Schematic/pkg/schema/context"

// The basic interface of our constraints.
type Constraint interface {
	Consume([]string, *context.ValidationContext) ([]string, error)
	String() string
	GetVariableName() string
}
