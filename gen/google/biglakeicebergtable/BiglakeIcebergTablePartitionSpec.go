package biglakeicebergtable


type BiglakeIcebergTablePartitionSpec struct {
	// fields block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/biglake_iceberg_table#fields BiglakeIcebergTable#fields}
	Fields interface{} `field:"required" json:"fields" yaml:"fields"`
}

