package googleprivatecacertificate


type GooglePrivatecaCertificateConfigSubjectConfig struct {
	// subject block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_privateca_certificate#subject GooglePrivatecaCertificate#subject}
	Subject *GooglePrivatecaCertificateConfigSubjectConfigSubject `field:"required" json:"subject" yaml:"subject"`
	// subject_alt_name block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_privateca_certificate#subject_alt_name GooglePrivatecaCertificate#subject_alt_name}
	SubjectAltName *GooglePrivatecaCertificateConfigSubjectConfigSubjectAltName `field:"optional" json:"subjectAltName" yaml:"subjectAltName"`
}

