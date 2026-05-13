package googledatalosspreventiondiscoveryconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googledatalosspreventiondiscoveryconfig/internal"
)

type GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetOutputReference interface {
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
	Conditions() GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetConditionsOutputReference
	ConditionsInput() *GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetConditions
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DataSourceType() GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetDataSourceTypeOutputReference
	DataSourceTypeInput() *GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetDataSourceType
	Disabled() GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetDisabledOutputReference
	DisabledInput() *GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetDisabled
	Filter() GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterOutputReference
	FilterInput() *GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilter
	// Experimental.
	Fqn() *string
	GenerationCadence() GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetGenerationCadenceOutputReference
	GenerationCadenceInput() *GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetGenerationCadence
	InternalValue() *GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTarget
	SetInternalValue(val *GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTarget)
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
	PutConditions(value *GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetConditions)
	PutDataSourceType(value *GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetDataSourceType)
	PutDisabled(value *GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetDisabled)
	PutFilter(value *GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilter)
	PutGenerationCadence(value *GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetGenerationCadence)
	ResetConditions()
	ResetDataSourceType()
	ResetDisabled()
	ResetGenerationCadence()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetOutputReference
type jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetOutputReference) Conditions() GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetConditionsOutputReference {
	var returns GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetConditionsOutputReference
	_jsii_.Get(
		j,
		"conditions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetOutputReference) ConditionsInput() *GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetConditions {
	var returns *GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetConditions
	_jsii_.Get(
		j,
		"conditionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetOutputReference) DataSourceType() GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetDataSourceTypeOutputReference {
	var returns GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetDataSourceTypeOutputReference
	_jsii_.Get(
		j,
		"dataSourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetOutputReference) DataSourceTypeInput() *GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetDataSourceType {
	var returns *GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetDataSourceType
	_jsii_.Get(
		j,
		"dataSourceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetOutputReference) Disabled() GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetDisabledOutputReference {
	var returns GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetDisabledOutputReference
	_jsii_.Get(
		j,
		"disabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetOutputReference) DisabledInput() *GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetDisabled {
	var returns *GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetDisabled
	_jsii_.Get(
		j,
		"disabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetOutputReference) Filter() GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterOutputReference {
	var returns GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterOutputReference
	_jsii_.Get(
		j,
		"filter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetOutputReference) FilterInput() *GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilter {
	var returns *GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilter
	_jsii_.Get(
		j,
		"filterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetOutputReference) GenerationCadence() GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetGenerationCadenceOutputReference {
	var returns GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetGenerationCadenceOutputReference
	_jsii_.Get(
		j,
		"generationCadence",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetOutputReference) GenerationCadenceInput() *GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetGenerationCadence {
	var returns *GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetGenerationCadence
	_jsii_.Get(
		j,
		"generationCadenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetOutputReference) InternalValue() *GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTarget {
	var returns *GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTarget
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDataLossPreventionDiscoveryConfig.GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetOutputReference_Override(g GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDataLossPreventionDiscoveryConfig.GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetOutputReference)SetInternalValue(val *GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTarget) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetOutputReference) PutConditions(value *GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetConditions) {
	if err := g.validatePutConditionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putConditions",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetOutputReference) PutDataSourceType(value *GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetDataSourceType) {
	if err := g.validatePutDataSourceTypeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDataSourceType",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetOutputReference) PutDisabled(value *GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetDisabled) {
	if err := g.validatePutDisabledParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDisabled",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetOutputReference) PutFilter(value *GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilter) {
	if err := g.validatePutFilterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putFilter",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetOutputReference) PutGenerationCadence(value *GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetGenerationCadence) {
	if err := g.validatePutGenerationCadenceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putGenerationCadence",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetOutputReference) ResetConditions() {
	_jsii_.InvokeVoid(
		g,
		"resetConditions",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetOutputReference) ResetDataSourceType() {
	_jsii_.InvokeVoid(
		g,
		"resetDataSourceType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetOutputReference) ResetDisabled() {
	_jsii_.InvokeVoid(
		g,
		"resetDisabled",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetOutputReference) ResetGenerationCadence() {
	_jsii_.InvokeVoid(
		g,
		"resetGenerationCadence",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

