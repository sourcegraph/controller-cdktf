package googleprivatecacertificateauthority

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/googlebeta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/googlebeta/googleprivatecacertificateauthority/internal"
)

type GooglePrivatecaCertificateAuthorityConfigX509ConfigKeyUsageOutputReference interface {
	cdktf.ComplexObject
	BaseKeyUsage() GooglePrivatecaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsageOutputReference
	BaseKeyUsageInput() *GooglePrivatecaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsage
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
	ExtendedKeyUsage() GooglePrivatecaCertificateAuthorityConfigX509ConfigKeyUsageExtendedKeyUsageOutputReference
	ExtendedKeyUsageInput() *GooglePrivatecaCertificateAuthorityConfigX509ConfigKeyUsageExtendedKeyUsage
	// Experimental.
	Fqn() *string
	InternalValue() *GooglePrivatecaCertificateAuthorityConfigX509ConfigKeyUsage
	SetInternalValue(val *GooglePrivatecaCertificateAuthorityConfigX509ConfigKeyUsage)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UnknownExtendedKeyUsages() GooglePrivatecaCertificateAuthorityConfigX509ConfigKeyUsageUnknownExtendedKeyUsagesList
	UnknownExtendedKeyUsagesInput() interface{}
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
	PutBaseKeyUsage(value *GooglePrivatecaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsage)
	PutExtendedKeyUsage(value *GooglePrivatecaCertificateAuthorityConfigX509ConfigKeyUsageExtendedKeyUsage)
	PutUnknownExtendedKeyUsages(value interface{})
	ResetUnknownExtendedKeyUsages()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GooglePrivatecaCertificateAuthorityConfigX509ConfigKeyUsageOutputReference
type jsiiProxy_GooglePrivatecaCertificateAuthorityConfigX509ConfigKeyUsageOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GooglePrivatecaCertificateAuthorityConfigX509ConfigKeyUsageOutputReference) BaseKeyUsage() GooglePrivatecaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsageOutputReference {
	var returns GooglePrivatecaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsageOutputReference
	_jsii_.Get(
		j,
		"baseKeyUsage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateAuthorityConfigX509ConfigKeyUsageOutputReference) BaseKeyUsageInput() *GooglePrivatecaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsage {
	var returns *GooglePrivatecaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsage
	_jsii_.Get(
		j,
		"baseKeyUsageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateAuthorityConfigX509ConfigKeyUsageOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateAuthorityConfigX509ConfigKeyUsageOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateAuthorityConfigX509ConfigKeyUsageOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateAuthorityConfigX509ConfigKeyUsageOutputReference) ExtendedKeyUsage() GooglePrivatecaCertificateAuthorityConfigX509ConfigKeyUsageExtendedKeyUsageOutputReference {
	var returns GooglePrivatecaCertificateAuthorityConfigX509ConfigKeyUsageExtendedKeyUsageOutputReference
	_jsii_.Get(
		j,
		"extendedKeyUsage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateAuthorityConfigX509ConfigKeyUsageOutputReference) ExtendedKeyUsageInput() *GooglePrivatecaCertificateAuthorityConfigX509ConfigKeyUsageExtendedKeyUsage {
	var returns *GooglePrivatecaCertificateAuthorityConfigX509ConfigKeyUsageExtendedKeyUsage
	_jsii_.Get(
		j,
		"extendedKeyUsageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateAuthorityConfigX509ConfigKeyUsageOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateAuthorityConfigX509ConfigKeyUsageOutputReference) InternalValue() *GooglePrivatecaCertificateAuthorityConfigX509ConfigKeyUsage {
	var returns *GooglePrivatecaCertificateAuthorityConfigX509ConfigKeyUsage
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateAuthorityConfigX509ConfigKeyUsageOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateAuthorityConfigX509ConfigKeyUsageOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateAuthorityConfigX509ConfigKeyUsageOutputReference) UnknownExtendedKeyUsages() GooglePrivatecaCertificateAuthorityConfigX509ConfigKeyUsageUnknownExtendedKeyUsagesList {
	var returns GooglePrivatecaCertificateAuthorityConfigX509ConfigKeyUsageUnknownExtendedKeyUsagesList
	_jsii_.Get(
		j,
		"unknownExtendedKeyUsages",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateAuthorityConfigX509ConfigKeyUsageOutputReference) UnknownExtendedKeyUsagesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"unknownExtendedKeyUsagesInput",
		&returns,
	)
	return returns
}


func NewGooglePrivatecaCertificateAuthorityConfigX509ConfigKeyUsageOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GooglePrivatecaCertificateAuthorityConfigX509ConfigKeyUsageOutputReference {
	_init_.Initialize()

	if err := validateNewGooglePrivatecaCertificateAuthorityConfigX509ConfigKeyUsageOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GooglePrivatecaCertificateAuthorityConfigX509ConfigKeyUsageOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-googlebeta.googlePrivatecaCertificateAuthority.GooglePrivatecaCertificateAuthorityConfigX509ConfigKeyUsageOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGooglePrivatecaCertificateAuthorityConfigX509ConfigKeyUsageOutputReference_Override(g GooglePrivatecaCertificateAuthorityConfigX509ConfigKeyUsageOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-googlebeta.googlePrivatecaCertificateAuthority.GooglePrivatecaCertificateAuthorityConfigX509ConfigKeyUsageOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GooglePrivatecaCertificateAuthorityConfigX509ConfigKeyUsageOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GooglePrivatecaCertificateAuthorityConfigX509ConfigKeyUsageOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GooglePrivatecaCertificateAuthorityConfigX509ConfigKeyUsageOutputReference)SetInternalValue(val *GooglePrivatecaCertificateAuthorityConfigX509ConfigKeyUsage) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GooglePrivatecaCertificateAuthorityConfigX509ConfigKeyUsageOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GooglePrivatecaCertificateAuthorityConfigX509ConfigKeyUsageOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GooglePrivatecaCertificateAuthorityConfigX509ConfigKeyUsageOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GooglePrivatecaCertificateAuthorityConfigX509ConfigKeyUsageOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GooglePrivatecaCertificateAuthorityConfigX509ConfigKeyUsageOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GooglePrivatecaCertificateAuthorityConfigX509ConfigKeyUsageOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GooglePrivatecaCertificateAuthorityConfigX509ConfigKeyUsageOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GooglePrivatecaCertificateAuthorityConfigX509ConfigKeyUsageOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GooglePrivatecaCertificateAuthorityConfigX509ConfigKeyUsageOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GooglePrivatecaCertificateAuthorityConfigX509ConfigKeyUsageOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GooglePrivatecaCertificateAuthorityConfigX509ConfigKeyUsageOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GooglePrivatecaCertificateAuthorityConfigX509ConfigKeyUsageOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GooglePrivatecaCertificateAuthorityConfigX509ConfigKeyUsageOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GooglePrivatecaCertificateAuthorityConfigX509ConfigKeyUsageOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GooglePrivatecaCertificateAuthorityConfigX509ConfigKeyUsageOutputReference) PutBaseKeyUsage(value *GooglePrivatecaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsage) {
	if err := g.validatePutBaseKeyUsageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putBaseKeyUsage",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GooglePrivatecaCertificateAuthorityConfigX509ConfigKeyUsageOutputReference) PutExtendedKeyUsage(value *GooglePrivatecaCertificateAuthorityConfigX509ConfigKeyUsageExtendedKeyUsage) {
	if err := g.validatePutExtendedKeyUsageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putExtendedKeyUsage",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GooglePrivatecaCertificateAuthorityConfigX509ConfigKeyUsageOutputReference) PutUnknownExtendedKeyUsages(value interface{}) {
	if err := g.validatePutUnknownExtendedKeyUsagesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putUnknownExtendedKeyUsages",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GooglePrivatecaCertificateAuthorityConfigX509ConfigKeyUsageOutputReference) ResetUnknownExtendedKeyUsages() {
	_jsii_.InvokeVoid(
		g,
		"resetUnknownExtendedKeyUsages",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePrivatecaCertificateAuthorityConfigX509ConfigKeyUsageOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GooglePrivatecaCertificateAuthorityConfigX509ConfigKeyUsageOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

