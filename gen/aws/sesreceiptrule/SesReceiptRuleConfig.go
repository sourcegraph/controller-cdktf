package sesreceiptrule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SesReceiptRuleConfig struct {
	// Experimental.
	Connection any `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count any `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]any `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/ses_receipt_rule#name SesReceiptRule#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/ses_receipt_rule#rule_set_name SesReceiptRule#rule_set_name}.
	RuleSetName *string `field:"required" json:"ruleSetName" yaml:"ruleSetName"`
	// add_header_action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/ses_receipt_rule#add_header_action SesReceiptRule#add_header_action}
	AddHeaderAction any `field:"optional" json:"addHeaderAction" yaml:"addHeaderAction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/ses_receipt_rule#after SesReceiptRule#after}.
	After *string `field:"optional" json:"after" yaml:"after"`
	// bounce_action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/ses_receipt_rule#bounce_action SesReceiptRule#bounce_action}
	BounceAction any `field:"optional" json:"bounceAction" yaml:"bounceAction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/ses_receipt_rule#enabled SesReceiptRule#enabled}.
	Enabled any `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/ses_receipt_rule#id SesReceiptRule#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// lambda_action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/ses_receipt_rule#lambda_action SesReceiptRule#lambda_action}
	LambdaAction any `field:"optional" json:"lambdaAction" yaml:"lambdaAction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/ses_receipt_rule#recipients SesReceiptRule#recipients}.
	Recipients *[]*string `field:"optional" json:"recipients" yaml:"recipients"`
	// s3_action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/ses_receipt_rule#s3_action SesReceiptRule#s3_action}
	S3Action any `field:"optional" json:"s3Action" yaml:"s3Action"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/ses_receipt_rule#scan_enabled SesReceiptRule#scan_enabled}.
	ScanEnabled any `field:"optional" json:"scanEnabled" yaml:"scanEnabled"`
	// sns_action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/ses_receipt_rule#sns_action SesReceiptRule#sns_action}
	SnsAction any `field:"optional" json:"snsAction" yaml:"snsAction"`
	// stop_action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/ses_receipt_rule#stop_action SesReceiptRule#stop_action}
	StopAction any `field:"optional" json:"stopAction" yaml:"stopAction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/ses_receipt_rule#tls_policy SesReceiptRule#tls_policy}.
	TlsPolicy *string `field:"optional" json:"tlsPolicy" yaml:"tlsPolicy"`
	// workmail_action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/ses_receipt_rule#workmail_action SesReceiptRule#workmail_action}
	WorkmailAction any `field:"optional" json:"workmailAction" yaml:"workmailAction"`
}
