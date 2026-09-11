package googlesaasruntimeunitkind

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlesaasruntimeunitkind/internal"
)

type GoogleSaasRuntimeUnitKindOutputVariableMappingsToOutputReference interface {
	cdktn.ComplexObject
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
	Dependency() *string
	SetDependency(val *string)
	DependencyInput() *string
	// Experimental.
	Fqn() *string
	IgnoreForLookup() interface{}
	SetIgnoreForLookup(val interface{})
	IgnoreForLookupInput() interface{}
	InputVariable() *string
	SetInputVariable(val *string)
	InputVariableInput() *string
	InternalValue() *GoogleSaasRuntimeUnitKindOutputVariableMappingsTo
	SetInternalValue(val *GoogleSaasRuntimeUnitKindOutputVariableMappingsTo)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable
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
	InterpolationAsList() cdktn.IResolvable
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	ResetIgnoreForLookup()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleSaasRuntimeUnitKindOutputVariableMappingsToOutputReference
type jsiiProxy_GoogleSaasRuntimeUnitKindOutputVariableMappingsToOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleSaasRuntimeUnitKindOutputVariableMappingsToOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSaasRuntimeUnitKindOutputVariableMappingsToOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSaasRuntimeUnitKindOutputVariableMappingsToOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSaasRuntimeUnitKindOutputVariableMappingsToOutputReference) Dependency() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dependency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSaasRuntimeUnitKindOutputVariableMappingsToOutputReference) DependencyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dependencyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSaasRuntimeUnitKindOutputVariableMappingsToOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSaasRuntimeUnitKindOutputVariableMappingsToOutputReference) IgnoreForLookup() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreForLookup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSaasRuntimeUnitKindOutputVariableMappingsToOutputReference) IgnoreForLookupInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreForLookupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSaasRuntimeUnitKindOutputVariableMappingsToOutputReference) InputVariable() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputVariable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSaasRuntimeUnitKindOutputVariableMappingsToOutputReference) InputVariableInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputVariableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSaasRuntimeUnitKindOutputVariableMappingsToOutputReference) InternalValue() *GoogleSaasRuntimeUnitKindOutputVariableMappingsTo {
	var returns *GoogleSaasRuntimeUnitKindOutputVariableMappingsTo
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSaasRuntimeUnitKindOutputVariableMappingsToOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSaasRuntimeUnitKindOutputVariableMappingsToOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleSaasRuntimeUnitKindOutputVariableMappingsToOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleSaasRuntimeUnitKindOutputVariableMappingsToOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleSaasRuntimeUnitKindOutputVariableMappingsToOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleSaasRuntimeUnitKindOutputVariableMappingsToOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleSaasRuntimeUnitKind.GoogleSaasRuntimeUnitKindOutputVariableMappingsToOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleSaasRuntimeUnitKindOutputVariableMappingsToOutputReference_Override(g GoogleSaasRuntimeUnitKindOutputVariableMappingsToOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleSaasRuntimeUnitKind.GoogleSaasRuntimeUnitKindOutputVariableMappingsToOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleSaasRuntimeUnitKindOutputVariableMappingsToOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleSaasRuntimeUnitKindOutputVariableMappingsToOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleSaasRuntimeUnitKindOutputVariableMappingsToOutputReference)SetDependency(val *string) {
	if err := j.validateSetDependencyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dependency",
		val,
	)
}

func (j *jsiiProxy_GoogleSaasRuntimeUnitKindOutputVariableMappingsToOutputReference)SetIgnoreForLookup(val interface{}) {
	if err := j.validateSetIgnoreForLookupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ignoreForLookup",
		val,
	)
}

func (j *jsiiProxy_GoogleSaasRuntimeUnitKindOutputVariableMappingsToOutputReference)SetInputVariable(val *string) {
	if err := j.validateSetInputVariableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inputVariable",
		val,
	)
}

func (j *jsiiProxy_GoogleSaasRuntimeUnitKindOutputVariableMappingsToOutputReference)SetInternalValue(val *GoogleSaasRuntimeUnitKindOutputVariableMappingsTo) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleSaasRuntimeUnitKindOutputVariableMappingsToOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleSaasRuntimeUnitKindOutputVariableMappingsToOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleSaasRuntimeUnitKindOutputVariableMappingsToOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleSaasRuntimeUnitKindOutputVariableMappingsToOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleSaasRuntimeUnitKindOutputVariableMappingsToOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleSaasRuntimeUnitKindOutputVariableMappingsToOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleSaasRuntimeUnitKindOutputVariableMappingsToOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleSaasRuntimeUnitKindOutputVariableMappingsToOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleSaasRuntimeUnitKindOutputVariableMappingsToOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleSaasRuntimeUnitKindOutputVariableMappingsToOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleSaasRuntimeUnitKindOutputVariableMappingsToOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleSaasRuntimeUnitKindOutputVariableMappingsToOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleSaasRuntimeUnitKindOutputVariableMappingsToOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleSaasRuntimeUnitKindOutputVariableMappingsToOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleSaasRuntimeUnitKindOutputVariableMappingsToOutputReference) ResetIgnoreForLookup() {
	_jsii_.InvokeVoid(
		g,
		"resetIgnoreForLookup",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSaasRuntimeUnitKindOutputVariableMappingsToOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := g.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleSaasRuntimeUnitKindOutputVariableMappingsToOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

