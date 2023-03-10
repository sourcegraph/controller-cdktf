package googlesecretmanagersecret

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleSecretManagerSecretConfig struct {
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
	// replication block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_secret_manager_secret#replication GoogleSecretManagerSecret#replication}
	Replication *GoogleSecretManagerSecretReplication `field:"required" json:"replication" yaml:"replication"`
	// This must be unique within the project.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_secret_manager_secret#secret_id GoogleSecretManagerSecret#secret_id}
	SecretId *string `field:"required" json:"secretId" yaml:"secretId"`
	// Timestamp in UTC when the Secret is scheduled to expire.
	//
	// This is always provided on output, regardless of what was sent on input.
	// A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_secret_manager_secret#expire_time GoogleSecretManagerSecret#expire_time}
	ExpireTime *string `field:"optional" json:"expireTime" yaml:"expireTime"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_secret_manager_secret#id GoogleSecretManagerSecret#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The labels assigned to this Secret.
	//
	// Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes,
	// and must conform to the following PCRE regular expression: [\p{Ll}\p{Lo}][\p{Ll}\p{Lo}\p{N}_-]{0,62}
	//
	// Label values must be between 0 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes,
	// and must conform to the following PCRE regular expression: [\p{Ll}\p{Lo}\p{N}_-]{0,63}
	//
	// No more than 64 labels can be assigned to a given resource.
	//
	// An object containing a list of "key": value pairs. Example:
	// { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_secret_manager_secret#labels GoogleSecretManagerSecret#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_secret_manager_secret#project GoogleSecretManagerSecret#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// rotation block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_secret_manager_secret#rotation GoogleSecretManagerSecret#rotation}
	Rotation *GoogleSecretManagerSecretRotation `field:"optional" json:"rotation" yaml:"rotation"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_secret_manager_secret#timeouts GoogleSecretManagerSecret#timeouts}
	Timeouts *GoogleSecretManagerSecretTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// topics block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_secret_manager_secret#topics GoogleSecretManagerSecret#topics}
	Topics interface{} `field:"optional" json:"topics" yaml:"topics"`
	// The TTL for the Secret. A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s".
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_secret_manager_secret#ttl GoogleSecretManagerSecret#ttl}
	Ttl *string `field:"optional" json:"ttl" yaml:"ttl"`
}

