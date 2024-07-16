package googlenotebooksenvironment

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleNotebooksEnvironmentConfig struct {
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
	// A reference to the zone where the machine resides.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_notebooks_environment#location GoogleNotebooksEnvironment#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The name specified for the Environment instance. Format: projects/{project_id}/locations/{location}/environments/{environmentId}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_notebooks_environment#name GoogleNotebooksEnvironment#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// container_image block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_notebooks_environment#container_image GoogleNotebooksEnvironment#container_image}
	ContainerImage *GoogleNotebooksEnvironmentContainerImage `field:"optional" json:"containerImage" yaml:"containerImage"`
	// A brief description of this environment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_notebooks_environment#description GoogleNotebooksEnvironment#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Display name of this environment for the UI.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_notebooks_environment#display_name GoogleNotebooksEnvironment#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_notebooks_environment#id GoogleNotebooksEnvironment#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Path to a Bash script that automatically runs after a notebook instance fully boots up.
	//
	// The path must be a URL or Cloud Storage path. Example: "gs://path-to-file/file-name"
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_notebooks_environment#post_startup_script GoogleNotebooksEnvironment#post_startup_script}
	PostStartupScript *string `field:"optional" json:"postStartupScript" yaml:"postStartupScript"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_notebooks_environment#project GoogleNotebooksEnvironment#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_notebooks_environment#timeouts GoogleNotebooksEnvironment#timeouts}
	Timeouts *GoogleNotebooksEnvironmentTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// vm_image block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_notebooks_environment#vm_image GoogleNotebooksEnvironment#vm_image}
	VmImage *GoogleNotebooksEnvironmentVmImage `field:"optional" json:"vmImage" yaml:"vmImage"`
}

