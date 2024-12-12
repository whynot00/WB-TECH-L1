package set

import (
	"fmt"
	"strings"
)

type (
	void struct{}

	Set struct {
		mp map[int]void
	}
)

func NewSet() *Set {
	return &Set{
		mp: make(map[int]void),
	}
}

func (s *Set) Add(data ...int) {
	for _, n := range data {
		s.mp[n] = void{}
	}
}

func (s *Set) Del(n int) {
	delete(s.mp, n)
}

func (s *Set) Intersection(setTwo *Set) *Set {

	tempMap := make(map[int]void)

	leadMap := s.mp
	slaveMap := setTwo.mp
	if len(leadMap) < len(slaveMap) {
		leadMap, slaveMap = slaveMap, leadMap
	}

	for item := range leadMap {
		if _, ok := slaveMap[item]; ok {
			tempMap[item] = void{}
		}
	}

	return &Set{
		mp: tempMap,
	}
}

func (s *Set) Show() {

	tempRes := make([]string, 0, len(s.mp))

	for item := range s.mp {
		tempRes = append(tempRes, fmt.Sprint(item))
	}

	fmt.Println(strings.Join(tempRes, ", "))
}
