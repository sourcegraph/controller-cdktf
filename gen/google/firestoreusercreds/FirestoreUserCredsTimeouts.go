package firestoreusercreds


type FirestoreUserCredsTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/firestore_user_creds#create FirestoreUserCreds#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/firestore_user_creds#delete FirestoreUserCreds#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

