package googledialogflowcxplaybook


type GoogleDialogflowCxPlaybookInstruction struct {
	// General guidelines for the playbook.
	//
	// These are unstructured instructions that are not directly part of the goal, e.g. "Always be polite". It's valid for this text to be long and used instead of steps altogether.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_playbook#guidelines GoogleDialogflowCxPlaybook#guidelines}
	Guidelines *string `field:"optional" json:"guidelines" yaml:"guidelines"`
	// steps block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_cx_playbook#steps GoogleDialogflowCxPlaybook#steps}
	Steps interface{} `field:"optional" json:"steps" yaml:"steps"`
}

