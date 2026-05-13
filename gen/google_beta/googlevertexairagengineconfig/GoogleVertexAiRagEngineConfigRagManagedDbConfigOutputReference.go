package googlevertexairagengineconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlevertexairagengineconfig/internal"
)

type GoogleVertexAiRagEngineConfigRagManagedDbConfigOutputReference interface {
	cdktf.ComplexObject
	Basic() GoogleVertexAiRagEngineConfigRagManagedDbConfigBasicOutputReference
	BasicInput() *GoogleVertexAiRagEngineConfigRagManagedDbConfigBasic
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
	InternalValue() *GoogleVertexAiRagEngineConfigRagManagedDbConfig
	SetInternalValue(val *GoogleVertexAiRagEngineConfigRagManagedDbConfig)
	Scaled() GoogleVertexAiRagEngineConfigRagManagedDbConfigScaledOutputReference
	ScaledInput() *GoogleVertexAiRagEngineConfigRagManagedDbConfigScaled
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Unprovisioned() GoogleVertexAiRagEngineConfigRagManagedDbConfigUnprovisionedOutputReference
	UnprovisionedInput() *GoogleVertexAiRagEngineConfigRagManagedDbConfigUnprovisioned
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
	PutBasic(value *GoogleVertexAiRagEngineConfigRagManagedDbConfigBasic)
	PutScaled(value *GoogleVertexAiRagEngineConfigRagManagedDbConfigScaled)
	PutUnprovisioned(value *GoogleVertexAiRagEngineConfigRagManagedDbConfigUnprovisioned)
	ResetBasic()
	ResetScaled()
	ResetUnprovisioned()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleVertexAiRagEngineConfigRagManagedDbConfigOutputReference
type jsiiProxy_GoogleVertexAiRagEngineConfigRagManagedDbConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleVertexAiRagEngineConfigRagManagedDbConfigOutputReference) Basic() GoogleVertexAiRagEngineConfigRagManagedDbConfigBasicOutputReference {
	var returns GoogleVertexAiRagEngineConfigRagManagedDbConfigBasicOutputReference
	_jsii_.Get(
		j,
		"basic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiRagEngineConfigRagManagedDbConfigOutputReference) BasicInput() *GoogleVertexAiRagEngineConfigRagManagedDbConfigBasic {
	var returns *GoogleVertexAiRagEngineConfigRagManagedDbConfigBasic
	_jsii_.Get(
		j,
		"basicInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiRagEngineConfigRagManagedDbConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiRagEngineConfigRagManagedDbConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiRagEngineConfigRagManagedDbConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiRagEngineConfigRagManagedDbConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiRagEngineConfigRagManagedDbConfigOutputReference) InternalValue() *GoogleVertexAiRagEngineConfigRagManagedDbConfig {
	var returns *GoogleVertexAiRagEngineConfigRagManagedDbConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiRagEngineConfigRagManagedDbConfigOutputReference) Scaled() GoogleVertexAiRagEngineConfigRagManagedDbConfigScaledOutputReference {
	var returns GoogleVertexAiRagEngineConfigRagManagedDbConfigScaledOutputReference
	_jsii_.Get(
		j,
		"scaled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiRagEngineConfigRagManagedDbConfigOutputReference) ScaledInput() *GoogleVertexAiRagEngineConfigRagManagedDbConfigScaled {
	var returns *GoogleVertexAiRagEngineConfigRagManagedDbConfigScaled
	_jsii_.Get(
		j,
		"scaledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiRagEngineConfigRagManagedDbConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiRagEngineConfigRagManagedDbConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiRagEngineConfigRagManagedDbConfigOutputReference) Unprovisioned() GoogleVertexAiRagEngineConfigRagManagedDbConfigUnprovisionedOutputReference {
	var returns GoogleVertexAiRagEngineConfigRagManagedDbConfigUnprovisionedOutputReference
	_jsii_.Get(
		j,
		"unprovisioned",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiRagEngineConfigRagManagedDbConfigOutputReference) UnprovisionedInput() *GoogleVertexAiRagEngineConfigRagManagedDbConfigUnprovisioned {
	var returns *GoogleVertexAiRagEngineConfigRagManagedDbConfigUnprovisioned
	_jsii_.Get(
		j,
		"unprovisionedInput",
		&returns,
	)
	return returns
}


func NewGoogleVertexAiRagEngineConfigRagManagedDbConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleVertexAiRagEngineConfigRagManagedDbConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleVertexAiRagEngineConfigRagManagedDbConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleVertexAiRagEngineConfigRagManagedDbConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleVertexAiRagEngineConfig.GoogleVertexAiRagEngineConfigRagManagedDbConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleVertexAiRagEngineConfigRagManagedDbConfigOutputReference_Override(g GoogleVertexAiRagEngineConfigRagManagedDbConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleVertexAiRagEngineConfig.GoogleVertexAiRagEngineConfigRagManagedDbConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleVertexAiRagEngineConfigRagManagedDbConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiRagEngineConfigRagManagedDbConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiRagEngineConfigRagManagedDbConfigOutputReference)SetInternalValue(val *GoogleVertexAiRagEngineConfigRagManagedDbConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiRagEngineConfigRagManagedDbConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiRagEngineConfigRagManagedDbConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleVertexAiRagEngineConfigRagManagedDbConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVertexAiRagEngineConfigRagManagedDbConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleVertexAiRagEngineConfigRagManagedDbConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleVertexAiRagEngineConfigRagManagedDbConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleVertexAiRagEngineConfigRagManagedDbConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleVertexAiRagEngineConfigRagManagedDbConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleVertexAiRagEngineConfigRagManagedDbConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleVertexAiRagEngineConfigRagManagedDbConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleVertexAiRagEngineConfigRagManagedDbConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleVertexAiRagEngineConfigRagManagedDbConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleVertexAiRagEngineConfigRagManagedDbConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVertexAiRagEngineConfigRagManagedDbConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleVertexAiRagEngineConfigRagManagedDbConfigOutputReference) PutBasic(value *GoogleVertexAiRagEngineConfigRagManagedDbConfigBasic) {
	if err := g.validatePutBasicParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putBasic",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleVertexAiRagEngineConfigRagManagedDbConfigOutputReference) PutScaled(value *GoogleVertexAiRagEngineConfigRagManagedDbConfigScaled) {
	if err := g.validatePutScaledParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putScaled",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleVertexAiRagEngineConfigRagManagedDbConfigOutputReference) PutUnprovisioned(value *GoogleVertexAiRagEngineConfigRagManagedDbConfigUnprovisioned) {
	if err := g.validatePutUnprovisionedParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putUnprovisioned",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleVertexAiRagEngineConfigRagManagedDbConfigOutputReference) ResetBasic() {
	_jsii_.InvokeVoid(
		g,
		"resetBasic",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiRagEngineConfigRagManagedDbConfigOutputReference) ResetScaled() {
	_jsii_.InvokeVoid(
		g,
		"resetScaled",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiRagEngineConfigRagManagedDbConfigOutputReference) ResetUnprovisioned() {
	_jsii_.InvokeVoid(
		g,
		"resetUnprovisioned",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiRagEngineConfigRagManagedDbConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleVertexAiRagEngineConfigRagManagedDbConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

