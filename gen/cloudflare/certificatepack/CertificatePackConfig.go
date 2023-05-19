package certificatepack

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CertificatePackConfig struct {
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
	// Which certificate authority to issue the certificate pack.
	//
	// Available values: `digicert`, `lets_encrypt`, `google`. **Modifying this attribute will force creation of a new resource.**
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/certificate_pack#certificate_authority CertificatePack#certificate_authority}
	CertificateAuthority *string `field:"required" json:"certificateAuthority" yaml:"certificateAuthority"`
	// List of hostnames to provision the certificate pack for.
	//
	// The zone name must be included as a host. Note: If using Let's Encrypt, you cannot use individual subdomains and only a wildcard for subdomain is available. **Modifying this attribute will force creation of a new resource.**
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/certificate_pack#hosts CertificatePack#hosts}
	Hosts *[]*string `field:"required" json:"hosts" yaml:"hosts"`
	// Certificate pack configuration type. Available values: `advanced`. **Modifying this attribute will force creation of a new resource.**.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/certificate_pack#type CertificatePack#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Which validation method to use in order to prove domain ownership.
	//
	// Available values: `txt`, `http`, `email`. **Modifying this attribute will force creation of a new resource.**
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/certificate_pack#validation_method CertificatePack#validation_method}
	ValidationMethod *string `field:"required" json:"validationMethod" yaml:"validationMethod"`
	// How long the certificate is valid for.
	//
	// Note: If using Let's Encrypt, this value can only be 90 days. Available values: `14`, `30`, `90`, `365`. **Modifying this attribute will force creation of a new resource.**
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/certificate_pack#validity_days CertificatePack#validity_days}
	ValidityDays *float64 `field:"required" json:"validityDays" yaml:"validityDays"`
	// The zone identifier to target for the resource. **Modifying this attribute will force creation of a new resource.**.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/certificate_pack#zone_id CertificatePack#zone_id}
	ZoneId *string `field:"required" json:"zoneId" yaml:"zoneId"`
	// Whether or not to include Cloudflare branding.
	//
	// This will add `sni.cloudflaressl.com` as the Common Name if set to `true`. **Modifying this attribute will force creation of a new resource.**
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/certificate_pack#cloudflare_branding CertificatePack#cloudflare_branding}
	CloudflareBranding interface{} `field:"optional" json:"cloudflareBranding" yaml:"cloudflareBranding"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/certificate_pack#id CertificatePack#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// validation_errors block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/certificate_pack#validation_errors CertificatePack#validation_errors}
	ValidationErrors interface{} `field:"optional" json:"validationErrors" yaml:"validationErrors"`
	// validation_records block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/certificate_pack#validation_records CertificatePack#validation_records}
	ValidationRecords interface{} `field:"optional" json:"validationRecords" yaml:"validationRecords"`
	// Whether or not to wait for a certificate pack to reach status `active` during creation.
	//
	// Defaults to `false`. **Modifying this attribute will force creation of a new resource.**
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/certificate_pack#wait_for_active_status CertificatePack#wait_for_active_status}
	WaitForActiveStatus interface{} `field:"optional" json:"waitForActiveStatus" yaml:"waitForActiveStatus"`
}

