package googleosconfigguestpolicies

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/googlebeta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/googlebeta/googleosconfigguestpolicies/internal"
)

type GoogleOsConfigGuestPoliciesRecipesArtifactsOutputReference interface {
	cdktf.ComplexObject
	AllowInsecure() interface{}
	SetAllowInsecure(val interface{})
	AllowInsecureInput() interface{}
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
	Gcs() GoogleOsConfigGuestPoliciesRecipesArtifactsGcsOutputReference
	GcsInput() *GoogleOsConfigGuestPoliciesRecipesArtifactsGcs
	Id() *string
	SetId(val *string)
	IdInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Remote() GoogleOsConfigGuestPoliciesRecipesArtifactsRemoteOutputReference
	RemoteInput() *GoogleOsConfigGuestPoliciesRecipesArtifactsRemote
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
	PutGcs(value *GoogleOsConfigGuestPoliciesRecipesArtifactsGcs)
	PutRemote(value *GoogleOsConfigGuestPoliciesRecipesArtifactsRemote)
	ResetAllowInsecure()
	ResetGcs()
	ResetRemote()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleOsConfigGuestPoliciesRecipesArtifactsOutputReference
type jsiiProxy_GoogleOsConfigGuestPoliciesRecipesArtifactsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesArtifactsOutputReference) AllowInsecure() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowInsecure",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesArtifactsOutputReference) AllowInsecureInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowInsecureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesArtifactsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesArtifactsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesArtifactsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesArtifactsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesArtifactsOutputReference) Gcs() GoogleOsConfigGuestPoliciesRecipesArtifactsGcsOutputReference {
	var returns GoogleOsConfigGuestPoliciesRecipesArtifactsGcsOutputReference
	_jsii_.Get(
		j,
		"gcs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesArtifactsOutputReference) GcsInput() *GoogleOsConfigGuestPoliciesRecipesArtifactsGcs {
	var returns *GoogleOsConfigGuestPoliciesRecipesArtifactsGcs
	_jsii_.Get(
		j,
		"gcsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesArtifactsOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesArtifactsOutputReference) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesArtifactsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesArtifactsOutputReference) Remote() GoogleOsConfigGuestPoliciesRecipesArtifactsRemoteOutputReference {
	var returns GoogleOsConfigGuestPoliciesRecipesArtifactsRemoteOutputReference
	_jsii_.Get(
		j,
		"remote",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesArtifactsOutputReference) RemoteInput() *GoogleOsConfigGuestPoliciesRecipesArtifactsRemote {
	var returns *GoogleOsConfigGuestPoliciesRecipesArtifactsRemote
	_jsii_.Get(
		j,
		"remoteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesArtifactsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesArtifactsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleOsConfigGuestPoliciesRecipesArtifactsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GoogleOsConfigGuestPoliciesRecipesArtifactsOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleOsConfigGuestPoliciesRecipesArtifactsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleOsConfigGuestPoliciesRecipesArtifactsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-googlebeta.googleOsConfigGuestPolicies.GoogleOsConfigGuestPoliciesRecipesArtifactsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGoogleOsConfigGuestPoliciesRecipesArtifactsOutputReference_Override(g GoogleOsConfigGuestPoliciesRecipesArtifactsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-googlebeta.googleOsConfigGuestPolicies.GoogleOsConfigGuestPoliciesRecipesArtifactsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesArtifactsOutputReference)SetAllowInsecure(val interface{}) {
	if err := j.validateSetAllowInsecureParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowInsecure",
		val,
	)
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesArtifactsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesArtifactsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesArtifactsOutputReference)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesArtifactsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesArtifactsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesArtifactsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesArtifactsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesArtifactsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesArtifactsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesArtifactsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesArtifactsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesArtifactsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesArtifactsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesArtifactsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesArtifactsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesArtifactsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesArtifactsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesArtifactsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesArtifactsOutputReference) PutGcs(value *GoogleOsConfigGuestPoliciesRecipesArtifactsGcs) {
	if err := g.validatePutGcsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putGcs",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesArtifactsOutputReference) PutRemote(value *GoogleOsConfigGuestPoliciesRecipesArtifactsRemote) {
	if err := g.validatePutRemoteParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putRemote",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesArtifactsOutputReference) ResetAllowInsecure() {
	_jsii_.InvokeVoid(
		g,
		"resetAllowInsecure",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesArtifactsOutputReference) ResetGcs() {
	_jsii_.InvokeVoid(
		g,
		"resetGcs",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesArtifactsOutputReference) ResetRemote() {
	_jsii_.InvokeVoid(
		g,
		"resetRemote",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesArtifactsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesArtifactsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

