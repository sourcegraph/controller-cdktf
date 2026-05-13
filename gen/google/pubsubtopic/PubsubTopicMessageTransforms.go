package pubsubtopic


type PubsubTopicMessageTransforms struct {
	// ai_inference block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/pubsub_topic#ai_inference PubsubTopic#ai_inference}
	AiInference *PubsubTopicMessageTransformsAiInference `field:"optional" json:"aiInference" yaml:"aiInference"`
	// Controls whether or not to use this transform.
	//
	// If not set or 'false',
	// the transform will be applied to messages. Default: 'true'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/pubsub_topic#disabled PubsubTopic#disabled}
	Disabled interface{} `field:"optional" json:"disabled" yaml:"disabled"`
	// javascript_udf block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/pubsub_topic#javascript_udf PubsubTopic#javascript_udf}
	JavascriptUdf *PubsubTopicMessageTransformsJavascriptUdf `field:"optional" json:"javascriptUdf" yaml:"javascriptUdf"`
}

