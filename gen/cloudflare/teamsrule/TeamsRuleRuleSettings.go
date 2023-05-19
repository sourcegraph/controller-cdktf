package teamsrule


type TeamsRuleRuleSettings struct {
	// Add custom headers to allowed requests in the form of key-value pairs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/teams_rule#add_headers TeamsRule#add_headers}
	AddHeaders *map[string]*string `field:"optional" json:"addHeaders" yaml:"addHeaders"`
	// Allow parent MSP accounts to enable bypass their children's rules.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/teams_rule#allow_child_bypass TeamsRule#allow_child_bypass}
	AllowChildBypass interface{} `field:"optional" json:"allowChildBypass" yaml:"allowChildBypass"`
	// audit_ssh block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/teams_rule#audit_ssh TeamsRule#audit_ssh}
	AuditSsh *TeamsRuleRuleSettingsAuditSsh `field:"optional" json:"auditSsh" yaml:"auditSsh"`
	// biso_admin_controls block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/teams_rule#biso_admin_controls TeamsRule#biso_admin_controls}
	BisoAdminControls *TeamsRuleRuleSettingsBisoAdminControls `field:"optional" json:"bisoAdminControls" yaml:"bisoAdminControls"`
	// Indicator of block page enablement.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/teams_rule#block_page_enabled TeamsRule#block_page_enabled}
	BlockPageEnabled interface{} `field:"optional" json:"blockPageEnabled" yaml:"blockPageEnabled"`
	// The displayed reason for a user being blocked.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/teams_rule#block_page_reason TeamsRule#block_page_reason}
	BlockPageReason *string `field:"optional" json:"blockPageReason" yaml:"blockPageReason"`
	// Allow child MSP accounts to bypass their parent's rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/teams_rule#bypass_parent_rule TeamsRule#bypass_parent_rule}
	BypassParentRule interface{} `field:"optional" json:"bypassParentRule" yaml:"bypassParentRule"`
	// check_session block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/teams_rule#check_session TeamsRule#check_session}
	CheckSession *TeamsRuleRuleSettingsCheckSession `field:"optional" json:"checkSession" yaml:"checkSession"`
	// egress block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/teams_rule#egress TeamsRule#egress}
	Egress *TeamsRuleRuleSettingsEgress `field:"optional" json:"egress" yaml:"egress"`
	// Disable DNSSEC validation (must be Allow rule).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/teams_rule#insecure_disable_dnssec_validation TeamsRule#insecure_disable_dnssec_validation}
	InsecureDisableDnssecValidation interface{} `field:"optional" json:"insecureDisableDnssecValidation" yaml:"insecureDisableDnssecValidation"`
	// Turns on IP category based filter on dns if the rule contains dns category checks.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/teams_rule#ip_categories TeamsRule#ip_categories}
	IpCategories interface{} `field:"optional" json:"ipCategories" yaml:"ipCategories"`
	// l4override block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/teams_rule#l4override TeamsRule#l4override}
	L4Override *TeamsRuleRuleSettingsL4Override `field:"optional" json:"l4Override" yaml:"l4Override"`
	// The host to override matching DNS queries with.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/teams_rule#override_host TeamsRule#override_host}
	OverrideHost *string `field:"optional" json:"overrideHost" yaml:"overrideHost"`
	// The IPs to override matching DNS queries with.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/teams_rule#override_ips TeamsRule#override_ips}
	OverrideIps *[]*string `field:"optional" json:"overrideIps" yaml:"overrideIps"`
	// payload_log block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/teams_rule#payload_log TeamsRule#payload_log}
	PayloadLog *TeamsRuleRuleSettingsPayloadLog `field:"optional" json:"payloadLog" yaml:"payloadLog"`
	// untrusted_cert block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/teams_rule#untrusted_cert TeamsRule#untrusted_cert}
	UntrustedCert *TeamsRuleRuleSettingsUntrustedCert `field:"optional" json:"untrustedCert" yaml:"untrustedCert"`
}

