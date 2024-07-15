package googlecontainercluster


type GoogleContainerClusterResourceUsageExportConfig struct {
	// bigquery_destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_cluster#bigquery_destination GoogleContainerCluster#bigquery_destination}
	BigqueryDestination *GoogleContainerClusterResourceUsageExportConfigBigqueryDestination `field:"required" json:"bigqueryDestination" yaml:"bigqueryDestination"`
	// Whether to enable network egress metering for this cluster.
	//
	// If enabled, a daemonset will be created in the cluster to meter network egress traffic.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_cluster#enable_network_egress_metering GoogleContainerCluster#enable_network_egress_metering}
	EnableNetworkEgressMetering interface{} `field:"optional" json:"enableNetworkEgressMetering" yaml:"enableNetworkEgressMetering"`
	// Whether to enable resource consumption metering on this cluster.
	//
	// When enabled, a table will be created in the resource export BigQuery dataset to store resource consumption data. The resulting table can be joined with the resource usage table or with BigQuery billing export. Defaults to true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_cluster#enable_resource_consumption_metering GoogleContainerCluster#enable_resource_consumption_metering}
	EnableResourceConsumptionMetering interface{} `field:"optional" json:"enableResourceConsumptionMetering" yaml:"enableResourceConsumptionMetering"`
}

