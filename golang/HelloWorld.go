package main

type UserState int32

const (
	Starting UserState = iota
	GroupStarted
	BothStarted
)

func main() {
	client := newClient()
	removeObject(client)
}
