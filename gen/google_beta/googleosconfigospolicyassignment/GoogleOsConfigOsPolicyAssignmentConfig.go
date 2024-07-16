package googleosconfigospolicyassignment

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleOsConfigOsPolicyAssignmentConfig struct {
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
	// instance_filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_os_config_os_policy_assignment#instance_filter GoogleOsConfigOsPolicyAssignment#instance_filter}
	InstanceFilter *GoogleOsConfigOsPolicyAssignmentInstanceFilter `field:"required" json:"instanceFilter" yaml:"instanceFilter"`
	// The location for the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_os_config_os_policy_assignment#location GoogleOsConfigOsPolicyAssignment#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// Resource name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_os_config_os_policy_assignment#name GoogleOsConfigOsPolicyAssignment#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// os_policies block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_os_config_os_policy_assignment#os_policies GoogleOsConfigOsPolicyAssignment#os_policies}
	OsPolicies interface{} `field:"required" json:"osPolicies" yaml:"osPolicies"`
	// rollout block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_os_config_os_policy_assignment#rollout GoogleOsConfigOsPolicyAssignment#rollout}
	Rollout *GoogleOsConfigOsPolicyAssignmentRollout `field:"required" json:"rollout" yaml:"rollout"`
	// OS policy assignment description. Length of the description is limited to 1024 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_os_config_os_policy_assignment#description GoogleOsConfigOsPolicyAssignment#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_os_config_os_policy_assignment#id GoogleOsConfigOsPolicyAssignment#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The project for the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_os_config_os_policy_assignment#project GoogleOsConfigOsPolicyAssignment#project}
	Project *string `field:"optional" json:"project" yaml:"project"`
	// Set to true to skip awaiting rollout during resource creation and update.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_os_config_os_policy_assignment#skip_await_rollout GoogleOsConfigOsPolicyAssignment#skip_await_rollout}
	SkipAwaitRollout interface{} `field:"optional" json:"skipAwaitRollout" yaml:"skipAwaitRollout"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_os_config_os_policy_assignment#timeouts GoogleOsConfigOsPolicyAssignment#timeouts}
	Timeouts *GoogleOsConfigOsPolicyAssignmentTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

