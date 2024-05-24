package regionalhostname

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type RegionalHostnameConfig struct {
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
	// The hostname to regionalize.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.20.0/docs/resources/regional_hostname#hostname RegionalHostname#hostname}
	Hostname *string `field:"required" json:"hostname" yaml:"hostname"`
	// The region key. See [the full region list](https://developers.cloudflare.com/data-localization/regional-services/get-started/).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.20.0/docs/resources/regional_hostname#region_key RegionalHostname#region_key}
	RegionKey *string `field:"required" json:"regionKey" yaml:"regionKey"`
	// The zone identifier to target for the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.20.0/docs/resources/regional_hostname#zone_id RegionalHostname#zone_id}
	ZoneId *string `field:"required" json:"zoneId" yaml:"zoneId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.20.0/docs/resources/regional_hostname#id RegionalHostname#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.20.0/docs/resources/regional_hostname#timeouts RegionalHostname#timeouts}
	Timeouts *RegionalHostnameTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

