package googledialogflowcxtestcase


type GoogleDialogflowCxTestCaseTestConfig struct {
	// Flow name to start the test case with.
	//
	// Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>/flows/<Flow ID>.
	// Only one of flow and page should be set to indicate the starting point of the test case. If neither is set, the test case will start with start page on the default start flow.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dialogflow_cx_test_case#flow GoogleDialogflowCxTestCase#flow}
	Flow *string `field:"optional" json:"flow" yaml:"flow"`
	// The page to start the test case with.
	//
	// Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>/flows/<Flow ID>/pages/<Page ID>.
	// Only one of flow and page should be set to indicate the starting point of the test case. If neither is set, the test case will start with start page on the default start flow.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dialogflow_cx_test_case#page GoogleDialogflowCxTestCase#page}
	Page *string `field:"optional" json:"page" yaml:"page"`
	// Session parameters to be compared when calculating differences.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dialogflow_cx_test_case#tracking_parameters GoogleDialogflowCxTestCase#tracking_parameters}
	TrackingParameters *[]*string `field:"optional" json:"trackingParameters" yaml:"trackingParameters"`
}

