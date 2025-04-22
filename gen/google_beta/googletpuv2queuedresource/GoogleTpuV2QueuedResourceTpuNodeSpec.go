package googletpuv2queuedresource


type GoogleTpuV2QueuedResourceTpuNodeSpec struct {
	// node block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_tpu_v2_queued_resource#node GoogleTpuV2QueuedResource#node}
	NodeAttribute *GoogleTpuV2QueuedResourceTpuNodeSpecNode `field:"required" json:"nodeAttribute" yaml:"nodeAttribute"`
	// The parent resource name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_tpu_v2_queued_resource#parent GoogleTpuV2QueuedResource#parent}
	Parent *string `field:"required" json:"parent" yaml:"parent"`
	// Unqualified node identifier used to identify the node in the project once provisioned.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_tpu_v2_queued_resource#node_id GoogleTpuV2QueuedResource#node_id}
	NodeId *string `field:"optional" json:"nodeId" yaml:"nodeId"`
}

