package googleapigeekeystoresaliaseskeycertfile


type GoogleApigeeKeystoresAliasesKeyCertFileCertsInfoCertInfo struct {
	// X.509 basic constraints extension.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_apigee_keystores_aliases_key_cert_file#basic_constraints GoogleApigeeKeystoresAliasesKeyCertFile#basic_constraints}
	BasicConstraints *string `field:"optional" json:"basicConstraints" yaml:"basicConstraints"`
	// X.509 notAfter validity period in milliseconds since epoch.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_apigee_keystores_aliases_key_cert_file#expiry_date GoogleApigeeKeystoresAliasesKeyCertFile#expiry_date}
	ExpiryDate *string `field:"optional" json:"expiryDate" yaml:"expiryDate"`
	// X.509 issuer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_apigee_keystores_aliases_key_cert_file#issuer GoogleApigeeKeystoresAliasesKeyCertFile#issuer}
	Issuer *string `field:"optional" json:"issuer" yaml:"issuer"`
	// Flag that specifies whether the certificate is valid.
	//
	// Flag is set to Yes if the certificate is valid, No if expired, or Not yet if not yet valid.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_apigee_keystores_aliases_key_cert_file#is_valid GoogleApigeeKeystoresAliasesKeyCertFile#is_valid}
	IsValid *string `field:"optional" json:"isValid" yaml:"isValid"`
	// Public key component of the X.509 subject public key info.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_apigee_keystores_aliases_key_cert_file#public_key GoogleApigeeKeystoresAliasesKeyCertFile#public_key}
	PublicKey *string `field:"optional" json:"publicKey" yaml:"publicKey"`
	// X.509 serial number.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_apigee_keystores_aliases_key_cert_file#serial_number GoogleApigeeKeystoresAliasesKeyCertFile#serial_number}
	SerialNumber *string `field:"optional" json:"serialNumber" yaml:"serialNumber"`
	// X.509 signatureAlgorithm.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_apigee_keystores_aliases_key_cert_file#sig_alg_name GoogleApigeeKeystoresAliasesKeyCertFile#sig_alg_name}
	SigAlgName *string `field:"optional" json:"sigAlgName" yaml:"sigAlgName"`
	// X.509 subject.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_apigee_keystores_aliases_key_cert_file#subject GoogleApigeeKeystoresAliasesKeyCertFile#subject}
	Subject *string `field:"optional" json:"subject" yaml:"subject"`
	// X.509 subject alternative names (SANs) extension.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_apigee_keystores_aliases_key_cert_file#subject_alternative_names GoogleApigeeKeystoresAliasesKeyCertFile#subject_alternative_names}
	SubjectAlternativeNames *[]*string `field:"optional" json:"subjectAlternativeNames" yaml:"subjectAlternativeNames"`
	// X.509 notBefore validity period in milliseconds since epoch.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_apigee_keystores_aliases_key_cert_file#valid_from GoogleApigeeKeystoresAliasesKeyCertFile#valid_from}
	ValidFrom *string `field:"optional" json:"validFrom" yaml:"validFrom"`
	// X.509 version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_apigee_keystores_aliases_key_cert_file#version GoogleApigeeKeystoresAliasesKeyCertFile#version}
	Version *float64 `field:"optional" json:"version" yaml:"version"`
}

