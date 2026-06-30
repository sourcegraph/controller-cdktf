package codebuildwebhook

type CodebuildWebhookFilterGroup struct {
	// filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/codebuild_webhook#filter CodebuildWebhook#filter}
	Filter any `field:"optional" json:"filter" yaml:"filter"`
}
