package googlevertexaitensorboard


type GoogleVertexAiTensorboardEncryptionSpec struct {
	// The Cloud KMS resource identifier of the customer managed encryption key used to protect a resource.
	//
	// Has the form: projects/my-project/locations/my-region/keyRings/my-kr/cryptoKeys/my-key. The key needs to be in the same region as where the resource is created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_vertex_ai_tensorboard#kms_key_name GoogleVertexAiTensorboard#kms_key_name}
	KmsKeyName *string `field:"required" json:"kmsKeyName" yaml:"kmsKeyName"`
}

