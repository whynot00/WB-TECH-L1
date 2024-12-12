package set

import (
	"fmt"
	"strings"
)

type (
	void struct{}

	Set struct {
		mp map[string]void
	}
)

func NewSet() *Set {
	return &Set{
		mp: make(map[string]void),
	}
}

func (s *Set) Add(data ...string) {
	for _, n := range data {
		s.mp[n] = void{}
	}
}

func (s *Set) Show() {

	tempRes := make([]string, 0, len(s.mp))

	for item := range s.mp {
		tempRes = append(tempRes, fmt.Sprint(item))
	}

	fmt.Println(strings.Join(tempRes, ", "))
}
