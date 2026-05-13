package googlecesevaluation

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleCesEvaluationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_evaluation#app GoogleCesEvaluation#app}.
	App *string `field:"required" json:"app" yaml:"app"`
	// User-defined display name of the evaluation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_evaluation#display_name GoogleCesEvaluation#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// The ID to use for the evaluation, which will become the final component of the evaluation's resource name.
	//
	// If not provided, a unique ID will be
	// automatically assigned for the evaluation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_evaluation#evaluation_id GoogleCesEvaluation#evaluation_id}
	EvaluationId *string `field:"required" json:"evaluationId" yaml:"evaluationId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_evaluation#location GoogleCesEvaluation#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// User-defined description of the evaluation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_evaluation#description GoogleCesEvaluation#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// golden block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_evaluation#golden GoogleCesEvaluation#golden}
	Golden *GoogleCesEvaluationGolden `field:"optional" json:"golden" yaml:"golden"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_evaluation#id GoogleCesEvaluation#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_evaluation#project GoogleCesEvaluation#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// User defined tags to categorize the evaluation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_evaluation#tags GoogleCesEvaluation#tags}
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_evaluation#timeouts GoogleCesEvaluation#timeouts}
	Timeouts *GoogleCesEvaluationTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

