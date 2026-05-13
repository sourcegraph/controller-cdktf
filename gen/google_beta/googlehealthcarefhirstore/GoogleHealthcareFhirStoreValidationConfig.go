package googlehealthcarefhirstore


type GoogleHealthcareFhirStoreValidationConfig struct {
	// Whether to disable FHIRPath validation for incoming resources.
	//
	// The default value is false. Set this to true to disable checking incoming resources for conformance against FHIRPath requirement defined in the FHIR specification. This property only affects resource types that do not have profiles configured for them, any rules in enabled implementation guides will still be enforced.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_healthcare_fhir_store#disable_fhirpath_validation GoogleHealthcareFhirStore#disable_fhirpath_validation}
	DisableFhirpathValidation interface{} `field:"optional" json:"disableFhirpathValidation" yaml:"disableFhirpathValidation"`
	// Whether to disable profile validation for this FHIR store.
	//
	// The default value is false. Set this to true to disable checking incoming resources for conformance against structure definitions in this FHIR store.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_healthcare_fhir_store#disable_profile_validation GoogleHealthcareFhirStore#disable_profile_validation}
	DisableProfileValidation interface{} `field:"optional" json:"disableProfileValidation" yaml:"disableProfileValidation"`
	// Whether to disable reference type validation for incoming resources.
	//
	// The default value is false. Set this to true to disable checking incoming resources for conformance against reference type requirement defined in the FHIR specification. This property only affects resource types that do not have profiles configured for them, any rules in enabled implementation guides will still be enforced.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_healthcare_fhir_store#disable_reference_type_validation GoogleHealthcareFhirStore#disable_reference_type_validation}
	DisableReferenceTypeValidation interface{} `field:"optional" json:"disableReferenceTypeValidation" yaml:"disableReferenceTypeValidation"`
	// Whether to disable required fields validation for incoming resources.
	//
	// The default value is false. Set this to true to disable checking incoming resources for conformance against required fields requirement defined in the FHIR specification. This property only affects resource types that do not have profiles configured for them, any rules in enabled implementation guides will still be enforced.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_healthcare_fhir_store#disable_required_field_validation GoogleHealthcareFhirStore#disable_required_field_validation}
	DisableRequiredFieldValidation interface{} `field:"optional" json:"disableRequiredFieldValidation" yaml:"disableRequiredFieldValidation"`
	// A list of implementation guide URLs in this FHIR store that are used to configure the profiles to use for validation.
	//
	// When a URL cannot be resolved (for example, in a type assertion), the server does not return an error.
	// For example, to use the US Core profiles for validation, set enabledImplementationGuides to ["http://hl7.org/fhir/us/core/ImplementationGuide/ig"]. If enabledImplementationGuides is empty or omitted, then incoming resources are only required to conform to the base FHIR profiles. Otherwise, a resource must conform to at least one profile listed in the global property of one of the enabled ImplementationGuides.
	// The Cloud Healthcare API does not currently enforce all of the rules in a StructureDefinition. The following rules are supported:
	// - min/max
	// - minValue/maxValue
	// - maxLength
	// - type
	// - fixed[x]
	// - pattern[x] on simple types
	// - slicing, when using "value" as the discriminator type
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_healthcare_fhir_store#enabled_implementation_guides GoogleHealthcareFhirStore#enabled_implementation_guides}
	EnabledImplementationGuides *[]*string `field:"optional" json:"enabledImplementationGuides" yaml:"enabledImplementationGuides"`
}

