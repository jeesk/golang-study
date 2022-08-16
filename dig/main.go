package main

import (
	"fmt"
	"go.uber.org/dig"
)

type User struct {
	ID int64
}

type Repo interface {
	Get(int64) (*User, error)
}

type repo struct{}

func (r *repo) Get(id int64) (*User, error) {
	return &User{ID: id}, nil
}

func NewRepo() Repo {
	return &repo{}
}

type Service struct {
	r Repo
}

func NewService(r Repo) *Service {
	return &Service{r: r}
}
func BuildContainer() *dig.Container {
	container := dig.New()
	container.Provide(NewService)
	container.Provide(NewRepo)
	return container
}
func main() {
	container := BuildContainer()
	container.Invoke(func(s *Service) {
		u, err := s.r.Get(1)
		fmt.Printf("user %+v err %+v", u, err)
	})
	err := container.Invoke(func(r Repo) {
		get, err := r.Get(3)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(get.ID)
	})
	if err != nil {
		return
	}
}
