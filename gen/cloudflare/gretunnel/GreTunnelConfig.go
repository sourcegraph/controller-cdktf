package gretunnel

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GreTunnelConfig struct {
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
	// The IP address assigned to the Cloudflare side of the GRE tunnel.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/gre_tunnel#cloudflare_gre_endpoint GreTunnel#cloudflare_gre_endpoint}
	CloudflareGreEndpoint *string `field:"required" json:"cloudflareGreEndpoint" yaml:"cloudflareGreEndpoint"`
	// The IP address assigned to the customer side of the GRE tunnel.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/gre_tunnel#customer_gre_endpoint GreTunnel#customer_gre_endpoint}
	CustomerGreEndpoint *string `field:"required" json:"customerGreEndpoint" yaml:"customerGreEndpoint"`
	// 31-bit prefix (/31 in CIDR notation) supporting 2 hosts, one for each side of the tunnel.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/gre_tunnel#interface_address GreTunnel#interface_address}
	InterfaceAddress *string `field:"required" json:"interfaceAddress" yaml:"interfaceAddress"`
	// Name of the GRE tunnel.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/gre_tunnel#name GreTunnel#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The account identifier to target for the resource. **Modifying this attribute will force creation of a new resource.**.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/gre_tunnel#account_id GreTunnel#account_id}
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
	// Description of the GRE tunnel intent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/gre_tunnel#description GreTunnel#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Specifies if ICMP tunnel health checks are enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/gre_tunnel#health_check_enabled GreTunnel#health_check_enabled}
	HealthCheckEnabled interface{} `field:"optional" json:"healthCheckEnabled" yaml:"healthCheckEnabled"`
	// The IP address of the customer endpoint that will receive tunnel health checks.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/gre_tunnel#health_check_target GreTunnel#health_check_target}
	HealthCheckTarget *string `field:"optional" json:"healthCheckTarget" yaml:"healthCheckTarget"`
	// Specifies the ICMP echo type for the health check. Available values: `request`, `reply`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/gre_tunnel#health_check_type GreTunnel#health_check_type}
	HealthCheckType *string `field:"optional" json:"healthCheckType" yaml:"healthCheckType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/gre_tunnel#id GreTunnel#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Maximum Transmission Unit (MTU) in bytes for the GRE tunnel.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/gre_tunnel#mtu GreTunnel#mtu}
	Mtu *float64 `field:"optional" json:"mtu" yaml:"mtu"`
	// Time To Live (TTL) in number of hops of the GRE tunnel.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/gre_tunnel#ttl GreTunnel#ttl}
	Ttl *float64 `field:"optional" json:"ttl" yaml:"ttl"`
}

