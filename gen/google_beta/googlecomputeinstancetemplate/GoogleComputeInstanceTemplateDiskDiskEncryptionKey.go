package googlecomputeinstancetemplate


type GoogleComputeInstanceTemplateDiskDiskEncryptionKey struct {
	// The self link of the encryption key that is stored in Google Cloud KMS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_instance_template#kms_key_self_link GoogleComputeInstanceTemplate#kms_key_self_link}
	KmsKeySelfLink *string `field:"required" json:"kmsKeySelfLink" yaml:"kmsKeySelfLink"`
}

