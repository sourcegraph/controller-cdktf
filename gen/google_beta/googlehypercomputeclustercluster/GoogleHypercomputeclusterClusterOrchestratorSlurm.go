package googlehypercomputeclustercluster


type GoogleHypercomputeclusterClusterOrchestratorSlurm struct {
	// login_nodes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_hypercomputecluster_cluster#login_nodes GoogleHypercomputeclusterCluster#login_nodes}
	LoginNodes *GoogleHypercomputeclusterClusterOrchestratorSlurmLoginNodes `field:"required" json:"loginNodes" yaml:"loginNodes"`
	// node_sets block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_hypercomputecluster_cluster#node_sets GoogleHypercomputeclusterCluster#node_sets}
	NodeSets interface{} `field:"required" json:"nodeSets" yaml:"nodeSets"`
	// partitions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_hypercomputecluster_cluster#partitions GoogleHypercomputeclusterCluster#partitions}
	Partitions interface{} `field:"required" json:"partitions" yaml:"partitions"`
	// Default partition to use for submitted jobs that do not explicitly specify a partition.
	//
	// Required if and only if there is more than one partition, in
	// which case it must match the id of one of the partitions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_hypercomputecluster_cluster#default_partition GoogleHypercomputeclusterCluster#default_partition}
	DefaultPartition *string `field:"optional" json:"defaultPartition" yaml:"defaultPartition"`
	// Slurm [epilog scripts](https://slurm.schedmd.com/prolog_epilog.html), which will be executed by compute nodes whenever a node finishes running a job. Values must not be empty.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_hypercomputecluster_cluster#epilog_bash_scripts GoogleHypercomputeclusterCluster#epilog_bash_scripts}
	EpilogBashScripts *[]*string `field:"optional" json:"epilogBashScripts" yaml:"epilogBashScripts"`
	// Slurm [prolog scripts](https://slurm.schedmd.com/prolog_epilog.html), which will be executed by compute nodes before a node begins running a new job. Values must not be empty.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_hypercomputecluster_cluster#prolog_bash_scripts GoogleHypercomputeclusterCluster#prolog_bash_scripts}
	PrologBashScripts *[]*string `field:"optional" json:"prologBashScripts" yaml:"prologBashScripts"`
}

