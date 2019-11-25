package base

type Service interface {
	SetOpt(opt *ClientOpt)
	Start()
}
