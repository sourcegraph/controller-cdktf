package googlebigqueryconnection


type GoogleBigqueryConnectionConfigurationEndpoint struct {
	// Host and port in the format of 'host:port' for the connector endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_bigquery_connection#host_port GoogleBigqueryConnection#host_port}
	HostPort *string `field:"optional" json:"hostPort" yaml:"hostPort"`
}

