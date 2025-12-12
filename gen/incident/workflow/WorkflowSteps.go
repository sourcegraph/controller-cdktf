package workflow


type WorkflowSteps struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/workflow#id Workflow#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/workflow#name Workflow#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Bindings for the operation parameters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/workflow#param_bindings Workflow#param_bindings}
	ParamBindings interface{} `field:"required" json:"paramBindings" yaml:"paramBindings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/workflow#for_each Workflow#for_each}.
	ForEach *string `field:"optional" json:"forEach" yaml:"forEach"`
}

