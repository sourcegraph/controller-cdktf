package certificatepack

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CertificatePackConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/certificate_pack#hosts CertificatePack#hosts}.
	Hosts *[]*string `field:"required" json:"hosts" yaml:"hosts"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/certificate_pack#type CertificatePack#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/certificate_pack#zone_id CertificatePack#zone_id}.
	ZoneId *string `field:"required" json:"zoneId" yaml:"zoneId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/certificate_pack#certificate_authority CertificatePack#certificate_authority}.
	CertificateAuthority *string `field:"optional" json:"certificateAuthority" yaml:"certificateAuthority"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/certificate_pack#cloudflare_branding CertificatePack#cloudflare_branding}.
	CloudflareBranding interface{} `field:"optional" json:"cloudflareBranding" yaml:"cloudflareBranding"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/certificate_pack#id CertificatePack#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// validation_errors block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/certificate_pack#validation_errors CertificatePack#validation_errors}
	ValidationErrors interface{} `field:"optional" json:"validationErrors" yaml:"validationErrors"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/certificate_pack#validation_method CertificatePack#validation_method}.
	ValidationMethod *string `field:"optional" json:"validationMethod" yaml:"validationMethod"`
	// validation_records block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/certificate_pack#validation_records CertificatePack#validation_records}
	ValidationRecords interface{} `field:"optional" json:"validationRecords" yaml:"validationRecords"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/certificate_pack#validity_days CertificatePack#validity_days}.
	ValidityDays *float64 `field:"optional" json:"validityDays" yaml:"validityDays"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/certificate_pack#wait_for_active_status CertificatePack#wait_for_active_status}.
	WaitForActiveStatus interface{} `field:"optional" json:"waitForActiveStatus" yaml:"waitForActiveStatus"`
}

