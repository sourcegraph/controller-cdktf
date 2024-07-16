package googlenotebooksruntime


type GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfigEncryptionConfig struct {
	// The Cloud KMS resource identifier of the customer-managed encryption key used to protect a resource, such as a disks.
	//
	// It has the following format:
	// 'projects/{PROJECT_ID}/locations/{REGION}/keyRings/
	// {KEY_RING_NAME}/cryptoKeys/{KEY_NAME}'
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_notebooks_runtime#kms_key GoogleNotebooksRuntime#kms_key}
	KmsKey *string `field:"optional" json:"kmsKey" yaml:"kmsKey"`
}

