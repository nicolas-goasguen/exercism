package dotdsl

import (
	"errors"
	"maps"
	"slices"
	"strconv"
	"strings"
	"unicode"
)

// Properties holds the properties of a node or edge.
// The values can be int, bool or string.
type Properties map[string]any

// Graph stores the parts of a dot graph.
// All entities are stored as a Properties map (`nil` Properties when none set)
// attrs is the Properties for the entire Graph, vs a specific node or edge.
type Graph struct {
	nodes map[string]Properties
	edges map[string]Properties
	attrs Properties
}

func NewGraph() *Graph {
	return &Graph{
		nodes: make(map[string]Properties),
		edges: make(map[string]Properties),
		attrs: make(Properties),
	}
}

const (
	graphKeyword = "graph"
	graphOpen    = "graph {"
	graphClose   = "}"

	propOpen  = "["
	propClose = "]"

	strValOpen  = "\""
	strValClose = "\""
)

var (
	ErrInvalidGraph         = errors.New("invalid graph")
	ErrInvalidLine          = errors.New("invalid line")
	ErrInvalidProperty      = errors.New("invalid property")
	ErrInvalidPropertyValue = errors.New("invalid property value")
	ErrInvalidNode          = errors.New("invalid node")
	ErrInvalidEdge          = errors.New("invalid edge")

	isLineSep     = func(r rune) bool { return r == ';' || r == '\n' }
	isPropertySep = func(r rune) bool { return r == '=' }
)

func isOnlyLetters(s string) bool {
	return len(s) > 0 && !strings.ContainsFunc(s, func(r rune) bool {
		return !unicode.IsLetter(r)
	})
}

func trim(s, prefix, suffix string, required bool) (string, error) {
	hasBounds := strings.HasPrefix(s, prefix) && strings.HasSuffix(s, suffix)

	if required && !hasBounds {
		return "", ErrInvalidLine
	}

	if hasBounds {
		s = s[len(prefix) : len(s)-len(suffix)]
	}

	return strings.TrimSpace(s), nil
}

func parseValue(s string) (any, error) {
	s, err := trim(s, strValOpen, strValClose, false)
	if err != nil {
		return nil, err
	}

	if len(s) == 0 {
		return nil, ErrInvalidPropertyValue
	}

	if i, err := strconv.Atoi(s); err == nil {
		return i, nil
	}
	if f, err := strconv.ParseFloat(s, 64); err == nil {
		return f, nil
	}
	if b, err := strconv.ParseBool(s); err == nil {
		return b, nil
	}
	return s, nil
}

func parseProperties(line string) (Properties, error) {
	cleaned, err := trim(line, propOpen, propClose, true)
	if err != nil {
		return nil, ErrInvalidProperty
	}

	lineFields := strings.FieldsFunc(cleaned, isPropertySep)

	if len(lineFields) != 2 || !isOnlyLetters(lineFields[0]) {
		return nil, ErrInvalidProperty
	}

	value, err := parseValue(lineFields[1])
	if err != nil {
		return nil, err
	}

	prop := make(Properties)
	prop[lineFields[0]] = value
	return prop, nil
}

func parseNode(line string) (string, Properties, error) {
	// node validation
	lineFields := strings.Fields(line)
	if len(lineFields) == 0 || len(lineFields) > 2 || !isOnlyLetters(lineFields[0]) {
		return "", nil, ErrInvalidNode
	}

	if len(lineFields) == 1 {
		return lineFields[0], nil, nil
	}

	// properties validation
	props, err := parseProperties(lineFields[1])
	if err != nil {
		return "", nil, err
	}
	return lineFields[0], props, nil
}

func parseEdge(line string) ([]string, map[string]Properties, error) {
	lineFields := strings.Fields(line)
	if len(lineFields) == 0 || len(lineFields) < 3 {
		return nil, nil, ErrInvalidEdge
	}

	propField := lineFields[len(lineFields)-1]
	prop, err := parseProperties(propField)
	if err == nil {
		line = line[:len(line)-len(propField)]
	}

	edgeNodes := strings.Split(line, "--")
	if len(edgeNodes) < 2 {
		return nil, nil, ErrInvalidEdge
	}

	nodes := make([]string, 0, len(edgeNodes))
	for i := range edgeNodes {
		edgeNodes[i] = strings.TrimSpace(edgeNodes[i])
		if !isOnlyLetters(edgeNodes[i]) {
			return nil, nil, ErrInvalidEdge
		}
		nodes = append(nodes, edgeNodes[i])
	}

	edges := make(map[string]Properties)
	for i := range len(edgeNodes) - 1 {
		pair := []string{edgeNodes[i], edgeNodes[i+1]}
		slices.Sort(pair)
		key := "{" + pair[0] + " " + pair[1] + "}"
		edges[key] = maps.Clone(prop)
	}

	return nodes, edges, nil
}

func parseComment(line string) bool {
	return strings.HasPrefix(line, "//") || strings.HasPrefix(line, "#")
}

func parseGraph(data string) (*Graph, error) {
	// "graph" validation
	cleaned, err := trim(data, graphOpen, graphClose, true)
	if err != nil {
		return nil, err
	}

	// "graph" data extraction
	graphLines := strings.FieldsFunc(cleaned, isLineSep)

	graph := NewGraph()

	for _, l := range graphLines {
		l := strings.TrimSpace(l)
		if ok := parseComment(l); ok {
			continue
		} else if props, err := parseProperties(l); err == nil {
			maps.Copy(graph.attrs, props)
			continue
		} else if node, props, err := parseNode(l); err == nil {
			graph.nodes[node] = props
			continue
		} else if nodes, edges, err := parseEdge(l); err == nil {
			for _, n := range nodes {
				if _, ok := graph.nodes[n]; !ok {
					graph.nodes[n] = make(Properties)
				}
			}
			maps.Copy(graph.edges, edges)
			continue
		} else {
			return nil, ErrInvalidGraph
		}
	}
	return graph, nil
}

// Parse creates a Graph from a text blob.
func Parse(data string) (*Graph, error) {
	graph, err := parseGraph(data)
	if err != nil {
		return nil, err
	}
	return graph, nil
}
