package googlemonitoringalertpolicy


type GoogleMonitoringAlertPolicyConditionsConditionThresholdForecastOptions struct {
	// The length of time into the future to forecast whether a timeseries will violate the threshold.
	//
	// If the predicted value is found to violate the
	// threshold, and the violation is observed in all
	// forecasts made for the Configured 'duration',
	// then the timeseries is considered to be failing.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_monitoring_alert_policy#forecast_horizon GoogleMonitoringAlertPolicy#forecast_horizon}
	ForecastHorizon *string `field:"required" json:"forecastHorizon" yaml:"forecastHorizon"`
}

