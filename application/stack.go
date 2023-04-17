package application

type StackInterface interface {
	User() UserAppInterface
	Compute() ComputeAppInterface
}
type stack struct {
	ComputeAppInterface
	UserAppInterface
}

func (s *stack) User() UserAppInterface {
	return s.UserAppInterface
}

func (s *stack) Compute() ComputeAppInterface {
	return s.ComputeAppInterface
}
