package googleapphubapplication

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googleapphubapplication/internal"
)

type GoogleApphubApplicationAttributesOutputReference interface {
	cdktf.ComplexObject
	BusinessOwners() GoogleApphubApplicationAttributesBusinessOwnersList
	BusinessOwnersInput() interface{}
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
	Criticality() GoogleApphubApplicationAttributesCriticalityOutputReference
	CriticalityInput() *GoogleApphubApplicationAttributesCriticality
	DeveloperOwners() GoogleApphubApplicationAttributesDeveloperOwnersList
	DeveloperOwnersInput() interface{}
	Environment() GoogleApphubApplicationAttributesEnvironmentOutputReference
	EnvironmentInput() *GoogleApphubApplicationAttributesEnvironment
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleApphubApplicationAttributes
	SetInternalValue(val *GoogleApphubApplicationAttributes)
	OperatorOwners() GoogleApphubApplicationAttributesOperatorOwnersList
	OperatorOwnersInput() interface{}
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
	PutBusinessOwners(value interface{})
	PutCriticality(value *GoogleApphubApplicationAttributesCriticality)
	PutDeveloperOwners(value interface{})
	PutEnvironment(value *GoogleApphubApplicationAttributesEnvironment)
	PutOperatorOwners(value interface{})
	ResetBusinessOwners()
	ResetCriticality()
	ResetDeveloperOwners()
	ResetEnvironment()
	ResetOperatorOwners()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleApphubApplicationAttributesOutputReference
type jsiiProxy_GoogleApphubApplicationAttributesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleApphubApplicationAttributesOutputReference) BusinessOwners() GoogleApphubApplicationAttributesBusinessOwnersList {
	var returns GoogleApphubApplicationAttributesBusinessOwnersList
	_jsii_.Get(
		j,
		"businessOwners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApphubApplicationAttributesOutputReference) BusinessOwnersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"businessOwnersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApphubApplicationAttributesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApphubApplicationAttributesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApphubApplicationAttributesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApphubApplicationAttributesOutputReference) Criticality() GoogleApphubApplicationAttributesCriticalityOutputReference {
	var returns GoogleApphubApplicationAttributesCriticalityOutputReference
	_jsii_.Get(
		j,
		"criticality",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApphubApplicationAttributesOutputReference) CriticalityInput() *GoogleApphubApplicationAttributesCriticality {
	var returns *GoogleApphubApplicationAttributesCriticality
	_jsii_.Get(
		j,
		"criticalityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApphubApplicationAttributesOutputReference) DeveloperOwners() GoogleApphubApplicationAttributesDeveloperOwnersList {
	var returns GoogleApphubApplicationAttributesDeveloperOwnersList
	_jsii_.Get(
		j,
		"developerOwners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApphubApplicationAttributesOutputReference) DeveloperOwnersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"developerOwnersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApphubApplicationAttributesOutputReference) Environment() GoogleApphubApplicationAttributesEnvironmentOutputReference {
	var returns GoogleApphubApplicationAttributesEnvironmentOutputReference
	_jsii_.Get(
		j,
		"environment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApphubApplicationAttributesOutputReference) EnvironmentInput() *GoogleApphubApplicationAttributesEnvironment {
	var returns *GoogleApphubApplicationAttributesEnvironment
	_jsii_.Get(
		j,
		"environmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApphubApplicationAttributesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApphubApplicationAttributesOutputReference) InternalValue() *GoogleApphubApplicationAttributes {
	var returns *GoogleApphubApplicationAttributes
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApphubApplicationAttributesOutputReference) OperatorOwners() GoogleApphubApplicationAttributesOperatorOwnersList {
	var returns GoogleApphubApplicationAttributesOperatorOwnersList
	_jsii_.Get(
		j,
		"operatorOwners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApphubApplicationAttributesOutputReference) OperatorOwnersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"operatorOwnersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApphubApplicationAttributesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApphubApplicationAttributesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleApphubApplicationAttributesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleApphubApplicationAttributesOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleApphubApplicationAttributesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleApphubApplicationAttributesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleApphubApplication.GoogleApphubApplicationAttributesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleApphubApplicationAttributesOutputReference_Override(g GoogleApphubApplicationAttributesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleApphubApplication.GoogleApphubApplicationAttributesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleApphubApplicationAttributesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleApphubApplicationAttributesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleApphubApplicationAttributesOutputReference)SetInternalValue(val *GoogleApphubApplicationAttributes) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleApphubApplicationAttributesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleApphubApplicationAttributesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleApphubApplicationAttributesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleApphubApplicationAttributesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleApphubApplicationAttributesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleApphubApplicationAttributesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleApphubApplicationAttributesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleApphubApplicationAttributesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleApphubApplicationAttributesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleApphubApplicationAttributesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleApphubApplicationAttributesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleApphubApplicationAttributesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleApphubApplicationAttributesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleApphubApplicationAttributesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleApphubApplicationAttributesOutputReference) PutBusinessOwners(value interface{}) {
	if err := g.validatePutBusinessOwnersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putBusinessOwners",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleApphubApplicationAttributesOutputReference) PutCriticality(value *GoogleApphubApplicationAttributesCriticality) {
	if err := g.validatePutCriticalityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putCriticality",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleApphubApplicationAttributesOutputReference) PutDeveloperOwners(value interface{}) {
	if err := g.validatePutDeveloperOwnersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDeveloperOwners",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleApphubApplicationAttributesOutputReference) PutEnvironment(value *GoogleApphubApplicationAttributesEnvironment) {
	if err := g.validatePutEnvironmentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putEnvironment",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleApphubApplicationAttributesOutputReference) PutOperatorOwners(value interface{}) {
	if err := g.validatePutOperatorOwnersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putOperatorOwners",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleApphubApplicationAttributesOutputReference) ResetBusinessOwners() {
	_jsii_.InvokeVoid(
		g,
		"resetBusinessOwners",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleApphubApplicationAttributesOutputReference) ResetCriticality() {
	_jsii_.InvokeVoid(
		g,
		"resetCriticality",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleApphubApplicationAttributesOutputReference) ResetDeveloperOwners() {
	_jsii_.InvokeVoid(
		g,
		"resetDeveloperOwners",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleApphubApplicationAttributesOutputReference) ResetEnvironment() {
	_jsii_.InvokeVoid(
		g,
		"resetEnvironment",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleApphubApplicationAttributesOutputReference) ResetOperatorOwners() {
	_jsii_.InvokeVoid(
		g,
		"resetOperatorOwners",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleApphubApplicationAttributesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleApphubApplicationAttributesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

