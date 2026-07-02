package workflow

type WorkflowDelay struct {
	// If this workflow is delayed, whether the conditions should be rechecked between trigger firing and execution.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/workflow#conditions_apply_over_delay Workflow#conditions_apply_over_delay}
	ConditionsApplyOverDelay any `field:"required" json:"conditionsApplyOverDelay" yaml:"conditionsApplyOverDelay"`
	// Delay in seconds between trigger firing and running the workflow.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/workflow#for_seconds Workflow#for_seconds}
	ForSeconds *float64 `field:"required" json:"forSeconds" yaml:"forSeconds"`
}
