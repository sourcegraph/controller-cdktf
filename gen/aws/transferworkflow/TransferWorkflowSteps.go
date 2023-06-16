package transferworkflow


type TransferWorkflowSteps struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/transfer_workflow#type TransferWorkflow#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// copy_step_details block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/transfer_workflow#copy_step_details TransferWorkflow#copy_step_details}
	CopyStepDetails *TransferWorkflowStepsCopyStepDetails `field:"optional" json:"copyStepDetails" yaml:"copyStepDetails"`
	// custom_step_details block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/transfer_workflow#custom_step_details TransferWorkflow#custom_step_details}
	CustomStepDetails *TransferWorkflowStepsCustomStepDetails `field:"optional" json:"customStepDetails" yaml:"customStepDetails"`
	// delete_step_details block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/transfer_workflow#delete_step_details TransferWorkflow#delete_step_details}
	DeleteStepDetails *TransferWorkflowStepsDeleteStepDetails `field:"optional" json:"deleteStepDetails" yaml:"deleteStepDetails"`
	// tag_step_details block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/transfer_workflow#tag_step_details TransferWorkflow#tag_step_details}
	TagStepDetails *TransferWorkflowStepsTagStepDetails `field:"optional" json:"tagStepDetails" yaml:"tagStepDetails"`
}

