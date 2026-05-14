package googlecesevaluation


type GoogleCesEvaluationScenarioScenarioExpectationsAgentResponseChunksImage struct {
	// Raw bytes of the image.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_ces_evaluation#data GoogleCesEvaluation#data}
	Data *string `field:"required" json:"data" yaml:"data"`
	// The IANA standard MIME type of the source data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_ces_evaluation#mime_type GoogleCesEvaluation#mime_type}
	MimeType *string `field:"required" json:"mimeType" yaml:"mimeType"`
}

