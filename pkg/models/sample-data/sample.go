package sampledata

import (
    "servicebinarytree/pkg/models"
)


var Binarytrees = map[string]*models.BinaryTree{
    "one": &models.BinaryTree{
        Name:"one",
        Root: &models.Node{
            Value: 20,
            Left: &models.Node{
                Value: 10,
                Left:  nil,
                Right: &models.Node{
                    Value: 9,
                    Left:  nil,
                    Right: nil,
                },
            },
            Right: &models.Node{
                Value: 25,
                Left:  nil,
                Right: nil,
            },
        },
    },
    "two": &models.BinaryTree{
        Name:"two",
        Root: &models.Node{
            Value: 200,
            Left: &models.Node{
                Value: 100,
                Left:  nil,
                Right: &models.Node{
                    Value: 90,
                    Left:  nil,
                    Right: nil,
                },
            },
            Right: &models.Node{
                Value: 250,
                Left:  nil,
                Right: nil,
            },
        },
    },
    "three": &models.BinaryTree{
        Name:"three",
        Root: &models.Node{
            Value: 2,
            Left: &models.Node{
                Value: 1,
                //Left:  nil,
                Right: &models.Node{
                    Value: 9,
                    //Left:  nil,
                    //Right: nil,
                },
            },
            Right: &models.Node{
                Value: 3,
                //Left:  nil,
                //Right: nil,
            },
        },
    },
    "four": &models.BinaryTree{
            Name:"four",
            Root: &models.Node{
                Value: 67,
                Left: &models.Node{
                    Value: 39,
                    Left: &models.Node{
                        Value: 28,
                        Right: &models.Node{
                            Value: 29,
                        },
                    },
                    Right: &models.Node{
                        Value: 44,
                    },
                },
                Right: &models.Node{
                Value: 76,
                Left: &models.Node{
                    Value: 74,
                },
                Right: &models.Node{
                    Value: 85,
                    Left: &models.Node{
                        Value: 83,
                    },
                    Right: &models.Node{
                        Value: 87,
                    },
                },
            },
        },
    },
}
var LowestAncResps = map[string]*models.LowestAncestorResp {
    "lowest ancestor found" : &models.LowestAncestorResp{
    Treename      : "four",
    Value1        : 29,
    Value2        : 44,
    Ancestor      : 39,
    },
    "node not found" : &models.LowestAncestorResp{
    Treename      : "four",
    Value1        : 100,
    Value2        : 44,
    },
    "value 1 not in params" : &models.LowestAncestorResp{
    Treename      : "four",
    Value2        : 44,
    },
    "value 1 equals to value 2" : &models.LowestAncestorResp{
    Treename      : "four",
    Value1        : 100,
    Value2        : 100,
    },
}

