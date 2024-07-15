package googledataformrepository

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleDataformRepositoryConfig struct {
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
	// The repository's name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataform_repository#name GoogleDataformRepository#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Optional. The repository's user-friendly name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataform_repository#display_name GoogleDataformRepository#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// git_remote_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataform_repository#git_remote_settings GoogleDataformRepository#git_remote_settings}
	GitRemoteSettings *GoogleDataformRepositoryGitRemoteSettings `field:"optional" json:"gitRemoteSettings" yaml:"gitRemoteSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataform_repository#id GoogleDataformRepository#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Optional.
	//
	// Repository user labels.
	// An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	//
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataform_repository#labels GoogleDataformRepository#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Optional.
	//
	// The name of the Secret Manager secret version to be used to interpolate variables into the .npmrc file for package installation operations. Must be in the format projects/* /secrets/* /versions/*. The file itself must be in a JSON format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataform_repository#npmrc_environment_variables_secret_version GoogleDataformRepository#npmrc_environment_variables_secret_version}
	//
	// Note: The above comment contained a comment block ending sequence (* followed by /). We have introduced a space between to prevent syntax errors. Please ignore the space.
	NpmrcEnvironmentVariablesSecretVersion *string `field:"optional" json:"npmrcEnvironmentVariablesSecretVersion" yaml:"npmrcEnvironmentVariablesSecretVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataform_repository#project GoogleDataformRepository#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// A reference to the region.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataform_repository#region GoogleDataformRepository#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// The service account to run workflow invocations under.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataform_repository#service_account GoogleDataformRepository#service_account}
	ServiceAccount *string `field:"optional" json:"serviceAccount" yaml:"serviceAccount"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataform_repository#timeouts GoogleDataformRepository#timeouts}
	Timeouts *GoogleDataformRepositoryTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// workspace_compilation_overrides block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataform_repository#workspace_compilation_overrides GoogleDataformRepository#workspace_compilation_overrides}
	WorkspaceCompilationOverrides *GoogleDataformRepositoryWorkspaceCompilationOverrides `field:"optional" json:"workspaceCompilationOverrides" yaml:"workspaceCompilationOverrides"`
}

