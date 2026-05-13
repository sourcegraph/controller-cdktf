package pubsubsubscription


type PubsubSubscriptionMessageTransforms struct {
	// ai_inference block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/pubsub_subscription#ai_inference PubsubSubscription#ai_inference}
	AiInference *PubsubSubscriptionMessageTransformsAiInference `field:"optional" json:"aiInference" yaml:"aiInference"`
	// Controls whether or not to use this transform.
	//
	// If not set or 'false',
	// the transform will be applied to messages. Default: 'true'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/pubsub_subscription#disabled PubsubSubscription#disabled}
	Disabled interface{} `field:"optional" json:"disabled" yaml:"disabled"`
	// javascript_udf block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/pubsub_subscription#javascript_udf PubsubSubscription#javascript_udf}
	JavascriptUdf *PubsubSubscriptionMessageTransformsJavascriptUdf `field:"optional" json:"javascriptUdf" yaml:"javascriptUdf"`
}

