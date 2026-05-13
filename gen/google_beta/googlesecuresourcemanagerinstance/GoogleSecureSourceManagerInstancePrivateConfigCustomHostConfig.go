package googlesecuresourcemanagerinstance


type GoogleSecureSourceManagerInstancePrivateConfigCustomHostConfig struct {
	// API hostname.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_secure_source_manager_instance#api GoogleSecureSourceManagerInstance#api}
	Api *string `field:"required" json:"api" yaml:"api"`
	// Git HTTP hostname.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_secure_source_manager_instance#git_http GoogleSecureSourceManagerInstance#git_http}
	GitHttp *string `field:"required" json:"gitHttp" yaml:"gitHttp"`
	// Git SSH hostname.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_secure_source_manager_instance#git_ssh GoogleSecureSourceManagerInstance#git_ssh}
	GitSsh *string `field:"required" json:"gitSsh" yaml:"gitSsh"`
	// HTML hostname.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_secure_source_manager_instance#html GoogleSecureSourceManagerInstance#html}
	Html *string `field:"required" json:"html" yaml:"html"`
}

