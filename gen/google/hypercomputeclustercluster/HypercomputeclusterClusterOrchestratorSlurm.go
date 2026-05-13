package hypercomputeclustercluster


type HypercomputeclusterClusterOrchestratorSlurm struct {
	// login_nodes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/hypercomputecluster_cluster#login_nodes HypercomputeclusterCluster#login_nodes}
	LoginNodes *HypercomputeclusterClusterOrchestratorSlurmLoginNodes `field:"required" json:"loginNodes" yaml:"loginNodes"`
	// node_sets block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/hypercomputecluster_cluster#node_sets HypercomputeclusterCluster#node_sets}
	NodeSets interface{} `field:"required" json:"nodeSets" yaml:"nodeSets"`
	// partitions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/hypercomputecluster_cluster#partitions HypercomputeclusterCluster#partitions}
	Partitions interface{} `field:"required" json:"partitions" yaml:"partitions"`
	// Default partition to use for submitted jobs that do not explicitly specify a partition.
	//
	// Required if and only if there is more than one partition, in
	// which case it must match the id of one of the partitions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/hypercomputecluster_cluster#default_partition HypercomputeclusterCluster#default_partition}
	DefaultPartition *string `field:"optional" json:"defaultPartition" yaml:"defaultPartition"`
	// Slurm [epilog scripts](https://slurm.schedmd.com/prolog_epilog.html), which will be executed by compute nodes whenever a node finishes running a job. Values must not be empty.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/hypercomputecluster_cluster#epilog_bash_scripts HypercomputeclusterCluster#epilog_bash_scripts}
	EpilogBashScripts *[]*string `field:"optional" json:"epilogBashScripts" yaml:"epilogBashScripts"`
	// Slurm [prolog scripts](https://slurm.schedmd.com/prolog_epilog.html), which will be executed by compute nodes before a node begins running a new job. Values must not be empty.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/hypercomputecluster_cluster#prolog_bash_scripts HypercomputeclusterCluster#prolog_bash_scripts}
	PrologBashScripts *[]*string `field:"optional" json:"prologBashScripts" yaml:"prologBashScripts"`
}

