package googledataprocworkflowtemplate


type GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigEncryptionConfig struct {
	// Optional. The Cloud KMS key name to use for PD disk encryption for all instances in the cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_workflow_template#gce_pd_kms_key_name GoogleDataprocWorkflowTemplate#gce_pd_kms_key_name}
	GcePdKmsKeyName *string `field:"optional" json:"gcePdKmsKeyName" yaml:"gcePdKmsKeyName"`
}

