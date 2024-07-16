package googlevpcaccessconnector


type GoogleVpcAccessConnectorSubnet struct {
	// Subnet name (relative, not fully qualified).
	//
	// E.g. if the full subnet selfLink is
	// https://compute.googleapis.com/compute/v1/projects/{project}/regions/{region}/subnetworks/{subnetName} the correct input for this field would be {subnetName}"
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_vpc_access_connector#name GoogleVpcAccessConnector#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Project in which the subnet exists.
	//
	// If not set, this project is assumed to be the project for which the connector create request was issued.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_vpc_access_connector#project_id GoogleVpcAccessConnector#project_id}
	ProjectId *string `field:"optional" json:"projectId" yaml:"projectId"`
}

