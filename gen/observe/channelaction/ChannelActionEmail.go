package channelaction


type ChannelActionEmail struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/channel_action#body ChannelAction#body}.
	Body *string `field:"required" json:"body" yaml:"body"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/channel_action#subject ChannelAction#subject}.
	Subject *string `field:"required" json:"subject" yaml:"subject"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/channel_action#to ChannelAction#to}.
	To *[]*string `field:"required" json:"to" yaml:"to"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/channel_action#is_html ChannelAction#is_html}.
	IsHtml interface{} `field:"optional" json:"isHtml" yaml:"isHtml"`
}

