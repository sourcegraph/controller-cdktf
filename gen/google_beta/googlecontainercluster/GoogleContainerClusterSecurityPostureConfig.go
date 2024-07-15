package googlecontainercluster


type GoogleContainerClusterSecurityPostureConfig struct {
	// Sets the mode of the Kubernetes security posture API's off-cluster features. Available options include DISABLED, BASIC, and ENTERPRISE.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_cluster#mode GoogleContainerCluster#mode}
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// Sets the mode of the Kubernetes security posture API's workload vulnerability scanning. Available options include VULNERABILITY_DISABLED, VULNERABILITY_BASIC and VULNERABILITY_ENTERPRISE.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_cluster#vulnerability_mode GoogleContainerCluster#vulnerability_mode}
	VulnerabilityMode *string `field:"optional" json:"vulnerabilityMode" yaml:"vulnerabilityMode"`
}

