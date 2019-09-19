package models

type Node struct {
    Value int
	Left  *Node `json:"node"`
    Right *Node `json:"node"`
}

type BinaryTree struct {
	Name string `json:"string"`
	Root Node `json:"node"`
}
