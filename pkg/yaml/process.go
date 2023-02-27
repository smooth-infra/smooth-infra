package yaml

type BaseStructure struct {
	Version int
	Input   struct {
		Terraform struct {
			OutputsFile string `yaml:"outputs_file"`
		}
	}
	Tests []Test
}

type Test struct {
	Name    string
	Type    string
	Params  map[string]interface{}
	Expects map[string]interface{}
}