package googledatastreamconnectionprofile


type GoogleDatastreamConnectionProfileMongodbProfileStandardConnectionFormat struct {
	// Specifies whether the client connects directly to the host[:port] in the connection URI.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_datastream_connection_profile#direct_connection GoogleDatastreamConnectionProfile#direct_connection}
	DirectConnection interface{} `field:"optional" json:"directConnection" yaml:"directConnection"`
}

