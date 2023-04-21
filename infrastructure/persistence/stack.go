package persistence

import "github.com/arturiamu/lplms-public_cloud/domain/repository"

var stk Stack

type stack interface {
	User() repository.StorageRepositoryInterface
	Compute() repository.ComputeRepositoryInterface
	Storage() repository.StorageRepositoryInterface
	Network() repository.NetworkRepositoryInterface
}

type Stack struct {
	U *UserRepo
	C *ComputeRepo
	S *StorageRepo
	N *NetworkRepo
}

func (s *Stack) User() repository.UserRepositoryInterface {
	return s.U
}

func (s *Stack) Compute() repository.ComputeRepositoryInterface {
	return s.C
}

func (s *Stack) Storage() repository.StorageRepositoryInterface {
	return s.S
}

func (s *Stack) Network() repository.NetworkRepositoryInterface {
	return s.N
}
