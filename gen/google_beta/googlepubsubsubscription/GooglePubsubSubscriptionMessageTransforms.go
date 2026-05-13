package googlepubsubsubscription


type GooglePubsubSubscriptionMessageTransforms struct {
	// ai_inference block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_pubsub_subscription#ai_inference GooglePubsubSubscription#ai_inference}
	AiInference *GooglePubsubSubscriptionMessageTransformsAiInference `field:"optional" json:"aiInference" yaml:"aiInference"`
	// Controls whether or not to use this transform.
	//
	// If not set or 'false',
	// the transform will be applied to messages. Default: 'true'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_pubsub_subscription#disabled GooglePubsubSubscription#disabled}
	Disabled interface{} `field:"optional" json:"disabled" yaml:"disabled"`
	// javascript_udf block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_pubsub_subscription#javascript_udf GooglePubsubSubscription#javascript_udf}
	JavascriptUdf *GooglePubsubSubscriptionMessageTransformsJavascriptUdf `field:"optional" json:"javascriptUdf" yaml:"javascriptUdf"`
}

