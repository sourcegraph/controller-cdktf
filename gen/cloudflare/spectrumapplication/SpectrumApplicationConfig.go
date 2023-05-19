package spectrumapplication

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SpectrumApplicationConfig struct {
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
	// dns block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/spectrum_application#dns SpectrumApplication#dns}
	Dns *SpectrumApplicationDns `field:"required" json:"dns" yaml:"dns"`
	// The port configuration at Cloudflare's edge. e.g. `tcp/22`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/spectrum_application#protocol SpectrumApplication#protocol}
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
	// The zone identifier to target for the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/spectrum_application#zone_id SpectrumApplication#zone_id}
	ZoneId *string `field:"required" json:"zoneId" yaml:"zoneId"`
	// Enables Argo Smart Routing. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/spectrum_application#argo_smart_routing SpectrumApplication#argo_smart_routing}
	ArgoSmartRouting interface{} `field:"optional" json:"argoSmartRouting" yaml:"argoSmartRouting"`
	// edge_ips block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/spectrum_application#edge_ips SpectrumApplication#edge_ips}
	EdgeIps *SpectrumApplicationEdgeIps `field:"optional" json:"edgeIps" yaml:"edgeIps"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/spectrum_application#id SpectrumApplication#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Enables the IP Firewall for this application. Defaults to `true`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/spectrum_application#ip_firewall SpectrumApplication#ip_firewall}
	IpFirewall interface{} `field:"optional" json:"ipFirewall" yaml:"ipFirewall"`
	// A list of destination addresses to the origin. e.g. `tcp://192.0.2.1:22`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/spectrum_application#origin_direct SpectrumApplication#origin_direct}
	OriginDirect *[]*string `field:"optional" json:"originDirect" yaml:"originDirect"`
	// origin_dns block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/spectrum_application#origin_dns SpectrumApplication#origin_dns}
	OriginDns *SpectrumApplicationOriginDns `field:"optional" json:"originDns" yaml:"originDns"`
	// Origin port to proxy traffice to. Conflicts with `origin_port_range`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/spectrum_application#origin_port SpectrumApplication#origin_port}
	OriginPort *float64 `field:"optional" json:"originPort" yaml:"originPort"`
	// origin_port_range block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/spectrum_application#origin_port_range SpectrumApplication#origin_port_range}
	OriginPortRange *SpectrumApplicationOriginPortRange `field:"optional" json:"originPortRange" yaml:"originPortRange"`
	// Enables a proxy protocol to the origin. Available values: `off`, `v1`, `v2`, `simple`. Defaults to `off`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/spectrum_application#proxy_protocol SpectrumApplication#proxy_protocol}
	ProxyProtocol *string `field:"optional" json:"proxyProtocol" yaml:"proxyProtocol"`
	// TLS configuration option for Cloudflare to connect to your origin. Available values: `off`, `flexible`, `full`, `strict`. Defaults to `off`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/spectrum_application#tls SpectrumApplication#tls}
	Tls *string `field:"optional" json:"tls" yaml:"tls"`
	// Sets application type. Available values: `direct`, `http`, `https`. Defaults to `direct`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/spectrum_application#traffic_type SpectrumApplication#traffic_type}
	TrafficType *string `field:"optional" json:"trafficType" yaml:"trafficType"`
}

