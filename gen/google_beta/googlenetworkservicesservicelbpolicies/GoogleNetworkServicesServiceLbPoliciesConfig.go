package googlenetworkservicesservicelbpolicies

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleNetworkServicesServiceLbPoliciesConfig struct {
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
	// The location of the service lb policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_services_service_lb_policies#location GoogleNetworkServicesServiceLbPolicies#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// Name of the ServiceLbPolicy resource. It matches pattern 'projects/{project}/locations/{location}/serviceLbPolicies/{service_lb_policy_name}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_services_service_lb_policies#name GoogleNetworkServicesServiceLbPolicies#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// auto_capacity_drain block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_services_service_lb_policies#auto_capacity_drain GoogleNetworkServicesServiceLbPolicies#auto_capacity_drain}
	AutoCapacityDrain *GoogleNetworkServicesServiceLbPoliciesAutoCapacityDrain `field:"optional" json:"autoCapacityDrain" yaml:"autoCapacityDrain"`
	// A free-text description of the resource. Max length 1024 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_services_service_lb_policies#description GoogleNetworkServicesServiceLbPolicies#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// failover_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_services_service_lb_policies#failover_config GoogleNetworkServicesServiceLbPolicies#failover_config}
	FailoverConfig *GoogleNetworkServicesServiceLbPoliciesFailoverConfig `field:"optional" json:"failoverConfig" yaml:"failoverConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_services_service_lb_policies#id GoogleNetworkServicesServiceLbPolicies#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Set of label tags associated with the ServiceLbPolicy resource.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_services_service_lb_policies#labels GoogleNetworkServicesServiceLbPolicies#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// The type of load balancing algorithm to be used.
	//
	// The default behavior is WATERFALL_BY_REGION. Possible values: ["SPRAY_TO_REGION", "SPRAY_TO_WORLD", "WATERFALL_BY_REGION", "WATERFALL_BY_ZONE"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_services_service_lb_policies#load_balancing_algorithm GoogleNetworkServicesServiceLbPolicies#load_balancing_algorithm}
	LoadBalancingAlgorithm *string `field:"optional" json:"loadBalancingAlgorithm" yaml:"loadBalancingAlgorithm"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_services_service_lb_policies#project GoogleNetworkServicesServiceLbPolicies#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_services_service_lb_policies#timeouts GoogleNetworkServicesServiceLbPolicies#timeouts}
	Timeouts *GoogleNetworkServicesServiceLbPoliciesTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

