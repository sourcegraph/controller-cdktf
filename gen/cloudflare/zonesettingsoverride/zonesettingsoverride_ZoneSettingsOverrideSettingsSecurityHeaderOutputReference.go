package zonesettingsoverride

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/cloudflare/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/cloudflare/zonesettingsoverride/internal"
)

type ZoneSettingsOverrideSettingsSecurityHeaderOutputReference interface {
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
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	// Experimental.
	Fqn() *string
	IncludeSubdomains() interface{}
	SetIncludeSubdomains(val interface{})
	IncludeSubdomainsInput() interface{}
	InternalValue() *ZoneSettingsOverrideSettingsSecurityHeader
	SetInternalValue(val *ZoneSettingsOverrideSettingsSecurityHeader)
	MaxAge() *float64
	SetMaxAge(val *float64)
	MaxAgeInput() *float64
	Nosniff() interface{}
	SetNosniff(val interface{})
	NosniffInput() interface{}
	Preload() interface{}
	SetPreload(val interface{})
	PreloadInput() interface{}
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
	ResetEnabled()
	ResetIncludeSubdomains()
	ResetMaxAge()
	ResetNosniff()
	ResetPreload()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ZoneSettingsOverrideSettingsSecurityHeaderOutputReference
type jsiiProxy_ZoneSettingsOverrideSettingsSecurityHeaderOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsSecurityHeaderOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsSecurityHeaderOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsSecurityHeaderOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsSecurityHeaderOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsSecurityHeaderOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsSecurityHeaderOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsSecurityHeaderOutputReference) IncludeSubdomains() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeSubdomains",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsSecurityHeaderOutputReference) IncludeSubdomainsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeSubdomainsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsSecurityHeaderOutputReference) InternalValue() *ZoneSettingsOverrideSettingsSecurityHeader {
	var returns *ZoneSettingsOverrideSettingsSecurityHeader
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsSecurityHeaderOutputReference) MaxAge() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxAge",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsSecurityHeaderOutputReference) MaxAgeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxAgeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsSecurityHeaderOutputReference) Nosniff() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"nosniff",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsSecurityHeaderOutputReference) NosniffInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"nosniffInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsSecurityHeaderOutputReference) Preload() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preload",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsSecurityHeaderOutputReference) PreloadInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preloadInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsSecurityHeaderOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsSecurityHeaderOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewZoneSettingsOverrideSettingsSecurityHeaderOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ZoneSettingsOverrideSettingsSecurityHeaderOutputReference {
	_init_.Initialize()

	if err := validateNewZoneSettingsOverrideSettingsSecurityHeaderOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ZoneSettingsOverrideSettingsSecurityHeaderOutputReference{}

	_jsii_.Create(
		"cloudflare.zoneSettingsOverride.ZoneSettingsOverrideSettingsSecurityHeaderOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewZoneSettingsOverrideSettingsSecurityHeaderOutputReference_Override(z ZoneSettingsOverrideSettingsSecurityHeaderOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"cloudflare.zoneSettingsOverride.ZoneSettingsOverrideSettingsSecurityHeaderOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		z,
	)
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsSecurityHeaderOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsSecurityHeaderOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsSecurityHeaderOutputReference)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsSecurityHeaderOutputReference)SetIncludeSubdomains(val interface{}) {
	if err := j.validateSetIncludeSubdomainsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeSubdomains",
		val,
	)
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsSecurityHeaderOutputReference)SetInternalValue(val *ZoneSettingsOverrideSettingsSecurityHeader) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsSecurityHeaderOutputReference)SetMaxAge(val *float64) {
	if err := j.validateSetMaxAgeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxAge",
		val,
	)
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsSecurityHeaderOutputReference)SetNosniff(val interface{}) {
	if err := j.validateSetNosniffParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nosniff",
		val,
	)
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsSecurityHeaderOutputReference)SetPreload(val interface{}) {
	if err := j.validateSetPreloadParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preload",
		val,
	)
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsSecurityHeaderOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ZoneSettingsOverrideSettingsSecurityHeaderOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (z *jsiiProxy_ZoneSettingsOverrideSettingsSecurityHeaderOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		z,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZoneSettingsOverrideSettingsSecurityHeaderOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := z.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		z,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZoneSettingsOverrideSettingsSecurityHeaderOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := z.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		z,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZoneSettingsOverrideSettingsSecurityHeaderOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := z.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		z,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZoneSettingsOverrideSettingsSecurityHeaderOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := z.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		z,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZoneSettingsOverrideSettingsSecurityHeaderOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := z.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		z,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZoneSettingsOverrideSettingsSecurityHeaderOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := z.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		z,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZoneSettingsOverrideSettingsSecurityHeaderOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := z.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		z,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZoneSettingsOverrideSettingsSecurityHeaderOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := z.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		z,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZoneSettingsOverrideSettingsSecurityHeaderOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := z.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		z,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZoneSettingsOverrideSettingsSecurityHeaderOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		z,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZoneSettingsOverrideSettingsSecurityHeaderOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := z.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		z,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZoneSettingsOverrideSettingsSecurityHeaderOutputReference) ResetEnabled() {
	_jsii_.InvokeVoid(
		z,
		"resetEnabled",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZoneSettingsOverrideSettingsSecurityHeaderOutputReference) ResetIncludeSubdomains() {
	_jsii_.InvokeVoid(
		z,
		"resetIncludeSubdomains",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZoneSettingsOverrideSettingsSecurityHeaderOutputReference) ResetMaxAge() {
	_jsii_.InvokeVoid(
		z,
		"resetMaxAge",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZoneSettingsOverrideSettingsSecurityHeaderOutputReference) ResetNosniff() {
	_jsii_.InvokeVoid(
		z,
		"resetNosniff",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZoneSettingsOverrideSettingsSecurityHeaderOutputReference) ResetPreload() {
	_jsii_.InvokeVoid(
		z,
		"resetPreload",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZoneSettingsOverrideSettingsSecurityHeaderOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := z.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		z,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZoneSettingsOverrideSettingsSecurityHeaderOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		z,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

