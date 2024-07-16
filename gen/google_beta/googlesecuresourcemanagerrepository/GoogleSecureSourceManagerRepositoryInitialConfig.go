package googlesecuresourcemanagerrepository


type GoogleSecureSourceManagerRepositoryInitialConfig struct {
	// Default branch name of the repository.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_secure_source_manager_repository#default_branch GoogleSecureSourceManagerRepository#default_branch}
	DefaultBranch *string `field:"optional" json:"defaultBranch" yaml:"defaultBranch"`
	// List of gitignore template names user can choose from. Valid values can be viewed at https://cloud.google.com/secure-source-manager/docs/reference/rest/v1/projects.locations.repositories#initialconfig.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_secure_source_manager_repository#gitignores GoogleSecureSourceManagerRepository#gitignores}
	Gitignores *[]*string `field:"optional" json:"gitignores" yaml:"gitignores"`
	// License template name user can choose from. Valid values can be viewed at https://cloud.google.com/secure-source-manager/docs/reference/rest/v1/projects.locations.repositories#initialconfig.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_secure_source_manager_repository#license GoogleSecureSourceManagerRepository#license}
	License *string `field:"optional" json:"license" yaml:"license"`
	// README template name. Valid values can be viewed at https://cloud.google.com/secure-source-manager/docs/reference/rest/v1/projects.locations.repositories#initialconfig.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_secure_source_manager_repository#readme GoogleSecureSourceManagerRepository#readme}
	Readme *string `field:"optional" json:"readme" yaml:"readme"`
}

