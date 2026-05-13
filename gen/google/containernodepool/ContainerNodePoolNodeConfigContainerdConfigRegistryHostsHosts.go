package containernodepool


type ContainerNodePoolNodeConfigContainerdConfigRegistryHostsHosts struct {
	// Configures the registry host/mirror.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/container_node_pool#host ContainerNodePool#host}
	Host *string `field:"required" json:"host" yaml:"host"`
	// ca block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/container_node_pool#ca ContainerNodePool#ca}
	Ca interface{} `field:"optional" json:"ca" yaml:"ca"`
	// Represent the capabilities of the registry host, specifying what operations a host is capable of performing.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/container_node_pool#capabilities ContainerNodePool#capabilities}
	Capabilities *[]*string `field:"optional" json:"capabilities" yaml:"capabilities"`
	// client block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/container_node_pool#client ContainerNodePool#client}
	Client interface{} `field:"optional" json:"client" yaml:"client"`
	// Specifies the maximum duration allowed for a connection attempt to complete.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/container_node_pool#dial_timeout ContainerNodePool#dial_timeout}
	DialTimeout *string `field:"optional" json:"dialTimeout" yaml:"dialTimeout"`
	// header block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/container_node_pool#header ContainerNodePool#header}
	Header interface{} `field:"optional" json:"header" yaml:"header"`
	// Indicate the host's API root endpoint is defined in the URL path rather than by the API specification.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/container_node_pool#override_path ContainerNodePool#override_path}
	OverridePath interface{} `field:"optional" json:"overridePath" yaml:"overridePath"`
}

