package googleiamworkforcepoolproviderscimtoken

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleIamWorkforcePoolProviderScimTokenConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_iam_workforce_pool_provider_scim_token#location GoogleIamWorkforcePoolProviderScimToken#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The ID of the Provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_iam_workforce_pool_provider_scim_token#provider_id GoogleIamWorkforcePoolProviderScimToken#provider_id}
	ProviderId *string `field:"required" json:"providerId" yaml:"providerId"`
	// The ID of the SCIM Tenant.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_iam_workforce_pool_provider_scim_token#scim_tenant_id GoogleIamWorkforcePoolProviderScimToken#scim_tenant_id}
	ScimTenantId *string `field:"required" json:"scimTenantId" yaml:"scimTenantId"`
	// The ID to use for the SCIM Token, which becomes the final component of the resource name.
	//
	// This value should be 4-32 characters and follow the pattern: '([a-z]([a-z0-9\\-]{2,30}[a-z0-9]))'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_iam_workforce_pool_provider_scim_token#scim_token_id GoogleIamWorkforcePoolProviderScimToken#scim_token_id}
	ScimTokenId *string `field:"required" json:"scimTokenId" yaml:"scimTokenId"`
	// The ID of the Workforce Pool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_iam_workforce_pool_provider_scim_token#workforce_pool_id GoogleIamWorkforcePoolProviderScimToken#workforce_pool_id}
	WorkforcePoolId *string `field:"required" json:"workforcePoolId" yaml:"workforcePoolId"`
	// A user-specified display name for the scim token. Cannot exceed 32 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_iam_workforce_pool_provider_scim_token#display_name GoogleIamWorkforcePoolProviderScimToken#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_iam_workforce_pool_provider_scim_token#id GoogleIamWorkforcePoolProviderScimToken#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_iam_workforce_pool_provider_scim_token#timeouts GoogleIamWorkforcePoolProviderScimToken#timeouts}
	Timeouts *GoogleIamWorkforcePoolProviderScimTokenTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

