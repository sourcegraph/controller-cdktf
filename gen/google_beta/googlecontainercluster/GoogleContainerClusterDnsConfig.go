package googlecontainercluster


type GoogleContainerClusterDnsConfig struct {
	// Enable additive VPC scope DNS in a GKE cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_cluster#additive_vpc_scope_dns_domain GoogleContainerCluster#additive_vpc_scope_dns_domain}
	AdditiveVpcScopeDnsDomain *string `field:"optional" json:"additiveVpcScopeDnsDomain" yaml:"additiveVpcScopeDnsDomain"`
	// Which in-cluster DNS provider should be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_cluster#cluster_dns GoogleContainerCluster#cluster_dns}
	ClusterDns *string `field:"optional" json:"clusterDns" yaml:"clusterDns"`
	// The suffix used for all cluster service records.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_cluster#cluster_dns_domain GoogleContainerCluster#cluster_dns_domain}
	ClusterDnsDomain *string `field:"optional" json:"clusterDnsDomain" yaml:"clusterDnsDomain"`
	// The scope of access to cluster DNS records.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_cluster#cluster_dns_scope GoogleContainerCluster#cluster_dns_scope}
	ClusterDnsScope *string `field:"optional" json:"clusterDnsScope" yaml:"clusterDnsScope"`
}

