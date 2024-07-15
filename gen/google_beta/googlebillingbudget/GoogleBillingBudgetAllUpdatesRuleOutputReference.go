package googlebillingbudget

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlebillingbudget/internal"
)

type GoogleBillingBudgetAllUpdatesRuleOutputReference interface {
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
	DisableDefaultIamRecipients() interface{}
	SetDisableDefaultIamRecipients(val interface{})
	DisableDefaultIamRecipientsInput() interface{}
	EnableProjectLevelRecipients() interface{}
	SetEnableProjectLevelRecipients(val interface{})
	EnableProjectLevelRecipientsInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleBillingBudgetAllUpdatesRule
	SetInternalValue(val *GoogleBillingBudgetAllUpdatesRule)
	MonitoringNotificationChannels() *[]*string
	SetMonitoringNotificationChannels(val *[]*string)
	MonitoringNotificationChannelsInput() *[]*string
	PubsubTopic() *string
	SetPubsubTopic(val *string)
	PubsubTopicInput() *string
	SchemaVersion() *string
	SetSchemaVersion(val *string)
	SchemaVersionInput() *string
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
	ResetDisableDefaultIamRecipients()
	ResetEnableProjectLevelRecipients()
	ResetMonitoringNotificationChannels()
	ResetPubsubTopic()
	ResetSchemaVersion()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleBillingBudgetAllUpdatesRuleOutputReference
type jsiiProxy_GoogleBillingBudgetAllUpdatesRuleOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleBillingBudgetAllUpdatesRuleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBillingBudgetAllUpdatesRuleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBillingBudgetAllUpdatesRuleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBillingBudgetAllUpdatesRuleOutputReference) DisableDefaultIamRecipients() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableDefaultIamRecipients",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBillingBudgetAllUpdatesRuleOutputReference) DisableDefaultIamRecipientsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableDefaultIamRecipientsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBillingBudgetAllUpdatesRuleOutputReference) EnableProjectLevelRecipients() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableProjectLevelRecipients",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBillingBudgetAllUpdatesRuleOutputReference) EnableProjectLevelRecipientsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableProjectLevelRecipientsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBillingBudgetAllUpdatesRuleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBillingBudgetAllUpdatesRuleOutputReference) InternalValue() *GoogleBillingBudgetAllUpdatesRule {
	var returns *GoogleBillingBudgetAllUpdatesRule
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBillingBudgetAllUpdatesRuleOutputReference) MonitoringNotificationChannels() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"monitoringNotificationChannels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBillingBudgetAllUpdatesRuleOutputReference) MonitoringNotificationChannelsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"monitoringNotificationChannelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBillingBudgetAllUpdatesRuleOutputReference) PubsubTopic() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pubsubTopic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBillingBudgetAllUpdatesRuleOutputReference) PubsubTopicInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pubsubTopicInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBillingBudgetAllUpdatesRuleOutputReference) SchemaVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schemaVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBillingBudgetAllUpdatesRuleOutputReference) SchemaVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schemaVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBillingBudgetAllUpdatesRuleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBillingBudgetAllUpdatesRuleOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleBillingBudgetAllUpdatesRuleOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleBillingBudgetAllUpdatesRuleOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleBillingBudgetAllUpdatesRuleOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleBillingBudgetAllUpdatesRuleOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleBillingBudget.GoogleBillingBudgetAllUpdatesRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleBillingBudgetAllUpdatesRuleOutputReference_Override(g GoogleBillingBudgetAllUpdatesRuleOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleBillingBudget.GoogleBillingBudgetAllUpdatesRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleBillingBudgetAllUpdatesRuleOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleBillingBudgetAllUpdatesRuleOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleBillingBudgetAllUpdatesRuleOutputReference)SetDisableDefaultIamRecipients(val interface{}) {
	if err := j.validateSetDisableDefaultIamRecipientsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableDefaultIamRecipients",
		val,
	)
}

func (j *jsiiProxy_GoogleBillingBudgetAllUpdatesRuleOutputReference)SetEnableProjectLevelRecipients(val interface{}) {
	if err := j.validateSetEnableProjectLevelRecipientsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableProjectLevelRecipients",
		val,
	)
}

func (j *jsiiProxy_GoogleBillingBudgetAllUpdatesRuleOutputReference)SetInternalValue(val *GoogleBillingBudgetAllUpdatesRule) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleBillingBudgetAllUpdatesRuleOutputReference)SetMonitoringNotificationChannels(val *[]*string) {
	if err := j.validateSetMonitoringNotificationChannelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"monitoringNotificationChannels",
		val,
	)
}

func (j *jsiiProxy_GoogleBillingBudgetAllUpdatesRuleOutputReference)SetPubsubTopic(val *string) {
	if err := j.validateSetPubsubTopicParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pubsubTopic",
		val,
	)
}

func (j *jsiiProxy_GoogleBillingBudgetAllUpdatesRuleOutputReference)SetSchemaVersion(val *string) {
	if err := j.validateSetSchemaVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"schemaVersion",
		val,
	)
}

func (j *jsiiProxy_GoogleBillingBudgetAllUpdatesRuleOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleBillingBudgetAllUpdatesRuleOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleBillingBudgetAllUpdatesRuleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBillingBudgetAllUpdatesRuleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleBillingBudgetAllUpdatesRuleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleBillingBudgetAllUpdatesRuleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleBillingBudgetAllUpdatesRuleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleBillingBudgetAllUpdatesRuleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleBillingBudgetAllUpdatesRuleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleBillingBudgetAllUpdatesRuleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleBillingBudgetAllUpdatesRuleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleBillingBudgetAllUpdatesRuleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleBillingBudgetAllUpdatesRuleOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBillingBudgetAllUpdatesRuleOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleBillingBudgetAllUpdatesRuleOutputReference) ResetDisableDefaultIamRecipients() {
	_jsii_.InvokeVoid(
		g,
		"resetDisableDefaultIamRecipients",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBillingBudgetAllUpdatesRuleOutputReference) ResetEnableProjectLevelRecipients() {
	_jsii_.InvokeVoid(
		g,
		"resetEnableProjectLevelRecipients",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBillingBudgetAllUpdatesRuleOutputReference) ResetMonitoringNotificationChannels() {
	_jsii_.InvokeVoid(
		g,
		"resetMonitoringNotificationChannels",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBillingBudgetAllUpdatesRuleOutputReference) ResetPubsubTopic() {
	_jsii_.InvokeVoid(
		g,
		"resetPubsubTopic",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBillingBudgetAllUpdatesRuleOutputReference) ResetSchemaVersion() {
	_jsii_.InvokeVoid(
		g,
		"resetSchemaVersion",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBillingBudgetAllUpdatesRuleOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleBillingBudgetAllUpdatesRuleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

