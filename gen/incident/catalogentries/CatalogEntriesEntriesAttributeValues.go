package catalogentries


type CatalogEntriesEntriesAttributeValues struct {
	// The value of this element of the array, in a format suitable for this attribute type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.17.1/docs/resources/catalog_entries#array_value CatalogEntries#array_value}
	ArrayValue *[]*string `field:"optional" json:"arrayValue" yaml:"arrayValue"`
	// The value of this attribute, in a format suitable for this attribute type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.17.1/docs/resources/catalog_entries#value CatalogEntries#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

