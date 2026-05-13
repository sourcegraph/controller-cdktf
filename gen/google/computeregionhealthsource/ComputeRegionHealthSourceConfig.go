package computeregionhealthsource

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComputeRegionHealthSourceConfig struct {
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
	// Provided by the client when the resource is created.
	// The name must be 1-63 characters long, and comply with RFC1035.
	// Specifically, the name must be 1-63 characters long and match the regular
	// expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first
	// character must be a lowercase letter, and all following characters must
	// be a dash, lowercase letter, or digit, except the last character, which
	// cannot be a dash.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/compute_region_health_source#name ComputeRegionHealthSource#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// URL of the region where the health source resides.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/compute_region_health_source#region ComputeRegionHealthSource#region}
	Region *string `field:"required" json:"region" yaml:"region"`
	// Specifies the type of the 'HealthSource'.
	//
	// The only allowed value
	// is 'BACKEND_SERVICE'. Must be specified when the
	// 'HealthSource' is created, and cannot be mutated. Possible values: ["BACKEND_SERVICE"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/compute_region_health_source#source_type ComputeRegionHealthSource#source_type}
	SourceType *string `field:"required" json:"sourceType" yaml:"sourceType"`
	// An optional description of this resource. Provide this property when you create the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/compute_region_health_source#description ComputeRegionHealthSource#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// URL to the 'HealthAggregationPolicy' resource.
	//
	// Must be set. Must
	// be regional and in the same region as the 'HealthSource'. Can be
	// mutated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/compute_region_health_source#health_aggregation_policy ComputeRegionHealthSource#health_aggregation_policy}
	HealthAggregationPolicy *string `field:"optional" json:"healthAggregationPolicy" yaml:"healthAggregationPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/compute_region_health_source#project ComputeRegionHealthSource#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// URLs to the source resources.
	//
	// Must be size 1. Must be a
	// 'BackendService' if the 'sourceType' is 'BACKEND_SERVICE'. The
	// 'BackendService' must have load balancing scheme
	// 'INTERNAL' or 'INTERNAL_MANAGED' and must be regional
	// and in the same region as the 'HealthSource' (cross-region
	// deployment for 'INTERNAL_MANAGED' is not supported). The
	// 'BackendService' may use only IGs, MIGs, or NEGs of type
	// 'GCE_VM_IP' or 'GCE_VM_IP_PORT'. The
	// 'BackendService' may not use 'haPolicy'. Can be
	// mutated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/compute_region_health_source#sources ComputeRegionHealthSource#sources}
	Sources *[]*string `field:"optional" json:"sources" yaml:"sources"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/compute_region_health_source#timeouts ComputeRegionHealthSource#timeouts}
	Timeouts *ComputeRegionHealthSourceTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

