package googlebiglakeicebergtable


type GoogleBiglakeIcebergTableSchemaFields struct {
	// The unique identifier of the field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_biglake_iceberg_table#id GoogleBiglakeIcebergTable#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *float64 `field:"required" json:"id" yaml:"id"`
	// The name of the field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_biglake_iceberg_table#name GoogleBiglakeIcebergTable#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Whether the field is required.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_biglake_iceberg_table#required GoogleBiglakeIcebergTable#required}
	Required interface{} `field:"required" json:"required" yaml:"required"`
	// The type of the field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_biglake_iceberg_table#type GoogleBiglakeIcebergTable#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// A description of the field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_biglake_iceberg_table#doc GoogleBiglakeIcebergTable#doc}
	Doc *string `field:"optional" json:"doc" yaml:"doc"`
}

