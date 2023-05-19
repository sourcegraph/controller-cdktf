package authenticatedoriginpulls

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AuthenticatedOriginPullsConfig struct {
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
	// Whether to enable Authenticated Origin Pulls on the given zone or hostname.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/authenticated_origin_pulls#enabled AuthenticatedOriginPulls#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// The zone identifier to target for the resource. **Modifying this attribute will force creation of a new resource.**.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/authenticated_origin_pulls#zone_id AuthenticatedOriginPulls#zone_id}
	ZoneId *string `field:"required" json:"zoneId" yaml:"zoneId"`
	// The ID of an uploaded Authenticated Origin Pulls certificate.
	//
	// If no hostname is provided, this certificate will be used zone wide as Per-Zone Authenticated Origin Pulls.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/authenticated_origin_pulls#authenticated_origin_pulls_certificate AuthenticatedOriginPulls#authenticated_origin_pulls_certificate}
	AuthenticatedOriginPullsCertificate *string `field:"optional" json:"authenticatedOriginPullsCertificate" yaml:"authenticatedOriginPullsCertificate"`
	// Specify a hostname to enable Per-Hostname Authenticated Origin Pulls on, using the provided certificate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/authenticated_origin_pulls#hostname AuthenticatedOriginPulls#hostname}
	Hostname *string `field:"optional" json:"hostname" yaml:"hostname"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/authenticated_origin_pulls#id AuthenticatedOriginPulls#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

