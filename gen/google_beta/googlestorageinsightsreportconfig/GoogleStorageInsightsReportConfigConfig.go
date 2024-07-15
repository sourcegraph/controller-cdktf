package googlestorageinsightsreportconfig

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleStorageInsightsReportConfigConfig struct {
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
	// csv_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_storage_insights_report_config#csv_options GoogleStorageInsightsReportConfig#csv_options}
	CsvOptions *GoogleStorageInsightsReportConfigCsvOptions `field:"required" json:"csvOptions" yaml:"csvOptions"`
	// The location of the ReportConfig. The source and destination buckets specified in the ReportConfig must be in the same location.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_storage_insights_report_config#location GoogleStorageInsightsReportConfig#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The editable display name of the inventory report configuration. Has a limit of 256 characters. Can be empty.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_storage_insights_report_config#display_name GoogleStorageInsightsReportConfig#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// frequency_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_storage_insights_report_config#frequency_options GoogleStorageInsightsReportConfig#frequency_options}
	FrequencyOptions *GoogleStorageInsightsReportConfigFrequencyOptions `field:"optional" json:"frequencyOptions" yaml:"frequencyOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_storage_insights_report_config#id GoogleStorageInsightsReportConfig#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// object_metadata_report_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_storage_insights_report_config#object_metadata_report_options GoogleStorageInsightsReportConfig#object_metadata_report_options}
	ObjectMetadataReportOptions *GoogleStorageInsightsReportConfigObjectMetadataReportOptions `field:"optional" json:"objectMetadataReportOptions" yaml:"objectMetadataReportOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_storage_insights_report_config#project GoogleStorageInsightsReportConfig#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_storage_insights_report_config#timeouts GoogleStorageInsightsReportConfig#timeouts}
	Timeouts *GoogleStorageInsightsReportConfigTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

