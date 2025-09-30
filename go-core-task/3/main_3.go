package main

func main() {

}

type keyValue struct {
	key   string
	value int
}

type StringIntMap struct {
	pairs []keyValue
}

func (m *StringIntMap) Add(key string, value int) {

}

func (m *StringIntMap) Remove(key string) {

}

func (m *StringIntMap) Copy() map[string]int {

}

func (m *StringIntMap) Exists(key string) bool {

}

func (m *StringIntMap) Get(key string) (int, bool) {

}
