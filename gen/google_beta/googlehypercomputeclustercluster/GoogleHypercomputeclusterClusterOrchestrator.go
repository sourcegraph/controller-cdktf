package googlehypercomputeclustercluster


type GoogleHypercomputeclusterClusterOrchestrator struct {
	// slurm block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_hypercomputecluster_cluster#slurm GoogleHypercomputeclusterCluster#slurm}
	Slurm *GoogleHypercomputeclusterClusterOrchestratorSlurm `field:"optional" json:"slurm" yaml:"slurm"`
}

