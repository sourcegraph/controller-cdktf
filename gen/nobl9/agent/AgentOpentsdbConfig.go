package agent


type AgentOpentsdbConfig struct {
	// OpenTSDB cluster URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/agent#url Agent#url}
	Url *string `field:"required" json:"url" yaml:"url"`
}

