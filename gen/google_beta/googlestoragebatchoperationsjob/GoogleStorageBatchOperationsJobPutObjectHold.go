package googlestoragebatchoperationsjob


type GoogleStorageBatchOperationsJobPutObjectHold struct {
	// set/unset to update event based hold for objects.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_storage_batch_operations_job#event_based_hold GoogleStorageBatchOperationsJob#event_based_hold}
	EventBasedHold *string `field:"optional" json:"eventBasedHold" yaml:"eventBasedHold"`
	// set/unset to update temporary based hold for objects.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_storage_batch_operations_job#temporary_hold GoogleStorageBatchOperationsJob#temporary_hold}
	TemporaryHold *string `field:"optional" json:"temporaryHold" yaml:"temporaryHold"`
}

