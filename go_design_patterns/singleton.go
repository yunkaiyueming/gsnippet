package main

type Single struct{}

var OneStance *Single

func GetInstance() *Single {
	if OneStance == nil {
		return new(Single)
	}
	return OneStance
}

func main() {
	GetInstance()
}
