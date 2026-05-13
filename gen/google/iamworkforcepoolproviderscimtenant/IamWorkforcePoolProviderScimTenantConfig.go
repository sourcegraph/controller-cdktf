package iamworkforcepoolproviderscimtenant

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IamWorkforcePoolProviderScimTenantConfig struct {
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
	// The location for the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/iam_workforce_pool_provider_scim_tenant#location IamWorkforcePoolProviderScimTenant#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The ID of the provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/iam_workforce_pool_provider_scim_tenant#provider_id IamWorkforcePoolProviderScimTenant#provider_id}
	ProviderId *string `field:"required" json:"providerId" yaml:"providerId"`
	// The ID to use for the SCIM tenant, which becomes the final component of the resource name.
	//
	// This value must be 4-32 characters, and may contain the characters [a-z0-9-].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/iam_workforce_pool_provider_scim_tenant#scim_tenant_id IamWorkforcePoolProviderScimTenant#scim_tenant_id}
	ScimTenantId *string `field:"required" json:"scimTenantId" yaml:"scimTenantId"`
	// The ID of the workforce pool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/iam_workforce_pool_provider_scim_tenant#workforce_pool_id IamWorkforcePoolProviderScimTenant#workforce_pool_id}
	WorkforcePoolId *string `field:"required" json:"workforcePoolId" yaml:"workforcePoolId"`
	// Maps BYOID claims to SCIM claims. This is a required field for new SCIM Tenants being created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/iam_workforce_pool_provider_scim_tenant#claim_mapping IamWorkforcePoolProviderScimTenant#claim_mapping}
	ClaimMapping *map[string]*string `field:"optional" json:"claimMapping" yaml:"claimMapping"`
	// A user-specified description of the provider. Cannot exceed 256 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/iam_workforce_pool_provider_scim_tenant#description IamWorkforcePoolProviderScimTenant#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A user-specified display name for the scim tenant. Cannot exceed 32 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/iam_workforce_pool_provider_scim_tenant#display_name IamWorkforcePoolProviderScimTenant#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Deletes the SCIM tenant immediately. This operation cannot be undone.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/iam_workforce_pool_provider_scim_tenant#hard_delete IamWorkforcePoolProviderScimTenant#hard_delete}
	HardDelete interface{} `field:"optional" json:"hardDelete" yaml:"hardDelete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/iam_workforce_pool_provider_scim_tenant#id IamWorkforcePoolProviderScimTenant#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/iam_workforce_pool_provider_scim_tenant#timeouts IamWorkforcePoolProviderScimTenant#timeouts}
	Timeouts *IamWorkforcePoolProviderScimTenantTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

