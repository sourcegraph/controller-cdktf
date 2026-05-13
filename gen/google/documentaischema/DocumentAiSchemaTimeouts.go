package documentaischema


type DocumentAiSchemaTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/document_ai_schema#create DocumentAiSchema#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/document_ai_schema#delete DocumentAiSchema#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/document_ai_schema#update DocumentAiSchema#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

