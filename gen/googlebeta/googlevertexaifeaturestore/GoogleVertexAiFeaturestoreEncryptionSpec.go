package googlevertexaifeaturestore


type GoogleVertexAiFeaturestoreEncryptionSpec struct {
	// The Cloud KMS resource identifier of the customer managed encryption key used to protect a resource.
	//
	// Has the form: projects/my-project/locations/my-region/keyRings/my-kr/cryptoKeys/my-key. The key needs to be in the same region as where the compute resource is created.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_vertex_ai_featurestore#kms_key_name GoogleVertexAiFeaturestore#kms_key_name}
	KmsKeyName *string `field:"required" json:"kmsKeyName" yaml:"kmsKeyName"`
}

