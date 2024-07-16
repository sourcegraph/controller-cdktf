package googlecomputeglobaladdress

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleComputeGlobalAddressConfig struct {
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
	// Name of the resource.
	//
	// Provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035.  Specifically, the name must be 1-63 characters long and
	// match the regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means
	// the first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_global_address#name GoogleComputeGlobalAddress#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The IP address or beginning of the address range represented by this resource.
	//
	// This can be supplied as an input to reserve a specific
	// address or omitted to allow GCP to choose a valid one for you.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_global_address#address GoogleComputeGlobalAddress#address}
	Address *string `field:"optional" json:"address" yaml:"address"`
	// The type of the address to reserve.
	//
	// * EXTERNAL indicates public/external single IP address.
	// * INTERNAL indicates internal IP ranges belonging to some network. Default value: "EXTERNAL" Possible values: ["EXTERNAL", "INTERNAL"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_global_address#address_type GoogleComputeGlobalAddress#address_type}
	AddressType *string `field:"optional" json:"addressType" yaml:"addressType"`
	// An optional description of this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_global_address#description GoogleComputeGlobalAddress#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_global_address#id GoogleComputeGlobalAddress#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The IP Version that will be used by this address. The default value is 'IPV4'. Possible values: ["IPV4", "IPV6"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_global_address#ip_version GoogleComputeGlobalAddress#ip_version}
	IpVersion *string `field:"optional" json:"ipVersion" yaml:"ipVersion"`
	// Labels to apply to this address.  A list of key->value pairs.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_global_address#labels GoogleComputeGlobalAddress#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// The URL of the network in which to reserve the IP range.
	//
	// The IP range
	// must be in RFC1918 space. The network cannot be deleted if there are
	// any reserved IP ranges referring to it.
	//
	// This should only be set when using an Internal address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_global_address#network GoogleComputeGlobalAddress#network}
	Network *string `field:"optional" json:"network" yaml:"network"`
	// The prefix length of the IP range. If not present, it means the address field is a single IP address.
	//
	// This field is not applicable to addresses with addressType=INTERNAL
	// when purpose=PRIVATE_SERVICE_CONNECT
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_global_address#prefix_length GoogleComputeGlobalAddress#prefix_length}
	PrefixLength *float64 `field:"optional" json:"prefixLength" yaml:"prefixLength"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_global_address#project GoogleComputeGlobalAddress#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// The purpose of the resource. Possible values include:.
	//
	// * VPC_PEERING - for peer networks
	//
	// * PRIVATE_SERVICE_CONNECT - for ([Beta](https://terraform.io/docs/providers/google/guides/provider_versions.html) only) Private Service Connect networks
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_global_address#purpose GoogleComputeGlobalAddress#purpose}
	Purpose *string `field:"optional" json:"purpose" yaml:"purpose"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_global_address#timeouts GoogleComputeGlobalAddress#timeouts}
	Timeouts *GoogleComputeGlobalAddressTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

