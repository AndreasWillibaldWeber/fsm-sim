package vizualization

import (
	"log"

	moore "github.com/AndreasWillibaldWeber/fsm-sim/machines/moore"
	graphviz "github.com/goccy/go-graphviz"
	cgraph "github.com/goccy/go-graphviz/cgraph"
)

const (
	commaSeparator string = ", "
	startString    string = "start"
	startSign      string = "_"
	linebreak      string = "\n"
	acceptSign     string = "1"
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

	// todo: integrate layout, make layout selectable via CLI flag
	//graph.SetLayout(string(graphviz.CIRCO))

	graph.SetRankDir(cgraph.LRRank)

	graph.SetOutputOrder(cgraph.NodesFirst)
	graph.SetOrdering(cgraph.OutOrdering)

	for n1, left := range moore.Transitions {
		for i, n2 := range left {
			m1, err := graph.CreateNode(n1)
			if err != nil {
				return err
			}
			m2, err := graph.CreateNode(n2)
			if err != nil {
				return err
			}
			e, err := graph.CreateEdge(string(i), m1, m2)
			if err != nil {
				return err
			}
			m1.SetShape(cgraph.CircleShape)
			if moore.Accept && moore.Mapping[n1] == acceptSign {
				m1.SetShape(cgraph.DoubleCircleShape)
			}
			m2.SetShape(cgraph.CircleShape)
			if moore.Accept && moore.Mapping[n2] == acceptSign {
				m2.SetShape(cgraph.DoubleCircleShape)
			}
			if !moore.Accept {
				m1.SetLabel(n1 + linebreak + moore.Mapping[n1])
				m2.SetLabel(n2 + linebreak + moore.Mapping[n2])
			}
			e.SetLabel(string(i))
		}
	}

	s1, err := graph.CreateNode(startSign)
	if err != nil {
		return err
	}
	s1.SetShape(cgraph.PointShape)
	s2, err := graph.CreateNode(moore.CurrentState)
	if err != nil {
		return err
	}
	e, err := graph.CreateEdge(startString, s1, s2)
	if err != nil {
		return err
	}
	e.SetLabel(startString)

	if err := g.RenderFilename(graph, format, path); err != nil {
		return err
	}

	return nil
}
