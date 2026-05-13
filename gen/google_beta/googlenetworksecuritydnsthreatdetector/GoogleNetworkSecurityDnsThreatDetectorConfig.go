package googlenetworksecuritydnsthreatdetector

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleNetworkSecurityDnsThreatDetectorConfig struct {
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
	// Name of the DnsThreatDetector resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_security_dns_threat_detector#name GoogleNetworkSecurityDnsThreatDetector#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// List of networks that are excluded from detection. Format: projects/{project}/global/networks/{name}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_security_dns_threat_detector#excluded_networks GoogleNetworkSecurityDnsThreatDetector#excluded_networks}
	ExcludedNetworks *[]*string `field:"optional" json:"excludedNetworks" yaml:"excludedNetworks"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_security_dns_threat_detector#id GoogleNetworkSecurityDnsThreatDetector#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Set of label tags associated with the DNS Threat Detector resource.
	//
	// An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	//
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_security_dns_threat_detector#labels GoogleNetworkSecurityDnsThreatDetector#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// The location of the DNS Threat Detector. The only supported value is 'global'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_security_dns_threat_detector#location GoogleNetworkSecurityDnsThreatDetector#location}
	Location *string `field:"optional" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_security_dns_threat_detector#project GoogleNetworkSecurityDnsThreatDetector#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// DNS Threat Detection provider. The only supported value is 'INFOBLOX'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_security_dns_threat_detector#threat_detector_provider GoogleNetworkSecurityDnsThreatDetector#threat_detector_provider}
	ThreatDetectorProvider *string `field:"optional" json:"threatDetectorProvider" yaml:"threatDetectorProvider"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_security_dns_threat_detector#timeouts GoogleNetworkSecurityDnsThreatDetector#timeouts}
	Timeouts *GoogleNetworkSecurityDnsThreatDetectorTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

