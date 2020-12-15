package service

import (
	"fmt"
	ERROR "github.com/godaner/goerrordemo/error"
	"testing"
)

func TestService_F1(t *testing.T) {
	s := &Service{}
	for i := 0; i <= 3; i++ {
		err := s.F1(i)
		if err != nil {
			switch err := err.(type) {
			case *ERROR.NotFoundErr:
				fmt.Println("not found err ,", err.Error())
			case *ERROR.ServerErr:
				fmt.Println("server err ,", err.Error())
			case *ERROR.RequestErr:
				fmt.Println("request err ,", err.Error())
			default:
				fmt.Println("no error")
			}
		}
	}

}
