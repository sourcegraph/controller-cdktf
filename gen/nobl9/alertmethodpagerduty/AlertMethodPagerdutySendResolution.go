package alertmethodpagerduty


type AlertMethodPagerdutySendResolution struct {
	// A message that will be attached to your 'all clear' notification.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/alert_method_pagerduty#message AlertMethodPagerduty#message}
	Message *string `field:"optional" json:"message" yaml:"message"`
}

