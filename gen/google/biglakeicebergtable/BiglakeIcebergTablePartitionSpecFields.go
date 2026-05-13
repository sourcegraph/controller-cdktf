package biglakeicebergtable


type BiglakeIcebergTablePartitionSpecFields struct {
	// The name of the partition field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/biglake_iceberg_table#name BiglakeIcebergTable#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The source field ID for the partition field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/biglake_iceberg_table#source_id BiglakeIcebergTable#source_id}
	SourceId *float64 `field:"required" json:"sourceId" yaml:"sourceId"`
	// The transform to apply to the source field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/biglake_iceberg_table#transform BiglakeIcebergTable#transform}
	Transform *string `field:"required" json:"transform" yaml:"transform"`
}

