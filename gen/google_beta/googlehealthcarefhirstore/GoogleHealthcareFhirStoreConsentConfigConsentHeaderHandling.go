package googlehealthcarefhirstore


type GoogleHealthcareFhirStoreConsentConfigConsentHeaderHandling struct {
	// Specifies the default server behavior when the header is empty.
	//
	// If not specified, the ScopeProfile.PERMIT_EMPTY_SCOPE option is used. Default value: "PERMIT_EMPTY_SCOPE" Possible values: ["SCOPE_PROFILE_UNSPECIFIED", "PERMIT_EMPTY_SCOPE", "REQUIRED_ON_READ"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_healthcare_fhir_store#profile GoogleHealthcareFhirStore#profile}
	Profile *string `field:"optional" json:"profile" yaml:"profile"`
}

