package yaml

type BaseStructure struct {
	Version int
	Input   map[string]Input `yaml:"input"`
	Tests   []Test
}

type Input struct {
	OutputsFile string `yaml:"outputs_file"`
}

type Test struct {
	Name    string
	Type    string
	Params  map[string]interface{}
	Expects map[string]interface{}
}
