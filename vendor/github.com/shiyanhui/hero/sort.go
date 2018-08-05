package hero

// sort implements topological sort.
type sort struct {
	vertices map[string]struct{}
	graph    map[string]map[string]struct{}
	v        map[string]int
}

func newSort() *sort {
	return &sort{
		vertices: make(map[string]struct{}),
		graph:    make(map[string]map[string]struct{}),
		v:        make(map[string]int),
	}
}

func (s *sort) addVertices(vertices ...string) {
	for _, vertex := range vertices {
		s.vertices[vertex] = struct{}{}
	}
}

func (s *sort) addEdge(from, to string) {
	if _, ok := s.graph[from]; !ok {
		s.graph[from] = map[string]struct{}{
			to: {},
		}
	} else {
		s.graph[from][to] = struct{}{}
	}
}

func (s *sort) collect(queue *[]string) {
	for vertex, count := range s.v {
		if count == 0 {
			*queue = append(*queue, vertex)
			delete(s.v, vertex)
		}
	}
}

// sort returns sorted string list.
func (s *sort) sort() []string {
	for vertex := range s.vertices {
		s.v[vertex] = 0
	}

	for _, tos := range s.graph {
		for to := range tos {
			s.v[to]++
		}
	}

	r := make([]string, 0, len(s.vertices))

	queue := make([]string, 0, len(s.vertices))
	s.collect(&queue)

	for len(queue) > 0 {
		r = append(r, queue[0])
		if tos, ok := s.graph[queue[0]]; ok {
			for to := range tos {
				s.v[to]--
			}
		}

		s.collect(&queue)

		delete(s.graph, queue[0])
		queue = queue[1:]
	}

	if len(s.v) > 0 {
		panic("import or include cycle")
	}

	return r
}
