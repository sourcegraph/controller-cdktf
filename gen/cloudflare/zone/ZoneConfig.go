package zone

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ZoneConfig struct {
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
	// Account ID to manage the zone resource in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/zone#account_id Zone#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// The DNS zone name which will be added. **Modifying this attribute will force creation of a new resource.**.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/zone#zone Zone#zone}
	Zone *string `field:"required" json:"zone" yaml:"zone"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/zone#id Zone#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Whether to scan for DNS records on creation. Ignored after zone is created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/zone#jump_start Zone#jump_start}
	JumpStart interface{} `field:"optional" json:"jumpStart" yaml:"jumpStart"`
	// Whether this zone is paused (traffic bypasses Cloudflare). Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/zone#paused Zone#paused}
	Paused interface{} `field:"optional" json:"paused" yaml:"paused"`
	// The name of the commercial plan to apply to the zone.
	//
	// Available values: `free`, `lite`, `pro`, `pro_plus`, `business`, `enterprise`, `partners_free`, `partners_pro`, `partners_business`, `partners_enterprise`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/zone#plan Zone#plan}
	Plan *string `field:"optional" json:"plan" yaml:"plan"`
	// A full zone implies that DNS is hosted with Cloudflare.
	//
	// A partial zone is typically a partner-hosted zone or a CNAME setup. Available values: `full`, `partial`. Defaults to `full`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/zone#type Zone#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

