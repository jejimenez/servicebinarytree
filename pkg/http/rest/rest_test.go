package rest

import (    
    "testing"
    "net/http"
    //"strconv"
    "io/ioutil"
    "servicebinarytree/pkg/models"
    sampledata "servicebinarytree/pkg/models/sample-data"
    "servicebinarytree/pkg/storage/inmemory"
    "fmt"

    "encoding/json"
    "net/http/httptest"
    "bytes"
)

type GetLowestCommonAncestorReq struct {
    binarytree *models.BinaryTree `json:"binarytree"`
    value1 int `json:"value1"`
    value2 int `json:"value2"`
}


// Test the tree creation
func TestHandler_CreateTree(t *testing.T) {

    // Table tests input and outputs
    testData := []struct {
        name   string
        bt      *models.BinaryTree
        status int
        err    string
    }{
        {name: "binarytree created", bt: binarytreeSample(), status: http.StatusCreated},
        {name: "binarytree already exists", bt: sampledata.Binarytrees["two"], status: http.StatusPreconditionFailed},
        {name: "binarytree already exists", bt: nil, status: http.StatusInternalServerError},
    }

    for _, tc := range testData {
        t.Run(tc.name, func(t *testing.T) {
            // Assetions 
            j, err := json.Marshal(tc.bt)

            if err != nil {
                t.Fatalf("could not created json request: %v", err)
            }

            jsonStr := []byte(j)
            //fmt.Println(bytes.NewBuffer(jsonStr))
            rq, err := http.NewRequest("POST", "/tree", bytes.NewBuffer(jsonStr))
            if err != nil {
                t.Fatalf("could not created request: %v", err)
            }

            repo := inmemory.InitBinarytreeRepository(sampledata.Binarytrees)
            s := New(repo)

            rc := httptest.NewRecorder()
            s.Handler().ServeHTTP(rc, rq)

            res := rc.Result()
            defer res.Body.Close()
            // if the status is not the expected
            if tc.status != res.StatusCode {
                t.Errorf("expected %d, got: %d", tc.status, res.StatusCode)
            }

        })
    }

}

// Test the function to get the lowest common ancestor
func TestHandler_GetLowestCommonAncestor(t *testing.T){

    // Table tests input and outputs
    testData := []struct {
        name    string
        treename      string
        val1    string
        val2    string
        ancestor int
        status  int
        err     string
    }{
        {name: "lowest ancestor found", treename: "four", val1: "29", val2:"44", ancestor:39, status: http.StatusFound},
        {name: "tree not found", treename: "zero", val1: "100", val2:"44", status: http.StatusNotFound, err:"StatusNotFound"},
        {name: "node not found", treename: "four", val1: "100", val2:"44", status: http.StatusNotFound, err:"StatusNotFound"},
        {name: "value 1 equals to value 2", treename: "four", val1: "12", val2:"12", status: http.StatusUnprocessableEntity, err:"StatusUnprocessableEntity"},
    }

    for _, tc := range testData {
        t.Run(tc.name, func(t *testing.T) {
            // Assetions 
            /*
            rq, err := http.NewRequest("GET", "/lowestancestor", nil)
            if err != nil {
                t.Fatalf("could not created request: %v", err)
            }
            q := rq.URL.Query()
            q.Add("treename", tc.treename)
            q.Add("value1", tc.val1)
            q.Add("value2", tc.val2)
            rq.URL.RawQuery = q.Encode()
            */

            uri := fmt.Sprintf("/lowestancestor/%s/%s/%s", tc.treename, tc.val1, tc.val2)
            rq, err := http.NewRequest("GET", uri, nil)
            if err != nil {
                t.Fatalf("could not created request: %v", err)
            }

            //fmt.Println("____")//rq.URL.String())

            repo := inmemory.InitBinarytreeRepository(sampledata.Binarytrees)
            s := New(repo)

            rc := httptest.NewRecorder()
            s.Handler().ServeHTTP(rc, rq)

            res := rc.Result()
            defer res.Body.Close()
            // if the status is not the expected
            if tc.status != res.StatusCode {
                t.Errorf("expected %d, got: %d", tc.status, res.StatusCode)
            }

            b, err := ioutil.ReadAll(res.Body)
            if err != nil {
                t.Fatalf("could not read response: %v", err)
            }
            // if is the expected ancestor
            if tc.err == "" {
                var lar *models.LowestAncestorResp
                err = json.Unmarshal(b, &lar)
                if err != nil {
                    t.Fatalf("could not unmarshall response %v", err)
                }

                if lar.Ancestor != tc.ancestor {
                    t.Fatalf("expected %v, got: %v", tc.ancestor, lar.Ancestor)
                }
            }

        })
    }

}

// Test data
func binarytreeSample() *models.BinaryTree {
    return &models.BinaryTree{
        Name:"samplebinarytree",
        Root: &models.Node{
            Value: 20,
            Left: &models.Node{
                Value: 10,
                Right: &models.Node{
                    Value: 15,
                },
            },
            Right: &models.Node{
            Value: 25,
            },
        },
    }
}
