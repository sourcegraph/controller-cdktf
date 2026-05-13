package googlecomputerouternamedset

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleComputeRouterNamedSetConfig struct {
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
	// The name of the Named Set, which must be a resource ID segment and unique within all named sets owned by the Router.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_router_named_set#name GoogleComputeRouterNamedSet#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The name of the Cloud Router in which this Named Set will be configured.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_router_named_set#router GoogleComputeRouterNamedSet#router}
	Router *string `field:"required" json:"router" yaml:"router"`
	// The type of the Named Set. Possible values: ["NAMED_SET_TYPE_PREFIX", "NAMED_SET_TYPE_COMMUNITY"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_router_named_set#type GoogleComputeRouterNamedSet#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// An optional description of the Named Set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_router_named_set#description GoogleComputeRouterNamedSet#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// elements block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_router_named_set#elements GoogleComputeRouterNamedSet#elements}
	Elements interface{} `field:"optional" json:"elements" yaml:"elements"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_router_named_set#id GoogleComputeRouterNamedSet#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_router_named_set#project GoogleComputeRouterNamedSet#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// Region where the router resides.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_router_named_set#region GoogleComputeRouterNamedSet#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_compute_router_named_set#timeouts GoogleComputeRouterNamedSet#timeouts}
	Timeouts *GoogleComputeRouterNamedSetTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

