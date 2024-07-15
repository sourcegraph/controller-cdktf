package googlepubsubtopiciammember


type GooglePubsubTopicIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_pubsub_topic_iam_member#expression GooglePubsubTopicIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_pubsub_topic_iam_member#title GooglePubsubTopicIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_pubsub_topic_iam_member#description GooglePubsubTopicIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

