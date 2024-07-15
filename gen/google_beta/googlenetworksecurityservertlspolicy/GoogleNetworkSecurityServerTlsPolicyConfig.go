package googlenetworksecurityservertlspolicy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleNetworkSecurityServerTlsPolicyConfig struct {
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
	// Name of the ServerTlsPolicy resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_security_server_tls_policy#name GoogleNetworkSecurityServerTlsPolicy#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// This field applies only for Traffic Director policies.
	//
	// It is must be set to false for external HTTPS load balancer policies.
	// Determines if server allows plaintext connections. If set to true, server allows plain text connections. By default, it is set to false. This setting is not exclusive of other encryption modes. For example, if allowOpen and mtlsPolicy are set, server allows both plain text and mTLS connections. See documentation of other encryption modes to confirm compatibility.
	// Consider using it if you wish to upgrade in place your deployment to TLS while having mixed TLS and non-TLS traffic reaching port :80.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_security_server_tls_policy#allow_open GoogleNetworkSecurityServerTlsPolicy#allow_open}
	AllowOpen interface{} `field:"optional" json:"allowOpen" yaml:"allowOpen"`
	// A free-text description of the resource. Max length 1024 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_security_server_tls_policy#description GoogleNetworkSecurityServerTlsPolicy#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_security_server_tls_policy#id GoogleNetworkSecurityServerTlsPolicy#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Set of label tags associated with the ServerTlsPolicy resource.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_security_server_tls_policy#labels GoogleNetworkSecurityServerTlsPolicy#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// The location of the server tls policy. The default value is 'global'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_security_server_tls_policy#location GoogleNetworkSecurityServerTlsPolicy#location}
	Location *string `field:"optional" json:"location" yaml:"location"`
	// mtls_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_security_server_tls_policy#mtls_policy GoogleNetworkSecurityServerTlsPolicy#mtls_policy}
	MtlsPolicy *GoogleNetworkSecurityServerTlsPolicyMtlsPolicy `field:"optional" json:"mtlsPolicy" yaml:"mtlsPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_security_server_tls_policy#project GoogleNetworkSecurityServerTlsPolicy#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// server_certificate block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_security_server_tls_policy#server_certificate GoogleNetworkSecurityServerTlsPolicy#server_certificate}
	ServerCertificate *GoogleNetworkSecurityServerTlsPolicyServerCertificate `field:"optional" json:"serverCertificate" yaml:"serverCertificate"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_security_server_tls_policy#timeouts GoogleNetworkSecurityServerTlsPolicy#timeouts}
	Timeouts *GoogleNetworkSecurityServerTlsPolicyTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

