package catalogentry


type CatalogEntryAttributeValues struct {
	// The ID of this attribute, usually loaded from the incident_catalog_type_attribute resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.17.1/docs/resources/catalog_entry#attribute CatalogEntry#attribute}
	Attribute *string `field:"required" json:"attribute" yaml:"attribute"`
	// The value of this element of the array, in a format suitable for this attribute type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.17.1/docs/resources/catalog_entry#array_value CatalogEntry#array_value}
	ArrayValue *[]*string `field:"optional" json:"arrayValue" yaml:"arrayValue"`
	// The value of this attribute, in a format suitable for this attribute type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.17.1/docs/resources/catalog_entry#value CatalogEntry#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

