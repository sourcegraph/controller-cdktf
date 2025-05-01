package googledataprocworkflowtemplate


type GoogleDataprocWorkflowTemplateEncryptionConfig struct {
	// Optional. The Cloud KMS key name to use for encryption.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_dataproc_workflow_template#kms_key GoogleDataprocWorkflowTemplate#kms_key}
	KmsKey *string `field:"optional" json:"kmsKey" yaml:"kmsKey"`
}

