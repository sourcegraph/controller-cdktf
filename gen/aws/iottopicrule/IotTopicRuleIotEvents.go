package iottopicrule


type IotTopicRuleIotEvents struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/iot_topic_rule#input_name IotTopicRule#input_name}.
	InputName *string `field:"required" json:"inputName" yaml:"inputName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/iot_topic_rule#role_arn IotTopicRule#role_arn}.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/iot_topic_rule#message_id IotTopicRule#message_id}.
	MessageId *string `field:"optional" json:"messageId" yaml:"messageId"`
}

