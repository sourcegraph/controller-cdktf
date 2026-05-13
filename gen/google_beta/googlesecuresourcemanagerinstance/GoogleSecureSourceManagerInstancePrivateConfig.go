package googlesecuresourcemanagerinstance


type GoogleSecureSourceManagerInstancePrivateConfig struct {
	// 'Indicate if it's private instance.'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_secure_source_manager_instance#is_private GoogleSecureSourceManagerInstance#is_private}
	IsPrivate interface{} `field:"required" json:"isPrivate" yaml:"isPrivate"`
	// CA pool resource, resource must in the format of 'projects/{project}/locations/{location}/caPools/{ca_pool}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_secure_source_manager_instance#ca_pool GoogleSecureSourceManagerInstance#ca_pool}
	CaPool *string `field:"optional" json:"caPool" yaml:"caPool"`
	// custom_host_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_secure_source_manager_instance#custom_host_config GoogleSecureSourceManagerInstance#custom_host_config}
	CustomHostConfig *GoogleSecureSourceManagerInstancePrivateConfigCustomHostConfig `field:"optional" json:"customHostConfig" yaml:"customHostConfig"`
}

