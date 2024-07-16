package googleedgecontainervpnconnection


type GoogleEdgecontainerVpnConnectionVpcProject struct {
	// The project of the VPC to connect to. If not specified, it is the same as the cluster project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_edgecontainer_vpn_connection#project_id GoogleEdgecontainerVpnConnection#project_id}
	ProjectId *string `field:"optional" json:"projectId" yaml:"projectId"`
}

