package service

import(
        "fmt"
        "servicebinarytree/pkg/models"
)


// Search in binary tree a value and return boolean, true if found
// Set the array with the path walked
func SearchInNodes(n *models.Node, key int, sl *[]int) bool {
    if n == nil {
        return false
    }
    if key < n.Value {
            *sl  = append(*sl, n.Value)
        return SearchInNodes(n.Left, key, sl)
    }
    if key > n.Value {
            *sl  = append(*sl, n.Value)
        return SearchInNodes(n.Right, key, sl)
    }
    *sl  = append(*sl, n.Value)
    return true
}

// Search in binary tree a value and return the path
func SearchInBinaryTree(n *models.BinaryTree, key int, sl *[]int) bool {
    return SearchInNodes(n.Root, key, sl)
}

// Get the last element of two slices which matches
// First take the last element of the shortest to start the comparison
func GetLastMatchOfSlices(sl1 []int, sl2 []int) (int, error){

    var ln int
    if ln= len(sl2); len(sl1) < len(sl2) {
        ln= len(sl1)
    }

    for i := range make([]int, ln) {
        if sl1[ln-i-1] == sl2[ln-i-1]{
        return sl1[ln-i-1], nil
        }
    }    
    return 0, fmt.Errorf("It's not a tree")
}


// Get the lowest common ancestor in binary tree of v1 and v2 value nodes
func GetLowestCommonAncestor(bt *models.BinaryTree, v1 int, v2 int) (int, error) {

    sl1 := []int{}
    sl2 := []int{}

    // Get the path of val1 in binary tree
    //if not found return the error
    if !SearchInBinaryTree(bt, v1, &sl1) {
        return 0, fmt.Errorf("Element %d not found in %s", v1, bt.Name)
    }

    if !SearchInBinaryTree(bt, v2, &sl2) {
        return 0, fmt.Errorf("Element %d not found in %s", v2, bt.Name)
    }

    return GetLastMatchOfSlices(sl1, sl2)

}

