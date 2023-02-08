package loadbalancerpool

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LoadBalancerPoolConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/load_balancer_pool#name LoadBalancerPool#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// origins block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/load_balancer_pool#origins LoadBalancerPool#origins}
	Origins interface{} `field:"required" json:"origins" yaml:"origins"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/load_balancer_pool#check_regions LoadBalancerPool#check_regions}.
	CheckRegions *[]*string `field:"optional" json:"checkRegions" yaml:"checkRegions"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/load_balancer_pool#description LoadBalancerPool#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/load_balancer_pool#enabled LoadBalancerPool#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/load_balancer_pool#id LoadBalancerPool#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/load_balancer_pool#latitude LoadBalancerPool#latitude}.
	Latitude *float64 `field:"optional" json:"latitude" yaml:"latitude"`
	// load_shedding block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/load_balancer_pool#load_shedding LoadBalancerPool#load_shedding}
	LoadShedding interface{} `field:"optional" json:"loadShedding" yaml:"loadShedding"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/load_balancer_pool#longitude LoadBalancerPool#longitude}.
	Longitude *float64 `field:"optional" json:"longitude" yaml:"longitude"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/load_balancer_pool#minimum_origins LoadBalancerPool#minimum_origins}.
	MinimumOrigins *float64 `field:"optional" json:"minimumOrigins" yaml:"minimumOrigins"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/load_balancer_pool#monitor LoadBalancerPool#monitor}.
	Monitor *string `field:"optional" json:"monitor" yaml:"monitor"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/load_balancer_pool#notification_email LoadBalancerPool#notification_email}.
	NotificationEmail *string `field:"optional" json:"notificationEmail" yaml:"notificationEmail"`
	// origin_steering block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/load_balancer_pool#origin_steering LoadBalancerPool#origin_steering}
	OriginSteering interface{} `field:"optional" json:"originSteering" yaml:"originSteering"`
}

