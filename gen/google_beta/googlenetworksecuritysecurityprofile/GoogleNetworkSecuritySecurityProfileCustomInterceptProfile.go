package googlenetworksecuritysecurityprofile


type GoogleNetworkSecuritySecurityProfileCustomInterceptProfile struct {
	// The Intercept Endpoint Group to which matching traffic should be intercepted. Format: projects/{project_id}/locations/global/interceptEndpointGroups/{endpoint_group_id}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_network_security_security_profile#intercept_endpoint_group GoogleNetworkSecuritySecurityProfile#intercept_endpoint_group}
	InterceptEndpointGroup *string `field:"required" json:"interceptEndpointGroup" yaml:"interceptEndpointGroup"`
}

