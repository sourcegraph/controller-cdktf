package googledocumentaiwarehousedocumentschema


type GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsPropertyTypeOptionsPropertyDefinitionsEnumTypeOptions struct {
	// List of possible enum values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_document_ai_warehouse_document_schema#possible_values GoogleDocumentAiWarehouseDocumentSchema#possible_values}
	PossibleValues *[]*string `field:"required" json:"possibleValues" yaml:"possibleValues"`
	// Make sure the enum property value provided in the document is in the possile value list during document creation.
	//
	// The validation check runs by default.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_document_ai_warehouse_document_schema#validation_check_disabled GoogleDocumentAiWarehouseDocumentSchema#validation_check_disabled}
	ValidationCheckDisabled interface{} `field:"optional" json:"validationCheckDisabled" yaml:"validationCheckDisabled"`
}

