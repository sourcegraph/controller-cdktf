package googlesecuresourcemanagerinstance


type GoogleSecureSourceManagerInstancePrivateConfig struct {
	// CA pool resource, resource must in the format of 'projects/{project}/locations/{location}/caPools/{ca_pool}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_secure_source_manager_instance#ca_pool GoogleSecureSourceManagerInstance#ca_pool}
	CaPool *string `field:"required" json:"caPool" yaml:"caPool"`
	// 'Indicate if it's private instance.'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_secure_source_manager_instance#is_private GoogleSecureSourceManagerInstance#is_private}
	IsPrivate interface{} `field:"required" json:"isPrivate" yaml:"isPrivate"`
}

