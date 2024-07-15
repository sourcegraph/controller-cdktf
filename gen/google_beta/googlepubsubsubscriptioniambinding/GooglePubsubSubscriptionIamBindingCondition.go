package googlepubsubsubscriptioniambinding


type GooglePubsubSubscriptionIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_pubsub_subscription_iam_binding#expression GooglePubsubSubscriptionIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_pubsub_subscription_iam_binding#title GooglePubsubSubscriptionIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_pubsub_subscription_iam_binding#description GooglePubsubSubscriptionIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

