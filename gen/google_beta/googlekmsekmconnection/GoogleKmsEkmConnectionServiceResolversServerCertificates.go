package googlekmsekmconnection


type GoogleKmsEkmConnectionServiceResolversServerCertificates struct {
	// Required. The raw certificate bytes in DER format. A base64-encoded string.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_kms_ekm_connection#raw_der GoogleKmsEkmConnection#raw_der}
	RawDer *string `field:"required" json:"rawDer" yaml:"rawDer"`
	// Output only. The subject Alternative DNS names. Only present if parsed is true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_kms_ekm_connection#subject_alternative_dns_names GoogleKmsEkmConnection#subject_alternative_dns_names}
	SubjectAlternativeDnsNames *[]*string `field:"optional" json:"subjectAlternativeDnsNames" yaml:"subjectAlternativeDnsNames"`
}

