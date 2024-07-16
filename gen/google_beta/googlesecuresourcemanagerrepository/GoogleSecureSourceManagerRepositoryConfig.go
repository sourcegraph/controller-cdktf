package googlesecuresourcemanagerrepository

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleSecureSourceManagerRepositoryConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The name of the instance in which the repository is hosted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_secure_source_manager_repository#instance GoogleSecureSourceManagerRepository#instance}
	Instance *string `field:"required" json:"instance" yaml:"instance"`
	// The location for the Repository.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_secure_source_manager_repository#location GoogleSecureSourceManagerRepository#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The ID for the Repository.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_secure_source_manager_repository#repository_id GoogleSecureSourceManagerRepository#repository_id}
	RepositoryId *string `field:"required" json:"repositoryId" yaml:"repositoryId"`
	// Description of the repository, which cannot exceed 500 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_secure_source_manager_repository#description GoogleSecureSourceManagerRepository#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_secure_source_manager_repository#id GoogleSecureSourceManagerRepository#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// initial_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_secure_source_manager_repository#initial_config GoogleSecureSourceManagerRepository#initial_config}
	InitialConfig *GoogleSecureSourceManagerRepositoryInitialConfig `field:"optional" json:"initialConfig" yaml:"initialConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_secure_source_manager_repository#project GoogleSecureSourceManagerRepository#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_secure_source_manager_repository#timeouts GoogleSecureSourceManagerRepository#timeouts}
	Timeouts *GoogleSecureSourceManagerRepositoryTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

