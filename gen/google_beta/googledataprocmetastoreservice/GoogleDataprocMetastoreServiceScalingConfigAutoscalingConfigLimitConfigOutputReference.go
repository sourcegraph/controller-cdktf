package googledataprocmetastoreservice

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googledataprocmetastoreservice/internal"
)

type GoogleDataprocMetastoreServiceScalingConfigAutoscalingConfigLimitConfigOutputReference interface {
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
	InternalValue() *GoogleDataprocMetastoreServiceScalingConfigAutoscalingConfigLimitConfig
	SetInternalValue(val *GoogleDataprocMetastoreServiceScalingConfigAutoscalingConfigLimitConfig)
	MaxScalingFactor() *float64
	SetMaxScalingFactor(val *float64)
	MaxScalingFactorInput() *float64
	MinScalingFactor() *float64
	SetMinScalingFactor(val *float64)
	MinScalingFactorInput() *float64
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
	ResetMaxScalingFactor()
	ResetMinScalingFactor()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleDataprocMetastoreServiceScalingConfigAutoscalingConfigLimitConfigOutputReference
type jsiiProxy_GoogleDataprocMetastoreServiceScalingConfigAutoscalingConfigLimitConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleDataprocMetastoreServiceScalingConfigAutoscalingConfigLimitConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocMetastoreServiceScalingConfigAutoscalingConfigLimitConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocMetastoreServiceScalingConfigAutoscalingConfigLimitConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocMetastoreServiceScalingConfigAutoscalingConfigLimitConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocMetastoreServiceScalingConfigAutoscalingConfigLimitConfigOutputReference) InternalValue() *GoogleDataprocMetastoreServiceScalingConfigAutoscalingConfigLimitConfig {
	var returns *GoogleDataprocMetastoreServiceScalingConfigAutoscalingConfigLimitConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocMetastoreServiceScalingConfigAutoscalingConfigLimitConfigOutputReference) MaxScalingFactor() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxScalingFactor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocMetastoreServiceScalingConfigAutoscalingConfigLimitConfigOutputReference) MaxScalingFactorInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxScalingFactorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocMetastoreServiceScalingConfigAutoscalingConfigLimitConfigOutputReference) MinScalingFactor() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minScalingFactor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocMetastoreServiceScalingConfigAutoscalingConfigLimitConfigOutputReference) MinScalingFactorInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minScalingFactorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocMetastoreServiceScalingConfigAutoscalingConfigLimitConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocMetastoreServiceScalingConfigAutoscalingConfigLimitConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleDataprocMetastoreServiceScalingConfigAutoscalingConfigLimitConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleDataprocMetastoreServiceScalingConfigAutoscalingConfigLimitConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDataprocMetastoreServiceScalingConfigAutoscalingConfigLimitConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDataprocMetastoreServiceScalingConfigAutoscalingConfigLimitConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDataprocMetastoreService.GoogleDataprocMetastoreServiceScalingConfigAutoscalingConfigLimitConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleDataprocMetastoreServiceScalingConfigAutoscalingConfigLimitConfigOutputReference_Override(g GoogleDataprocMetastoreServiceScalingConfigAutoscalingConfigLimitConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDataprocMetastoreService.GoogleDataprocMetastoreServiceScalingConfigAutoscalingConfigLimitConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleDataprocMetastoreServiceScalingConfigAutoscalingConfigLimitConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocMetastoreServiceScalingConfigAutoscalingConfigLimitConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocMetastoreServiceScalingConfigAutoscalingConfigLimitConfigOutputReference)SetInternalValue(val *GoogleDataprocMetastoreServiceScalingConfigAutoscalingConfigLimitConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocMetastoreServiceScalingConfigAutoscalingConfigLimitConfigOutputReference)SetMaxScalingFactor(val *float64) {
	if err := j.validateSetMaxScalingFactorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxScalingFactor",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocMetastoreServiceScalingConfigAutoscalingConfigLimitConfigOutputReference)SetMinScalingFactor(val *float64) {
	if err := j.validateSetMinScalingFactorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minScalingFactor",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocMetastoreServiceScalingConfigAutoscalingConfigLimitConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocMetastoreServiceScalingConfigAutoscalingConfigLimitConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleDataprocMetastoreServiceScalingConfigAutoscalingConfigLimitConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataprocMetastoreServiceScalingConfigAutoscalingConfigLimitConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDataprocMetastoreServiceScalingConfigAutoscalingConfigLimitConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDataprocMetastoreServiceScalingConfigAutoscalingConfigLimitConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDataprocMetastoreServiceScalingConfigAutoscalingConfigLimitConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDataprocMetastoreServiceScalingConfigAutoscalingConfigLimitConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDataprocMetastoreServiceScalingConfigAutoscalingConfigLimitConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDataprocMetastoreServiceScalingConfigAutoscalingConfigLimitConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDataprocMetastoreServiceScalingConfigAutoscalingConfigLimitConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDataprocMetastoreServiceScalingConfigAutoscalingConfigLimitConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDataprocMetastoreServiceScalingConfigAutoscalingConfigLimitConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataprocMetastoreServiceScalingConfigAutoscalingConfigLimitConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDataprocMetastoreServiceScalingConfigAutoscalingConfigLimitConfigOutputReference) ResetMaxScalingFactor() {
	_jsii_.InvokeVoid(
		g,
		"resetMaxScalingFactor",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocMetastoreServiceScalingConfigAutoscalingConfigLimitConfigOutputReference) ResetMinScalingFactor() {
	_jsii_.InvokeVoid(
		g,
		"resetMinScalingFactor",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocMetastoreServiceScalingConfigAutoscalingConfigLimitConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleDataprocMetastoreServiceScalingConfigAutoscalingConfigLimitConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

