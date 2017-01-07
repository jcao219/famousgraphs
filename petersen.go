package famousgraphs

import "github.com/gonum/graph/simple"

// PetersenGraph returns the Petersen graph of 10 vertices and 15 edges.
func PetersenGraph() *simple.UndirectedGraph {
	g := simple.NewUndirectedGraph(0.0, 0.0)
	for i := 2; i <= 10; i++ {
		g.SetEdge(simple.Edge{F: simple.Node(i), T: simple.Node(i - 1), W: 0.0})
	}
	g.SetEdge(simple.Edge{F: simple.Node(2), T: simple.Node(10), W: 0.0})
	g.SetEdge(simple.Edge{F: simple.Node(3), T: simple.Node(7), W: 0.0})
	g.SetEdge(simple.Edge{F: simple.Node(4), T: simple.Node(9), W: 0.0})
	g.SetEdge(simple.Edge{F: simple.Node(6), T: simple.Node(10), W: 0.0})
	g.SetEdge(simple.Edge{F: simple.Node(1), T: simple.Node(5), W: 0.0})
	g.SetEdge(simple.Edge{F: simple.Node(1), T: simple.Node(8), W: 0.0})
	return g
}
