package googlecomputeserviceattachment


type GoogleComputeServiceAttachmentConsumerAcceptLists struct {
	// The number of consumer forwarding rules the consumer project can create.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_compute_service_attachment#connection_limit GoogleComputeServiceAttachment#connection_limit}
	ConnectionLimit *float64 `field:"required" json:"connectionLimit" yaml:"connectionLimit"`
	// A project that is allowed to connect to this service attachment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_compute_service_attachment#project_id_or_num GoogleComputeServiceAttachment#project_id_or_num}
	ProjectIdOrNum *string `field:"required" json:"projectIdOrNum" yaml:"projectIdOrNum"`
}

