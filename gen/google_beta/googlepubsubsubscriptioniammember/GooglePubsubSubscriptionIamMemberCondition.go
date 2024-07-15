package googlepubsubsubscriptioniammember


type GooglePubsubSubscriptionIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_pubsub_subscription_iam_member#expression GooglePubsubSubscriptionIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_pubsub_subscription_iam_member#title GooglePubsubSubscriptionIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_pubsub_subscription_iam_member#description GooglePubsubSubscriptionIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

