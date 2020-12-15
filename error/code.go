package error

type Code interface {
	Msg() (m string)
	MsgID() (mID string)
}
