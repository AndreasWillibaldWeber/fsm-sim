package vizualization

import (
	"log"

	moore "github.com/AndreasWillibaldWeber/fsm-sim/machines/moore"
	graphviz "github.com/goccy/go-graphviz"
)

func DrawMoore(moore *moore.Moore, format graphviz.Format, layout graphviz.Layout, path string) error {
	g := graphviz.New()
	graph, err := g.Graph()
	if err != nil {
		return err
	}
	defer func() {
		if err := graph.Close(); err != nil {
			log.Fatal(err)
		}
		g.Close()
	}()

	graph.SetLayout(string(graphviz.CIRCO))

	for n1, left := range moore.Transitions {
		for i, n2 := range left {
			m1, err := graph.CreateNode(n1 + ", " + moore.Mapping[n1])
			if err != nil {
				return err
			}
			m2, err := graph.CreateNode(n2 + ", " + moore.Mapping[n2])
			if err != nil {
				return err
			}
			e, err := graph.CreateEdge(string(i), m1, m2)
			if err != nil {
				return err
			}
			e.SetLabel(string(i))
		}
	}

	if err := g.RenderFilename(graph, format, path); err != nil {
		return err
	}

	return nil
}
