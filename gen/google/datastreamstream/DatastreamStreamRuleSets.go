package datastreamstream


type DatastreamStreamRuleSets struct {
	// customization_rules block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/datastream_stream#customization_rules DatastreamStream#customization_rules}
	CustomizationRules interface{} `field:"required" json:"customizationRules" yaml:"customizationRules"`
	// object_filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/datastream_stream#object_filter DatastreamStream#object_filter}
	ObjectFilter *DatastreamStreamRuleSetsObjectFilter `field:"required" json:"objectFilter" yaml:"objectFilter"`
}

