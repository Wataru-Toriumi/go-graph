# go-graph

This is the package for manipulating graphs.

## How to use

### install

```bash
go get github.com/Wataru-Toriumi/go-graph
```

### use

```go
package main

import "github.com/Wataru-Toriumi/go-graph/src/graph"

func main(){
    var nodes []string = []string{"1","2","3"}
    var edges [][]string = [][]string{{"1","2"},{"1","3"}}
    g := graph.New(nodes, edges)
    graph.Show(g)
}
```

## Test

```bash
go test -v src/graph/graph.go src/graph/graph_test.go
```
