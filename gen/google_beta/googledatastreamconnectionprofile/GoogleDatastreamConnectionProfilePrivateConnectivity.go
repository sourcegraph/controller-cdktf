package googledatastreamconnectionprofile


type GoogleDatastreamConnectionProfilePrivateConnectivity struct {
	// A reference to a private connection resource. Format: 'projects/{project}/locations/{location}/privateConnections/{name}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_datastream_connection_profile#private_connection GoogleDatastreamConnectionProfile#private_connection}
	PrivateConnection *string `field:"required" json:"privateConnection" yaml:"privateConnection"`
}

