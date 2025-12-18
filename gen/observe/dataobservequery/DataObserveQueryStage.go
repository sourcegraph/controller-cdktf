package dataobservequery


type DataObserveQueryStage struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/data-sources/query#alias DataObserveQuery#alias}.
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/data-sources/query#input DataObserveQuery#input}.
	Input *string `field:"optional" json:"input" yaml:"input"`
	// A boolean flag used to specify the output stage.
	//
	// Should be used only for
	// a stage preceding the last stage. The last stage is an output stage by default.
	//
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/data-sources/query#output_stage DataObserveQuery#output_stage}
	OutputStage interface{} `field:"optional" json:"outputStage" yaml:"outputStage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/data-sources/query#pipeline DataObserveQuery#pipeline}.
	Pipeline *string `field:"optional" json:"pipeline" yaml:"pipeline"`
}

