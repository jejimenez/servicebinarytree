package rest

import (    
    "testing"
    "net/http"
    "strconv"
    "servicebinarytree/pkg/models"
    sampledata "servicebinarytree/pkg/models/sample-data"
    "servicebinarytree/pkg/storage/inmemory"
    //"fmt"

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
        {name: "binarytree already exists", bt: sampledata.Binarytrees["two"], status: http.StatusPreconditionFailed, err: "Binarytree already exists"},
    }

    for _, tc := range testData {
        t.Run(tc.name, func(t *testing.T) {
            // Assetions 
            j, err := json.Marshal(tc.bt)

            if err != nil {
                t.Fatalf("could not created json request: %v", err)
            }

            jsonStr := []byte(j)
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
        val1    int
        val2    int
        status  int
        err     string
    }{
        {name: "lowest ancestor found", treename: "four", val1: 29, val2:44, status: http.StatusAccepted},
    }

    for _, tc := range testData {
        t.Run(tc.name, func(t *testing.T) {
            // Assetions 

            rq, err := http.NewRequest("GET", "/lowestancestor", nil)
            if err != nil {
                t.Fatalf("could not created request: %v", err)
            }
            q := rq.URL.Query()
            q.Add("treename", tc.treename)
            q.Add("value1", strconv.Itoa(tc.val1))
            q.Add("value2", strconv.Itoa(tc.val2))
            rq.URL.RawQuery = q.Encode()

            //fmt.Println(rq.URL.String())

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

// Test data
func binarytreeSample() *models.BinaryTree {
    return &models.BinaryTree{
        Name:"samplebinarytree",
        Root: &models.Node{
            Value: 20,
            Left: &models.Node{
                Value: 10,
                Right: &models.Node{
                    Value: 9,
                },
            },
            Right: &models.Node{
            Value: 25,
            },
        },
    }
}
