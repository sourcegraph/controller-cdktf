package workflow


type WorkflowExpressionsOperationsBranchesBranches struct {
	// Groups of prerequisite conditions. All conditions in at least one group must be satisfied.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/workflow#condition_groups Workflow#condition_groups}
	ConditionGroups interface{} `field:"required" json:"conditionGroups" yaml:"conditionGroups"`
	// The result assumed if the condition groups are satisfied.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/workflow#result Workflow#result}
	Result *WorkflowExpressionsOperationsBranchesBranchesResult `field:"required" json:"result" yaml:"result"`
}

