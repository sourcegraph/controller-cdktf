package googlechroniclereferencelist


type GoogleChronicleReferenceListScopeInfoReferenceListScope struct {
	// Optional.
	//
	// The list of scope names of the reference list. The scope names should be
	// full resource names and should be of the format:
	// "projects/{project}/locations/{location}/instances/{instance}/dataAccessScopes/{scope_name}".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_reference_list#scope_names GoogleChronicleReferenceList#scope_names}
	ScopeNames *[]*string `field:"optional" json:"scopeNames" yaml:"scopeNames"`
}

