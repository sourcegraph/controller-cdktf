package googledataproccluster


type GoogleDataprocClusterClusterConfigEncryptionConfig struct {
	// The Cloud KMS key name to use for PD disk encryption for all instances in the cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataproc_cluster#kms_key_name GoogleDataprocCluster#kms_key_name}
	KmsKeyName *string `field:"required" json:"kmsKeyName" yaml:"kmsKeyName"`
}

