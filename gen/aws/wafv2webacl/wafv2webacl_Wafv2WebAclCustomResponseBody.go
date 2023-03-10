package wafv2webacl


type Wafv2WebAclCustomResponseBody struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/wafv2_web_acl#content Wafv2WebAcl#content}.
	Content *string `field:"required" json:"content" yaml:"content"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/wafv2_web_acl#content_type Wafv2WebAcl#content_type}.
	ContentType *string `field:"required" json:"contentType" yaml:"contentType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/wafv2_web_acl#key Wafv2WebAcl#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
}

