package googlebiglakeicebergtable


type GoogleBiglakeIcebergTableSchema struct {
	// fields block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_biglake_iceberg_table#fields GoogleBiglakeIcebergTable#fields}
	Fields interface{} `field:"required" json:"fields" yaml:"fields"`
	// The field IDs that make up the identifier for the table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_biglake_iceberg_table#identifier_field_ids GoogleBiglakeIcebergTable#identifier_field_ids}
	IdentifierFieldIds *[]*float64 `field:"optional" json:"identifierFieldIds" yaml:"identifierFieldIds"`
	// The type of the schema.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_biglake_iceberg_table#type GoogleBiglakeIcebergTable#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

