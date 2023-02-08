package googlepubsubschema


type GooglePubsubSchemaTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_pubsub_schema#create GooglePubsubSchema#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_pubsub_schema#delete GooglePubsubSchema#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

