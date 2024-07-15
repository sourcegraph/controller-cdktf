package googledatastreamprivateconnection


type GoogleDatastreamPrivateConnectionVpcPeeringConfig struct {
	// A free subnet for peering. (CIDR of /29).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_datastream_private_connection#subnet GoogleDatastreamPrivateConnection#subnet}
	Subnet *string `field:"required" json:"subnet" yaml:"subnet"`
	// Fully qualified name of the VPC that Datastream will peer to. Format: projects/{project}/global/{networks}/{name}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_datastream_private_connection#vpc GoogleDatastreamPrivateConnection#vpc}
	Vpc *string `field:"required" json:"vpc" yaml:"vpc"`
}

