package error

import "fmt"

type NotFoundErr struct {
}

func (r *NotFoundErr) Msg() (m string) {
	return "notfound err"
}

func (r *NotFoundErr) MsgID() (mID string) {
	return "E0033"
}

func (r *NotFoundErr) Error() string {
	return fmt.Sprintf("%v : %v", r.MsgID(), r.Msg())
}
