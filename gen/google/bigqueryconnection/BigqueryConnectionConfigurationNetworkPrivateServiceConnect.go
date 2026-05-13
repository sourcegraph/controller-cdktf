package bigqueryconnection


type BigqueryConnectionConfigurationNetworkPrivateServiceConnect struct {
	// The resource name of a network attachment in the format of 'projects/{project}/regions/{region}/networkAttachments/{networkAttachment}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/bigquery_connection#network_attachment BigqueryConnection#network_attachment}
	NetworkAttachment *string `field:"required" json:"networkAttachment" yaml:"networkAttachment"`
}

