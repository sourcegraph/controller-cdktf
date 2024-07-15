package googlecomputeorganizationsecuritypolicy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleComputeOrganizationSecurityPolicyConfig struct {
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
	// A textual name of the security policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_organization_security_policy#display_name GoogleComputeOrganizationSecurityPolicy#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// The parent of this OrganizationSecurityPolicy in the Cloud Resource Hierarchy. Format: organizations/{organization_id} or folders/{folder_id}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_organization_security_policy#parent GoogleComputeOrganizationSecurityPolicy#parent}
	Parent *string `field:"required" json:"parent" yaml:"parent"`
	// A textual description for the organization security policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_organization_security_policy#description GoogleComputeOrganizationSecurityPolicy#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_organization_security_policy#id GoogleComputeOrganizationSecurityPolicy#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_organization_security_policy#timeouts GoogleComputeOrganizationSecurityPolicy#timeouts}
	Timeouts *GoogleComputeOrganizationSecurityPolicyTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// The type indicates the intended use of the security policy.
	//
	// For organization security policies, the only supported type
	// is "FIREWALL". Default value: "FIREWALL" Possible values: ["FIREWALL"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_organization_security_policy#type GoogleComputeOrganizationSecurityPolicy#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

