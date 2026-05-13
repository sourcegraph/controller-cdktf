package iamworkloadidentitypoolmanagedidentity

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IamWorkloadIdentityPoolManagedIdentityConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/iam_workload_identity_pool_managed_identity#workload_identity_pool_id IamWorkloadIdentityPoolManagedIdentity#workload_identity_pool_id}
	WorkloadIdentityPoolId *string `field:"required" json:"workloadIdentityPoolId" yaml:"workloadIdentityPoolId"`
	// The ID to use for the managed identity.
	//
	// This value must:
	// * contain at most 63 characters
	// * contain only lowercase alphanumeric characters or '-'
	// * start with an alphanumeric character
	// * end with an alphanumeric character
	//
	//
	// The prefix 'gcp-' will be reserved for future uses.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/iam_workload_identity_pool_managed_identity#workload_identity_pool_managed_identity_id IamWorkloadIdentityPoolManagedIdentity#workload_identity_pool_managed_identity_id}
	WorkloadIdentityPoolManagedIdentityId *string `field:"required" json:"workloadIdentityPoolManagedIdentityId" yaml:"workloadIdentityPoolManagedIdentityId"`
	// The ID to use for the namespace.
	//
	// This value must:
	// * contain at most 63 characters
	// * contain only lowercase alphanumeric characters or '-'
	// * start with an alphanumeric character
	// * end with an alphanumeric character
	//
	//
	// The prefix 'gcp-' will be reserved for future uses.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/iam_workload_identity_pool_managed_identity#workload_identity_pool_namespace_id IamWorkloadIdentityPoolManagedIdentity#workload_identity_pool_namespace_id}
	WorkloadIdentityPoolNamespaceId *string `field:"required" json:"workloadIdentityPoolNamespaceId" yaml:"workloadIdentityPoolNamespaceId"`
	// attestation_rules block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/iam_workload_identity_pool_managed_identity#attestation_rules IamWorkloadIdentityPoolManagedIdentity#attestation_rules}
	AttestationRules interface{} `field:"optional" json:"attestationRules" yaml:"attestationRules"`
	// A description of the managed identity. Cannot exceed 256 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/iam_workload_identity_pool_managed_identity#description IamWorkloadIdentityPoolManagedIdentity#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Whether the managed identity is disabled.
	//
	// If disabled, credentials may no longer be issued for
	// the identity, however existing credentials will still be accepted until they expire.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/iam_workload_identity_pool_managed_identity#disabled IamWorkloadIdentityPoolManagedIdentity#disabled}
	Disabled interface{} `field:"optional" json:"disabled" yaml:"disabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/iam_workload_identity_pool_managed_identity#id IamWorkloadIdentityPoolManagedIdentity#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/iam_workload_identity_pool_managed_identity#project IamWorkloadIdentityPoolManagedIdentity#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/iam_workload_identity_pool_managed_identity#timeouts IamWorkloadIdentityPoolManagedIdentity#timeouts}
	Timeouts *IamWorkloadIdentityPoolManagedIdentityTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

