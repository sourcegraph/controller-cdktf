package googlecontainercluster


type GoogleContainerClusterNodeConfigKubeletConfigTopologyManager struct {
	// The Topology Manager policy to use. This policy dictates how resource alignment is handled on the node.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_container_cluster#policy GoogleContainerCluster#policy}
	Policy *string `field:"optional" json:"policy" yaml:"policy"`
	// The Topology Manager scope, defining the granularity at which policy decisions are applied.
	//
	// Valid values are "container" (resources are aligned per container within a pod) or "pod" (resources are aligned for the entire pod).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_container_cluster#scope GoogleContainerCluster#scope}
	Scope *string `field:"optional" json:"scope" yaml:"scope"`
}

