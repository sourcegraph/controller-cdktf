package googlenetworksecuritysecurityprofile


type GoogleNetworkSecuritySecurityProfileCustomMirroringProfile struct {
	// The target Mirroring Endpoint Group.
	//
	// When a mirroring rule with this security profile attached matches a packet,
	// a replica will be mirrored to the location-local target in this group.
	// Format: projects/{project_id}/locations/global/mirroringEndpointGroups/{endpoint_group_id}
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_security_security_profile#mirroring_endpoint_group GoogleNetworkSecuritySecurityProfile#mirroring_endpoint_group}
	MirroringEndpointGroup *string `field:"required" json:"mirroringEndpointGroup" yaml:"mirroringEndpointGroup"`
	// The target downstream Mirroring Deployment Groups.
	//
	// This field is used for Packet Broker mirroring endpoint groups to specify
	// the deployment groups that the packet should be mirrored to by the broker.
	// Format: projects/{project_id}/locations/global/mirroringDeploymentGroups/{deployment_group_id}
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_security_security_profile#mirroring_deployment_groups GoogleNetworkSecuritySecurityProfile#mirroring_deployment_groups}
	MirroringDeploymentGroups *[]*string `field:"optional" json:"mirroringDeploymentGroups" yaml:"mirroringDeploymentGroups"`
}

