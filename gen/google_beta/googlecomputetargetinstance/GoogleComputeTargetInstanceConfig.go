package googlecomputetargetinstance

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleComputeTargetInstanceConfig struct {
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
	// The Compute instance VM handling traffic for this target instance.
	//
	// Accepts the instance self-link, relative path
	// (e.g. 'projects/project/zones/zone/instances/instance') or name. If
	// name is given, the zone will default to the given zone or
	// the provider-default zone and the project will default to the
	// provider-level project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_target_instance#instance GoogleComputeTargetInstance#instance}
	Instance *string `field:"required" json:"instance" yaml:"instance"`
	// Name of the resource.
	//
	// Provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035. Specifically, the name must be 1-63 characters long and match
	// the regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means the
	// first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_target_instance#name GoogleComputeTargetInstance#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// An optional description of this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_target_instance#description GoogleComputeTargetInstance#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_target_instance#id GoogleComputeTargetInstance#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// NAT option controlling how IPs are NAT'ed to the instance.
	//
	// Currently only NO_NAT (default value) is supported. Default value: "NO_NAT" Possible values: ["NO_NAT"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_target_instance#nat_policy GoogleComputeTargetInstance#nat_policy}
	NatPolicy *string `field:"optional" json:"natPolicy" yaml:"natPolicy"`
	// The URL of the network this target instance uses to forward traffic.
	//
	// If not specified, the traffic will be forwarded to the network that the default network interface belongs to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_target_instance#network GoogleComputeTargetInstance#network}
	Network *string `field:"optional" json:"network" yaml:"network"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_target_instance#project GoogleComputeTargetInstance#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// The resource URL for the security policy associated with this target instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_target_instance#security_policy GoogleComputeTargetInstance#security_policy}
	SecurityPolicy *string `field:"optional" json:"securityPolicy" yaml:"securityPolicy"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_target_instance#timeouts GoogleComputeTargetInstance#timeouts}
	Timeouts *GoogleComputeTargetInstanceTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// URL of the zone where the target instance resides.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_target_instance#zone GoogleComputeTargetInstance#zone}
	Zone *string `field:"optional" json:"zone" yaml:"zone"`
}

