package error

import "fmt"

type ServerErr struct {
}

func (r *ServerErr) Msg() (m string) {
	return "server err"
}

func (r *ServerErr) MsgID() (mID string) {
	return "E0031"
}

func (r *ServerErr) Error() string {
	return fmt.Sprintf("%v : %v", r.MsgID(), r.Msg())
}
