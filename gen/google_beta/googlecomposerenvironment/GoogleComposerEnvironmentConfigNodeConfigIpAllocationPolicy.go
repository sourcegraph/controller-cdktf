package googlecomposerenvironment


type GoogleComposerEnvironmentConfigNodeConfigIpAllocationPolicy struct {
	// The IP address range used to allocate IP addresses to pods in the cluster.
	//
	// For Cloud Composer environments in versions composer-1.*.*-airflow-*.*.*, this field is applicable only when use_ip_aliases is true. Set to blank to have GKE choose a range with the default size. Set to /netmask (e.g. /14) to have GKE choose a range with a specific netmask. Set to a CIDR notation (e.g. 10.96.0.0/14) from the RFC-1918 private networks (e.g. 10.0.0.0/8, 172.16.0.0/12, 192.168.0.0/16) to pick a specific range to use. Specify either cluster_secondary_range_name or cluster_ipv4_cidr_block but not both.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_composer_environment#cluster_ipv4_cidr_block GoogleComposerEnvironment#cluster_ipv4_cidr_block}
	ClusterIpv4CidrBlock *string `field:"optional" json:"clusterIpv4CidrBlock" yaml:"clusterIpv4CidrBlock"`
	// The name of the cluster's secondary range used to allocate IP addresses to pods.
	//
	// Specify either cluster_secondary_range_name or cluster_ipv4_cidr_block but not both. For Cloud Composer environments in versions composer-1.*.*-airflow-*.*.*, this field is applicable only when use_ip_aliases is true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_composer_environment#cluster_secondary_range_name GoogleComposerEnvironment#cluster_secondary_range_name}
	ClusterSecondaryRangeName *string `field:"optional" json:"clusterSecondaryRangeName" yaml:"clusterSecondaryRangeName"`
	// The IP address range used to allocate IP addresses in this cluster.
	//
	// For Cloud Composer environments in versions composer-1.*.*-airflow-*.*.*, this field is applicable only when use_ip_aliases is true. Set to blank to have GKE choose a range with the default size. Set to /netmask (e.g. /14) to have GKE choose a range with a specific netmask. Set to a CIDR notation (e.g. 10.96.0.0/14) from the RFC-1918 private networks (e.g. 10.0.0.0/8, 172.16.0.0/12, 192.168.0.0/16) to pick a specific range to use. Specify either services_secondary_range_name or services_ipv4_cidr_block but not both.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_composer_environment#services_ipv4_cidr_block GoogleComposerEnvironment#services_ipv4_cidr_block}
	ServicesIpv4CidrBlock *string `field:"optional" json:"servicesIpv4CidrBlock" yaml:"servicesIpv4CidrBlock"`
	// The name of the services' secondary range used to allocate IP addresses to the cluster.
	//
	// Specify either services_secondary_range_name or services_ipv4_cidr_block but not both. For Cloud Composer environments in versions composer-1.*.*-airflow-*.*.*, this field is applicable only when use_ip_aliases is true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_composer_environment#services_secondary_range_name GoogleComposerEnvironment#services_secondary_range_name}
	ServicesSecondaryRangeName *string `field:"optional" json:"servicesSecondaryRangeName" yaml:"servicesSecondaryRangeName"`
	// Whether or not to enable Alias IPs in the GKE cluster.
	//
	// If true, a VPC-native cluster is created. Defaults to true if the ip_allocation_policy block is present in config. This field is only supported for Cloud Composer environments in versions composer-1.*.*-airflow-*.*.*. Environments in newer versions always use VPC-native GKE clusters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_composer_environment#use_ip_aliases GoogleComposerEnvironment#use_ip_aliases}
	UseIpAliases interface{} `field:"optional" json:"useIpAliases" yaml:"useIpAliases"`
}

