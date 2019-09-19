package servicebinarytree

import(
	"servicebinarytree/pkg/models"
)

// Binary tree storage
type BinarytreeRepository interface {
	CreateBinarytree(g *models.BinaryTree) error
}
