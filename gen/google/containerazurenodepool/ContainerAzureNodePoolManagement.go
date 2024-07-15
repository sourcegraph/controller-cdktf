package containerazurenodepool


type ContainerAzureNodePoolManagement struct {
	// Optional. Whether or not the nodes will be automatically repaired.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.38.0/docs/resources/container_azure_node_pool#auto_repair ContainerAzureNodePool#auto_repair}
	AutoRepair interface{} `field:"optional" json:"autoRepair" yaml:"autoRepair"`
}

