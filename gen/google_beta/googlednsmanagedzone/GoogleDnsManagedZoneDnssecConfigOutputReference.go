package googlednsmanagedzone

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlednsmanagedzone/internal"
)

type GoogleDnsManagedZoneDnssecConfigOutputReference interface {
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
	DefaultKeySpecs() GoogleDnsManagedZoneDnssecConfigDefaultKeySpecsList
	DefaultKeySpecsInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleDnsManagedZoneDnssecConfig
	SetInternalValue(val *GoogleDnsManagedZoneDnssecConfig)
	Kind() *string
	SetKind(val *string)
	KindInput() *string
	NonExistence() *string
	SetNonExistence(val *string)
	NonExistenceInput() *string
	State() *string
	SetState(val *string)
	StateInput() *string
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
	PutDefaultKeySpecs(value interface{})
	ResetDefaultKeySpecs()
	ResetKind()
	ResetNonExistence()
	ResetState()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleDnsManagedZoneDnssecConfigOutputReference
type jsiiProxy_GoogleDnsManagedZoneDnssecConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleDnsManagedZoneDnssecConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDnsManagedZoneDnssecConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDnsManagedZoneDnssecConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDnsManagedZoneDnssecConfigOutputReference) DefaultKeySpecs() GoogleDnsManagedZoneDnssecConfigDefaultKeySpecsList {
	var returns GoogleDnsManagedZoneDnssecConfigDefaultKeySpecsList
	_jsii_.Get(
		j,
		"defaultKeySpecs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDnsManagedZoneDnssecConfigOutputReference) DefaultKeySpecsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"defaultKeySpecsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDnsManagedZoneDnssecConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDnsManagedZoneDnssecConfigOutputReference) InternalValue() *GoogleDnsManagedZoneDnssecConfig {
	var returns *GoogleDnsManagedZoneDnssecConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDnsManagedZoneDnssecConfigOutputReference) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDnsManagedZoneDnssecConfigOutputReference) KindInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kindInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDnsManagedZoneDnssecConfigOutputReference) NonExistence() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nonExistence",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDnsManagedZoneDnssecConfigOutputReference) NonExistenceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nonExistenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDnsManagedZoneDnssecConfigOutputReference) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDnsManagedZoneDnssecConfigOutputReference) StateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDnsManagedZoneDnssecConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDnsManagedZoneDnssecConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleDnsManagedZoneDnssecConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleDnsManagedZoneDnssecConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDnsManagedZoneDnssecConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDnsManagedZoneDnssecConfigOutputReference{}

	_jsii_.Create(
		"google-beta.googleDnsManagedZone.GoogleDnsManagedZoneDnssecConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleDnsManagedZoneDnssecConfigOutputReference_Override(g GoogleDnsManagedZoneDnssecConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"google-beta.googleDnsManagedZone.GoogleDnsManagedZoneDnssecConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleDnsManagedZoneDnssecConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDnsManagedZoneDnssecConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDnsManagedZoneDnssecConfigOutputReference)SetInternalValue(val *GoogleDnsManagedZoneDnssecConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDnsManagedZoneDnssecConfigOutputReference)SetKind(val *string) {
	if err := j.validateSetKindParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kind",
		val,
	)
}

func (j *jsiiProxy_GoogleDnsManagedZoneDnssecConfigOutputReference)SetNonExistence(val *string) {
	if err := j.validateSetNonExistenceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nonExistence",
		val,
	)
}

func (j *jsiiProxy_GoogleDnsManagedZoneDnssecConfigOutputReference)SetState(val *string) {
	if err := j.validateSetStateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"state",
		val,
	)
}

func (j *jsiiProxy_GoogleDnsManagedZoneDnssecConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDnsManagedZoneDnssecConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleDnsManagedZoneDnssecConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDnsManagedZoneDnssecConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDnsManagedZoneDnssecConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDnsManagedZoneDnssecConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDnsManagedZoneDnssecConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDnsManagedZoneDnssecConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDnsManagedZoneDnssecConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDnsManagedZoneDnssecConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDnsManagedZoneDnssecConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDnsManagedZoneDnssecConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDnsManagedZoneDnssecConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDnsManagedZoneDnssecConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDnsManagedZoneDnssecConfigOutputReference) PutDefaultKeySpecs(value interface{}) {
	if err := g.validatePutDefaultKeySpecsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDefaultKeySpecs",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDnsManagedZoneDnssecConfigOutputReference) ResetDefaultKeySpecs() {
	_jsii_.InvokeVoid(
		g,
		"resetDefaultKeySpecs",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDnsManagedZoneDnssecConfigOutputReference) ResetKind() {
	_jsii_.InvokeVoid(
		g,
		"resetKind",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDnsManagedZoneDnssecConfigOutputReference) ResetNonExistence() {
	_jsii_.InvokeVoid(
		g,
		"resetNonExistence",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDnsManagedZoneDnssecConfigOutputReference) ResetState() {
	_jsii_.InvokeVoid(
		g,
		"resetState",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDnsManagedZoneDnssecConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleDnsManagedZoneDnssecConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

