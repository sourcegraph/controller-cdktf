package googlecontainerattachedcluster


type GoogleContainerAttachedClusterProxyConfig struct {
	// kubernetes_secret block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_attached_cluster#kubernetes_secret GoogleContainerAttachedCluster#kubernetes_secret}
	KubernetesSecret *GoogleContainerAttachedClusterProxyConfigKubernetesSecret `field:"optional" json:"kubernetesSecret" yaml:"kubernetesSecret"`
}

