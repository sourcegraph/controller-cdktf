package googledataplexentrylink


type GoogleDataplexEntryLinkEntryReferences struct {
	// The relative resource name of the referenced Entry, of the form: projects/{project_id_or_number}/locations/{location_id}/entryGroups/{entry_group_id}/entries/{entry_id}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dataplex_entry_link#name GoogleDataplexEntryLink#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The path in the Entry that is referenced in the Entry Link.
	//
	// Empty path denotes that the Entry itself is referenced in the Entry Link.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dataplex_entry_link#path GoogleDataplexEntryLink#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
	// The reference type of the Entry. Possible values: ["SOURCE", "TARGET"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dataplex_entry_link#type GoogleDataplexEntryLink#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

