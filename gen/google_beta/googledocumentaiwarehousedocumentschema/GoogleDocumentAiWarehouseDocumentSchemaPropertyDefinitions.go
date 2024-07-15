package googledocumentaiwarehousedocumentschema


type GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitions struct {
	// The name of the metadata property.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_document_ai_warehouse_document_schema#name GoogleDocumentAiWarehouseDocumentSchema#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// date_time_type_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_document_ai_warehouse_document_schema#date_time_type_options GoogleDocumentAiWarehouseDocumentSchema#date_time_type_options}
	DateTimeTypeOptions *GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsDateTimeTypeOptions `field:"optional" json:"dateTimeTypeOptions" yaml:"dateTimeTypeOptions"`
	// The display-name for the property, used for front-end.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_document_ai_warehouse_document_schema#display_name GoogleDocumentAiWarehouseDocumentSchema#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// enum_type_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_document_ai_warehouse_document_schema#enum_type_options GoogleDocumentAiWarehouseDocumentSchema#enum_type_options}
	EnumTypeOptions *GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsEnumTypeOptions `field:"optional" json:"enumTypeOptions" yaml:"enumTypeOptions"`
	// float_type_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_document_ai_warehouse_document_schema#float_type_options GoogleDocumentAiWarehouseDocumentSchema#float_type_options}
	FloatTypeOptions *GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsFloatTypeOptions `field:"optional" json:"floatTypeOptions" yaml:"floatTypeOptions"`
	// integer_type_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_document_ai_warehouse_document_schema#integer_type_options GoogleDocumentAiWarehouseDocumentSchema#integer_type_options}
	IntegerTypeOptions *GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsIntegerTypeOptions `field:"optional" json:"integerTypeOptions" yaml:"integerTypeOptions"`
	// Whether the property can be filtered. If this is a sub-property, all the parent properties must be marked filterable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_document_ai_warehouse_document_schema#is_filterable GoogleDocumentAiWarehouseDocumentSchema#is_filterable}
	IsFilterable interface{} `field:"optional" json:"isFilterable" yaml:"isFilterable"`
	// Whether the property is user supplied metadata.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_document_ai_warehouse_document_schema#is_metadata GoogleDocumentAiWarehouseDocumentSchema#is_metadata}
	IsMetadata interface{} `field:"optional" json:"isMetadata" yaml:"isMetadata"`
	// Whether the property can have multiple values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_document_ai_warehouse_document_schema#is_repeatable GoogleDocumentAiWarehouseDocumentSchema#is_repeatable}
	IsRepeatable interface{} `field:"optional" json:"isRepeatable" yaml:"isRepeatable"`
	// Whether the property is mandatory.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_document_ai_warehouse_document_schema#is_required GoogleDocumentAiWarehouseDocumentSchema#is_required}
	IsRequired interface{} `field:"optional" json:"isRequired" yaml:"isRequired"`
	// Indicates that the property should be included in a global search.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_document_ai_warehouse_document_schema#is_searchable GoogleDocumentAiWarehouseDocumentSchema#is_searchable}
	IsSearchable interface{} `field:"optional" json:"isSearchable" yaml:"isSearchable"`
	// map_type_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_document_ai_warehouse_document_schema#map_type_options GoogleDocumentAiWarehouseDocumentSchema#map_type_options}
	MapTypeOptions *GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsMapTypeOptions `field:"optional" json:"mapTypeOptions" yaml:"mapTypeOptions"`
	// property_type_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_document_ai_warehouse_document_schema#property_type_options GoogleDocumentAiWarehouseDocumentSchema#property_type_options}
	PropertyTypeOptions *GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsPropertyTypeOptions `field:"optional" json:"propertyTypeOptions" yaml:"propertyTypeOptions"`
	// Stores the retrieval importance. Possible values: ["HIGHEST", "HIGHER", "HIGH", "MEDIUM", "LOW", "LOWEST"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_document_ai_warehouse_document_schema#retrieval_importance GoogleDocumentAiWarehouseDocumentSchema#retrieval_importance}
	RetrievalImportance *string `field:"optional" json:"retrievalImportance" yaml:"retrievalImportance"`
	// schema_sources block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_document_ai_warehouse_document_schema#schema_sources GoogleDocumentAiWarehouseDocumentSchema#schema_sources}
	SchemaSources interface{} `field:"optional" json:"schemaSources" yaml:"schemaSources"`
	// text_type_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_document_ai_warehouse_document_schema#text_type_options GoogleDocumentAiWarehouseDocumentSchema#text_type_options}
	TextTypeOptions *GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsTextTypeOptions `field:"optional" json:"textTypeOptions" yaml:"textTypeOptions"`
	// timestamp_type_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_document_ai_warehouse_document_schema#timestamp_type_options GoogleDocumentAiWarehouseDocumentSchema#timestamp_type_options}
	TimestampTypeOptions *GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsTimestampTypeOptions `field:"optional" json:"timestampTypeOptions" yaml:"timestampTypeOptions"`
}

