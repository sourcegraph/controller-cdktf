package googlecontainercluster


type GoogleContainerClusterGatewayApiConfig struct {
	// The Gateway API release channel to use for Gateway API.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_container_cluster#channel GoogleContainerCluster#channel}
	Channel *string `field:"required" json:"channel" yaml:"channel"`
}

