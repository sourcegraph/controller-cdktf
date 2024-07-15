package googledatacatalogtagtemplate


type GoogleDataCatalogTagTemplateFieldsType struct {
	// enum_type block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_catalog_tag_template#enum_type GoogleDataCatalogTagTemplate#enum_type}
	EnumType *GoogleDataCatalogTagTemplateFieldsTypeEnumType `field:"optional" json:"enumType" yaml:"enumType"`
	// Represents primitive types - string, bool etc.
	//
	// Exactly one of 'primitive_type' or 'enum_type' must be set Possible values: ["DOUBLE", "STRING", "BOOL", "TIMESTAMP"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_catalog_tag_template#primitive_type GoogleDataCatalogTagTemplate#primitive_type}
	PrimitiveType *string `field:"optional" json:"primitiveType" yaml:"primitiveType"`
}

