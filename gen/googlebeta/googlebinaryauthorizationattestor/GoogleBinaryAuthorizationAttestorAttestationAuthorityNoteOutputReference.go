package googlebinaryauthorizationattestor

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/googlebeta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/googlebeta/googlebinaryauthorizationattestor/internal"
)

type GoogleBinaryAuthorizationAttestorAttestationAuthorityNoteOutputReference interface {
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
	DelegationServiceAccountEmail() *string
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleBinaryAuthorizationAttestorAttestationAuthorityNote
	SetInternalValue(val *GoogleBinaryAuthorizationAttestorAttestationAuthorityNote)
	NoteReference() *string
	SetNoteReference(val *string)
	NoteReferenceInput() *string
	PublicKeys() GoogleBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysList
	PublicKeysInput() interface{}
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
	PutPublicKeys(value interface{})
	ResetPublicKeys()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleBinaryAuthorizationAttestorAttestationAuthorityNoteOutputReference
type jsiiProxy_GoogleBinaryAuthorizationAttestorAttestationAuthorityNoteOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleBinaryAuthorizationAttestorAttestationAuthorityNoteOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBinaryAuthorizationAttestorAttestationAuthorityNoteOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBinaryAuthorizationAttestorAttestationAuthorityNoteOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBinaryAuthorizationAttestorAttestationAuthorityNoteOutputReference) DelegationServiceAccountEmail() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delegationServiceAccountEmail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBinaryAuthorizationAttestorAttestationAuthorityNoteOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBinaryAuthorizationAttestorAttestationAuthorityNoteOutputReference) InternalValue() *GoogleBinaryAuthorizationAttestorAttestationAuthorityNote {
	var returns *GoogleBinaryAuthorizationAttestorAttestationAuthorityNote
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBinaryAuthorizationAttestorAttestationAuthorityNoteOutputReference) NoteReference() *string {
	var returns *string
	_jsii_.Get(
		j,
		"noteReference",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBinaryAuthorizationAttestorAttestationAuthorityNoteOutputReference) NoteReferenceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"noteReferenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBinaryAuthorizationAttestorAttestationAuthorityNoteOutputReference) PublicKeys() GoogleBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysList {
	var returns GoogleBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysList
	_jsii_.Get(
		j,
		"publicKeys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBinaryAuthorizationAttestorAttestationAuthorityNoteOutputReference) PublicKeysInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publicKeysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBinaryAuthorizationAttestorAttestationAuthorityNoteOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBinaryAuthorizationAttestorAttestationAuthorityNoteOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleBinaryAuthorizationAttestorAttestationAuthorityNoteOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleBinaryAuthorizationAttestorAttestationAuthorityNoteOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleBinaryAuthorizationAttestorAttestationAuthorityNoteOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleBinaryAuthorizationAttestorAttestationAuthorityNoteOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-googlebeta.googleBinaryAuthorizationAttestor.GoogleBinaryAuthorizationAttestorAttestationAuthorityNoteOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleBinaryAuthorizationAttestorAttestationAuthorityNoteOutputReference_Override(g GoogleBinaryAuthorizationAttestorAttestationAuthorityNoteOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-googlebeta.googleBinaryAuthorizationAttestor.GoogleBinaryAuthorizationAttestorAttestationAuthorityNoteOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleBinaryAuthorizationAttestorAttestationAuthorityNoteOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleBinaryAuthorizationAttestorAttestationAuthorityNoteOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleBinaryAuthorizationAttestorAttestationAuthorityNoteOutputReference)SetInternalValue(val *GoogleBinaryAuthorizationAttestorAttestationAuthorityNote) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleBinaryAuthorizationAttestorAttestationAuthorityNoteOutputReference)SetNoteReference(val *string) {
	if err := j.validateSetNoteReferenceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noteReference",
		val,
	)
}

func (j *jsiiProxy_GoogleBinaryAuthorizationAttestorAttestationAuthorityNoteOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleBinaryAuthorizationAttestorAttestationAuthorityNoteOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleBinaryAuthorizationAttestorAttestationAuthorityNoteOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBinaryAuthorizationAttestorAttestationAuthorityNoteOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleBinaryAuthorizationAttestorAttestationAuthorityNoteOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleBinaryAuthorizationAttestorAttestationAuthorityNoteOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleBinaryAuthorizationAttestorAttestationAuthorityNoteOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleBinaryAuthorizationAttestorAttestationAuthorityNoteOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleBinaryAuthorizationAttestorAttestationAuthorityNoteOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleBinaryAuthorizationAttestorAttestationAuthorityNoteOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleBinaryAuthorizationAttestorAttestationAuthorityNoteOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleBinaryAuthorizationAttestorAttestationAuthorityNoteOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleBinaryAuthorizationAttestorAttestationAuthorityNoteOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBinaryAuthorizationAttestorAttestationAuthorityNoteOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleBinaryAuthorizationAttestorAttestationAuthorityNoteOutputReference) PutPublicKeys(value interface{}) {
	if err := g.validatePutPublicKeysParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPublicKeys",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleBinaryAuthorizationAttestorAttestationAuthorityNoteOutputReference) ResetPublicKeys() {
	_jsii_.InvokeVoid(
		g,
		"resetPublicKeys",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBinaryAuthorizationAttestorAttestationAuthorityNoteOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleBinaryAuthorizationAttestorAttestationAuthorityNoteOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

