package computeregioncompositehealthcheck

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComputeRegionCompositeHealthCheckConfig struct {
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
	// URL to the destination resource.
	//
	// Must be set. Must be a
	// ForwardingRule. The ForwardingRule must have
	// load balancing scheme INTERNAL or
	// INTERNAL_MANAGED and must be regional and in the same region
	// as the CompositeHealthCheck (cross-region deployment for
	// INTERNAL_MANAGED is not supported). Can be mutated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/compute_region_composite_health_check#health_destination ComputeRegionCompositeHealthCheck#health_destination}
	HealthDestination *string `field:"required" json:"healthDestination" yaml:"healthDestination"`
	// Name of the resource.
	//
	// Provided by the client when the resource is created.
	// The name must be 1-63 characters long, and comply with RFC1035.
	// Specifically, the name must be 1-63 characters long and match the regular
	// expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first
	// character must be a lowercase letter, and all following characters must
	// be a dash, lowercase letter, or digit, except the last character, which
	// cannot be a dash.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/compute_region_composite_health_check#name ComputeRegionCompositeHealthCheck#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// URL of the region where the composite health check resides.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/compute_region_composite_health_check#region ComputeRegionCompositeHealthCheck#region}
	Region *string `field:"required" json:"region" yaml:"region"`
	// An optional description of this resource. Provide this property when you create the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/compute_region_composite_health_check#description ComputeRegionCompositeHealthCheck#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// URLs to the HealthSource resources whose results are AND'ed.
	//
	// I.e. he aggregated result is is HEALTHY only if all sources
	// are HEALTHY. Must have at least 1. Must not have more than 10.
	// Must be regional and in the same region as the
	// CompositeHealthCheck. Can be mutated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/compute_region_composite_health_check#health_sources ComputeRegionCompositeHealthCheck#health_sources}
	HealthSources *[]*string `field:"optional" json:"healthSources" yaml:"healthSources"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/compute_region_composite_health_check#project ComputeRegionCompositeHealthCheck#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/compute_region_composite_health_check#timeouts ComputeRegionCompositeHealthCheck#timeouts}
	Timeouts *ComputeRegionCompositeHealthCheckTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

