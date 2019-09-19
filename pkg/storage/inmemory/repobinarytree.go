package inmemory

import (
    "sync"
    binarytree "servicebinarytree/pkg"
    "servicebinarytree/pkg/models"
)


type binarytreeRepository struct {
    mtx     sync.RWMutex
    binarytrees map[string]*models.BinaryTree
}



func InitBinarytreeRepository(binarytrees map[string]*models.BinaryTree) binarytree.BinarytreeRepository {
    if binarytrees == nil {
        binarytrees = make(map[string]*models.BinaryTree)
    }
    return &binarytreeRepository{
        binarytrees: binarytrees,
    }
}


func (r *binarytreeRepository) CreateBinarytree(g *models.BinaryTree) error {
    return nil
}


