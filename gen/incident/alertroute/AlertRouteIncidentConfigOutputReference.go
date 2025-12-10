package alertroute

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/incident/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/incident/alertroute/internal"
)

type AlertRouteIncidentConfigOutputReference interface {
	cdktf.ComplexObject
	AutoDeclineEnabled() interface{}
	SetAutoDeclineEnabled(val interface{})
	AutoDeclineEnabledInput() interface{}
	AutoRelateGroupedAlerts() interface{}
	SetAutoRelateGroupedAlerts(val interface{})
	AutoRelateGroupedAlertsInput() interface{}
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
	ConditionGroups() AlertRouteIncidentConfigConditionGroupsList
	ConditionGroupsInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DeferTimeSeconds() *float64
	SetDeferTimeSeconds(val *float64)
	DeferTimeSecondsInput() *float64
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	// Experimental.
	Fqn() *string
	GroupingKeys() AlertRouteIncidentConfigGroupingKeysList
	GroupingKeysInput() interface{}
	GroupingWindowSeconds() *float64
	SetGroupingWindowSeconds(val *float64)
	GroupingWindowSecondsInput() *float64
	InternalValue() interface{}
	SetInternalValue(val interface{})
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
	PutConditionGroups(value interface{})
	PutGroupingKeys(value interface{})
	ResetAutoRelateGroupedAlerts()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AlertRouteIncidentConfigOutputReference
type jsiiProxy_AlertRouteIncidentConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AlertRouteIncidentConfigOutputReference) AutoDeclineEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoDeclineEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertRouteIncidentConfigOutputReference) AutoDeclineEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoDeclineEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertRouteIncidentConfigOutputReference) AutoRelateGroupedAlerts() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoRelateGroupedAlerts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertRouteIncidentConfigOutputReference) AutoRelateGroupedAlertsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoRelateGroupedAlertsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertRouteIncidentConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertRouteIncidentConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertRouteIncidentConfigOutputReference) ConditionGroups() AlertRouteIncidentConfigConditionGroupsList {
	var returns AlertRouteIncidentConfigConditionGroupsList
	_jsii_.Get(
		j,
		"conditionGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertRouteIncidentConfigOutputReference) ConditionGroupsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"conditionGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertRouteIncidentConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertRouteIncidentConfigOutputReference) DeferTimeSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"deferTimeSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertRouteIncidentConfigOutputReference) DeferTimeSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"deferTimeSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertRouteIncidentConfigOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertRouteIncidentConfigOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertRouteIncidentConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertRouteIncidentConfigOutputReference) GroupingKeys() AlertRouteIncidentConfigGroupingKeysList {
	var returns AlertRouteIncidentConfigGroupingKeysList
	_jsii_.Get(
		j,
		"groupingKeys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertRouteIncidentConfigOutputReference) GroupingKeysInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"groupingKeysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertRouteIncidentConfigOutputReference) GroupingWindowSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"groupingWindowSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertRouteIncidentConfigOutputReference) GroupingWindowSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"groupingWindowSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertRouteIncidentConfigOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertRouteIncidentConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertRouteIncidentConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAlertRouteIncidentConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AlertRouteIncidentConfigOutputReference {
	_init_.Initialize()

	if err := validateNewAlertRouteIncidentConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AlertRouteIncidentConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-incident.alertRoute.AlertRouteIncidentConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAlertRouteIncidentConfigOutputReference_Override(a AlertRouteIncidentConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-incident.alertRoute.AlertRouteIncidentConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AlertRouteIncidentConfigOutputReference)SetAutoDeclineEnabled(val interface{}) {
	if err := j.validateSetAutoDeclineEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoDeclineEnabled",
		val,
	)
}

func (j *jsiiProxy_AlertRouteIncidentConfigOutputReference)SetAutoRelateGroupedAlerts(val interface{}) {
	if err := j.validateSetAutoRelateGroupedAlertsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoRelateGroupedAlerts",
		val,
	)
}

func (j *jsiiProxy_AlertRouteIncidentConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AlertRouteIncidentConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AlertRouteIncidentConfigOutputReference)SetDeferTimeSeconds(val *float64) {
	if err := j.validateSetDeferTimeSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deferTimeSeconds",
		val,
	)
}

func (j *jsiiProxy_AlertRouteIncidentConfigOutputReference)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_AlertRouteIncidentConfigOutputReference)SetGroupingWindowSeconds(val *float64) {
	if err := j.validateSetGroupingWindowSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"groupingWindowSeconds",
		val,
	)
}

func (j *jsiiProxy_AlertRouteIncidentConfigOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AlertRouteIncidentConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AlertRouteIncidentConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AlertRouteIncidentConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertRouteIncidentConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertRouteIncidentConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertRouteIncidentConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertRouteIncidentConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertRouteIncidentConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertRouteIncidentConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertRouteIncidentConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertRouteIncidentConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertRouteIncidentConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertRouteIncidentConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertRouteIncidentConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertRouteIncidentConfigOutputReference) PutConditionGroups(value interface{}) {
	if err := a.validatePutConditionGroupsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putConditionGroups",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AlertRouteIncidentConfigOutputReference) PutGroupingKeys(value interface{}) {
	if err := a.validatePutGroupingKeysParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putGroupingKeys",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AlertRouteIncidentConfigOutputReference) ResetAutoRelateGroupedAlerts() {
	_jsii_.InvokeVoid(
		a,
		"resetAutoRelateGroupedAlerts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertRouteIncidentConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := a.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertRouteIncidentConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

