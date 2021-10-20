//go:generate go run github.com/99designs/gqlgen
package graph

import (
	"gql-demo/graph/model"
	"sync"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
	todos []*model.Todo
	mu sync.Mutex
	DurationMapSubscription map[int64]*chan []*model.Todo
}
