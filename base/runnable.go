package base

type Runner interface {
	Run(data JobData)
}

type BaseRunner struct {
}
