package googlesecurityscannerscanconfig

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleSecurityScannerScanConfigConfig struct {
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
	// The user provider display name of the ScanConfig.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_security_scanner_scan_config#display_name GoogleSecurityScannerScanConfig#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// The starting URLs from which the scanner finds site pages.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_security_scanner_scan_config#starting_urls GoogleSecurityScannerScanConfig#starting_urls}
	StartingUrls *[]*string `field:"required" json:"startingUrls" yaml:"startingUrls"`
	// authentication block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_security_scanner_scan_config#authentication GoogleSecurityScannerScanConfig#authentication}
	Authentication *GoogleSecurityScannerScanConfigAuthentication `field:"optional" json:"authentication" yaml:"authentication"`
	// The blacklist URL patterns as described in https://cloud.google.com/security-scanner/docs/excluded-urls.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_security_scanner_scan_config#blacklist_patterns GoogleSecurityScannerScanConfig#blacklist_patterns}
	BlacklistPatterns *[]*string `field:"optional" json:"blacklistPatterns" yaml:"blacklistPatterns"`
	// Controls export of scan configurations and results to Cloud Security Command Center. Default value: "ENABLED" Possible values: ["ENABLED", "DISABLED"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_security_scanner_scan_config#export_to_security_command_center GoogleSecurityScannerScanConfig#export_to_security_command_center}
	ExportToSecurityCommandCenter *string `field:"optional" json:"exportToSecurityCommandCenter" yaml:"exportToSecurityCommandCenter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_security_scanner_scan_config#id GoogleSecurityScannerScanConfig#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The maximum QPS during scanning. A valid value ranges from 5 to 20 inclusively. Defaults to 15.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_security_scanner_scan_config#max_qps GoogleSecurityScannerScanConfig#max_qps}
	MaxQps *float64 `field:"optional" json:"maxQps" yaml:"maxQps"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_security_scanner_scan_config#project GoogleSecurityScannerScanConfig#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// schedule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_security_scanner_scan_config#schedule GoogleSecurityScannerScanConfig#schedule}
	Schedule *GoogleSecurityScannerScanConfigSchedule `field:"optional" json:"schedule" yaml:"schedule"`
	// Set of Cloud Platforms targeted by the scan.
	//
	// If empty, APP_ENGINE will be used as a default. Possible values: ["APP_ENGINE", "COMPUTE"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_security_scanner_scan_config#target_platforms GoogleSecurityScannerScanConfig#target_platforms}
	TargetPlatforms *[]*string `field:"optional" json:"targetPlatforms" yaml:"targetPlatforms"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_security_scanner_scan_config#timeouts GoogleSecurityScannerScanConfig#timeouts}
	Timeouts *GoogleSecurityScannerScanConfigTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Type of the user agents used for scanning Default value: "CHROME_LINUX" Possible values: ["USER_AGENT_UNSPECIFIED", "CHROME_LINUX", "CHROME_ANDROID", "SAFARI_IPHONE"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_security_scanner_scan_config#user_agent GoogleSecurityScannerScanConfig#user_agent}
	UserAgent *string `field:"optional" json:"userAgent" yaml:"userAgent"`
}

