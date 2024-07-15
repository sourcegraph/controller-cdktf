package googlekmskeyring


type GoogleKmsKeyRingTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_kms_key_ring#create GoogleKmsKeyRing#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_kms_key_ring#delete GoogleKmsKeyRing#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

