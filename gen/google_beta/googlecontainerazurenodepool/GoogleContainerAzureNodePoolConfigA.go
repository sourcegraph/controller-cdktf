package googlecontainerazurenodepool


type GoogleContainerAzureNodePoolConfigA struct {
	// ssh_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_azure_node_pool#ssh_config GoogleContainerAzureNodePool#ssh_config}
	SshConfig *GoogleContainerAzureNodePoolConfigSshConfig `field:"required" json:"sshConfig" yaml:"sshConfig"`
	// The OS image type to use on node pool instances.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_azure_node_pool#image_type GoogleContainerAzureNodePool#image_type}
	ImageType *string `field:"optional" json:"imageType" yaml:"imageType"`
	// Optional.
	//
	// The initial labels assigned to nodes of this node pool. An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_azure_node_pool#labels GoogleContainerAzureNodePool#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// proxy_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_azure_node_pool#proxy_config GoogleContainerAzureNodePool#proxy_config}
	ProxyConfig *GoogleContainerAzureNodePoolConfigProxyConfig `field:"optional" json:"proxyConfig" yaml:"proxyConfig"`
	// root_volume block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_azure_node_pool#root_volume GoogleContainerAzureNodePool#root_volume}
	RootVolume *GoogleContainerAzureNodePoolConfigRootVolume `field:"optional" json:"rootVolume" yaml:"rootVolume"`
	// Optional.
	//
	// A set of tags to apply to all underlying Azure resources for this node pool. This currently only includes Virtual Machine Scale Sets. Specify at most 50 pairs containing alphanumerics, spaces, and symbols (.+-=_:@/). Keys can be up to 127 Unicode characters. Values can be up to 255 Unicode characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_azure_node_pool#tags GoogleContainerAzureNodePool#tags}
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Optional. The Azure VM size name. Example: `Standard_DS2_v2`. See (/anthos/clusters/docs/azure/reference/supported-vms) for options. When unspecified, it defaults to `Standard_DS2_v2`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_azure_node_pool#vm_size GoogleContainerAzureNodePool#vm_size}
	VmSize *string `field:"optional" json:"vmSize" yaml:"vmSize"`
}

