package googledatalosspreventionjobtrigger


type GoogleDataLossPreventionJobTriggerInspectJobStorageConfigDatastoreOptions struct {
	// kind block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_data_loss_prevention_job_trigger#kind GoogleDataLossPreventionJobTrigger#kind}
	Kind *GoogleDataLossPreventionJobTriggerInspectJobStorageConfigDatastoreOptionsKind `field:"required" json:"kind" yaml:"kind"`
	// partition_id block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_data_loss_prevention_job_trigger#partition_id GoogleDataLossPreventionJobTrigger#partition_id}
	PartitionId *GoogleDataLossPreventionJobTriggerInspectJobStorageConfigDatastoreOptionsPartitionId `field:"required" json:"partitionId" yaml:"partitionId"`
}

