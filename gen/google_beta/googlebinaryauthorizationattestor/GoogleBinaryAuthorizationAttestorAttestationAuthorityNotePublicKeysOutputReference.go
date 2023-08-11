package googlebinaryauthorizationattestor

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlebinaryauthorizationattestor/internal"
)

type GoogleBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysOutputReference interface {
	cdktf.ComplexObject
	AsciiArmoredPgpPublicKey() *string
	SetAsciiArmoredPgpPublicKey(val *string)
	AsciiArmoredPgpPublicKeyInput() *string
	Comment() *string
	SetComment(val *string)
	CommentInput() *string
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
	Id() *string
	SetId(val *string)
	IdInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	PkixPublicKey() GoogleBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysPkixPublicKeyOutputReference
	PkixPublicKeyInput() *GoogleBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysPkixPublicKey
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
	PutPkixPublicKey(value *GoogleBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysPkixPublicKey)
	ResetAsciiArmoredPgpPublicKey()
	ResetComment()
	ResetId()
	ResetPkixPublicKey()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysOutputReference
type jsiiProxy_GoogleBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysOutputReference) AsciiArmoredPgpPublicKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"asciiArmoredPgpPublicKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysOutputReference) AsciiArmoredPgpPublicKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"asciiArmoredPgpPublicKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysOutputReference) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysOutputReference) CommentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysOutputReference) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysOutputReference) PkixPublicKey() GoogleBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysPkixPublicKeyOutputReference {
	var returns GoogleBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysPkixPublicKeyOutputReference
	_jsii_.Get(
		j,
		"pkixPublicKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysOutputReference) PkixPublicKeyInput() *GoogleBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysPkixPublicKey {
	var returns *GoogleBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysPkixPublicKey
	_jsii_.Get(
		j,
		"pkixPublicKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GoogleBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleBinaryAuthorizationAttestor.GoogleBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGoogleBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysOutputReference_Override(g GoogleBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleBinaryAuthorizationAttestor.GoogleBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GoogleBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysOutputReference)SetAsciiArmoredPgpPublicKey(val *string) {
	if err := j.validateSetAsciiArmoredPgpPublicKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"asciiArmoredPgpPublicKey",
		val,
	)
}

func (j *jsiiProxy_GoogleBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysOutputReference)SetComment(val *string) {
	if err := j.validateSetCommentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"comment",
		val,
	)
}

func (j *jsiiProxy_GoogleBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysOutputReference)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysOutputReference) PutPkixPublicKey(value *GoogleBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysPkixPublicKey) {
	if err := g.validatePutPkixPublicKeyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPkixPublicKey",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysOutputReference) ResetAsciiArmoredPgpPublicKey() {
	_jsii_.InvokeVoid(
		g,
		"resetAsciiArmoredPgpPublicKey",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysOutputReference) ResetComment() {
	_jsii_.InvokeVoid(
		g,
		"resetComment",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysOutputReference) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysOutputReference) ResetPkixPublicKey() {
	_jsii_.InvokeVoid(
		g,
		"resetPkixPublicKey",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

