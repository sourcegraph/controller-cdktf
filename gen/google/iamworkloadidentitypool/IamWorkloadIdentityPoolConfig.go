package iamworkloadidentitypool

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IamWorkloadIdentityPoolConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/iam_workload_identity_pool#workload_identity_pool_id IamWorkloadIdentityPool#workload_identity_pool_id}
	WorkloadIdentityPoolId *string `field:"required" json:"workloadIdentityPoolId" yaml:"workloadIdentityPoolId"`
	// attestation_rules block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/iam_workload_identity_pool#attestation_rules IamWorkloadIdentityPool#attestation_rules}
	AttestationRules interface{} `field:"optional" json:"attestationRules" yaml:"attestationRules"`
	// A description of the pool. Cannot exceed 256 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/iam_workload_identity_pool#description IamWorkloadIdentityPool#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Whether the pool is disabled.
	//
	// You cannot use a disabled pool to exchange tokens, or use
	// existing tokens to access resources. If the pool is re-enabled, existing tokens grant
	// access again.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/iam_workload_identity_pool#disabled IamWorkloadIdentityPool#disabled}
	Disabled interface{} `field:"optional" json:"disabled" yaml:"disabled"`
	// A display name for the pool. Cannot exceed 32 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/iam_workload_identity_pool#display_name IamWorkloadIdentityPool#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/iam_workload_identity_pool#id IamWorkloadIdentityPool#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// inline_certificate_issuance_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/iam_workload_identity_pool#inline_certificate_issuance_config IamWorkloadIdentityPool#inline_certificate_issuance_config}
	InlineCertificateIssuanceConfig *IamWorkloadIdentityPoolInlineCertificateIssuanceConfig `field:"optional" json:"inlineCertificateIssuanceConfig" yaml:"inlineCertificateIssuanceConfig"`
	// inline_trust_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/iam_workload_identity_pool#inline_trust_config IamWorkloadIdentityPool#inline_trust_config}
	InlineTrustConfig *IamWorkloadIdentityPoolInlineTrustConfig `field:"optional" json:"inlineTrustConfig" yaml:"inlineTrustConfig"`
	// The mode for the pool is operating in.
	//
	// Pools with an unspecified mode will operate as if they
	// are in 'FEDERATION_ONLY' mode.
	//
	//
	// ~> **Note** This field cannot be changed after the Workload Identity Pool is created. While
	// 'terraform plan' may show an update if you change this field's value, 'terraform apply'
	// **will fail with an API error** (such as 'Error 400: Attempted to update an immutable field.').
	// To specify a different 'mode', please create a new Workload Identity Pool resource.
	//
	// * 'FEDERATION_ONLY': Pools can only be used for federating external workload identities into
	// Google Cloud. Unless otherwise noted, no structure or format constraints are applied to
	// workload identities in a 'FEDERATION_ONLY' mode pool, and you may not create any resources
	// within the pool besides providers.
	// * 'TRUST_DOMAIN': Pools can be used to assign identities to Google Cloud workloads. All
	// identities within a 'TRUST_DOMAIN' mode pool must consist of a single namespace and individual
	// workload identifier. The subject identifier for all identities must conform to the following
	// format: 'ns/<namespace>/sa/<workload_identifier>'.
	// 'google_iam_workload_identity_pool_provider's cannot be created within 'TRUST_DOMAIN'
	// mode pools.
	// * 'SYSTEM_TRUST_DOMAIN': Pools are managed by Google Cloud services. Neither
	// 'google_iam_workload_identity_pool_namespace's nor 'google_iam_workload_identity_pool_provider's
	// can be created within 'SYSTEM_TRUST_DOMAIN' mode pools. All identities within a
	// 'SYSTEM_TRUST_DOMAIN' mode pool are in one of the following formats:
	//
	//     * 'spiffe://<trust-domain>/ns/<kubernetes-namespace>/sa/<kubernetes-service-account>'
	//     * 'spiffe://<trust-domain>/resources/<resource-scope>/<resource-name>' Possible values: ["FEDERATION_ONLY", "TRUST_DOMAIN", "SYSTEM_TRUST_DOMAIN"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/iam_workload_identity_pool#mode IamWorkloadIdentityPool#mode}
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/iam_workload_identity_pool#project IamWorkloadIdentityPool#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/iam_workload_identity_pool#timeouts IamWorkloadIdentityPool#timeouts}
	Timeouts *IamWorkloadIdentityPoolTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

