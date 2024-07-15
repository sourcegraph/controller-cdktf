package googlenetworksecurityclienttlspolicy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleNetworkSecurityClientTlsPolicyConfig struct {
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
	// Name of the ClientTlsPolicy resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_security_client_tls_policy#name GoogleNetworkSecurityClientTlsPolicy#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// client_certificate block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_security_client_tls_policy#client_certificate GoogleNetworkSecurityClientTlsPolicy#client_certificate}
	ClientCertificate *GoogleNetworkSecurityClientTlsPolicyClientCertificate `field:"optional" json:"clientCertificate" yaml:"clientCertificate"`
	// A free-text description of the resource. Max length 1024 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_security_client_tls_policy#description GoogleNetworkSecurityClientTlsPolicy#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_security_client_tls_policy#id GoogleNetworkSecurityClientTlsPolicy#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Set of label tags associated with the ClientTlsPolicy resource.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_security_client_tls_policy#labels GoogleNetworkSecurityClientTlsPolicy#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// The location of the client tls policy. The default value is 'global'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_security_client_tls_policy#location GoogleNetworkSecurityClientTlsPolicy#location}
	Location *string `field:"optional" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_security_client_tls_policy#project GoogleNetworkSecurityClientTlsPolicy#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// server_validation_ca block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_security_client_tls_policy#server_validation_ca GoogleNetworkSecurityClientTlsPolicy#server_validation_ca}
	ServerValidationCa interface{} `field:"optional" json:"serverValidationCa" yaml:"serverValidationCa"`
	// Server Name Indication string to present to the server during TLS handshake. E.g: "secure.example.com".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_security_client_tls_policy#sni GoogleNetworkSecurityClientTlsPolicy#sni}
	Sni *string `field:"optional" json:"sni" yaml:"sni"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_security_client_tls_policy#timeouts GoogleNetworkSecurityClientTlsPolicy#timeouts}
	Timeouts *GoogleNetworkSecurityClientTlsPolicyTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

