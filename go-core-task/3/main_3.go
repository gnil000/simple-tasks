package main

import "fmt"

func main() {
	namesAge := NewStringIntMap()

	namesAge.Add("john", 22)
	namesAge.Add("bob", 13)
	namesAge.Add("alex", 56)

	namesAge.Remove("alex")

	secondNamesAge := namesAge.Copy()
	fmt.Println(secondNamesAge)

	fmt.Println(namesAge.Exists("john"))

	fmt.Println(namesAge.Get("bob"))
}

type StringIntMap struct {
	pairs map[string]int
}

func NewStringIntMap() *StringIntMap {
	return &StringIntMap{
		pairs: make(map[string]int),
	}
}

func (m *StringIntMap) Add(key string, value int) {
	m.pairs[key] = value
}

func (m *StringIntMap) Remove(key string) {
	delete(m.pairs, key)
}

func (m *StringIntMap) Copy() map[string]int {
	copyMap := make(map[string]int, len(m.pairs))
	for key, value := range m.pairs {
		copyMap[key] = value
	}
	return copyMap
}

func (m *StringIntMap) Exists(key string) bool {
	_, ok := m.pairs[key]
	return ok
}

func (m *StringIntMap) Get(key string) (int, bool) {
	v, ok := m.pairs[key]
	return v, ok
}
