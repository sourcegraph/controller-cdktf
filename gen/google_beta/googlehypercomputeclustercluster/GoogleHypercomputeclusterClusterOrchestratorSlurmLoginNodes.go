package googlehypercomputeclustercluster


type GoogleHypercomputeclusterClusterOrchestratorSlurmLoginNodes struct {
	// Number of login node instances to create.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_hypercomputecluster_cluster#count GoogleHypercomputeclusterCluster#count}
	Count *string `field:"required" json:"count" yaml:"count"`
	// Name of the Compute Engine [machine type](https://cloud.google.com/compute/docs/machine-resource) to use for login nodes, e.g. 'n2-standard-2'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_hypercomputecluster_cluster#machine_type GoogleHypercomputeclusterCluster#machine_type}
	MachineType *string `field:"required" json:"machineType" yaml:"machineType"`
	// Name of the zone in which login nodes should run, e.g., 'us-central1-a'. Must be in the same region as the cluster, and must match the zone of any other resources specified in the cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_hypercomputecluster_cluster#zone GoogleHypercomputeclusterCluster#zone}
	Zone *string `field:"required" json:"zone" yaml:"zone"`
	// boot_disk block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_hypercomputecluster_cluster#boot_disk GoogleHypercomputeclusterCluster#boot_disk}
	BootDisk *GoogleHypercomputeclusterClusterOrchestratorSlurmLoginNodesBootDisk `field:"optional" json:"bootDisk" yaml:"bootDisk"`
	// Whether [OS Login](https://cloud.google.com/compute/docs/oslogin) should be enabled on login node instances.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_hypercomputecluster_cluster#enable_os_login GoogleHypercomputeclusterCluster#enable_os_login}
	EnableOsLogin interface{} `field:"optional" json:"enableOsLogin" yaml:"enableOsLogin"`
	// Whether login node instances should be assigned [external IP addresses](https://cloud.google.com/compute/docs/ip-addresses#externaladdresses).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_hypercomputecluster_cluster#enable_public_ips GoogleHypercomputeclusterCluster#enable_public_ips}
	EnablePublicIps interface{} `field:"optional" json:"enablePublicIps" yaml:"enablePublicIps"`
	// [Labels](https://cloud.google.com/compute/docs/labeling-resources) that should be applied to each login node instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_hypercomputecluster_cluster#labels GoogleHypercomputeclusterCluster#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// [Startup script](https://cloud.google.com/compute/docs/instances/startup-scripts/linux) to be run on each login node instance. Max 256KB. The script must complete within the system-defined default timeout of 5 minutes. For tasks that require more time, consider running them in the background using methods such as '&' or 'nohup'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_hypercomputecluster_cluster#startup_script GoogleHypercomputeclusterCluster#startup_script}
	StartupScript *string `field:"optional" json:"startupScript" yaml:"startupScript"`
	// storage_configs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_hypercomputecluster_cluster#storage_configs GoogleHypercomputeclusterCluster#storage_configs}
	StorageConfigs interface{} `field:"optional" json:"storageConfigs" yaml:"storageConfigs"`
}

