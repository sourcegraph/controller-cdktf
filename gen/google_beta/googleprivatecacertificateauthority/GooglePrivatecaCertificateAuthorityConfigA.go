package googleprivatecacertificateauthority


type GooglePrivatecaCertificateAuthorityConfigA struct {
	// subject_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_privateca_certificate_authority#subject_config GooglePrivatecaCertificateAuthority#subject_config}
	SubjectConfig *GooglePrivatecaCertificateAuthorityConfigSubjectConfig `field:"required" json:"subjectConfig" yaml:"subjectConfig"`
	// x509_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_privateca_certificate_authority#x509_config GooglePrivatecaCertificateAuthority#x509_config}
	X509Config *GooglePrivatecaCertificateAuthorityConfigX509Config `field:"required" json:"x509Config" yaml:"x509Config"`
	// subject_key_id block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_privateca_certificate_authority#subject_key_id GooglePrivatecaCertificateAuthority#subject_key_id}
	SubjectKeyId *GooglePrivatecaCertificateAuthorityConfigSubjectKeyId `field:"optional" json:"subjectKeyId" yaml:"subjectKeyId"`
}

