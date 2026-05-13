package googledatastreamprivateconnection


type GoogleDatastreamPrivateConnectionPscInterfaceConfig struct {
	// Fully qualified name of the network attachment that Datastream will connect to. Format: projects/{project}/regions/{region}/networkAttachments/{name}.
	//
	// To get Datastream project for the accepted list:
	// 'gcloud datastream private-connections create [PC ID] --location=[LOCATION] --network-attachment=[NA URI] --validate-only --display-name=[ANY STRING]'
	// Add Datastream project to the attachment accepted list:
	// 'gcloud compute network-attachments update [NA URI] --region=[NA region] --producer-accept-list=[TP from prev command]'
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_datastream_private_connection#network_attachment GoogleDatastreamPrivateConnection#network_attachment}
	NetworkAttachment *string `field:"required" json:"networkAttachment" yaml:"networkAttachment"`
}

