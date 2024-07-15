package googlehealthcarehl7v2store


type GoogleHealthcareHl7V2StoreParserConfig struct {
	// Determines whether messages with no header are allowed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_healthcare_hl7_v2_store#allow_null_header GoogleHealthcareHl7V2Store#allow_null_header}
	AllowNullHeader interface{} `field:"optional" json:"allowNullHeader" yaml:"allowNullHeader"`
	// JSON encoded string for schemas used to parse messages in this store if schematized parsing is desired.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_healthcare_hl7_v2_store#schema GoogleHealthcareHl7V2Store#schema}
	Schema *string `field:"optional" json:"schema" yaml:"schema"`
	// Byte(s) to be used as the segment terminator. If this is unset, '\r' will be used as segment terminator.
	//
	// A base64-encoded string.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_healthcare_hl7_v2_store#segment_terminator GoogleHealthcareHl7V2Store#segment_terminator}
	SegmentTerminator *string `field:"optional" json:"segmentTerminator" yaml:"segmentTerminator"`
	// The version of the unschematized parser to be used when a custom 'schema' is not set.
	//
	// Default value: "V1" Possible values: ["V1", "V2", "V3"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_healthcare_hl7_v2_store#version GoogleHealthcareHl7V2Store#version}
	Version *string `field:"optional" json:"version" yaml:"version"`
}

