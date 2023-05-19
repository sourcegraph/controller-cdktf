package teamsaccount

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type TeamsAccountConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/teams_account#account_id TeamsAccount#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// Whether to enable the activity log.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/teams_account#activity_log_enabled TeamsAccount#activity_log_enabled}
	ActivityLogEnabled interface{} `field:"optional" json:"activityLogEnabled" yaml:"activityLogEnabled"`
	// antivirus block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/teams_account#antivirus TeamsAccount#antivirus}
	Antivirus *TeamsAccountAntivirus `field:"optional" json:"antivirus" yaml:"antivirus"`
	// block_page block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/teams_account#block_page TeamsAccount#block_page}
	BlockPage *TeamsAccountBlockPage `field:"optional" json:"blockPage" yaml:"blockPage"`
	// fips block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/teams_account#fips TeamsAccount#fips}
	Fips *TeamsAccountFips `field:"optional" json:"fips" yaml:"fips"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/teams_account#id TeamsAccount#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// logging block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/teams_account#logging TeamsAccount#logging}
	Logging *TeamsAccountLogging `field:"optional" json:"logging" yaml:"logging"`
	// payload_log block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/teams_account#payload_log TeamsAccount#payload_log}
	PayloadLog *TeamsAccountPayloadLog `field:"optional" json:"payloadLog" yaml:"payloadLog"`
	// proxy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/teams_account#proxy TeamsAccount#proxy}
	Proxy *TeamsAccountProxy `field:"optional" json:"proxy" yaml:"proxy"`
	// Indicator that decryption of TLS traffic is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/teams_account#tls_decrypt_enabled TeamsAccount#tls_decrypt_enabled}
	TlsDecryptEnabled interface{} `field:"optional" json:"tlsDecryptEnabled" yaml:"tlsDecryptEnabled"`
	// Safely browse websites in Browser Isolation through a URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/teams_account#url_browser_isolation_enabled TeamsAccount#url_browser_isolation_enabled}
	UrlBrowserIsolationEnabled interface{} `field:"optional" json:"urlBrowserIsolationEnabled" yaml:"urlBrowserIsolationEnabled"`
}

