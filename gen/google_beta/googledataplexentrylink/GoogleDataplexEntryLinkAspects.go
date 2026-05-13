package googledataplexentrylink


type GoogleDataplexEntryLinkAspects struct {
	// aspect block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dataplex_entry_link#aspect GoogleDataplexEntryLink#aspect}
	Aspect *GoogleDataplexEntryLinkAspectsAspect `field:"required" json:"aspect" yaml:"aspect"`
	// The map keys of the Aspects which the service should modify.
	//
	// It should be the aspect type reference in the format '{project_number}.{location_id}.{aspect_type_id}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dataplex_entry_link#aspect_key GoogleDataplexEntryLink#aspect_key}
	AspectKey *string `field:"required" json:"aspectKey" yaml:"aspectKey"`
}

