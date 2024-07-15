package googlecontainerazurenodepool


type GoogleContainerAzureNodePoolManagement struct {
	// Optional. Whether or not the nodes will be automatically repaired.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_azure_node_pool#auto_repair GoogleContainerAzureNodePool#auto_repair}
	AutoRepair interface{} `field:"optional" json:"autoRepair" yaml:"autoRepair"`
}

