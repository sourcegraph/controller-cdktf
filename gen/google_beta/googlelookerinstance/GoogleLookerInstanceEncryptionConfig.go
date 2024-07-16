package googlelookerinstance


type GoogleLookerInstanceEncryptionConfig struct {
	// Name of the customer managed encryption key (CMEK) in KMS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_looker_instance#kms_key_name GoogleLookerInstance#kms_key_name}
	KmsKeyName *string `field:"optional" json:"kmsKeyName" yaml:"kmsKeyName"`
}

