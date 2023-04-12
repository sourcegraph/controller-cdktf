package notificationpolicy


type NotificationPolicyFilters struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/notification_policy#enabled NotificationPolicy#enabled}.
	Enabled *[]*string `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/notification_policy#health_check_id NotificationPolicy#health_check_id}.
	HealthCheckId *[]*string `field:"optional" json:"healthCheckId" yaml:"healthCheckId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/notification_policy#limit NotificationPolicy#limit}.
	Limit *[]*string `field:"optional" json:"limit" yaml:"limit"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/notification_policy#pool_id NotificationPolicy#pool_id}.
	PoolId *[]*string `field:"optional" json:"poolId" yaml:"poolId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/notification_policy#product NotificationPolicy#product}.
	Product *[]*string `field:"optional" json:"product" yaml:"product"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/notification_policy#services NotificationPolicy#services}.
	Services *[]*string `field:"optional" json:"services" yaml:"services"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/notification_policy#slo NotificationPolicy#slo}.
	Slo *[]*string `field:"optional" json:"slo" yaml:"slo"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/notification_policy#status NotificationPolicy#status}.
	Status *[]*string `field:"optional" json:"status" yaml:"status"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/notification_policy#zones NotificationPolicy#zones}.
	Zones *[]*string `field:"optional" json:"zones" yaml:"zones"`
}

