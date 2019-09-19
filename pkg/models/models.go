package models

type Node struct {
    Value int  `json:"value"`
    Left  *Node `json:"left"`
    Right *Node `json:"right"`
}

type BinaryTree struct {
    Name string `json:"name"`
    Root *Node `json:"root"`
}
