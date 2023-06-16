package customhostname

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CustomHostnameConfig struct {
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
	// Hostname you intend to request a certificate for. **Modifying this attribute will force creation of a new resource.**.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/custom_hostname#hostname CustomHostname#hostname}
	Hostname *string `field:"required" json:"hostname" yaml:"hostname"`
	// The zone identifier to target for the resource. **Modifying this attribute will force creation of a new resource.**.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/custom_hostname#zone_id CustomHostname#zone_id}
	ZoneId *string `field:"required" json:"zoneId" yaml:"zoneId"`
	// Custom metadata associated with custom hostname.
	//
	// Only supports primitive string values, all other values are accessible via the API directly.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/custom_hostname#custom_metadata CustomHostname#custom_metadata}
	CustomMetadata *map[string]*string `field:"optional" json:"customMetadata" yaml:"customMetadata"`
	// The custom origin server used for certificates.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/custom_hostname#custom_origin_server CustomHostname#custom_origin_server}
	CustomOriginServer *string `field:"optional" json:"customOriginServer" yaml:"customOriginServer"`
	// The [custom origin SNI](https://developers.cloudflare.com/ssl/ssl-for-saas/hostname-specific-behavior/custom-origin) used for certificates.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/custom_hostname#custom_origin_sni CustomHostname#custom_origin_sni}
	CustomOriginSni *string `field:"optional" json:"customOriginSni" yaml:"customOriginSni"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/custom_hostname#id CustomHostname#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// ssl block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/custom_hostname#ssl CustomHostname#ssl}
	Ssl interface{} `field:"optional" json:"ssl" yaml:"ssl"`
	// Whether to wait for a custom hostname SSL sub-object to reach status `pending_validation` during creation. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/custom_hostname#wait_for_ssl_pending_validation CustomHostname#wait_for_ssl_pending_validation}
	WaitForSslPendingValidation interface{} `field:"optional" json:"waitForSslPendingValidation" yaml:"waitForSslPendingValidation"`
}

