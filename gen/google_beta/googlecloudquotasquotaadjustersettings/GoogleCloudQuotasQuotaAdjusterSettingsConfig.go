package googlecloudquotasquotaadjustersettings

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleCloudQuotasQuotaAdjusterSettingsConfig struct {
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
	// Required. The configured value of the enablement at the given resource. Possible values: ["ENABLED", "DISABLED"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_cloud_quotas_quota_adjuster_settings#enablement GoogleCloudQuotasQuotaAdjusterSettings#enablement}
	Enablement *string `field:"required" json:"enablement" yaml:"enablement"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_cloud_quotas_quota_adjuster_settings#id GoogleCloudQuotasQuotaAdjusterSettings#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The parent of the quota preference. Allowed parent format is "projects/[project-id / number]".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_cloud_quotas_quota_adjuster_settings#parent GoogleCloudQuotasQuotaAdjusterSettings#parent}
	Parent *string `field:"optional" json:"parent" yaml:"parent"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_cloud_quotas_quota_adjuster_settings#timeouts GoogleCloudQuotasQuotaAdjusterSettings#timeouts}
	Timeouts *GoogleCloudQuotasQuotaAdjusterSettingsTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

