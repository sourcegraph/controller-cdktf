package googlecomputeinstancegroupmanager

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleComputeInstanceGroupManagerConfig struct {
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
	// The base instance name to use for instances in this group.
	//
	// The value must be a valid RFC1035 name. Supported characters are lowercase letters, numbers, and hyphens (-). Instances are named by appending a hyphen and a random four-character string to the base instance name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_instance_group_manager#base_instance_name GoogleComputeInstanceGroupManager#base_instance_name}
	BaseInstanceName *string `field:"required" json:"baseInstanceName" yaml:"baseInstanceName"`
	// The name of the instance group manager.
	//
	// Must be 1-63 characters long and comply with RFC1035. Supported characters include lowercase letters, numbers, and hyphens.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_instance_group_manager#name GoogleComputeInstanceGroupManager#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// version block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_instance_group_manager#version GoogleComputeInstanceGroupManager#version}
	Version interface{} `field:"required" json:"version" yaml:"version"`
	// all_instances_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_instance_group_manager#all_instances_config GoogleComputeInstanceGroupManager#all_instances_config}
	AllInstancesConfig *GoogleComputeInstanceGroupManagerAllInstancesConfig `field:"optional" json:"allInstancesConfig" yaml:"allInstancesConfig"`
	// auto_healing_policies block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_instance_group_manager#auto_healing_policies GoogleComputeInstanceGroupManager#auto_healing_policies}
	AutoHealingPolicies *GoogleComputeInstanceGroupManagerAutoHealingPolicies `field:"optional" json:"autoHealingPolicies" yaml:"autoHealingPolicies"`
	// An optional textual description of the instance group manager.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_instance_group_manager#description GoogleComputeInstanceGroupManager#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_instance_group_manager#id GoogleComputeInstanceGroupManager#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// instance_lifecycle_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_instance_group_manager#instance_lifecycle_policy GoogleComputeInstanceGroupManager#instance_lifecycle_policy}
	InstanceLifecyclePolicy *GoogleComputeInstanceGroupManagerInstanceLifecyclePolicy `field:"optional" json:"instanceLifecyclePolicy" yaml:"instanceLifecyclePolicy"`
	// Pagination behavior of the listManagedInstances API method for this managed instance group.
	//
	// Valid values are: "PAGELESS", "PAGINATED". If PAGELESS (default), Pagination is disabled for the group's listManagedInstances API method. maxResults and pageToken query parameters are ignored and all instances are returned in a single response. If PAGINATED, pagination is enabled, maxResults and pageToken query parameters are respected.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_instance_group_manager#list_managed_instances_results GoogleComputeInstanceGroupManager#list_managed_instances_results}
	ListManagedInstancesResults *string `field:"optional" json:"listManagedInstancesResults" yaml:"listManagedInstancesResults"`
	// named_port block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_instance_group_manager#named_port GoogleComputeInstanceGroupManager#named_port}
	NamedPort interface{} `field:"optional" json:"namedPort" yaml:"namedPort"`
	// params block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_instance_group_manager#params GoogleComputeInstanceGroupManager#params}
	Params *GoogleComputeInstanceGroupManagerParams `field:"optional" json:"params" yaml:"params"`
	// The ID of the project in which the resource belongs.
	//
	// If it is not provided, the provider project is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_instance_group_manager#project GoogleComputeInstanceGroupManager#project}
	Project *string `field:"optional" json:"project" yaml:"project"`
	// standby_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_instance_group_manager#standby_policy GoogleComputeInstanceGroupManager#standby_policy}
	StandbyPolicy *GoogleComputeInstanceGroupManagerStandbyPolicy `field:"optional" json:"standbyPolicy" yaml:"standbyPolicy"`
	// stateful_disk block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_instance_group_manager#stateful_disk GoogleComputeInstanceGroupManager#stateful_disk}
	StatefulDisk interface{} `field:"optional" json:"statefulDisk" yaml:"statefulDisk"`
	// stateful_external_ip block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_instance_group_manager#stateful_external_ip GoogleComputeInstanceGroupManager#stateful_external_ip}
	StatefulExternalIp interface{} `field:"optional" json:"statefulExternalIp" yaml:"statefulExternalIp"`
	// stateful_internal_ip block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_instance_group_manager#stateful_internal_ip GoogleComputeInstanceGroupManager#stateful_internal_ip}
	StatefulInternalIp interface{} `field:"optional" json:"statefulInternalIp" yaml:"statefulInternalIp"`
	// The full URL of all target pools to which new instances in the group are added.
	//
	// Updating the target pools attribute does not affect existing instances.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_instance_group_manager#target_pools GoogleComputeInstanceGroupManager#target_pools}
	TargetPools *[]*string `field:"optional" json:"targetPools" yaml:"targetPools"`
	// The target number of running instances for this managed instance group.
	//
	// This value should always be explicitly set unless this resource is attached to an autoscaler, in which case it should never be set. Defaults to 0.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_instance_group_manager#target_size GoogleComputeInstanceGroupManager#target_size}
	TargetSize *float64 `field:"optional" json:"targetSize" yaml:"targetSize"`
	// The target number of stopped instances for this managed instance group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_instance_group_manager#target_stopped_size GoogleComputeInstanceGroupManager#target_stopped_size}
	TargetStoppedSize *float64 `field:"optional" json:"targetStoppedSize" yaml:"targetStoppedSize"`
	// The target number of suspended instances for this managed instance group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_instance_group_manager#target_suspended_size GoogleComputeInstanceGroupManager#target_suspended_size}
	TargetSuspendedSize *float64 `field:"optional" json:"targetSuspendedSize" yaml:"targetSuspendedSize"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_instance_group_manager#timeouts GoogleComputeInstanceGroupManager#timeouts}
	Timeouts *GoogleComputeInstanceGroupManagerTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// update_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_instance_group_manager#update_policy GoogleComputeInstanceGroupManager#update_policy}
	UpdatePolicy *GoogleComputeInstanceGroupManagerUpdatePolicy `field:"optional" json:"updatePolicy" yaml:"updatePolicy"`
	// Whether to wait for all instances to be created/updated before returning.
	//
	// Note that if this is set to true and the operation does not succeed, Terraform will continue trying until it times out.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_instance_group_manager#wait_for_instances GoogleComputeInstanceGroupManager#wait_for_instances}
	WaitForInstances interface{} `field:"optional" json:"waitForInstances" yaml:"waitForInstances"`
	// When used with wait_for_instances specifies the status to wait for.
	//
	// When STABLE is specified this resource will wait until the instances are stable before returning. When UPDATED is set, it will wait for the version target to be reached and any per instance configs to be effective and all instances configs to be effective as well as all instances to be stable before returning.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_instance_group_manager#wait_for_instances_status GoogleComputeInstanceGroupManager#wait_for_instances_status}
	WaitForInstancesStatus *string `field:"optional" json:"waitForInstancesStatus" yaml:"waitForInstancesStatus"`
	// The zone that instances in this group should be created in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_instance_group_manager#zone GoogleComputeInstanceGroupManager#zone}
	Zone *string `field:"optional" json:"zone" yaml:"zone"`
}

