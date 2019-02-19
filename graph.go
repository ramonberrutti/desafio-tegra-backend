package main

import (
	"sync"
)

// Node each airport will be a node
type Node Airport

// Edge is the flight
type Edge Flight

// Graph is the main
type Graph struct {
	nodes map[string]*Node
	edges map[string]map[string]map[*Edge]bool // Use map to prevent duplicated Edges
	lock  sync.RWMutex
}

// AddNode will add a node to the existing graph
func (g *Graph) AddNode(name string, n *Node) *Node {
	g.lock.Lock()
	if g.nodes == nil {
		g.nodes = make(map[string]*Node)
	}

	g.nodes[name] = n
	g.lock.Unlock()
	return n
}

// AddEdge ...
func (g *Graph) AddEdge(nameFrom, nameTo string, edge *Edge) {
	g.lock.Lock()

	if g.edges == nil {
		g.edges = make(map[string]map[string]map[*Edge]bool)
	}

	if g.edges[nameFrom] == nil {
		g.edges[nameFrom] = make(map[string]map[*Edge]bool)
	}

	if g.edges[nameFrom][nameTo] == nil {
		g.edges[nameFrom][nameTo] = make(map[*Edge]bool)
	}

	g.edges[nameFrom][nameTo][edge] = true

	g.lock.Unlock()
}

// GetNode ....
func (g *Graph) GetNode(name string) *Node {
	g.lock.RLock()
	node, ok := g.nodes[name]
	g.lock.RUnlock()

	if !ok {
		return nil
	}

	return node
}

// GetEdges ...
func (g *Graph) GetEdges(name string) []*Edge {
	edges := make([]*Edge, 0)

	if g.edges == nil || g.edges[name] == nil {
		return edges
	}

	for _, from := range g.edges[name] {
		for edge := range from {
			edges = append(edges, edge)
		}
	}

	return edges
}

// GetEdgesTo ...
func (g *Graph) GetEdgesTo(nameFrom, nameTo string) []*Edge {
	edges := make([]*Edge, 0)

	if g.edges == nil || g.edges[nameFrom] == nil || g.edges[nameFrom][nameTo] == nil {
		return edges
	}

	for edge := range g.edges[nameFrom][nameTo] {
		edges = append(edges, edge)
	}

	return edges
}

// FoundRoute ...
func (g *Graph) FoundRoute(nodeFrom, nodeTo string, initialFilter func(nodeFrom string, edge Edge) bool, routeFilter func(nodeFrom string, edgeFrom, edge Edge) bool) [][]Edge {
	routes := make([][]Edge, 0)

	for _, edge := range g.GetEdges(nodeFrom) {
		if initialFilter(nodeFrom, *edge) {
			list := make([]Edge, 0)
			g.foundRoute(edge.Destination, nodeTo, map[string]bool{}, append(list, *edge), &routes, routeFilter)
		}
	}

	return routes
}

func (g *Graph) foundRoute(from, to string, visited map[string]bool, list []Edge, globalList *[][]Edge, routeFilter func(nodeFrom string, edgeFrom, edge Edge) bool) {
	if from == to {
		*globalList = append(*globalList, list)
	}

	if visited[from] == true {
		return
	}
	visited[from] = true

	for _, edge := range g.GetEdges(from) {
		if routeFilter(from, list[len(list)-1], *edge) {
			g.foundRoute(edge.Destination, to, visited, append(list, *edge), globalList, routeFilter)
		}

	}
}

/*
   // Prints all paths from
   // 's' to 'd'
   public void printAllPaths(int s, int d)
   {
       boolean[] isVisited = new boolean[v];
       ArrayList pathList = new ArrayList<>();

       //add source to path[]
       pathList.add(s);

       //Call recursive utility
       printAllPathsUtil(s, d, isVisited, pathList);
   }

   // A recursive function to print
   // all paths from 'u' to 'd'.
   // isVisited[] keeps track of
   // vertices in current path.
   // localPathList<> stores actual
   // vertices in the current path
   private void printAllPathsUtil(Integer u, Integer d,
                                   boolean[] isVisited,
                           List localPathList) {

       // Mark the current node
       isVisited[u] = true;

       if (u.equals(d))
       {
           System.out.println(localPathList);
       }

       // Recur for all the vertices
       // adjacent to current vertex
       for (Integer i : adjList[u])
       {
           if (!isVisited[i])
           {
               // store current node
               // in path[]
               localPathList.add(i);
               printAllPathsUtil(i, d, isVisited, localPathList);

               // remove current node
               // in path[]
               localPathList.remove(i);
           }
       }

       // Mark the current node
       isVisited[u] = false;
   }
*/
