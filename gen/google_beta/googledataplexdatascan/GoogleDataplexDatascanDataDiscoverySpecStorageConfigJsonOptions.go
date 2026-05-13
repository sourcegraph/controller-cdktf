package googledataplexdatascan


type GoogleDataplexDatascanDataDiscoverySpecStorageConfigJsonOptions struct {
	// The character encoding of the data. The default is UTF-8.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dataplex_datascan#encoding GoogleDataplexDatascan#encoding}
	Encoding *string `field:"optional" json:"encoding" yaml:"encoding"`
	// Whether to disable the inference of data types for JSON data.
	//
	// If true, all columns are registered as their primitive types (strings, number, or boolean).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dataplex_datascan#type_inference_disabled GoogleDataplexDatascan#type_inference_disabled}
	TypeInferenceDisabled interface{} `field:"optional" json:"typeInferenceDisabled" yaml:"typeInferenceDisabled"`
}

