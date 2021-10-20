package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"gql-demo/graph/generated"
	"gql-demo/graph/model"
	"math/rand"
	"time"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	todo := &model.Todo{
		Text:   input.Text,
		ID:     fmt.Sprintf("T%d", rand.Int()),
		UserID: input.UserID,
	}
	r.todos = append(r.todos, todo)
	go func() {
		for id, todoSubItem := range r.DurationMapSubscription {
			fmt.Printf("%d \n", id)
			*todoSubItem <- r.todos
		}
	}()

	return todo, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	return r.todos, nil
}

func (r *subscriptionResolver) Todos(ctx context.Context) (<-chan []*model.Todo, error) {
	id := time.Now().UnixNano()
	todos := make(chan []*model.Todo)
	r.mu.Lock()
	if r.DurationMapSubscription == nil {
		r.DurationMapSubscription = make(map[int64]*chan []*model.Todo)
	}
	r.DurationMapSubscription[id] = &todos
	r.mu.Unlock()
	fmt.Printf("connect id: %d\n", id)
	go func() {
		<-ctx.Done()
		r.mu.Lock()
		delete(r.DurationMapSubscription, id)
		r.mu.Unlock()
		fmt.Printf("disconnect id: %d\n", id)
		fmt.Println(r.DurationMapSubscription)
	}()
	//go func() {
	//	ticker := time.NewTicker(time.Second)
	//	for {
	//		<-ticker.C
	//		if r.todos != nil {
	//			*r.DurationMapSubscription[id] <- r.todos
	//		}
	//	}
	//}()

	return *r.DurationMapSubscription[id], nil
}

func (r *todoResolver) User(ctx context.Context, obj *model.Todo) ( *model.User, error) {
	user := model.User{}

	return &user, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Subscription returns generated.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }

// Todo returns generated.TodoResolver implementation.
func (r *Resolver) Todo() generated.TodoResolver { return &todoResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
type todoResolver struct{ *Resolver }
