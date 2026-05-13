package googlebigqueryconnection


type GoogleBigqueryConnectionConfigurationNetworkPrivateServiceConnect struct {
	// The resource name of a network attachment in the format of 'projects/{project}/regions/{region}/networkAttachments/{networkAttachment}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_bigquery_connection#network_attachment GoogleBigqueryConnection#network_attachment}
	NetworkAttachment *string `field:"required" json:"networkAttachment" yaml:"networkAttachment"`
}

