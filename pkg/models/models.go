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

type LowestAncestorResp struct{
    Treename      string `json:"treename"`
    Value1          int    `json:"value1"`
    Value2          int  	 `json:"value2"`
    Ancestor      int	 `json:"ancestor"`
}
