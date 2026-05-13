package securesourcemanagerinstance


type SecureSourceManagerInstancePrivateConfigCustomHostConfig struct {
	// API hostname.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/secure_source_manager_instance#api SecureSourceManagerInstance#api}
	Api *string `field:"required" json:"api" yaml:"api"`
	// Git HTTP hostname.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/secure_source_manager_instance#git_http SecureSourceManagerInstance#git_http}
	GitHttp *string `field:"required" json:"gitHttp" yaml:"gitHttp"`
	// Git SSH hostname.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/secure_source_manager_instance#git_ssh SecureSourceManagerInstance#git_ssh}
	GitSsh *string `field:"required" json:"gitSsh" yaml:"gitSsh"`
	// HTML hostname.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/secure_source_manager_instance#html SecureSourceManagerInstance#html}
	Html *string `field:"required" json:"html" yaml:"html"`
}

