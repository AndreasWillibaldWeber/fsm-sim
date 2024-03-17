package vizualization

import (
	"log"

	moore "github.com/AndreasWillibaldWeber/fsm-sim/machines/moore"
	graphviz "github.com/goccy/go-graphviz"
	"github.com/goccy/go-graphviz/cgraph"
)

const (
	commaSeparator string = ", "
	startString    string = "start"
	startSign      string = "_"
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

	// todo: check if sorting transitions makes the output repeatable for LRRank, make rankDir selectable via CLI flag
	graph.SetRankDir(cgraph.LRRank)

	for n1, left := range moore.Transitions {
		for i, n2 := range left {
			m1, err := graph.CreateNode(n1 + commaSeparator + moore.Mapping[n1])
			if err != nil {
				return err
			}
			m2, err := graph.CreateNode(n2 + commaSeparator + moore.Mapping[n2])
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

	s1, err := graph.CreateNode(startSign)
	if err != nil {
		return err
	}
	s1.SetShape(cgraph.PointShape)
	s2, err := graph.CreateNode(moore.CurrentState + commaSeparator + moore.Mapping[moore.CurrentState])
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
