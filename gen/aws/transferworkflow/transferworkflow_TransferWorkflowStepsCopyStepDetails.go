package transferworkflow


type TransferWorkflowStepsCopyStepDetails struct {
	// destination_file_location block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/transfer_workflow#destination_file_location TransferWorkflow#destination_file_location}
	DestinationFileLocation *TransferWorkflowStepsCopyStepDetailsDestinationFileLocation `field:"optional" json:"destinationFileLocation" yaml:"destinationFileLocation"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/transfer_workflow#name TransferWorkflow#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/transfer_workflow#overwrite_existing TransferWorkflow#overwrite_existing}.
	OverwriteExisting *string `field:"optional" json:"overwriteExisting" yaml:"overwriteExisting"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/transfer_workflow#source_file_location TransferWorkflow#source_file_location}.
	SourceFileLocation *string `field:"optional" json:"sourceFileLocation" yaml:"sourceFileLocation"`
}

