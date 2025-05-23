package apphubservice


type ApphubServiceAttributesCriticality struct {
	// Criticality type. Possible values: ["MISSION_CRITICAL", "HIGH", "MEDIUM", "LOW"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.32.0/docs/resources/apphub_service#type ApphubService#type}
	Type *string `field:"required" json:"type" yaml:"type"`
}

