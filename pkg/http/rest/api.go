package rest

import (
    "net/http"

    servicebinarytree "servicebinarytree/pkg"

    "github.com/gorilla/mux"
)

func (a *api) Handler() http.Handler {
    return a.handler
}

type api struct {
    handler     http.Handler
    repositorybt servicebinarytree.BinarytreeRepository
}

type Rest interface {
    Handler() http.Handler
    CreateBinarytree(w http.ResponseWriter, r *http.Request)
    LowestAncestor(w http.ResponseWriter, r *http.Request)
}

// Config the end points 
func New(repobt servicebinarytree.BinarytreeRepository) Rest {
    a := &api{repositorybt: repobt}

    r := mux.NewRouter()
    r.HandleFunc("/tree", a.CreateBinarytree).Methods(http.MethodPost)
    r.HandleFunc("/lowestancestor", a.LowestAncestor).Methods(http.MethodGet)

    a.handler = r
    return a
}


func (a *api) CreateBinarytree(w http.ResponseWriter, r *http.Request) {

}


func (a *api) LowestAncestor(w http.ResponseWriter, r *http.Request){
    
}

