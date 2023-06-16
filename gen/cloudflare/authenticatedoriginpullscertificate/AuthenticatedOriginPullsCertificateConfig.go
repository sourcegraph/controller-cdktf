package authenticatedoriginpullscertificate

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AuthenticatedOriginPullsCertificateConfig struct {
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
	// The public client certificate. **Modifying this attribute will force creation of a new resource.**.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/authenticated_origin_pulls_certificate#certificate AuthenticatedOriginPullsCertificate#certificate}
	Certificate *string `field:"required" json:"certificate" yaml:"certificate"`
	// The private key of the client certificate. **Modifying this attribute will force creation of a new resource.**.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/authenticated_origin_pulls_certificate#private_key AuthenticatedOriginPullsCertificate#private_key}
	PrivateKey *string `field:"required" json:"privateKey" yaml:"privateKey"`
	// The form of Authenticated Origin Pulls to upload the certificate to.
	//
	// Available values: `per-zone`, `per-hostname`. **Modifying this attribute will force creation of a new resource.**
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/authenticated_origin_pulls_certificate#type AuthenticatedOriginPullsCertificate#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// The zone identifier to target for the resource. **Modifying this attribute will force creation of a new resource.**.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/authenticated_origin_pulls_certificate#zone_id AuthenticatedOriginPullsCertificate#zone_id}
	ZoneId *string `field:"required" json:"zoneId" yaml:"zoneId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/authenticated_origin_pulls_certificate#id AuthenticatedOriginPullsCertificate#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/authenticated_origin_pulls_certificate#timeouts AuthenticatedOriginPullsCertificate#timeouts}
	Timeouts *AuthenticatedOriginPullsCertificateTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

