package googleprivatecacertificateauthority


type GooglePrivatecaCertificateAuthorityConfigSubjectConfig struct {
	// subject block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_privateca_certificate_authority#subject GooglePrivatecaCertificateAuthority#subject}
	Subject *GooglePrivatecaCertificateAuthorityConfigSubjectConfigSubject `field:"required" json:"subject" yaml:"subject"`
	// subject_alt_name block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_privateca_certificate_authority#subject_alt_name GooglePrivatecaCertificateAuthority#subject_alt_name}
	SubjectAltName *GooglePrivatecaCertificateAuthorityConfigSubjectConfigSubjectAltName `field:"optional" json:"subjectAltName" yaml:"subjectAltName"`
}

