package googlesccmanagementorganizationsecurityhealthanalyticscustommodule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleSccManagementOrganizationSecurityHealthAnalyticsCustomModuleConfig struct {
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
	// Numerical ID of the parent organization.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_scc_management_organization_security_health_analytics_custom_module#organization GoogleSccManagementOrganizationSecurityHealthAnalyticsCustomModule#organization}
	Organization *string `field:"required" json:"organization" yaml:"organization"`
	// custom_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_scc_management_organization_security_health_analytics_custom_module#custom_config GoogleSccManagementOrganizationSecurityHealthAnalyticsCustomModule#custom_config}
	CustomConfig *GoogleSccManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfig `field:"optional" json:"customConfig" yaml:"customConfig"`
	// The display name of the Security Health Analytics custom module.
	//
	// This
	// display name becomes the finding category for all findings that are
	// returned by this custom module. The display name must be between 1 and
	// 128 characters, start with a lowercase letter, and contain alphanumeric
	// characters or underscores only.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_scc_management_organization_security_health_analytics_custom_module#display_name GoogleSccManagementOrganizationSecurityHealthAnalyticsCustomModule#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// The enablement state of the custom module. Possible values: ["ENABLED", "DISABLED"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_scc_management_organization_security_health_analytics_custom_module#enablement_state GoogleSccManagementOrganizationSecurityHealthAnalyticsCustomModule#enablement_state}
	EnablementState *string `field:"optional" json:"enablementState" yaml:"enablementState"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_scc_management_organization_security_health_analytics_custom_module#id GoogleSccManagementOrganizationSecurityHealthAnalyticsCustomModule#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Location ID of the parent organization. If not provided, 'global' will be used as the default location.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_scc_management_organization_security_health_analytics_custom_module#location GoogleSccManagementOrganizationSecurityHealthAnalyticsCustomModule#location}
	Location *string `field:"optional" json:"location" yaml:"location"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_scc_management_organization_security_health_analytics_custom_module#timeouts GoogleSccManagementOrganizationSecurityHealthAnalyticsCustomModule#timeouts}
	Timeouts *GoogleSccManagementOrganizationSecurityHealthAnalyticsCustomModuleTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

