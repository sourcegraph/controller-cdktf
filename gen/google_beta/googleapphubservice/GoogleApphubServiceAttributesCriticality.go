package googleapphubservice


type GoogleApphubServiceAttributesCriticality struct {
	// Criticality type. Possible values: ["MISSION_CRITICAL", "HIGH", "MEDIUM", "LOW"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_apphub_service#type GoogleApphubService#type}
	Type *string `field:"required" json:"type" yaml:"type"`
}

