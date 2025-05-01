package googleiamprojectspolicybinding

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleIamProjectsPolicyBindingConfig struct {
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
	// The location of the Policy Binding.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_iam_projects_policy_binding#location GoogleIamProjectsPolicyBinding#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// Required.
	//
	// Immutable. The resource name of the policy to be bound. The binding parent and policy must belong to the same Organization (or Project).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_iam_projects_policy_binding#policy GoogleIamProjectsPolicyBinding#policy}
	Policy *string `field:"required" json:"policy" yaml:"policy"`
	// The Policy Binding ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_iam_projects_policy_binding#policy_binding_id GoogleIamProjectsPolicyBinding#policy_binding_id}
	PolicyBindingId *string `field:"required" json:"policyBindingId" yaml:"policyBindingId"`
	// target block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_iam_projects_policy_binding#target GoogleIamProjectsPolicyBinding#target}
	Target *GoogleIamProjectsPolicyBindingTarget `field:"required" json:"target" yaml:"target"`
	// Optional. User defined annotations. See https://google.aip.dev/148#annotations for more details such as format and size limitations.
	//
	// **Note**: This field is non-authoritative, and will only manage the annotations present in your configuration.
	// Please refer to the field 'effective_annotations' for all of the annotations present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_iam_projects_policy_binding#annotations GoogleIamProjectsPolicyBinding#annotations}
	Annotations *map[string]*string `field:"optional" json:"annotations" yaml:"annotations"`
	// condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_iam_projects_policy_binding#condition GoogleIamProjectsPolicyBinding#condition}
	Condition *GoogleIamProjectsPolicyBindingCondition `field:"optional" json:"condition" yaml:"condition"`
	// Optional. The description of the policy binding. Must be less than or equal to 63 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_iam_projects_policy_binding#display_name GoogleIamProjectsPolicyBinding#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_iam_projects_policy_binding#id GoogleIamProjectsPolicyBinding#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Immutable.
	//
	// The kind of the policy to attach in this binding. This
	// field must be one of the following:  - Left empty (will be automatically set
	// to the policy kind) - The input policy kind   Possible values:  POLICY_KIND_UNSPECIFIED PRINCIPAL_ACCESS_BOUNDARY ACCESS
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_iam_projects_policy_binding#policy_kind GoogleIamProjectsPolicyBinding#policy_kind}
	PolicyKind *string `field:"optional" json:"policyKind" yaml:"policyKind"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_iam_projects_policy_binding#project GoogleIamProjectsPolicyBinding#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_iam_projects_policy_binding#timeouts GoogleIamProjectsPolicyBinding#timeouts}
	Timeouts *GoogleIamProjectsPolicyBindingTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

