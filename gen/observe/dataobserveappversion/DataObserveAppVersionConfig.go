package dataobserveappversion

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataObserveAppVersionConfig struct {
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
	// The app module name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/data-sources/app_version#module_id DataObserveAppVersion#module_id}
	ModuleId *string `field:"required" json:"moduleId" yaml:"moduleId"`
	// The version constraint rules which are used to find the newest acceptable version.
	//
	// Multiple constraints should be comma separated.
	//
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/data-sources/app_version#version_constraint DataObserveAppVersion#version_constraint}
	VersionConstraint *string `field:"required" json:"versionConstraint" yaml:"versionConstraint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/data-sources/app_version#id DataObserveAppVersion#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Whether to include prerelease versions in the version search. Defaults to false (don't include).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/data-sources/app_version#include_prerelease DataObserveAppVersion#include_prerelease}
	IncludePrerelease interface{} `field:"optional" json:"includePrerelease" yaml:"includePrerelease"`
}

