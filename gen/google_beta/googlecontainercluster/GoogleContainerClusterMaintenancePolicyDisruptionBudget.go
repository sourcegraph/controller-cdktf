package googlecontainercluster


type GoogleContainerClusterMaintenancePolicyDisruptionBudget struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_container_cluster#minor_version_disruption_interval GoogleContainerCluster#minor_version_disruption_interval}.
	MinorVersionDisruptionInterval *string `field:"optional" json:"minorVersionDisruptionInterval" yaml:"minorVersionDisruptionInterval"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_container_cluster#patch_version_disruption_interval GoogleContainerCluster#patch_version_disruption_interval}.
	PatchVersionDisruptionInterval *string `field:"optional" json:"patchVersionDisruptionInterval" yaml:"patchVersionDisruptionInterval"`
}

