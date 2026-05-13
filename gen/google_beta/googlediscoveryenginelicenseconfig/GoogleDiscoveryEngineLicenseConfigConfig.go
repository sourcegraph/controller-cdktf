package googlediscoveryenginelicenseconfig

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleDiscoveryEngineLicenseConfigConfig struct {
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
	// The unique id of the license config.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_license_config#license_config_id GoogleDiscoveryEngineLicenseConfig#license_config_id}
	LicenseConfigId *string `field:"required" json:"licenseConfigId" yaml:"licenseConfigId"`
	// Number of licenses purchased.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_license_config#license_count GoogleDiscoveryEngineLicenseConfig#license_count}
	LicenseCount *float64 `field:"required" json:"licenseCount" yaml:"licenseCount"`
	// The geographic location where the data store should reside. The value can only be one of "global", "us" and "eu".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_license_config#location GoogleDiscoveryEngineLicenseConfig#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// start_date block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_license_config#start_date GoogleDiscoveryEngineLicenseConfig#start_date}
	StartDate *GoogleDiscoveryEngineLicenseConfigStartDate `field:"required" json:"startDate" yaml:"startDate"`
	// Subscription term. Possible values: ["SUBSCRIPTION_TERM_UNSPECIFIED", "SUBSCRIPTION_TERM_ONE_MONTH", "SUBSCRIPTION_TERM_ONE_YEAR", "SUBSCRIPTION_TERM_THREE_YEARS", "SUBSCRIPTION_TERM_THREE_MONTHS", "SUBSCRIPTION_TERM_FOURTEEN_DAYS", "SUBSCRIPTION_TERM_CUSTOM"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_license_config#subscription_term GoogleDiscoveryEngineLicenseConfig#subscription_term}
	SubscriptionTerm *string `field:"required" json:"subscriptionTerm" yaml:"subscriptionTerm"`
	// Subscription tier information for the license config. Possible values: ["SUBSCRIPTION_TIER_UNSPECIFIED", "SUBSCRIPTION_TIER_SEARCH", "SUBSCRIPTION_TIER_SEARCH_AND_ASSISTANT", "SUBSCRIPTION_TIER_NOTEBOOK_LM", "SUBSCRIPTION_TIER_FRONTLINE_WORKER", "SUBSCRIPTION_TIER_AGENTSPACE_STARTER", "SUBSCRIPTION_TIER_AGENTSPACE_BUSINESS", "SUBSCRIPTION_TIER_ENTERPRISE", "SUBSCRIPTION_TIER_EDU", "SUBSCRIPTION_TIER_EDU_PRO"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_license_config#subscription_tier GoogleDiscoveryEngineLicenseConfig#subscription_tier}
	SubscriptionTier *string `field:"required" json:"subscriptionTier" yaml:"subscriptionTier"`
	// Whether the license config should be auto renewed when it reaches the end date.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_license_config#auto_renew GoogleDiscoveryEngineLicenseConfig#auto_renew}
	AutoRenew interface{} `field:"optional" json:"autoRenew" yaml:"autoRenew"`
	// end_date block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_license_config#end_date GoogleDiscoveryEngineLicenseConfig#end_date}
	EndDate *GoogleDiscoveryEngineLicenseConfigEndDate `field:"optional" json:"endDate" yaml:"endDate"`
	// Whether the license config is for free trial.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_license_config#free_trial GoogleDiscoveryEngineLicenseConfig#free_trial}
	FreeTrial interface{} `field:"optional" json:"freeTrial" yaml:"freeTrial"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_license_config#id GoogleDiscoveryEngineLicenseConfig#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_license_config#project GoogleDiscoveryEngineLicenseConfig#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_license_config#timeouts GoogleDiscoveryEngineLicenseConfig#timeouts}
	Timeouts *GoogleDiscoveryEngineLicenseConfigTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

