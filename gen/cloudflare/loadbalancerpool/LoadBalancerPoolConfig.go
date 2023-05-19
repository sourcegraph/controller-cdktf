package loadbalancerpool

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LoadBalancerPoolConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/load_balancer_pool#account_id LoadBalancerPool#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// A short name (tag) for the pool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/load_balancer_pool#name LoadBalancerPool#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// origins block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/load_balancer_pool#origins LoadBalancerPool#origins}
	Origins interface{} `field:"required" json:"origins" yaml:"origins"`
	// A list of regions (specified by region code) from which to run health checks.
	//
	// Empty means every Cloudflare data center (the default), but requires an Enterprise plan. Region codes can be found [here](https://developers.cloudflare.com/load-balancing/reference/region-mapping-api).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/load_balancer_pool#check_regions LoadBalancerPool#check_regions}
	CheckRegions *[]*string `field:"optional" json:"checkRegions" yaml:"checkRegions"`
	// Free text description.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/load_balancer_pool#description LoadBalancerPool#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Whether to enable (the default) this pool.
	//
	// Disabled pools will not receive traffic and are excluded from health checks. Disabling a pool will cause any load balancers using it to failover to the next pool (if any). Defaults to `true`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/load_balancer_pool#enabled LoadBalancerPool#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/load_balancer_pool#id LoadBalancerPool#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The latitude this pool is physically located at; used for proximity steering.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/load_balancer_pool#latitude LoadBalancerPool#latitude}
	Latitude *float64 `field:"optional" json:"latitude" yaml:"latitude"`
	// load_shedding block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/load_balancer_pool#load_shedding LoadBalancerPool#load_shedding}
	LoadShedding interface{} `field:"optional" json:"loadShedding" yaml:"loadShedding"`
	// The longitude this pool is physically located at; used for proximity steering.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/load_balancer_pool#longitude LoadBalancerPool#longitude}
	Longitude *float64 `field:"optional" json:"longitude" yaml:"longitude"`
	// The minimum number of origins that must be healthy for this pool to serve traffic.
	//
	// If the number of healthy origins falls below this number, the pool will be marked unhealthy and we will failover to the next available pool. Defaults to `1`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/load_balancer_pool#minimum_origins LoadBalancerPool#minimum_origins}
	MinimumOrigins *float64 `field:"optional" json:"minimumOrigins" yaml:"minimumOrigins"`
	// The ID of the Monitor to use for health checking origins within this pool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/load_balancer_pool#monitor LoadBalancerPool#monitor}
	Monitor *string `field:"optional" json:"monitor" yaml:"monitor"`
	// The email address to send health status notifications to.
	//
	// This can be an individual mailbox or a mailing list. Multiple emails can be supplied as a comma delimited list.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/load_balancer_pool#notification_email LoadBalancerPool#notification_email}
	NotificationEmail *string `field:"optional" json:"notificationEmail" yaml:"notificationEmail"`
	// origin_steering block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/load_balancer_pool#origin_steering LoadBalancerPool#origin_steering}
	OriginSteering interface{} `field:"optional" json:"originSteering" yaml:"originSteering"`
}

