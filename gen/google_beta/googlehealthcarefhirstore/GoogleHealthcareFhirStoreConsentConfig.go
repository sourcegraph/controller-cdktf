package googlehealthcarefhirstore


type GoogleHealthcareFhirStoreConsentConfig struct {
	// Specifies which consent enforcement version is being used for this FHIR store.
	//
	// This field can only be set once by either [fhirStores.create][] or [fhirStores.patch][]. After that, you must call [fhirStores.applyConsents][] to change the version. Possible values: ["CONSENT_ENFORCEMENT_VERSION_UNSPECIFIED", "V1"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_healthcare_fhir_store#version GoogleHealthcareFhirStore#version}
	Version *string `field:"required" json:"version" yaml:"version"`
	// access_determination_log_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_healthcare_fhir_store#access_determination_log_config GoogleHealthcareFhirStore#access_determination_log_config}
	AccessDeterminationLogConfig *GoogleHealthcareFhirStoreConsentConfigAccessDeterminationLogConfig `field:"optional" json:"accessDeterminationLogConfig" yaml:"accessDeterminationLogConfig"`
	// The default value is false.
	//
	// If set to true, when accessing FHIR resources, the consent headers will be verified against consents given by patients. See the ConsentEnforcementVersion for the supported consent headers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_healthcare_fhir_store#access_enforced GoogleHealthcareFhirStore#access_enforced}
	AccessEnforced interface{} `field:"optional" json:"accessEnforced" yaml:"accessEnforced"`
	// consent_header_handling block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_healthcare_fhir_store#consent_header_handling GoogleHealthcareFhirStore#consent_header_handling}
	ConsentHeaderHandling *GoogleHealthcareFhirStoreConsentConfigConsentHeaderHandling `field:"optional" json:"consentHeaderHandling" yaml:"consentHeaderHandling"`
}

