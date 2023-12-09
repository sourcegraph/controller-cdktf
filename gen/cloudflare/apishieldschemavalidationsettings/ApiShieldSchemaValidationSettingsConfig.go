package apishieldschemavalidationsettings

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApiShieldSchemaValidationSettingsConfig struct {
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
	// The default mitigation action used when there is no mitigation action defined on the operation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.20.0/docs/resources/api_shield_schema_validation_settings#validation_default_mitigation_action ApiShieldSchemaValidationSettings#validation_default_mitigation_action}
	ValidationDefaultMitigationAction *string `field:"required" json:"validationDefaultMitigationAction" yaml:"validationDefaultMitigationAction"`
	// The zone identifier to target for the resource. **Modifying this attribute will force creation of a new resource.**.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.20.0/docs/resources/api_shield_schema_validation_settings#zone_id ApiShieldSchemaValidationSettings#zone_id}
	ZoneId *string `field:"required" json:"zoneId" yaml:"zoneId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.20.0/docs/resources/api_shield_schema_validation_settings#id ApiShieldSchemaValidationSettings#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// When set, this overrides both zone level and operation level mitigation actions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.20.0/docs/resources/api_shield_schema_validation_settings#validation_override_mitigation_action ApiShieldSchemaValidationSettings#validation_override_mitigation_action}
	ValidationOverrideMitigationAction *string `field:"optional" json:"validationOverrideMitigationAction" yaml:"validationOverrideMitigationAction"`
}

