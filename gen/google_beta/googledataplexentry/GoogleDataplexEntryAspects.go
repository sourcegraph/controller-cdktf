package googledataplexentry


type GoogleDataplexEntryAspects struct {
	// aspect block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dataplex_entry#aspect GoogleDataplexEntry#aspect}
	Aspect *GoogleDataplexEntryAspectsAspect `field:"required" json:"aspect" yaml:"aspect"`
	// Depending on how the aspect is attached to the entry, the format of the aspect key can be one of the following:.
	//
	// If the aspect is attached directly to the entry: {project_number}.{locationId}.{aspectTypeId}
	// If the aspect is attached to an entry's path: {project_number}.{locationId}.{aspectTypeId}@{path}
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dataplex_entry#aspect_key GoogleDataplexEntry#aspect_key}
	AspectKey *string `field:"required" json:"aspectKey" yaml:"aspectKey"`
}

