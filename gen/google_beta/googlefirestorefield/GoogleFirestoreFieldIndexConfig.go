package googlefirestorefield


type GoogleFirestoreFieldIndexConfig struct {
	// indexes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_firestore_field#indexes GoogleFirestoreField#indexes}
	Indexes interface{} `field:"optional" json:"indexes" yaml:"indexes"`
}

