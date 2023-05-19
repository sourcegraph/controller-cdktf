package googleosconfigguestpolicies

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googleosconfigguestpolicies/internal"
)

type GoogleOsConfigGuestPoliciesPackageRepositoriesAptOutputReference interface {
	cdktf.ComplexObject
	ArchiveType() *string
	SetArchiveType(val *string)
	ArchiveTypeInput() *string
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
	Components() *[]*string
	SetComponents(val *[]*string)
	ComponentsInput() *[]*string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Distribution() *string
	SetDistribution(val *string)
	DistributionInput() *string
	// Experimental.
	Fqn() *string
	GpgKey() *string
	SetGpgKey(val *string)
	GpgKeyInput() *string
	InternalValue() *GoogleOsConfigGuestPoliciesPackageRepositoriesApt
	SetInternalValue(val *GoogleOsConfigGuestPoliciesPackageRepositoriesApt)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Uri() *string
	SetUri(val *string)
	UriInput() *string
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
	ResetArchiveType()
	ResetGpgKey()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleOsConfigGuestPoliciesPackageRepositoriesAptOutputReference
type jsiiProxy_GoogleOsConfigGuestPoliciesPackageRepositoriesAptOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesPackageRepositoriesAptOutputReference) ArchiveType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"archiveType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesPackageRepositoriesAptOutputReference) ArchiveTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"archiveTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesPackageRepositoriesAptOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesPackageRepositoriesAptOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesPackageRepositoriesAptOutputReference) Components() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"components",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesPackageRepositoriesAptOutputReference) ComponentsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"componentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesPackageRepositoriesAptOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesPackageRepositoriesAptOutputReference) Distribution() *string {
	var returns *string
	_jsii_.Get(
		j,
		"distribution",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesPackageRepositoriesAptOutputReference) DistributionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"distributionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesPackageRepositoriesAptOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesPackageRepositoriesAptOutputReference) GpgKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gpgKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesPackageRepositoriesAptOutputReference) GpgKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gpgKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesPackageRepositoriesAptOutputReference) InternalValue() *GoogleOsConfigGuestPoliciesPackageRepositoriesApt {
	var returns *GoogleOsConfigGuestPoliciesPackageRepositoriesApt
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesPackageRepositoriesAptOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesPackageRepositoriesAptOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesPackageRepositoriesAptOutputReference) Uri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesPackageRepositoriesAptOutputReference) UriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uriInput",
		&returns,
	)
	return returns
}


func NewGoogleOsConfigGuestPoliciesPackageRepositoriesAptOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleOsConfigGuestPoliciesPackageRepositoriesAptOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleOsConfigGuestPoliciesPackageRepositoriesAptOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleOsConfigGuestPoliciesPackageRepositoriesAptOutputReference{}

	_jsii_.Create(
		"google-beta.googleOsConfigGuestPolicies.GoogleOsConfigGuestPoliciesPackageRepositoriesAptOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleOsConfigGuestPoliciesPackageRepositoriesAptOutputReference_Override(g GoogleOsConfigGuestPoliciesPackageRepositoriesAptOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"google-beta.googleOsConfigGuestPolicies.GoogleOsConfigGuestPoliciesPackageRepositoriesAptOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesPackageRepositoriesAptOutputReference)SetArchiveType(val *string) {
	if err := j.validateSetArchiveTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"archiveType",
		val,
	)
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesPackageRepositoriesAptOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesPackageRepositoriesAptOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesPackageRepositoriesAptOutputReference)SetComponents(val *[]*string) {
	if err := j.validateSetComponentsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"components",
		val,
	)
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesPackageRepositoriesAptOutputReference)SetDistribution(val *string) {
	if err := j.validateSetDistributionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"distribution",
		val,
	)
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesPackageRepositoriesAptOutputReference)SetGpgKey(val *string) {
	if err := j.validateSetGpgKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gpgKey",
		val,
	)
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesPackageRepositoriesAptOutputReference)SetInternalValue(val *GoogleOsConfigGuestPoliciesPackageRepositoriesApt) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesPackageRepositoriesAptOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesPackageRepositoriesAptOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesPackageRepositoriesAptOutputReference)SetUri(val *string) {
	if err := j.validateSetUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"uri",
		val,
	)
}

func (g *jsiiProxy_GoogleOsConfigGuestPoliciesPackageRepositoriesAptOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleOsConfigGuestPoliciesPackageRepositoriesAptOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleOsConfigGuestPoliciesPackageRepositoriesAptOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleOsConfigGuestPoliciesPackageRepositoriesAptOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleOsConfigGuestPoliciesPackageRepositoriesAptOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleOsConfigGuestPoliciesPackageRepositoriesAptOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleOsConfigGuestPoliciesPackageRepositoriesAptOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleOsConfigGuestPoliciesPackageRepositoriesAptOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleOsConfigGuestPoliciesPackageRepositoriesAptOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleOsConfigGuestPoliciesPackageRepositoriesAptOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleOsConfigGuestPoliciesPackageRepositoriesAptOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleOsConfigGuestPoliciesPackageRepositoriesAptOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleOsConfigGuestPoliciesPackageRepositoriesAptOutputReference) ResetArchiveType() {
	_jsii_.InvokeVoid(
		g,
		"resetArchiveType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOsConfigGuestPoliciesPackageRepositoriesAptOutputReference) ResetGpgKey() {
	_jsii_.InvokeVoid(
		g,
		"resetGpgKey",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOsConfigGuestPoliciesPackageRepositoriesAptOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleOsConfigGuestPoliciesPackageRepositoriesAptOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

