package googleapigeetargetserver


type GoogleApigeeTargetServerSSlInfoCommonName struct {
	// The TLS Common Name string of the certificate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_apigee_target_server#value GoogleApigeeTargetServer#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
	// Indicates whether the cert should be matched against as a wildcard cert.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_apigee_target_server#wildcard_match GoogleApigeeTargetServer#wildcard_match}
	WildcardMatch interface{} `field:"optional" json:"wildcardMatch" yaml:"wildcardMatch"`
}

