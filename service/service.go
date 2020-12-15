package service

import (
	ERROR "github.com/godaner/goerrordemo/error"
)

type Service struct {
}

func (s *Service) F1(i int) (err error) {
	if i == 0 {
		return ERROR.ErrServer
	}
	if i == 1 {
		return ERROR.ErrNotFound
	}
	if i == 2 {
		return ERROR.ErrRequest
	}
	return nil
}
