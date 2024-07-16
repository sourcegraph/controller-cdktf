package googleapigeeendpointattachment

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleApigeeEndpointAttachmentConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// ID of the endpoint attachment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_apigee_endpoint_attachment#endpoint_attachment_id GoogleApigeeEndpointAttachment#endpoint_attachment_id}
	EndpointAttachmentId *string `field:"required" json:"endpointAttachmentId" yaml:"endpointAttachmentId"`
	// Location of the endpoint attachment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_apigee_endpoint_attachment#location GoogleApigeeEndpointAttachment#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The Apigee Organization associated with the Apigee instance, in the format 'organizations/{{org_name}}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_apigee_endpoint_attachment#org_id GoogleApigeeEndpointAttachment#org_id}
	OrgId *string `field:"required" json:"orgId" yaml:"orgId"`
	// Format: projects/* /regions/* /serviceAttachments/*.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_apigee_endpoint_attachment#service_attachment GoogleApigeeEndpointAttachment#service_attachment}
	//
	// Note: The above comment contained a comment block ending sequence (* followed by /). We have introduced a space between to prevent syntax errors. Please ignore the space.
	ServiceAttachment *string `field:"required" json:"serviceAttachment" yaml:"serviceAttachment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_apigee_endpoint_attachment#id GoogleApigeeEndpointAttachment#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_apigee_endpoint_attachment#timeouts GoogleApigeeEndpointAttachment#timeouts}
	Timeouts *GoogleApigeeEndpointAttachmentTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

