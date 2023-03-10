package gluetrigger


type GlueTriggerTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_trigger#create GlueTrigger#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_trigger#delete GlueTrigger#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

