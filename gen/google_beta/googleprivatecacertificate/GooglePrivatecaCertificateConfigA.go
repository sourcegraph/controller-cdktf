package googleprivatecacertificate


type GooglePrivatecaCertificateConfigA struct {
	// public_key block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_privateca_certificate#public_key GooglePrivatecaCertificate#public_key}
	PublicKey *GooglePrivatecaCertificateConfigPublicKey `field:"required" json:"publicKey" yaml:"publicKey"`
	// subject_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_privateca_certificate#subject_config GooglePrivatecaCertificate#subject_config}
	SubjectConfig *GooglePrivatecaCertificateConfigSubjectConfig `field:"required" json:"subjectConfig" yaml:"subjectConfig"`
	// x509_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_privateca_certificate#x509_config GooglePrivatecaCertificate#x509_config}
	X509Config *GooglePrivatecaCertificateConfigX509Config `field:"required" json:"x509Config" yaml:"x509Config"`
	// subject_key_id block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_privateca_certificate#subject_key_id GooglePrivatecaCertificate#subject_key_id}
	SubjectKeyId *GooglePrivatecaCertificateConfigSubjectKeyId `field:"optional" json:"subjectKeyId" yaml:"subjectKeyId"`
}

