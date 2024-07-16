package googlenetworkservicesendpointpolicy


type GoogleNetworkServicesEndpointPolicyEndpointMatcherMetadataLabelMatcher struct {
	// Specifies how matching should be done. Possible values: ["MATCH_ANY", "MATCH_ALL"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_services_endpoint_policy#metadata_label_match_criteria GoogleNetworkServicesEndpointPolicy#metadata_label_match_criteria}
	MetadataLabelMatchCriteria *string `field:"required" json:"metadataLabelMatchCriteria" yaml:"metadataLabelMatchCriteria"`
	// metadata_labels block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_services_endpoint_policy#metadata_labels GoogleNetworkServicesEndpointPolicy#metadata_labels}
	MetadataLabels interface{} `field:"optional" json:"metadataLabels" yaml:"metadataLabels"`
}

