package googlememorystoreinstance


type GoogleMemorystoreInstanceDesiredPscAutoConnections struct {
	// Required. The consumer network where the IP address resides, in the form of projects/{project_id}/global/networks/{network_id}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_memorystore_instance#network GoogleMemorystoreInstance#network}
	Network *string `field:"required" json:"network" yaml:"network"`
	// Required. The consumer project_id where the forwarding rule is created from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_memorystore_instance#project_id GoogleMemorystoreInstance#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
}

