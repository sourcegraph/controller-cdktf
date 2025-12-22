package dataset


type DatasetStage struct {
	// The stage alias is the label by which subsequent stages can refer to the results of this stage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/dataset#alias Dataset#alias}
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// The stage input defines what input should be used as a starting point for the stage pipeline.
	//
	// It must refer to a label contained in `inputs`, or a
	// previous stage `alias`. The stage input can be omitted if `inputs`
	// contains a single element.
	//
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/dataset#input Dataset#input}
	Input *string `field:"optional" json:"input" yaml:"input"`
	// A boolean flag used to specify the output stage.
	//
	// Should be used only for
	// a stage preceding the last stage. The last stage is an output stage by default.
	//
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/dataset#output_stage Dataset#output_stage}
	OutputStage interface{} `field:"optional" json:"outputStage" yaml:"outputStage"`
	// An OPAL snippet defining a transformation on the selected input.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/dataset#pipeline Dataset#pipeline}
	Pipeline *string `field:"optional" json:"pipeline" yaml:"pipeline"`
}

