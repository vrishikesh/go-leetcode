package leetcode

import "fmt"

type Monarchy interface {
	birth(child, parent string)
	death(name string)
	getOrderOfSuccession() []string
}

type monarchNode struct {
	name     string
	children []*monarchNode
	isDead   bool
}

type monarch struct {
	head    *monarchNode
	persons map[string]*monarchNode
}

func NewMonarch(name string) *monarch {
	person := &monarchNode{name: name}
	persons := map[string]*monarchNode{}
	persons[name] = person

	return &monarch{
		head:    person,
		persons: persons,
	}
}
func (m *monarch) birth(child, parent string) {
	parentNode := m.persons[parent]
	childNode := &monarchNode{name: child}
	parentNode.children = append(parentNode.children, childNode)
	m.persons[child] = childNode
}
func (m *monarch) death(name string) {
	node := m.persons[name]
	node.isDead = true
}
func (m *monarch) getOrderOfSuccession() []string {
	ans := []string{}
	var dfs func(node *monarchNode)
	dfs = func(node *monarchNode) {
		if !node.isDead {
			ans = append(ans, node.name)
		}
		for _, child := range node.children {
			dfs(child)
		}
	}
	dfs(m.head)
	return ans
}

func MonarchySolver() {
	m := NewMonarch("Jake")
	m.birth("Catherine", "Jake")
	m.birth("Tom", "Jake")
	m.birth("Celine", "Jake")
	m.birth("Jane", "Catherine")
	m.birth("Mark", "Catherine")
	m.birth("Farah", "Jane")
	m.birth("Peter", "Celine")
	fmt.Println(m.getOrderOfSuccession())

	m.death("Jake")
	m.death("Jane")
	fmt.Println(m.getOrderOfSuccession())
}
