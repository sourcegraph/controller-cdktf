package googledataplexdatascan

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googledataplexdatascan/internal"
)

type GoogleDataplexDatascanDataQualitySpecPostScanActionsNotificationReportOutputReference interface {
	cdktf.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleDataplexDatascanDataQualitySpecPostScanActionsNotificationReport
	SetInternalValue(val *GoogleDataplexDatascanDataQualitySpecPostScanActionsNotificationReport)
	JobEndTrigger() GoogleDataplexDatascanDataQualitySpecPostScanActionsNotificationReportJobEndTriggerOutputReference
	JobEndTriggerInput() *GoogleDataplexDatascanDataQualitySpecPostScanActionsNotificationReportJobEndTrigger
	JobFailureTrigger() GoogleDataplexDatascanDataQualitySpecPostScanActionsNotificationReportJobFailureTriggerOutputReference
	JobFailureTriggerInput() *GoogleDataplexDatascanDataQualitySpecPostScanActionsNotificationReportJobFailureTrigger
	Recipients() GoogleDataplexDatascanDataQualitySpecPostScanActionsNotificationReportRecipientsOutputReference
	RecipientsInput() *GoogleDataplexDatascanDataQualitySpecPostScanActionsNotificationReportRecipients
	ScoreThresholdTrigger() GoogleDataplexDatascanDataQualitySpecPostScanActionsNotificationReportScoreThresholdTriggerOutputReference
	ScoreThresholdTriggerInput() *GoogleDataplexDatascanDataQualitySpecPostScanActionsNotificationReportScoreThresholdTrigger
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	PutJobEndTrigger(value *GoogleDataplexDatascanDataQualitySpecPostScanActionsNotificationReportJobEndTrigger)
	PutJobFailureTrigger(value *GoogleDataplexDatascanDataQualitySpecPostScanActionsNotificationReportJobFailureTrigger)
	PutRecipients(value *GoogleDataplexDatascanDataQualitySpecPostScanActionsNotificationReportRecipients)
	PutScoreThresholdTrigger(value *GoogleDataplexDatascanDataQualitySpecPostScanActionsNotificationReportScoreThresholdTrigger)
	ResetJobEndTrigger()
	ResetJobFailureTrigger()
	ResetScoreThresholdTrigger()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleDataplexDatascanDataQualitySpecPostScanActionsNotificationReportOutputReference
type jsiiProxy_GoogleDataplexDatascanDataQualitySpecPostScanActionsNotificationReportOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleDataplexDatascanDataQualitySpecPostScanActionsNotificationReportOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexDatascanDataQualitySpecPostScanActionsNotificationReportOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexDatascanDataQualitySpecPostScanActionsNotificationReportOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexDatascanDataQualitySpecPostScanActionsNotificationReportOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexDatascanDataQualitySpecPostScanActionsNotificationReportOutputReference) InternalValue() *GoogleDataplexDatascanDataQualitySpecPostScanActionsNotificationReport {
	var returns *GoogleDataplexDatascanDataQualitySpecPostScanActionsNotificationReport
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexDatascanDataQualitySpecPostScanActionsNotificationReportOutputReference) JobEndTrigger() GoogleDataplexDatascanDataQualitySpecPostScanActionsNotificationReportJobEndTriggerOutputReference {
	var returns GoogleDataplexDatascanDataQualitySpecPostScanActionsNotificationReportJobEndTriggerOutputReference
	_jsii_.Get(
		j,
		"jobEndTrigger",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexDatascanDataQualitySpecPostScanActionsNotificationReportOutputReference) JobEndTriggerInput() *GoogleDataplexDatascanDataQualitySpecPostScanActionsNotificationReportJobEndTrigger {
	var returns *GoogleDataplexDatascanDataQualitySpecPostScanActionsNotificationReportJobEndTrigger
	_jsii_.Get(
		j,
		"jobEndTriggerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexDatascanDataQualitySpecPostScanActionsNotificationReportOutputReference) JobFailureTrigger() GoogleDataplexDatascanDataQualitySpecPostScanActionsNotificationReportJobFailureTriggerOutputReference {
	var returns GoogleDataplexDatascanDataQualitySpecPostScanActionsNotificationReportJobFailureTriggerOutputReference
	_jsii_.Get(
		j,
		"jobFailureTrigger",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexDatascanDataQualitySpecPostScanActionsNotificationReportOutputReference) JobFailureTriggerInput() *GoogleDataplexDatascanDataQualitySpecPostScanActionsNotificationReportJobFailureTrigger {
	var returns *GoogleDataplexDatascanDataQualitySpecPostScanActionsNotificationReportJobFailureTrigger
	_jsii_.Get(
		j,
		"jobFailureTriggerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexDatascanDataQualitySpecPostScanActionsNotificationReportOutputReference) Recipients() GoogleDataplexDatascanDataQualitySpecPostScanActionsNotificationReportRecipientsOutputReference {
	var returns GoogleDataplexDatascanDataQualitySpecPostScanActionsNotificationReportRecipientsOutputReference
	_jsii_.Get(
		j,
		"recipients",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexDatascanDataQualitySpecPostScanActionsNotificationReportOutputReference) RecipientsInput() *GoogleDataplexDatascanDataQualitySpecPostScanActionsNotificationReportRecipients {
	var returns *GoogleDataplexDatascanDataQualitySpecPostScanActionsNotificationReportRecipients
	_jsii_.Get(
		j,
		"recipientsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexDatascanDataQualitySpecPostScanActionsNotificationReportOutputReference) ScoreThresholdTrigger() GoogleDataplexDatascanDataQualitySpecPostScanActionsNotificationReportScoreThresholdTriggerOutputReference {
	var returns GoogleDataplexDatascanDataQualitySpecPostScanActionsNotificationReportScoreThresholdTriggerOutputReference
	_jsii_.Get(
		j,
		"scoreThresholdTrigger",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexDatascanDataQualitySpecPostScanActionsNotificationReportOutputReference) ScoreThresholdTriggerInput() *GoogleDataplexDatascanDataQualitySpecPostScanActionsNotificationReportScoreThresholdTrigger {
	var returns *GoogleDataplexDatascanDataQualitySpecPostScanActionsNotificationReportScoreThresholdTrigger
	_jsii_.Get(
		j,
		"scoreThresholdTriggerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexDatascanDataQualitySpecPostScanActionsNotificationReportOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexDatascanDataQualitySpecPostScanActionsNotificationReportOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleDataplexDatascanDataQualitySpecPostScanActionsNotificationReportOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleDataplexDatascanDataQualitySpecPostScanActionsNotificationReportOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDataplexDatascanDataQualitySpecPostScanActionsNotificationReportOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDataplexDatascanDataQualitySpecPostScanActionsNotificationReportOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDataplexDatascan.GoogleDataplexDatascanDataQualitySpecPostScanActionsNotificationReportOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleDataplexDatascanDataQualitySpecPostScanActionsNotificationReportOutputReference_Override(g GoogleDataplexDatascanDataQualitySpecPostScanActionsNotificationReportOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDataplexDatascan.GoogleDataplexDatascanDataQualitySpecPostScanActionsNotificationReportOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleDataplexDatascanDataQualitySpecPostScanActionsNotificationReportOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDataplexDatascanDataQualitySpecPostScanActionsNotificationReportOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDataplexDatascanDataQualitySpecPostScanActionsNotificationReportOutputReference)SetInternalValue(val *GoogleDataplexDatascanDataQualitySpecPostScanActionsNotificationReport) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDataplexDatascanDataQualitySpecPostScanActionsNotificationReportOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDataplexDatascanDataQualitySpecPostScanActionsNotificationReportOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleDataplexDatascanDataQualitySpecPostScanActionsNotificationReportOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataplexDatascanDataQualitySpecPostScanActionsNotificationReportOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := g.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataplexDatascanDataQualitySpecPostScanActionsNotificationReportOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataplexDatascanDataQualitySpecPostScanActionsNotificationReportOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := g.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		g,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataplexDatascanDataQualitySpecPostScanActionsNotificationReportOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := g.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataplexDatascanDataQualitySpecPostScanActionsNotificationReportOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := g.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		g,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataplexDatascanDataQualitySpecPostScanActionsNotificationReportOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := g.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		g,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataplexDatascanDataQualitySpecPostScanActionsNotificationReportOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := g.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		g,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataplexDatascanDataQualitySpecPostScanActionsNotificationReportOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := g.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataplexDatascanDataQualitySpecPostScanActionsNotificationReportOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := g.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		g,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataplexDatascanDataQualitySpecPostScanActionsNotificationReportOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataplexDatascanDataQualitySpecPostScanActionsNotificationReportOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataplexDatascanDataQualitySpecPostScanActionsNotificationReportOutputReference) PutJobEndTrigger(value *GoogleDataplexDatascanDataQualitySpecPostScanActionsNotificationReportJobEndTrigger) {
	if err := g.validatePutJobEndTriggerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putJobEndTrigger",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataplexDatascanDataQualitySpecPostScanActionsNotificationReportOutputReference) PutJobFailureTrigger(value *GoogleDataplexDatascanDataQualitySpecPostScanActionsNotificationReportJobFailureTrigger) {
	if err := g.validatePutJobFailureTriggerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putJobFailureTrigger",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataplexDatascanDataQualitySpecPostScanActionsNotificationReportOutputReference) PutRecipients(value *GoogleDataplexDatascanDataQualitySpecPostScanActionsNotificationReportRecipients) {
	if err := g.validatePutRecipientsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putRecipients",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataplexDatascanDataQualitySpecPostScanActionsNotificationReportOutputReference) PutScoreThresholdTrigger(value *GoogleDataplexDatascanDataQualitySpecPostScanActionsNotificationReportScoreThresholdTrigger) {
	if err := g.validatePutScoreThresholdTriggerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putScoreThresholdTrigger",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataplexDatascanDataQualitySpecPostScanActionsNotificationReportOutputReference) ResetJobEndTrigger() {
	_jsii_.InvokeVoid(
		g,
		"resetJobEndTrigger",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataplexDatascanDataQualitySpecPostScanActionsNotificationReportOutputReference) ResetJobFailureTrigger() {
	_jsii_.InvokeVoid(
		g,
		"resetJobFailureTrigger",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataplexDatascanDataQualitySpecPostScanActionsNotificationReportOutputReference) ResetScoreThresholdTrigger() {
	_jsii_.InvokeVoid(
		g,
		"resetScoreThresholdTrigger",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataplexDatascanDataQualitySpecPostScanActionsNotificationReportOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := g.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataplexDatascanDataQualitySpecPostScanActionsNotificationReportOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

