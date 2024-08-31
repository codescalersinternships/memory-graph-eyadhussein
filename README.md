# mem-graphdb

## Overview

This is a Go implementation of a straightforward graph database. It provides essential functionalities like adding nodes and edges, as well as querying relationships such as siblings, cousins, and grandparents.

## Features

- AddVertices: Add vertices to the graph.
- AddEdges: Add edge between two vertices.
- QuerySiblings: Retrieve sibling vertices of a given node.
- QueryParents: Retrieve parent vertices of a given node.
- QueryGrandparents: Retrieve grandparent vertices of a given node.
- FilterVertices: Filter vertices based on specific filter func.

## How to Install and Run the Project

1. import it into your Go project:
   ```go
   import graph "github.com/codescalersinternships/memory-graph-eyadhussein/pkg"
   ```

## Usage Example

```go
	// Initialize a new graph
	g := graph.NewGraph()

    // Add Vertices
	g.AddVertex("A")
	g.AddVertex("B")
	g.AddVertex("C")
	g.AddVertex("D")
	g.AddVertex("E")


    // Add Edges
	g.AddEdge("A", "B")
	g.AddEdge("A", "C")
	g.AddEdge("B", "D")
	g.AddEdge("C", "E")
	g.AddEdge("D", "E")


	// Query Siblings
	siblings = g.QuerySiblings("B")
	fmt.Println("Siblings of B:")
	for _, sibling := range siblings {
		fmt.Println(sibling)
	}

    // Query Cousins
	cousins := g.QueryCousins("A")
	fmt.Println("Cousins of A:")
	for _, cousin := range cousins {
		fmt.Println(cousin)
	}

    // Query Grandparents
	grandparents = g.QueryGrandparents("E")
	fmt.Println("Grandparents of E:")
	for _, grandparent := range grandparents {
		fmt.Println(grandparent)
	}

    // Filter Vertices
	filteredVertices := g.FilterVertices(func(v *graph.Vertex) bool {
		return v.Name == "A" || v.Name == "C"
	})

	for _, vertex := range filteredVertices {
		fmt.Println(vertex.Name)
	}

```

## Testing

```sh
make test
```

## Linting

```sh
make lint
```
