package containercluster


type ContainerClusterAutopilotClusterPolicyConfig struct {
	// If true, prevents standard node pools and requires only autopilot node pools.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/container_cluster#no_standard_node_pools ContainerCluster#no_standard_node_pools}
	NoStandardNodePools interface{} `field:"optional" json:"noStandardNodePools" yaml:"noStandardNodePools"`
	// If true, prevents impersonation and CSRs for GKE System users.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/container_cluster#no_system_impersonation ContainerCluster#no_system_impersonation}
	NoSystemImpersonation interface{} `field:"optional" json:"noSystemImpersonation" yaml:"noSystemImpersonation"`
	// If true, prevents creation and mutation of resources in GKE managed namespaces and cluster-scoped GKE managed resources.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/container_cluster#no_system_mutation ContainerCluster#no_system_mutation}
	NoSystemMutation interface{} `field:"optional" json:"noSystemMutation" yaml:"noSystemMutation"`
	// If true, unsafe webhooks are not allowed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/container_cluster#no_unsafe_webhooks ContainerCluster#no_unsafe_webhooks}
	NoUnsafeWebhooks interface{} `field:"optional" json:"noUnsafeWebhooks" yaml:"noUnsafeWebhooks"`
}

