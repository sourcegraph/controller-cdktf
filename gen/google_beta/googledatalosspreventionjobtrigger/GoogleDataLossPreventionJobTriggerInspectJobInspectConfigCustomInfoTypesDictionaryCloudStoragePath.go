package googledatalosspreventionjobtrigger


type GoogleDataLossPreventionJobTriggerInspectJobInspectConfigCustomInfoTypesDictionaryCloudStoragePath struct {
	// A url representing a file or path (no wildcards) in Cloud Storage. Example: 'gs://[BUCKET_NAME]/dictionary.txt'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_job_trigger#path GoogleDataLossPreventionJobTrigger#path}
	Path *string `field:"required" json:"path" yaml:"path"`
}

