package googlemodelarmortemplate


type GoogleModelArmorTemplateFilterConfigRaiSettingsRaiFilters struct {
	// Possible values: SEXUALLY_EXPLICIT HATE_SPEECH HARASSMENT DANGEROUS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_model_armor_template#filter_type GoogleModelArmorTemplate#filter_type}
	FilterType *string `field:"required" json:"filterType" yaml:"filterType"`
	// Possible values: LOW_AND_ABOVE MEDIUM_AND_ABOVE HIGH.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_model_armor_template#confidence_level GoogleModelArmorTemplate#confidence_level}
	ConfidenceLevel *string `field:"optional" json:"confidenceLevel" yaml:"confidenceLevel"`
}

