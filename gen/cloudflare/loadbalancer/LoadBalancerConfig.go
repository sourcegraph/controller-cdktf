package loadbalancer

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LoadBalancerConfig struct {
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
	// A list of pool IDs ordered by their failover priority. Used whenever [`pop_pools`](#pop_pools)/[`country_pools`](#country_pools)/[`region_pools`](#region_pools) are not defined.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/load_balancer#default_pool_ids LoadBalancer#default_pool_ids}
	DefaultPoolIds *[]*string `field:"required" json:"defaultPoolIds" yaml:"defaultPoolIds"`
	// The pool ID to use when all other pools are detected as unhealthy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/load_balancer#fallback_pool_id LoadBalancer#fallback_pool_id}
	FallbackPoolId *string `field:"required" json:"fallbackPoolId" yaml:"fallbackPoolId"`
	// The DNS hostname to associate with your load balancer.
	//
	// If this hostname already exists as a DNS record in Cloudflare's DNS, the load balancer will take precedence and the DNS record will not be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/load_balancer#name LoadBalancer#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The zone ID to add the load balancer to. **Modifying this attribute will force creation of a new resource.**.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/load_balancer#zone_id LoadBalancer#zone_id}
	ZoneId *string `field:"required" json:"zoneId" yaml:"zoneId"`
	// adaptive_routing block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/load_balancer#adaptive_routing LoadBalancer#adaptive_routing}
	AdaptiveRouting interface{} `field:"optional" json:"adaptiveRouting" yaml:"adaptiveRouting"`
	// country_pools block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/load_balancer#country_pools LoadBalancer#country_pools}
	CountryPools interface{} `field:"optional" json:"countryPools" yaml:"countryPools"`
	// Free text description.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/load_balancer#description LoadBalancer#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Enable or disable the load balancer. Defaults to `true`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/load_balancer#enabled LoadBalancer#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/load_balancer#id LoadBalancer#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// location_strategy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/load_balancer#location_strategy LoadBalancer#location_strategy}
	LocationStrategy interface{} `field:"optional" json:"locationStrategy" yaml:"locationStrategy"`
	// pop_pools block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/load_balancer#pop_pools LoadBalancer#pop_pools}
	PopPools interface{} `field:"optional" json:"popPools" yaml:"popPools"`
	// Whether the hostname gets Cloudflare's origin protection. Defaults to `false`. Conflicts with `ttl`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/load_balancer#proxied LoadBalancer#proxied}
	Proxied interface{} `field:"optional" json:"proxied" yaml:"proxied"`
	// random_steering block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/load_balancer#random_steering LoadBalancer#random_steering}
	RandomSteering interface{} `field:"optional" json:"randomSteering" yaml:"randomSteering"`
	// region_pools block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/load_balancer#region_pools LoadBalancer#region_pools}
	RegionPools interface{} `field:"optional" json:"regionPools" yaml:"regionPools"`
	// rules block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/load_balancer#rules LoadBalancer#rules}
	Rules interface{} `field:"optional" json:"rules" yaml:"rules"`
	// Specifies the type of session affinity the load balancer should use unless specified as `none` or `""` (default).
	//
	// With value `cookie`, on the first request to a proxied load balancer, a cookie is generated, encoding information of which origin the request will be forwarded to. Subsequent requests, by the same client to the same load balancer, will be sent to the origin server the cookie encodes, for the duration of the cookie and as long as the origin server remains healthy. If the cookie has expired or the origin server is unhealthy then a new origin server is calculated and used. Value `ip_cookie` behaves the same as `cookie` except the initial origin selection is stable and based on the client's IP address. Available values: `""`, `none`, `cookie`, `ip_cookie`. Defaults to `none`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/load_balancer#session_affinity LoadBalancer#session_affinity}
	SessionAffinity *string `field:"optional" json:"sessionAffinity" yaml:"sessionAffinity"`
	// session_affinity_attributes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/load_balancer#session_affinity_attributes LoadBalancer#session_affinity_attributes}
	SessionAffinityAttributes interface{} `field:"optional" json:"sessionAffinityAttributes" yaml:"sessionAffinityAttributes"`
	// Time, in seconds, until this load balancer's session affinity cookie expires after being created.
	//
	// This parameter is ignored unless a supported session affinity policy is set. The current default of `82800` (23 hours) will be used unless [`session_affinity_ttl`](#session_affinity_ttl) is explicitly set. Once the expiry time has been reached, subsequent requests may get sent to a different origin server. Valid values are between `1800` and `604800`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/load_balancer#session_affinity_ttl LoadBalancer#session_affinity_ttl}
	SessionAffinityTtl *float64 `field:"optional" json:"sessionAffinityTtl" yaml:"sessionAffinityTtl"`
	// The method the load balancer uses to determine the route to your origin.
	//
	// Value `off` uses [`default_pool_ids`](#default_pool_ids). Value `geo` uses [`pop_pools`](#pop_pools)/[`country_pools`](#country_pools)/[`region_pools`](#region_pools). For non-proxied requests, the [`country`](#country) for [`country_pools`](#country_pools) is determined by [`location_strategy`](#location_strategy). Value `random` selects a pool randomly. Value `dynamic_latency` uses round trip time to select the closest pool in [`default_pool_ids`](#default_pool_ids) (requires pool health checks). Value `proximity` uses the pools' latitude and longitude to select the closest pool using the Cloudflare PoP location for proxied requests or the location determined by [`location_strategy`](#location_strategy) for non-proxied requests. Value `""` maps to `geo` if you use [`pop_pools`](#pop_pools)/[`country_pools`](#country_pools)/[`region_pools`](#region_pools) otherwise `off`. Available values: `off`, `geo`, `dynamic_latency`, `random`, `proximity`, `""` Defaults to `""`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/load_balancer#steering_policy LoadBalancer#steering_policy}
	SteeringPolicy *string `field:"optional" json:"steeringPolicy" yaml:"steeringPolicy"`
	// Time to live (TTL) of the DNS entry for the IP address returned by this load balancer.
	//
	// This cannot be set for proxied load balancers. Defaults to `30`. Conflicts with `proxied`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/load_balancer#ttl LoadBalancer#ttl}
	Ttl *float64 `field:"optional" json:"ttl" yaml:"ttl"`
}

