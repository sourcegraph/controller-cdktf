package loadbalancer

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LoadBalancerConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/load_balancer#default_pool_ids LoadBalancer#default_pool_ids}.
	DefaultPoolIds *[]*string `field:"required" json:"defaultPoolIds" yaml:"defaultPoolIds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/load_balancer#fallback_pool_id LoadBalancer#fallback_pool_id}.
	FallbackPoolId *string `field:"required" json:"fallbackPoolId" yaml:"fallbackPoolId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/load_balancer#name LoadBalancer#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/load_balancer#zone_id LoadBalancer#zone_id}.
	ZoneId *string `field:"required" json:"zoneId" yaml:"zoneId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/load_balancer#description LoadBalancer#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/load_balancer#enabled LoadBalancer#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/load_balancer#id LoadBalancer#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// pop_pools block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/load_balancer#pop_pools LoadBalancer#pop_pools}
	PopPools interface{} `field:"optional" json:"popPools" yaml:"popPools"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/load_balancer#proxied LoadBalancer#proxied}.
	Proxied interface{} `field:"optional" json:"proxied" yaml:"proxied"`
	// region_pools block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/load_balancer#region_pools LoadBalancer#region_pools}
	RegionPools interface{} `field:"optional" json:"regionPools" yaml:"regionPools"`
	// rules block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/load_balancer#rules LoadBalancer#rules}
	Rules interface{} `field:"optional" json:"rules" yaml:"rules"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/load_balancer#session_affinity LoadBalancer#session_affinity}.
	SessionAffinity *string `field:"optional" json:"sessionAffinity" yaml:"sessionAffinity"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/load_balancer#session_affinity_attributes LoadBalancer#session_affinity_attributes}.
	SessionAffinityAttributes *map[string]*string `field:"optional" json:"sessionAffinityAttributes" yaml:"sessionAffinityAttributes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/load_balancer#session_affinity_ttl LoadBalancer#session_affinity_ttl}.
	SessionAffinityTtl *float64 `field:"optional" json:"sessionAffinityTtl" yaml:"sessionAffinityTtl"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/load_balancer#steering_policy LoadBalancer#steering_policy}.
	SteeringPolicy *string `field:"optional" json:"steeringPolicy" yaml:"steeringPolicy"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/load_balancer#ttl LoadBalancer#ttl}.
	Ttl *float64 `field:"optional" json:"ttl" yaml:"ttl"`
}

