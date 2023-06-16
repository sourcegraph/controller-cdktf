package transferserver


type TransferServerWorkflowDetails struct {
	// on_upload block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/transfer_server#on_upload TransferServer#on_upload}
	OnUpload *TransferServerWorkflowDetailsOnUpload `field:"optional" json:"onUpload" yaml:"onUpload"`
}

