package googlestoragebatchoperationsjob


type GoogleStorageBatchOperationsJobDeleteObject struct {
	// enable flag to permanently delete object and all object versions if versioning is enabled on bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_storage_batch_operations_job#permanent_object_deletion_enabled GoogleStorageBatchOperationsJob#permanent_object_deletion_enabled}
	PermanentObjectDeletionEnabled interface{} `field:"required" json:"permanentObjectDeletionEnabled" yaml:"permanentObjectDeletionEnabled"`
}

