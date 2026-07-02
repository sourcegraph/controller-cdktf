package biglakeicebergtable

type BiglakeIcebergTablePartitionSpec struct {
	// fields block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/biglake_iceberg_table#fields BiglakeIcebergTable#fields}
	Fields any `field:"required" json:"fields" yaml:"fields"`
}
