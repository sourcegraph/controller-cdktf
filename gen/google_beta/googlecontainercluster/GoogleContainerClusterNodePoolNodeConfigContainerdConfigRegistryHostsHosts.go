package googlecontainercluster


type GoogleContainerClusterNodePoolNodeConfigContainerdConfigRegistryHostsHosts struct {
	// Configures the registry host/mirror.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_container_cluster#host GoogleContainerCluster#host}
	Host *string `field:"required" json:"host" yaml:"host"`
	// ca block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_container_cluster#ca GoogleContainerCluster#ca}
	Ca interface{} `field:"optional" json:"ca" yaml:"ca"`
	// Represent the capabilities of the registry host, specifying what operations a host is capable of performing.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_container_cluster#capabilities GoogleContainerCluster#capabilities}
	Capabilities *[]*string `field:"optional" json:"capabilities" yaml:"capabilities"`
	// client block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_container_cluster#client GoogleContainerCluster#client}
	Client interface{} `field:"optional" json:"client" yaml:"client"`
	// Specifies the maximum duration allowed for a connection attempt to complete.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_container_cluster#dial_timeout GoogleContainerCluster#dial_timeout}
	DialTimeout *string `field:"optional" json:"dialTimeout" yaml:"dialTimeout"`
	// header block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_container_cluster#header GoogleContainerCluster#header}
	Header interface{} `field:"optional" json:"header" yaml:"header"`
	// Indicate the host's API root endpoint is defined in the URL path rather than by the API specification.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_container_cluster#override_path GoogleContainerCluster#override_path}
	OverridePath interface{} `field:"optional" json:"overridePath" yaml:"overridePath"`
}

