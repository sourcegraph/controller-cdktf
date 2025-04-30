package googlechroniclereferencelist


type GoogleChronicleReferenceListEntries struct {
	// Required. The value of the entry. Maximum length is 512 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_chronicle_reference_list#value GoogleChronicleReferenceList#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

