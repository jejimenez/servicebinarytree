package service

import(
	"servicebinarytree/pkg/models"
)

func getPath(bt *models.Node, a [] ){

	//func Walk(t *tree.Tree, ch chan int) {
 if t == nil {
  return
 } else if t.Left == nil {
  ch <- t.Value
  if t.Right != nil {
   Walk(t.Right, ch)
  }
  return
 } else {
  Walk(t.Left, ch)
 }
 ch <- t.Value
 if t.Right != nil {
  Walk(t.Right, ch)
 }

}