package googlecontainercluster


type GoogleContainerClusterNodePoolAutoConfigNodeKubeletConfig struct {
	// Controls whether the kubelet read-only port is enabled.
	//
	// It is strongly recommended to set this to `FALSE`. Possible values: `TRUE`, `FALSE`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_container_cluster#insecure_kubelet_readonly_port_enabled GoogleContainerCluster#insecure_kubelet_readonly_port_enabled}
	InsecureKubeletReadonlyPortEnabled *string `field:"optional" json:"insecureKubeletReadonlyPortEnabled" yaml:"insecureKubeletReadonlyPortEnabled"`
}

