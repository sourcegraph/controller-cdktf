package googlecolabruntimetemplate


type GoogleColabRuntimeTemplateEncryptionSpec struct {
	// The Cloud KMS encryption key (customer-managed encryption key) used to protect the runtime.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_colab_runtime_template#kms_key_name GoogleColabRuntimeTemplate#kms_key_name}
	KmsKeyName *string `field:"optional" json:"kmsKeyName" yaml:"kmsKeyName"`
}

