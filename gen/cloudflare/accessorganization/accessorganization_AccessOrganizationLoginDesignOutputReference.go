package accessorganization

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/cloudflare/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/cloudflare/accessorganization/internal"
)

type AccessOrganizationLoginDesignOutputReference interface {
	cdktf.ComplexObject
	BackgroundColor() *string
	SetBackgroundColor(val *string)
	BackgroundColorInput() *string
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
	FooterText() *string
	SetFooterText(val *string)
	FooterTextInput() *string
	// Experimental.
	Fqn() *string
	HeaderText() *string
	SetHeaderText(val *string)
	HeaderTextInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LogoPath() *string
	SetLogoPath(val *string)
	LogoPathInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TextColor() *string
	SetTextColor(val *string)
	TextColorInput() *string
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
	ResetBackgroundColor()
	ResetFooterText()
	ResetHeaderText()
	ResetLogoPath()
	ResetTextColor()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AccessOrganizationLoginDesignOutputReference
type jsiiProxy_AccessOrganizationLoginDesignOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AccessOrganizationLoginDesignOutputReference) BackgroundColor() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backgroundColor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessOrganizationLoginDesignOutputReference) BackgroundColorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backgroundColorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessOrganizationLoginDesignOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessOrganizationLoginDesignOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessOrganizationLoginDesignOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessOrganizationLoginDesignOutputReference) FooterText() *string {
	var returns *string
	_jsii_.Get(
		j,
		"footerText",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessOrganizationLoginDesignOutputReference) FooterTextInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"footerTextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessOrganizationLoginDesignOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessOrganizationLoginDesignOutputReference) HeaderText() *string {
	var returns *string
	_jsii_.Get(
		j,
		"headerText",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessOrganizationLoginDesignOutputReference) HeaderTextInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"headerTextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessOrganizationLoginDesignOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessOrganizationLoginDesignOutputReference) LogoPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logoPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessOrganizationLoginDesignOutputReference) LogoPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logoPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessOrganizationLoginDesignOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessOrganizationLoginDesignOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessOrganizationLoginDesignOutputReference) TextColor() *string {
	var returns *string
	_jsii_.Get(
		j,
		"textColor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessOrganizationLoginDesignOutputReference) TextColorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"textColorInput",
		&returns,
	)
	return returns
}


func NewAccessOrganizationLoginDesignOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AccessOrganizationLoginDesignOutputReference {
	_init_.Initialize()

	if err := validateNewAccessOrganizationLoginDesignOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AccessOrganizationLoginDesignOutputReference{}

	_jsii_.Create(
		"cloudflare.accessOrganization.AccessOrganizationLoginDesignOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewAccessOrganizationLoginDesignOutputReference_Override(a AccessOrganizationLoginDesignOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"cloudflare.accessOrganization.AccessOrganizationLoginDesignOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AccessOrganizationLoginDesignOutputReference)SetBackgroundColor(val *string) {
	if err := j.validateSetBackgroundColorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backgroundColor",
		val,
	)
}

func (j *jsiiProxy_AccessOrganizationLoginDesignOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AccessOrganizationLoginDesignOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AccessOrganizationLoginDesignOutputReference)SetFooterText(val *string) {
	if err := j.validateSetFooterTextParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"footerText",
		val,
	)
}

func (j *jsiiProxy_AccessOrganizationLoginDesignOutputReference)SetHeaderText(val *string) {
	if err := j.validateSetHeaderTextParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"headerText",
		val,
	)
}

func (j *jsiiProxy_AccessOrganizationLoginDesignOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AccessOrganizationLoginDesignOutputReference)SetLogoPath(val *string) {
	if err := j.validateSetLogoPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logoPath",
		val,
	)
}

func (j *jsiiProxy_AccessOrganizationLoginDesignOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AccessOrganizationLoginDesignOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AccessOrganizationLoginDesignOutputReference)SetTextColor(val *string) {
	if err := j.validateSetTextColorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"textColor",
		val,
	)
}

func (a *jsiiProxy_AccessOrganizationLoginDesignOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessOrganizationLoginDesignOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessOrganizationLoginDesignOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessOrganizationLoginDesignOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessOrganizationLoginDesignOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessOrganizationLoginDesignOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessOrganizationLoginDesignOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessOrganizationLoginDesignOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessOrganizationLoginDesignOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessOrganizationLoginDesignOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessOrganizationLoginDesignOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessOrganizationLoginDesignOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessOrganizationLoginDesignOutputReference) ResetBackgroundColor() {
	_jsii_.InvokeVoid(
		a,
		"resetBackgroundColor",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessOrganizationLoginDesignOutputReference) ResetFooterText() {
	_jsii_.InvokeVoid(
		a,
		"resetFooterText",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessOrganizationLoginDesignOutputReference) ResetHeaderText() {
	_jsii_.InvokeVoid(
		a,
		"resetHeaderText",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessOrganizationLoginDesignOutputReference) ResetLogoPath() {
	_jsii_.InvokeVoid(
		a,
		"resetLogoPath",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessOrganizationLoginDesignOutputReference) ResetTextColor() {
	_jsii_.InvokeVoid(
		a,
		"resetTextColor",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessOrganizationLoginDesignOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := a.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessOrganizationLoginDesignOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

