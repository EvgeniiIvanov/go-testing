package users

import "errors"

type User struct {
	ID   int
	Name string
}

var ErrNotFound = errors.New("user not found")
var ErrAlreadyExists = errors.New("user already exists")

type UserStorage interface {
	Save(u User) error
	ByID(id int) (User, error)
}

type Service struct{ s UserStorage }

func New(s UserStorage) *Service { return &Service{s: s} }

func (svc *Service) Add(u User) error { return svc.s.Save(u) }

func (svc *Service) Get(id int) (User, error) { return svc.s.ByID(id) }
