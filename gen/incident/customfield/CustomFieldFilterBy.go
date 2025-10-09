package customfield


type CustomFieldFilterBy struct {
	// This must be an attribute of the catalog type of this custom field.
	//
	// It must be an attribute that points to another catalog type (so not a plain string, number, or boolean attribute).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.17.1/docs/resources/custom_field#catalog_attribute_id CustomField#catalog_attribute_id}
	CatalogAttributeId *string `field:"required" json:"catalogAttributeId" yaml:"catalogAttributeId"`
	// This must be the ID of a custom field, which must have values of the same type as the attribute you are filtering by.
	//
	// When this filtering field is set on an incident, the options for this custom field will be filtered to only those with the attribute value that matches the value of the filtering field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.17.1/docs/resources/custom_field#custom_field_id CustomField#custom_field_id}
	CustomFieldId *string `field:"required" json:"customFieldId" yaml:"customFieldId"`
}

