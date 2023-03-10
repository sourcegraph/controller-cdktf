package googleassuredworkloadsworkload

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleAssuredWorkloadsWorkloadConfig struct {
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
	// Required.
	//
	// Input only. The billing account used for the resources which are direct children of workload. This billing account is initially associated with the resources created as part of Workload creation. After the initial creation of these resources, the customer can change the assigned billing account. The resource name has the form `billingAccounts/{billing_account_id}`. For example, 'billingAccounts/012345-567890-ABCDEF`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_assured_workloads_workload#billing_account GoogleAssuredWorkloadsWorkload#billing_account}
	BillingAccount *string `field:"required" json:"billingAccount" yaml:"billingAccount"`
	// Required. Immutable. Compliance Regime associated with this workload. Possible values: COMPLIANCE_REGIME_UNSPECIFIED, IL4, CJIS, FEDRAMP_HIGH, FEDRAMP_MODERATE, US_REGIONAL_ACCESS.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_assured_workloads_workload#compliance_regime GoogleAssuredWorkloadsWorkload#compliance_regime}
	ComplianceRegime *string `field:"required" json:"complianceRegime" yaml:"complianceRegime"`
	// Required.
	//
	// The user-assigned display name of the Workload. When present it must be between 4 to 30 characters. Allowed characters are: lowercase and uppercase letters, numbers, hyphen, and spaces. Example: My Workload
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_assured_workloads_workload#display_name GoogleAssuredWorkloadsWorkload#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// The location for the resource.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_assured_workloads_workload#location GoogleAssuredWorkloadsWorkload#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The organization for the resource.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_assured_workloads_workload#organization GoogleAssuredWorkloadsWorkload#organization}
	Organization *string `field:"required" json:"organization" yaml:"organization"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_assured_workloads_workload#id GoogleAssuredWorkloadsWorkload#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// kms_settings block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_assured_workloads_workload#kms_settings GoogleAssuredWorkloadsWorkload#kms_settings}
	KmsSettings *GoogleAssuredWorkloadsWorkloadKmsSettings `field:"optional" json:"kmsSettings" yaml:"kmsSettings"`
	// Optional. Labels applied to the workload.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_assured_workloads_workload#labels GoogleAssuredWorkloadsWorkload#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Input only.
	//
	// The parent resource for the resources managed by this Assured Workload. May be either an organization or a folder. Must be the same or a child of the Workload parent. If not specified all resources are created under the Workload parent. Formats: folders/{folder_id}, organizations/{organization_id}
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_assured_workloads_workload#provisioned_resources_parent GoogleAssuredWorkloadsWorkload#provisioned_resources_parent}
	ProvisionedResourcesParent *string `field:"optional" json:"provisionedResourcesParent" yaml:"provisionedResourcesParent"`
	// resource_settings block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_assured_workloads_workload#resource_settings GoogleAssuredWorkloadsWorkload#resource_settings}
	ResourceSettings interface{} `field:"optional" json:"resourceSettings" yaml:"resourceSettings"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_assured_workloads_workload#timeouts GoogleAssuredWorkloadsWorkload#timeouts}
	Timeouts *GoogleAssuredWorkloadsWorkloadTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

