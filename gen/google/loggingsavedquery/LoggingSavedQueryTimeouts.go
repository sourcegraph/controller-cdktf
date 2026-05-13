package loggingsavedquery


type LoggingSavedQueryTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/logging_saved_query#create LoggingSavedQuery#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/logging_saved_query#delete LoggingSavedQuery#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/logging_saved_query#update LoggingSavedQuery#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

