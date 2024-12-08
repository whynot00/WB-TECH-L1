package main

import "fmt"

type Human struct {
	HumanID int64
	name    string
}

func NewHuman() *Human {
	return &Human{}
}

func (h *Human) SetName(name string) {
	h.name = name
}

func (h *Human) GetName() string {
	return h.name
}

// action

type Action struct {
	actionID int64
	Human
}

func NewAction() *Action {
	return &Action{}
}

func main() {
	action := NewAction()

	action.SetName("Ivan")

	action.HumanID = 1119
	action.actionID = 9222

	fmt.Printf("%s : %d : %d\n", action.GetName(), action.HumanID, action.actionID)

}
