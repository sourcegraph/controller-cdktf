package googlefirestorefield


type GoogleFirestoreFieldTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_firestore_field#create GoogleFirestoreField#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_firestore_field#delete GoogleFirestoreField#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_firestore_field#update GoogleFirestoreField#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

