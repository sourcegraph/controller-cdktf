package googlesecretmanagersecret


type GoogleSecretManagerSecretTopics struct {
	// The resource name of the Pub/Sub topic that will be published to, in the following format: projects/* /topics/*.
	//
	// For publication to succeed, the Secret Manager Service Agent service account must have pubsub.publisher permissions on the topic.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_secret_manager_secret#name GoogleSecretManagerSecret#name}
	//
	// Note: The above comment contained a comment block ending sequence (* followed by /). We have introduced a space between to prevent syntax errors. Please ignore the space.
	Name *string `field:"required" json:"name" yaml:"name"`
}

