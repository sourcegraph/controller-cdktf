package googledatacatalogtag


type GoogleDataCatalogTagFields struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_catalog_tag#field_name GoogleDataCatalogTag#field_name}.
	FieldName *string `field:"required" json:"fieldName" yaml:"fieldName"`
	// Holds the value for a tag field with boolean type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_catalog_tag#bool_value GoogleDataCatalogTag#bool_value}
	BoolValue interface{} `field:"optional" json:"boolValue" yaml:"boolValue"`
	// Holds the value for a tag field with double type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_catalog_tag#double_value GoogleDataCatalogTag#double_value}
	DoubleValue *float64 `field:"optional" json:"doubleValue" yaml:"doubleValue"`
	// The display name of the enum value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_catalog_tag#enum_value GoogleDataCatalogTag#enum_value}
	EnumValue *string `field:"optional" json:"enumValue" yaml:"enumValue"`
	// Holds the value for a tag field with string type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_catalog_tag#string_value GoogleDataCatalogTag#string_value}
	StringValue *string `field:"optional" json:"stringValue" yaml:"stringValue"`
	// Holds the value for a tag field with timestamp type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_catalog_tag#timestamp_value GoogleDataCatalogTag#timestamp_value}
	TimestampValue *string `field:"optional" json:"timestampValue" yaml:"timestampValue"`
}

