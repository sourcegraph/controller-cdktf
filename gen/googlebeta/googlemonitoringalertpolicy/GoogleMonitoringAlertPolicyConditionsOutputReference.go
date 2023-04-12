package googlemonitoringalertpolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/googlebeta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/googlebeta/googlemonitoringalertpolicy/internal"
)

type GoogleMonitoringAlertPolicyConditionsOutputReference interface {
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
	ConditionAbsent() GoogleMonitoringAlertPolicyConditionsConditionAbsentOutputReference
	ConditionAbsentInput() *GoogleMonitoringAlertPolicyConditionsConditionAbsent
	ConditionMatchedLog() GoogleMonitoringAlertPolicyConditionsConditionMatchedLogOutputReference
	ConditionMatchedLogInput() *GoogleMonitoringAlertPolicyConditionsConditionMatchedLog
	ConditionMonitoringQueryLanguage() GoogleMonitoringAlertPolicyConditionsConditionMonitoringQueryLanguageOutputReference
	ConditionMonitoringQueryLanguageInput() *GoogleMonitoringAlertPolicyConditionsConditionMonitoringQueryLanguage
	ConditionThreshold() GoogleMonitoringAlertPolicyConditionsConditionThresholdOutputReference
	ConditionThresholdInput() *GoogleMonitoringAlertPolicyConditionsConditionThreshold
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DisplayName() *string
	SetDisplayName(val *string)
	DisplayNameInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Name() *string
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
	PutConditionAbsent(value *GoogleMonitoringAlertPolicyConditionsConditionAbsent)
	PutConditionMatchedLog(value *GoogleMonitoringAlertPolicyConditionsConditionMatchedLog)
	PutConditionMonitoringQueryLanguage(value *GoogleMonitoringAlertPolicyConditionsConditionMonitoringQueryLanguage)
	PutConditionThreshold(value *GoogleMonitoringAlertPolicyConditionsConditionThreshold)
	ResetConditionAbsent()
	ResetConditionMatchedLog()
	ResetConditionMonitoringQueryLanguage()
	ResetConditionThreshold()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleMonitoringAlertPolicyConditionsOutputReference
type jsiiProxy_GoogleMonitoringAlertPolicyConditionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleMonitoringAlertPolicyConditionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringAlertPolicyConditionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringAlertPolicyConditionsOutputReference) ConditionAbsent() GoogleMonitoringAlertPolicyConditionsConditionAbsentOutputReference {
	var returns GoogleMonitoringAlertPolicyConditionsConditionAbsentOutputReference
	_jsii_.Get(
		j,
		"conditionAbsent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringAlertPolicyConditionsOutputReference) ConditionAbsentInput() *GoogleMonitoringAlertPolicyConditionsConditionAbsent {
	var returns *GoogleMonitoringAlertPolicyConditionsConditionAbsent
	_jsii_.Get(
		j,
		"conditionAbsentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringAlertPolicyConditionsOutputReference) ConditionMatchedLog() GoogleMonitoringAlertPolicyConditionsConditionMatchedLogOutputReference {
	var returns GoogleMonitoringAlertPolicyConditionsConditionMatchedLogOutputReference
	_jsii_.Get(
		j,
		"conditionMatchedLog",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringAlertPolicyConditionsOutputReference) ConditionMatchedLogInput() *GoogleMonitoringAlertPolicyConditionsConditionMatchedLog {
	var returns *GoogleMonitoringAlertPolicyConditionsConditionMatchedLog
	_jsii_.Get(
		j,
		"conditionMatchedLogInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringAlertPolicyConditionsOutputReference) ConditionMonitoringQueryLanguage() GoogleMonitoringAlertPolicyConditionsConditionMonitoringQueryLanguageOutputReference {
	var returns GoogleMonitoringAlertPolicyConditionsConditionMonitoringQueryLanguageOutputReference
	_jsii_.Get(
		j,
		"conditionMonitoringQueryLanguage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringAlertPolicyConditionsOutputReference) ConditionMonitoringQueryLanguageInput() *GoogleMonitoringAlertPolicyConditionsConditionMonitoringQueryLanguage {
	var returns *GoogleMonitoringAlertPolicyConditionsConditionMonitoringQueryLanguage
	_jsii_.Get(
		j,
		"conditionMonitoringQueryLanguageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringAlertPolicyConditionsOutputReference) ConditionThreshold() GoogleMonitoringAlertPolicyConditionsConditionThresholdOutputReference {
	var returns GoogleMonitoringAlertPolicyConditionsConditionThresholdOutputReference
	_jsii_.Get(
		j,
		"conditionThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringAlertPolicyConditionsOutputReference) ConditionThresholdInput() *GoogleMonitoringAlertPolicyConditionsConditionThreshold {
	var returns *GoogleMonitoringAlertPolicyConditionsConditionThreshold
	_jsii_.Get(
		j,
		"conditionThresholdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringAlertPolicyConditionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringAlertPolicyConditionsOutputReference) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringAlertPolicyConditionsOutputReference) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringAlertPolicyConditionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringAlertPolicyConditionsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringAlertPolicyConditionsOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringAlertPolicyConditionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringAlertPolicyConditionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleMonitoringAlertPolicyConditionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GoogleMonitoringAlertPolicyConditionsOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleMonitoringAlertPolicyConditionsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleMonitoringAlertPolicyConditionsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-googlebeta.googleMonitoringAlertPolicy.GoogleMonitoringAlertPolicyConditionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGoogleMonitoringAlertPolicyConditionsOutputReference_Override(g GoogleMonitoringAlertPolicyConditionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-googlebeta.googleMonitoringAlertPolicy.GoogleMonitoringAlertPolicyConditionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GoogleMonitoringAlertPolicyConditionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleMonitoringAlertPolicyConditionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleMonitoringAlertPolicyConditionsOutputReference)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_GoogleMonitoringAlertPolicyConditionsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleMonitoringAlertPolicyConditionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleMonitoringAlertPolicyConditionsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleMonitoringAlertPolicyConditionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleMonitoringAlertPolicyConditionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleMonitoringAlertPolicyConditionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleMonitoringAlertPolicyConditionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleMonitoringAlertPolicyConditionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleMonitoringAlertPolicyConditionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleMonitoringAlertPolicyConditionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleMonitoringAlertPolicyConditionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleMonitoringAlertPolicyConditionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleMonitoringAlertPolicyConditionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleMonitoringAlertPolicyConditionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleMonitoringAlertPolicyConditionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleMonitoringAlertPolicyConditionsOutputReference) PutConditionAbsent(value *GoogleMonitoringAlertPolicyConditionsConditionAbsent) {
	if err := g.validatePutConditionAbsentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putConditionAbsent",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleMonitoringAlertPolicyConditionsOutputReference) PutConditionMatchedLog(value *GoogleMonitoringAlertPolicyConditionsConditionMatchedLog) {
	if err := g.validatePutConditionMatchedLogParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putConditionMatchedLog",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleMonitoringAlertPolicyConditionsOutputReference) PutConditionMonitoringQueryLanguage(value *GoogleMonitoringAlertPolicyConditionsConditionMonitoringQueryLanguage) {
	if err := g.validatePutConditionMonitoringQueryLanguageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putConditionMonitoringQueryLanguage",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleMonitoringAlertPolicyConditionsOutputReference) PutConditionThreshold(value *GoogleMonitoringAlertPolicyConditionsConditionThreshold) {
	if err := g.validatePutConditionThresholdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putConditionThreshold",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleMonitoringAlertPolicyConditionsOutputReference) ResetConditionAbsent() {
	_jsii_.InvokeVoid(
		g,
		"resetConditionAbsent",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleMonitoringAlertPolicyConditionsOutputReference) ResetConditionMatchedLog() {
	_jsii_.InvokeVoid(
		g,
		"resetConditionMatchedLog",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleMonitoringAlertPolicyConditionsOutputReference) ResetConditionMonitoringQueryLanguage() {
	_jsii_.InvokeVoid(
		g,
		"resetConditionMonitoringQueryLanguage",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleMonitoringAlertPolicyConditionsOutputReference) ResetConditionThreshold() {
	_jsii_.InvokeVoid(
		g,
		"resetConditionThreshold",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleMonitoringAlertPolicyConditionsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleMonitoringAlertPolicyConditionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

