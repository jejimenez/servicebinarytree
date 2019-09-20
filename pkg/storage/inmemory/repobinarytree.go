package inmemory

import (
    "sync"
    "fmt"
    //"encoding/json"
    binarytree "servicebinarytree/pkg"
    "servicebinarytree/pkg/models"
)

type binarytreeRepository struct {
    mtx     sync.RWMutex
    binarytrees map[string]*models.BinaryTree
}


// Init the binary tree storage with the map recieved
func InitBinarytreeRepository(binarytrees map[string]*models.BinaryTree) binarytree.BinarytreeRepository {
    if binarytrees == nil {
        binarytrees = make(map[string]*models.BinaryTree)
    }
    return &binarytreeRepository{
        binarytrees: binarytrees,
    }
}

// Check if the binary tree named like Name already exists
func (r *binarytreeRepository) CheckIfExists(Name string) error {
    for _, v := range r.binarytrees {
        if v.Name == Name {
            return fmt.Errorf("The binarytree %s already exist", Name)
        }
    }
    return nil
}

// Create the binarytree element in the memory storage map
func (r *binarytreeRepository) CreateBinarytree(g *models.BinaryTree) error {

    r.mtx.Lock()
    defer r.mtx.Unlock()
    //s, _ := json.MarshalIndent(g, "", "\t")
    //fmt.Println(string(s))

    r.binarytrees[g.Name] = g
    return nil

}

// Get all binary trees
func (r *binarytreeRepository) GetAllBinarytrees() ([]*models.BinaryTree, error) {
    r.mtx.Lock()
    defer r.mtx.Unlock()
    bts := make([]*models.BinaryTree, 0, len(r.binarytrees))
    for _, bt := range r.binarytrees {
        bts = append(bts, bt)
    }
    return bts, nil
}

// Get binary tree by name
func (r *binarytreeRepository) GetBinarytreeByName(n string) (*models.BinaryTree, error) {
    r.mtx.Lock()
    defer r.mtx.Unlock()

    for _, v := range r.binarytrees {
        if v.Name == n {
            return v, nil
        }
    }
    return nil, fmt.Errorf("The Binarytree %s doesn't exist", n)
}

