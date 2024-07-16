package googlenetworkservicesendpointpolicy


type GoogleNetworkServicesEndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabels struct {
	// Required. Label name presented as key in xDS Node Metadata.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_services_endpoint_policy#label_name GoogleNetworkServicesEndpointPolicy#label_name}
	LabelName *string `field:"required" json:"labelName" yaml:"labelName"`
	// Required. Label value presented as value corresponding to the above key, in xDS Node Metadata.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_services_endpoint_policy#label_value GoogleNetworkServicesEndpointPolicy#label_value}
	LabelValue *string `field:"required" json:"labelValue" yaml:"labelValue"`
}

