package googlecontainerattachedcluster


type GoogleContainerAttachedClusterSecurityPostureConfig struct {
	// Sets the mode of the Kubernetes security posture API's workload vulnerability scanning. Possible values: ["VULNERABILITY_DISABLED", "VULNERABILITY_ENTERPRISE"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_container_attached_cluster#vulnerability_mode GoogleContainerAttachedCluster#vulnerability_mode}
	VulnerabilityMode *string `field:"required" json:"vulnerabilityMode" yaml:"vulnerabilityMode"`
}

