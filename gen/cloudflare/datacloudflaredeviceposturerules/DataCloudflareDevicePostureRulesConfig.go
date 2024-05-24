package datacloudflaredeviceposturerules

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataCloudflareDevicePostureRulesConfig struct {
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
	// The account identifier to target for the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.20.0/docs/data-sources/device_posture_rules#account_id DataCloudflareDevicePostureRules#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.20.0/docs/data-sources/device_posture_rules#id DataCloudflareDevicePostureRules#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Name of the Device Posture Rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.20.0/docs/data-sources/device_posture_rules#name DataCloudflareDevicePostureRules#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The device posture rule type.
	//
	// Available values: `serial_number`, `file`, `application`, `gateway`, `warp`, `domain_joined`, `os_version`, `disk_encryption`, `firewall`, `client_certificate`, `workspace_one`, `unique_client_id`, `crowdstrike_s2s`, `sentinelone`, `kolide`, `tanium_s2s`, `intune`, `sentinelone_s2s`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.20.0/docs/data-sources/device_posture_rules#type DataCloudflareDevicePostureRules#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

