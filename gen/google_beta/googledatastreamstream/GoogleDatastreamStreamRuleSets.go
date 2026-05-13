package googledatastreamstream


type GoogleDatastreamStreamRuleSets struct {
	// customization_rules block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_datastream_stream#customization_rules GoogleDatastreamStream#customization_rules}
	CustomizationRules interface{} `field:"required" json:"customizationRules" yaml:"customizationRules"`
	// object_filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_datastream_stream#object_filter GoogleDatastreamStream#object_filter}
	ObjectFilter *GoogleDatastreamStreamRuleSetsObjectFilter `field:"required" json:"objectFilter" yaml:"objectFilter"`
}

