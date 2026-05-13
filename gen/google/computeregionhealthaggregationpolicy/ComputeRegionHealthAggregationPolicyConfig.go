package computeregionhealthaggregationpolicy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComputeRegionHealthAggregationPolicyConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/compute_region_health_aggregation_policy#name ComputeRegionHealthAggregationPolicy#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// URL of the region where the health aggregation policy resides.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/compute_region_health_aggregation_policy#region ComputeRegionHealthAggregationPolicy#region}
	Region *string `field:"required" json:"region" yaml:"region"`
	// An optional description of this resource. Provide this property when you create the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/compute_region_health_aggregation_policy#description ComputeRegionHealthAggregationPolicy#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Can only be set if the 'policyType' field is 'BACKEND_SERVICE_POLICY'.
	//
	// Specifies the threshold (as a
	// percentage) of healthy endpoints required in order to consider the
	// aggregated health result HEALTHY. Defaults to '60'. Must be in
	// range [0, 100]. Not applicable if the 'policyType' field is
	// 'DNB_PUBLIC_IP_POLICY'. Can be mutated. This field is optional,
	// and will be set to the default if unspecified. Note that both this
	// threshold and 'minHealthyThreshold' must be satisfied in order
	// for HEALTHY to be the aggregated result. "Endpoints" refers to network
	// endpoints within a Network Endpoint Group or instances within an Instance
	// Group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/compute_region_health_aggregation_policy#healthy_percent_threshold ComputeRegionHealthAggregationPolicy#healthy_percent_threshold}
	HealthyPercentThreshold *float64 `field:"optional" json:"healthyPercentThreshold" yaml:"healthyPercentThreshold"`
	// Can only be set if the 'policyType' field is 'BACKEND_SERVICE_POLICY'.
	//
	// Specifies the minimum number of
	// healthy endpoints required in order to consider the aggregated health
	// result HEALTHY. Defaults to '1'. Must be positive. Not
	// applicable if the 'policyType' field is
	// 'DNB_PUBLIC_IP_POLICY'. Can be mutated. This field is optional,
	// and will be set to the default if unspecified. Note that both this
	// threshold and 'healthyPercentThreshold' must be satisfied in
	// order for HEALTHY to be the aggregated result. "Endpoints" refers to
	// network endpoints within a Network Endpoint Group or instances within an
	// Instance Group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/compute_region_health_aggregation_policy#min_healthy_threshold ComputeRegionHealthAggregationPolicy#min_healthy_threshold}
	MinHealthyThreshold *float64 `field:"optional" json:"minHealthyThreshold" yaml:"minHealthyThreshold"`
	// Specifies the type of the healthAggregationPolicy.
	//
	// The only allowed value
	// for global resources is 'DNS_PUBLIC_IP_POLICY'. The only allowed
	// value for regional resources is 'BACKEND_SERVICE_POLICY'. Must
	// be specified when the healthAggregationPolicy is created, and cannot be
	// mutated. Default value: "BACKEND_SERVICE_POLICY" Possible values: ["DNS_PUBLIC_IP_POLICY", "BACKEND_SERVICE_POLICY"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/compute_region_health_aggregation_policy#policy_type ComputeRegionHealthAggregationPolicy#policy_type}
	PolicyType *string `field:"optional" json:"policyType" yaml:"policyType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/compute_region_health_aggregation_policy#project ComputeRegionHealthAggregationPolicy#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/compute_region_health_aggregation_policy#timeouts ComputeRegionHealthAggregationPolicy#timeouts}
	Timeouts *ComputeRegionHealthAggregationPolicyTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

