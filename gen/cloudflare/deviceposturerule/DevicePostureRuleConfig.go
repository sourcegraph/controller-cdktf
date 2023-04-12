package deviceposturerule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DevicePostureRuleConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/device_posture_rule#account_id DevicePostureRule#account_id}.
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/device_posture_rule#type DevicePostureRule#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/device_posture_rule#description DevicePostureRule#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/device_posture_rule#expiration DevicePostureRule#expiration}.
	Expiration *string `field:"optional" json:"expiration" yaml:"expiration"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/device_posture_rule#id DevicePostureRule#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// input block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/device_posture_rule#input DevicePostureRule#input}
	Input interface{} `field:"optional" json:"input" yaml:"input"`
	// match block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/device_posture_rule#match DevicePostureRule#match}
	Match interface{} `field:"optional" json:"match" yaml:"match"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/device_posture_rule#name DevicePostureRule#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/device_posture_rule#schedule DevicePostureRule#schedule}.
	Schedule *string `field:"optional" json:"schedule" yaml:"schedule"`
}

