package keyspacestable

type KeyspacesTableSchemaDefinition struct {
	// column block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/keyspaces_table#column KeyspacesTable#column}
	Column any `field:"required" json:"column" yaml:"column"`
	// partition_key block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/keyspaces_table#partition_key KeyspacesTable#partition_key}
	PartitionKey any `field:"required" json:"partitionKey" yaml:"partitionKey"`
	// clustering_key block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/keyspaces_table#clustering_key KeyspacesTable#clustering_key}
	ClusteringKey any `field:"optional" json:"clusteringKey" yaml:"clusteringKey"`
	// static_column block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/keyspaces_table#static_column KeyspacesTable#static_column}
	StaticColumn any `field:"optional" json:"staticColumn" yaml:"staticColumn"`
}
