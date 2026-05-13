package googlecomputepublicadvertisedprefix

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleComputePublicAdvertisedPrefixConfig struct {
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
	// The address range, in CIDR format, represented by this public advertised prefix.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_public_advertised_prefix#ip_cidr_range GoogleComputePublicAdvertisedPrefix#ip_cidr_range}
	IpCidrRange *string `field:"required" json:"ipCidrRange" yaml:"ipCidrRange"`
	// Name of the resource.
	//
	// The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters
	// long and match the regular expression '[a-z]([-a-z0-9]*[a-z0-9])?'
	// which means the first character must be a lowercase letter, and all
	// following characters must be a dash, lowercase letter, or digit,
	// except the last character, which cannot be a dash.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_public_advertised_prefix#name GoogleComputePublicAdvertisedPrefix#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// An optional description of this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_public_advertised_prefix#description GoogleComputePublicAdvertisedPrefix#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The IPv4 address to be used for reverse DNS verification.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_public_advertised_prefix#dns_verification_ip GoogleComputePublicAdvertisedPrefix#dns_verification_ip}
	DnsVerificationIp *string `field:"optional" json:"dnsVerificationIp" yaml:"dnsVerificationIp"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_public_advertised_prefix#id GoogleComputePublicAdvertisedPrefix#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The internet access type for IPv6 Public Advertised Prefixes.
	//
	// It can be
	// set to one of following:
	//   * EXTERNAL: Default access type. The prefix will be announced to the
	//   internet. All children PDPs will have access type as EXTERNAL.
	//   * INTERNAL: The prefix won’t be announced to the internet. Prefix will
	//   be used privately within Google Cloud. All children PDPs will have
	//   access type as INTERNAL. Possible values: ["EXTERNAL", "INTERNAL"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_public_advertised_prefix#ipv6_access_type GoogleComputePublicAdvertisedPrefix#ipv6_access_type}
	Ipv6AccessType *string `field:"optional" json:"ipv6AccessType" yaml:"ipv6AccessType"`
	// Specifies how child public delegated prefix will be scoped.
	//
	// pdpScope
	// must be one of: GLOBAL, REGIONAL
	// * REGIONAL: The public delegated prefix is regional only. The
	// provisioning will take a few minutes.
	// * GLOBAL: The public delegated prefix is global only. The provisioning
	// will take ~4 weeks. Possible values: ["GLOBAL", "REGIONAL"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_public_advertised_prefix#pdp_scope GoogleComputePublicAdvertisedPrefix#pdp_scope}
	PdpScope *string `field:"optional" json:"pdpScope" yaml:"pdpScope"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_public_advertised_prefix#project GoogleComputePublicAdvertisedPrefix#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_public_advertised_prefix#timeouts GoogleComputePublicAdvertisedPrefix#timeouts}
	Timeouts *GoogleComputePublicAdvertisedPrefixTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

