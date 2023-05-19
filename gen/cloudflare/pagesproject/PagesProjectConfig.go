package pagesproject

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PagesProjectConfig struct {
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
	// The account identifier to target for the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/pages_project#account_id PagesProject#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// Name of the project. **Modifying this attribute will force creation of a new resource.**.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/pages_project#name PagesProject#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The name of the branch that is used for the production environment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/pages_project#production_branch PagesProject#production_branch}
	ProductionBranch *string `field:"required" json:"productionBranch" yaml:"productionBranch"`
	// build_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/pages_project#build_config PagesProject#build_config}
	BuildConfig *PagesProjectBuildConfig `field:"optional" json:"buildConfig" yaml:"buildConfig"`
	// deployment_configs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/pages_project#deployment_configs PagesProject#deployment_configs}
	DeploymentConfigs *PagesProjectDeploymentConfigs `field:"optional" json:"deploymentConfigs" yaml:"deploymentConfigs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/pages_project#id PagesProject#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/pages_project#source PagesProject#source}
	Source *PagesProjectSource `field:"optional" json:"source" yaml:"source"`
}

