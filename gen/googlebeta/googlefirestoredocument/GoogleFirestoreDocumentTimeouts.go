package googlefirestoredocument


type GoogleFirestoreDocumentTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_firestore_document#create GoogleFirestoreDocument#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_firestore_document#delete GoogleFirestoreDocument#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_firestore_document#update GoogleFirestoreDocument#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

