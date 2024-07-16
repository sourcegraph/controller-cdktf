package googleidentityplatformconfig

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleIdentityPlatformConfigConfig struct {
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
	// List of domains authorized for OAuth redirects.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_identity_platform_config#authorized_domains GoogleIdentityPlatformConfig#authorized_domains}
	AuthorizedDomains *[]*string `field:"optional" json:"authorizedDomains" yaml:"authorizedDomains"`
	// Whether anonymous users will be auto-deleted after a period of 30 days.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_identity_platform_config#autodelete_anonymous_users GoogleIdentityPlatformConfig#autodelete_anonymous_users}
	AutodeleteAnonymousUsers interface{} `field:"optional" json:"autodeleteAnonymousUsers" yaml:"autodeleteAnonymousUsers"`
	// blocking_functions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_identity_platform_config#blocking_functions GoogleIdentityPlatformConfig#blocking_functions}
	BlockingFunctions *GoogleIdentityPlatformConfigBlockingFunctions `field:"optional" json:"blockingFunctions" yaml:"blockingFunctions"`
	// client block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_identity_platform_config#client GoogleIdentityPlatformConfig#client}
	Client *GoogleIdentityPlatformConfigClient `field:"optional" json:"client" yaml:"client"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_identity_platform_config#id GoogleIdentityPlatformConfig#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// mfa block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_identity_platform_config#mfa GoogleIdentityPlatformConfig#mfa}
	Mfa *GoogleIdentityPlatformConfigMfa `field:"optional" json:"mfa" yaml:"mfa"`
	// monitoring block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_identity_platform_config#monitoring GoogleIdentityPlatformConfig#monitoring}
	Monitoring *GoogleIdentityPlatformConfigMonitoring `field:"optional" json:"monitoring" yaml:"monitoring"`
	// multi_tenant block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_identity_platform_config#multi_tenant GoogleIdentityPlatformConfig#multi_tenant}
	MultiTenant *GoogleIdentityPlatformConfigMultiTenant `field:"optional" json:"multiTenant" yaml:"multiTenant"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_identity_platform_config#project GoogleIdentityPlatformConfig#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// quota block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_identity_platform_config#quota GoogleIdentityPlatformConfig#quota}
	Quota *GoogleIdentityPlatformConfigQuota `field:"optional" json:"quota" yaml:"quota"`
	// sign_in block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_identity_platform_config#sign_in GoogleIdentityPlatformConfig#sign_in}
	SignIn *GoogleIdentityPlatformConfigSignIn `field:"optional" json:"signIn" yaml:"signIn"`
	// sms_region_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_identity_platform_config#sms_region_config GoogleIdentityPlatformConfig#sms_region_config}
	SmsRegionConfig *GoogleIdentityPlatformConfigSmsRegionConfig `field:"optional" json:"smsRegionConfig" yaml:"smsRegionConfig"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_identity_platform_config#timeouts GoogleIdentityPlatformConfig#timeouts}
	Timeouts *GoogleIdentityPlatformConfigTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

