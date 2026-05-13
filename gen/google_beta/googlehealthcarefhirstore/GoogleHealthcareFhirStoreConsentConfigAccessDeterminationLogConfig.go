package googlehealthcarefhirstore


type GoogleHealthcareFhirStoreConsentConfigAccessDeterminationLogConfig struct {
	// Controls the amount of detail to include as part of the audit logs.
	//
	// Default value: "MINIMUM" Possible values: ["LOG_LEVEL_UNSPECIFIED", "DISABLED", "MINIMUM", "VERBOSE"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_healthcare_fhir_store#log_level GoogleHealthcareFhirStore#log_level}
	LogLevel *string `field:"optional" json:"logLevel" yaml:"logLevel"`
}

