package cesapp


type CesAppDefaultChannelProfilePersonaProperty struct {
	// The persona of the channel. Possible values: UNKNOWN CONCISE CHATTY.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_app#persona CesApp#persona}
	Persona *string `field:"optional" json:"persona" yaml:"persona"`
}

