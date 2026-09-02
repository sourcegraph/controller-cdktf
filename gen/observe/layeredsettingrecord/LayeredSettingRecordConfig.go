package layeredsettingrecord

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type LayeredSettingRecordConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/layered_setting_record#name LayeredSettingRecord#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/layered_setting_record#setting LayeredSettingRecord#setting}.
	Setting *string `field:"required" json:"setting" yaml:"setting"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/layered_setting_record#target LayeredSettingRecord#target}.
	Target *string `field:"required" json:"target" yaml:"target"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/layered_setting_record#workspace LayeredSettingRecord#workspace}.
	Workspace *string `field:"required" json:"workspace" yaml:"workspace"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/layered_setting_record#id LayeredSettingRecord#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/layered_setting_record#value_bool LayeredSettingRecord#value_bool}.
	ValueBool interface{} `field:"optional" json:"valueBool" yaml:"valueBool"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/layered_setting_record#value_duration LayeredSettingRecord#value_duration}.
	ValueDuration *string `field:"optional" json:"valueDuration" yaml:"valueDuration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/layered_setting_record#value_float64 LayeredSettingRecord#value_float64}.
	ValueFloat64 *float64 `field:"optional" json:"valueFloat64" yaml:"valueFloat64"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/layered_setting_record#value_int64 LayeredSettingRecord#value_int64}.
	ValueInt64 *float64 `field:"optional" json:"valueInt64" yaml:"valueInt64"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/layered_setting_record#value_string LayeredSettingRecord#value_string}.
	ValueString *string `field:"optional" json:"valueString" yaml:"valueString"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/layered_setting_record#value_timestamp LayeredSettingRecord#value_timestamp}.
	ValueTimestamp *string `field:"optional" json:"valueTimestamp" yaml:"valueTimestamp"`
}

