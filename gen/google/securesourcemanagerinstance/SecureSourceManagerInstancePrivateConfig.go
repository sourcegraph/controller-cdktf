package securesourcemanagerinstance


type SecureSourceManagerInstancePrivateConfig struct {
	// 'Indicate if it's private instance.'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/secure_source_manager_instance#is_private SecureSourceManagerInstance#is_private}
	IsPrivate interface{} `field:"required" json:"isPrivate" yaml:"isPrivate"`
	// CA pool resource, resource must in the format of 'projects/{project}/locations/{location}/caPools/{ca_pool}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/secure_source_manager_instance#ca_pool SecureSourceManagerInstance#ca_pool}
	CaPool *string `field:"optional" json:"caPool" yaml:"caPool"`
	// custom_host_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/secure_source_manager_instance#custom_host_config SecureSourceManagerInstance#custom_host_config}
	CustomHostConfig *SecureSourceManagerInstancePrivateConfigCustomHostConfig `field:"optional" json:"customHostConfig" yaml:"customHostConfig"`
}

