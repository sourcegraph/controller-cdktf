package privatecacertificate


type PrivatecaCertificateConfigSubjectKeyId struct {
	// The value of the KeyId in lowercase hexidecimal.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.38.0/docs/resources/privateca_certificate#key_id PrivatecaCertificate#key_id}
	KeyId *string `field:"optional" json:"keyId" yaml:"keyId"`
}

