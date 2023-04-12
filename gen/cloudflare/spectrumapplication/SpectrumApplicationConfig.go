package spectrumapplication

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SpectrumApplicationConfig struct {
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
	// dns block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/spectrum_application#dns SpectrumApplication#dns}
	Dns *SpectrumApplicationDns `field:"required" json:"dns" yaml:"dns"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/spectrum_application#protocol SpectrumApplication#protocol}.
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/spectrum_application#zone_id SpectrumApplication#zone_id}.
	ZoneId *string `field:"required" json:"zoneId" yaml:"zoneId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/spectrum_application#argo_smart_routing SpectrumApplication#argo_smart_routing}.
	ArgoSmartRouting interface{} `field:"optional" json:"argoSmartRouting" yaml:"argoSmartRouting"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/spectrum_application#edge_ip_connectivity SpectrumApplication#edge_ip_connectivity}.
	EdgeIpConnectivity *string `field:"optional" json:"edgeIpConnectivity" yaml:"edgeIpConnectivity"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/spectrum_application#edge_ips SpectrumApplication#edge_ips}.
	EdgeIps *[]*string `field:"optional" json:"edgeIps" yaml:"edgeIps"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/spectrum_application#id SpectrumApplication#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/spectrum_application#ip_firewall SpectrumApplication#ip_firewall}.
	IpFirewall interface{} `field:"optional" json:"ipFirewall" yaml:"ipFirewall"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/spectrum_application#origin_direct SpectrumApplication#origin_direct}.
	OriginDirect *[]*string `field:"optional" json:"originDirect" yaml:"originDirect"`
	// origin_dns block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/spectrum_application#origin_dns SpectrumApplication#origin_dns}
	OriginDns *SpectrumApplicationOriginDns `field:"optional" json:"originDns" yaml:"originDns"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/spectrum_application#origin_port SpectrumApplication#origin_port}.
	OriginPort *float64 `field:"optional" json:"originPort" yaml:"originPort"`
	// origin_port_range block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/spectrum_application#origin_port_range SpectrumApplication#origin_port_range}
	OriginPortRange *SpectrumApplicationOriginPortRange `field:"optional" json:"originPortRange" yaml:"originPortRange"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/spectrum_application#proxy_protocol SpectrumApplication#proxy_protocol}.
	ProxyProtocol *string `field:"optional" json:"proxyProtocol" yaml:"proxyProtocol"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/spectrum_application#tls SpectrumApplication#tls}.
	Tls *string `field:"optional" json:"tls" yaml:"tls"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/spectrum_application#traffic_type SpectrumApplication#traffic_type}.
	TrafficType *string `field:"optional" json:"trafficType" yaml:"trafficType"`
}

