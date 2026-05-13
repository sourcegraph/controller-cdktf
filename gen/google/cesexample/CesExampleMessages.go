package cesexample


type CesExampleMessages struct {
	// chunks block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_example#chunks CesExample#chunks}
	Chunks interface{} `field:"optional" json:"chunks" yaml:"chunks"`
	// The role within the conversation, e.g., user, agent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_example#role CesExample#role}
	Role *string `field:"optional" json:"role" yaml:"role"`
}

