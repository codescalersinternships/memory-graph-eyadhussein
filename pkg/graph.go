// Package pkg provides a simple graph implementation with the ability to query grandparents, siblings, and cousins of a vertex.
package pkg

// Vertex represents a vertex in a graph
type Vertex struct {
	Name  string
	Edges map[string]*Vertex
}

// NewVertex creates a new vertex with the given name
func NewVertex(name string) *Vertex {
	return &Vertex{
		Name:  name,
		Edges: make(map[string]*Vertex),
	}
}

// Graph represents a graph
type Graph struct {
	vertices map[string]*Vertex
}

// NewGraph creates a new graph
func NewGraph() *Graph {
	return &Graph{
		vertices: make(map[string]*Vertex),
	}
}

// AddVertex adds a vertex to the graph
func (g *Graph) AddVertex(name string) {
	if _, ok := g.vertices[name]; !ok {
		g.vertices[name] = NewVertex(name)
	}
}

// AddEdge adds an edge between two vertices in the graph
func (g *Graph) AddEdge(from, to string) {
	if _, ok := g.vertices[from]; !ok {
		g.AddVertex(from)
	}
	if _, ok := g.vertices[to]; !ok {
		g.AddVertex(to)
	}
	g.vertices[from].Edges[to] = g.vertices[to]
}

// QueryGrandparents returns the grandparents of a vertex
func (g *Graph) QueryGrandparents(name string) []string {
	grandparents := make(map[string]bool)

	if _, ok := g.vertices[name]; !ok {
		return []string{}
	}

	vertex := g.vertices[name]
	for _, parent := range vertex.Edges {
		for _, grandparent := range parent.Edges {
			grandparents[grandparent.Name] = true
		}
	}

	result := make([]string, 0, len(grandparents))
	for grandparent := range grandparents {
		result = append(result, grandparent)
	}
	return result
}

// QuerySiblings returns the siblings of a vertex
func (g *Graph) QuerySiblings(name string) []string {
	siblings := make(map[string]bool)

	if _, ok := g.vertices[name]; !ok {
		return []string{}
	}

	vertex := g.vertices[name]
	for _, parent := range vertex.Edges {
		for _, sibling := range parent.Edges {
			if sibling.Name != name {
				siblings[sibling.Name] = true
			}
		}
	}

	result := make([]string, 0, len(siblings))
	for sibling := range siblings {
		result = append(result, sibling)
	}
	return result
}

// QueryCousins returns the cousins of a vertex
func (g *Graph) QueryCousins(name string) []string {
	cousins := make(map[string]bool)

	if _, ok := g.vertices[name]; !ok {
		return []string{}
	}

	vertex := g.vertices[name]
	for _, parent := range vertex.Edges {
		for _, sibling := range parent.Edges {
			if sibling.Name == name {
				continue
			}
			for _, cousin := range sibling.Edges {
				if cousin.Name != name {
					cousins[cousin.Name] = true
				}
			}

		}
	}

	result := make([]string, 0, len(cousins))
	for cousin := range cousins {
		result = append(result, cousin)
	}
	return result
}

// FilterVertices returns the vertices that satisfy the given filter
func (g *Graph) FilterVertices(filter func(v *Vertex) bool) []*Vertex {

	vertices := make([]*Vertex, 0)

	for _, vertex := range g.vertices {
		if filter(vertex) {
			vertices = append(vertices, vertex)
		}
	}

	return vertices
}
