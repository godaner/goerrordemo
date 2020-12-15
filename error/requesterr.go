package error

import "fmt"

type RequestErr struct {
}

func (r *RequestErr) Msg() (m string) {
	return "request err"
}

func (r *RequestErr) MsgID() (mID string) {
	return "E0032"
}

func (r *RequestErr) Error() string {
	return fmt.Sprintf("%v : %v", r.MsgID(), r.Msg())
}
