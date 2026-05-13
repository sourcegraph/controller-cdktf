package computeorganizationsecuritypolicyassociation

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComputeOrganizationSecurityPolicyAssociationConfig struct {
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
	// The resource that the security policy is attached to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/compute_organization_security_policy_association#attachment_id ComputeOrganizationSecurityPolicyAssociation#attachment_id}
	AttachmentId *string `field:"required" json:"attachmentId" yaml:"attachmentId"`
	// The name for an association.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/compute_organization_security_policy_association#name ComputeOrganizationSecurityPolicyAssociation#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The security policy ID of the association.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/compute_organization_security_policy_association#policy_id ComputeOrganizationSecurityPolicyAssociation#policy_id}
	PolicyId *string `field:"required" json:"policyId" yaml:"policyId"`
	// A list of folders to exclude from the security policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/compute_organization_security_policy_association#excluded_folders ComputeOrganizationSecurityPolicyAssociation#excluded_folders}
	ExcludedFolders *[]*string `field:"optional" json:"excludedFolders" yaml:"excludedFolders"`
	// A list of projects to exclude from the security policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/compute_organization_security_policy_association#excluded_projects ComputeOrganizationSecurityPolicyAssociation#excluded_projects}
	ExcludedProjects *[]*string `field:"optional" json:"excludedProjects" yaml:"excludedProjects"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/compute_organization_security_policy_association#id ComputeOrganizationSecurityPolicyAssociation#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/compute_organization_security_policy_association#timeouts ComputeOrganizationSecurityPolicyAssociation#timeouts}
	Timeouts *ComputeOrganizationSecurityPolicyAssociationTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

