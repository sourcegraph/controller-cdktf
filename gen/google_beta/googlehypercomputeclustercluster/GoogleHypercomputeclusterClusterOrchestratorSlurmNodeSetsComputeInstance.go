package googlehypercomputeclustercluster


type GoogleHypercomputeclusterClusterOrchestratorSlurmNodeSetsComputeInstance struct {
	// boot_disk block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_hypercomputecluster_cluster#boot_disk GoogleHypercomputeclusterCluster#boot_disk}
	BootDisk *GoogleHypercomputeclusterClusterOrchestratorSlurmNodeSetsComputeInstanceBootDisk `field:"optional" json:"bootDisk" yaml:"bootDisk"`
	// [Labels](https://cloud.google.com/compute/docs/labeling-resources) that should be applied to each VM instance in the nodeset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_hypercomputecluster_cluster#labels GoogleHypercomputeclusterCluster#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// [Startup script](https://cloud.google.com/compute/docs/instances/startup-scripts/linux) to be run on each VM instance in the nodeset. Max 256KB.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_hypercomputecluster_cluster#startup_script GoogleHypercomputeclusterCluster#startup_script}
	StartupScript *string `field:"optional" json:"startupScript" yaml:"startupScript"`
}

