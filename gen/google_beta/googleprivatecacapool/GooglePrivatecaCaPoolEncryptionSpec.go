package googleprivatecacapool


type GooglePrivatecaCaPoolEncryptionSpec struct {
	// The resource name for an existing Cloud KMS key in the format 'projects/* /locations/* /keyRings/* /cryptoKeys/*'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_privateca_ca_pool#cloud_kms_key GooglePrivatecaCaPool#cloud_kms_key}
	//
	// Note: The above comment contained a comment block ending sequence (* followed by /). We have introduced a space between to prevent syntax errors. Please ignore the space.
	CloudKmsKey *string `field:"optional" json:"cloudKmsKey" yaml:"cloudKmsKey"`
}

