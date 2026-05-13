package containercluster


type ContainerClusterClusterAutoscaling struct {
	// auto_provisioning_defaults block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/container_cluster#auto_provisioning_defaults ContainerCluster#auto_provisioning_defaults}
	AutoProvisioningDefaults *ContainerClusterClusterAutoscalingAutoProvisioningDefaults `field:"optional" json:"autoProvisioningDefaults" yaml:"autoProvisioningDefaults"`
	// The list of Google Compute Engine zones in which the NodePool's nodes can be created by NAP.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/container_cluster#auto_provisioning_locations ContainerCluster#auto_provisioning_locations}
	AutoProvisioningLocations *[]*string `field:"optional" json:"autoProvisioningLocations" yaml:"autoProvisioningLocations"`
	// Configuration options for the Autoscaling profile feature, which lets you choose whether the cluster autoscaler should optimize for resource utilization or resource availability when deciding to remove nodes from a cluster.
	//
	// Can be BALANCED or OPTIMIZE_UTILIZATION. Defaults to BALANCED.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/container_cluster#autoscaling_profile ContainerCluster#autoscaling_profile}
	AutoscalingProfile *string `field:"optional" json:"autoscalingProfile" yaml:"autoscalingProfile"`
	// Specifies whether default compute class behaviour is enabled.
	//
	// If enabled, cluster autoscaler will use Compute Class with name default for all the workloads, if not overriden.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/container_cluster#default_compute_class_enabled ContainerCluster#default_compute_class_enabled}
	DefaultComputeClassEnabled interface{} `field:"optional" json:"defaultComputeClassEnabled" yaml:"defaultComputeClassEnabled"`
	// Whether node auto-provisioning is enabled. Resource limits for cpu and memory must be defined to enable node auto-provisioning.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/container_cluster#enabled ContainerCluster#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// resource_limits block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/container_cluster#resource_limits ContainerCluster#resource_limits}
	ResourceLimits interface{} `field:"optional" json:"resourceLimits" yaml:"resourceLimits"`
}

