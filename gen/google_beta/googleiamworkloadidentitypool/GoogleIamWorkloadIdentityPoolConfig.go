package googleiamworkloadidentitypool

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleIamWorkloadIdentityPoolConfig struct {
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
	// The ID to use for the pool, which becomes the final component of the resource name.
	//
	// This
	// value should be 4-32 characters, and may contain the characters [a-z0-9-]. The prefix
	// 'gcp-' is reserved for use by Google, and may not be specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_iam_workload_identity_pool#workload_identity_pool_id GoogleIamWorkloadIdentityPool#workload_identity_pool_id}
	WorkloadIdentityPoolId *string `field:"required" json:"workloadIdentityPoolId" yaml:"workloadIdentityPoolId"`
	// A description of the pool. Cannot exceed 256 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_iam_workload_identity_pool#description GoogleIamWorkloadIdentityPool#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Whether the pool is disabled.
	//
	// You cannot use a disabled pool to exchange tokens, or use
	// existing tokens to access resources. If the pool is re-enabled, existing tokens grant
	// access again.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_iam_workload_identity_pool#disabled GoogleIamWorkloadIdentityPool#disabled}
	Disabled interface{} `field:"optional" json:"disabled" yaml:"disabled"`
	// A display name for the pool. Cannot exceed 32 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_iam_workload_identity_pool#display_name GoogleIamWorkloadIdentityPool#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_iam_workload_identity_pool#id GoogleIamWorkloadIdentityPool#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_iam_workload_identity_pool#project GoogleIamWorkloadIdentityPool#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_iam_workload_identity_pool#timeouts GoogleIamWorkloadIdentityPool#timeouts}
	Timeouts *GoogleIamWorkloadIdentityPoolTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

