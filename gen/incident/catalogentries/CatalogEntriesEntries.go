package catalogentries


type CatalogEntriesEntries struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.17.1/docs/resources/catalog_entries#attribute_values CatalogEntries#attribute_values}.
	AttributeValues interface{} `field:"required" json:"attributeValues" yaml:"attributeValues"`
	// Name is the human readable name of this entry.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.17.1/docs/resources/catalog_entries#name CatalogEntries#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Optional aliases that can be used to reference this entry.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.17.1/docs/resources/catalog_entries#aliases CatalogEntries#aliases}
	Aliases *[]*string `field:"optional" json:"aliases" yaml:"aliases"`
	// When catalog type is ranked, this is used to help order things.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.17.1/docs/resources/catalog_entries#rank CatalogEntries#rank}
	Rank *float64 `field:"optional" json:"rank" yaml:"rank"`
}

