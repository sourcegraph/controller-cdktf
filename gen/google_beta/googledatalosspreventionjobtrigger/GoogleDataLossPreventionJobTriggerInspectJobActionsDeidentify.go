package googledatalosspreventionjobtrigger


type GoogleDataLossPreventionJobTriggerInspectJobActionsDeidentify struct {
	// User settable Cloud Storage bucket and folders to store de-identified files.
	//
	// This field must be set for cloud storage deidentification.
	//
	// The output Cloud Storage bucket must be different from the input bucket.
	//
	// De-identified files will overwrite files in the output path.
	//
	// Form of: gs://bucket/folder/ or gs://bucket
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_job_trigger#cloud_storage_output GoogleDataLossPreventionJobTrigger#cloud_storage_output}
	CloudStorageOutput *string `field:"required" json:"cloudStorageOutput" yaml:"cloudStorageOutput"`
	// List of user-specified file type groups to transform. If specified, only the files with these filetypes will be transformed.
	//
	// If empty, all supported files will be transformed. Supported types may be automatically added over time.
	//
	// If a file type is set in this field that isn't supported by the Deidentify action then the job will fail and will not be successfully created/started. Possible values: ["IMAGE", "TEXT_FILE", "CSV", "TSV"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_job_trigger#file_types_to_transform GoogleDataLossPreventionJobTrigger#file_types_to_transform}
	FileTypesToTransform *[]*string `field:"optional" json:"fileTypesToTransform" yaml:"fileTypesToTransform"`
	// transformation_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_job_trigger#transformation_config GoogleDataLossPreventionJobTrigger#transformation_config}
	TransformationConfig *GoogleDataLossPreventionJobTriggerInspectJobActionsDeidentifyTransformationConfig `field:"optional" json:"transformationConfig" yaml:"transformationConfig"`
	// transformation_details_storage_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_job_trigger#transformation_details_storage_config GoogleDataLossPreventionJobTrigger#transformation_details_storage_config}
	TransformationDetailsStorageConfig *GoogleDataLossPreventionJobTriggerInspectJobActionsDeidentifyTransformationDetailsStorageConfig `field:"optional" json:"transformationDetailsStorageConfig" yaml:"transformationDetailsStorageConfig"`
}

