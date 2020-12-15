package code

import "fmt"

type Code interface {
	Msg() (m string)
	MsgID() (mID string)
	String() (str string)
}

// common
var (
	CodeSuccess = &SuccessCode{}
)

type SuccessCode struct {
}

func (s *SuccessCode) String() (str string) {
	return fmt.Sprintf("%v : %v", s.MsgID(), s.Msg())
}

func (s *SuccessCode) Msg() (m string) {
	return "S0001"
}

func (s *SuccessCode) MsgID() (mID string) {
	return "success"
}
