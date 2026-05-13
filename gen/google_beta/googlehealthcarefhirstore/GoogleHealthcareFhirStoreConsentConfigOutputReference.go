package googlehealthcarefhirstore

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlehealthcarefhirstore/internal"
)

type GoogleHealthcareFhirStoreConsentConfigOutputReference interface {
	cdktf.ComplexObject
	AccessDeterminationLogConfig() GoogleHealthcareFhirStoreConsentConfigAccessDeterminationLogConfigOutputReference
	AccessDeterminationLogConfigInput() *GoogleHealthcareFhirStoreConsentConfigAccessDeterminationLogConfig
	AccessEnforced() interface{}
	SetAccessEnforced(val interface{})
	AccessEnforcedInput() interface{}
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
	ConsentHeaderHandling() GoogleHealthcareFhirStoreConsentConfigConsentHeaderHandlingOutputReference
	ConsentHeaderHandlingInput() *GoogleHealthcareFhirStoreConsentConfigConsentHeaderHandling
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	EnforcedAdminConsents() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleHealthcareFhirStoreConsentConfig
	SetInternalValue(val *GoogleHealthcareFhirStoreConsentConfig)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Version() *string
	SetVersion(val *string)
	VersionInput() *string
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
	PutAccessDeterminationLogConfig(value *GoogleHealthcareFhirStoreConsentConfigAccessDeterminationLogConfig)
	PutConsentHeaderHandling(value *GoogleHealthcareFhirStoreConsentConfigConsentHeaderHandling)
	ResetAccessDeterminationLogConfig()
	ResetAccessEnforced()
	ResetConsentHeaderHandling()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleHealthcareFhirStoreConsentConfigOutputReference
type jsiiProxy_GoogleHealthcareFhirStoreConsentConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleHealthcareFhirStoreConsentConfigOutputReference) AccessDeterminationLogConfig() GoogleHealthcareFhirStoreConsentConfigAccessDeterminationLogConfigOutputReference {
	var returns GoogleHealthcareFhirStoreConsentConfigAccessDeterminationLogConfigOutputReference
	_jsii_.Get(
		j,
		"accessDeterminationLogConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHealthcareFhirStoreConsentConfigOutputReference) AccessDeterminationLogConfigInput() *GoogleHealthcareFhirStoreConsentConfigAccessDeterminationLogConfig {
	var returns *GoogleHealthcareFhirStoreConsentConfigAccessDeterminationLogConfig
	_jsii_.Get(
		j,
		"accessDeterminationLogConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHealthcareFhirStoreConsentConfigOutputReference) AccessEnforced() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accessEnforced",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHealthcareFhirStoreConsentConfigOutputReference) AccessEnforcedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accessEnforcedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHealthcareFhirStoreConsentConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHealthcareFhirStoreConsentConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHealthcareFhirStoreConsentConfigOutputReference) ConsentHeaderHandling() GoogleHealthcareFhirStoreConsentConfigConsentHeaderHandlingOutputReference {
	var returns GoogleHealthcareFhirStoreConsentConfigConsentHeaderHandlingOutputReference
	_jsii_.Get(
		j,
		"consentHeaderHandling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHealthcareFhirStoreConsentConfigOutputReference) ConsentHeaderHandlingInput() *GoogleHealthcareFhirStoreConsentConfigConsentHeaderHandling {
	var returns *GoogleHealthcareFhirStoreConsentConfigConsentHeaderHandling
	_jsii_.Get(
		j,
		"consentHeaderHandlingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHealthcareFhirStoreConsentConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHealthcareFhirStoreConsentConfigOutputReference) EnforcedAdminConsents() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"enforcedAdminConsents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHealthcareFhirStoreConsentConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHealthcareFhirStoreConsentConfigOutputReference) InternalValue() *GoogleHealthcareFhirStoreConsentConfig {
	var returns *GoogleHealthcareFhirStoreConsentConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHealthcareFhirStoreConsentConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHealthcareFhirStoreConsentConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHealthcareFhirStoreConsentConfigOutputReference) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHealthcareFhirStoreConsentConfigOutputReference) VersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionInput",
		&returns,
	)
	return returns
}


func NewGoogleHealthcareFhirStoreConsentConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleHealthcareFhirStoreConsentConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleHealthcareFhirStoreConsentConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleHealthcareFhirStoreConsentConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleHealthcareFhirStore.GoogleHealthcareFhirStoreConsentConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleHealthcareFhirStoreConsentConfigOutputReference_Override(g GoogleHealthcareFhirStoreConsentConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleHealthcareFhirStore.GoogleHealthcareFhirStoreConsentConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleHealthcareFhirStoreConsentConfigOutputReference)SetAccessEnforced(val interface{}) {
	if err := j.validateSetAccessEnforcedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accessEnforced",
		val,
	)
}

func (j *jsiiProxy_GoogleHealthcareFhirStoreConsentConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleHealthcareFhirStoreConsentConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleHealthcareFhirStoreConsentConfigOutputReference)SetInternalValue(val *GoogleHealthcareFhirStoreConsentConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleHealthcareFhirStoreConsentConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleHealthcareFhirStoreConsentConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleHealthcareFhirStoreConsentConfigOutputReference)SetVersion(val *string) {
	if err := j.validateSetVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"version",
		val,
	)
}

func (g *jsiiProxy_GoogleHealthcareFhirStoreConsentConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleHealthcareFhirStoreConsentConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleHealthcareFhirStoreConsentConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleHealthcareFhirStoreConsentConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleHealthcareFhirStoreConsentConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleHealthcareFhirStoreConsentConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleHealthcareFhirStoreConsentConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleHealthcareFhirStoreConsentConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleHealthcareFhirStoreConsentConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleHealthcareFhirStoreConsentConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleHealthcareFhirStoreConsentConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleHealthcareFhirStoreConsentConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleHealthcareFhirStoreConsentConfigOutputReference) PutAccessDeterminationLogConfig(value *GoogleHealthcareFhirStoreConsentConfigAccessDeterminationLogConfig) {
	if err := g.validatePutAccessDeterminationLogConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAccessDeterminationLogConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleHealthcareFhirStoreConsentConfigOutputReference) PutConsentHeaderHandling(value *GoogleHealthcareFhirStoreConsentConfigConsentHeaderHandling) {
	if err := g.validatePutConsentHeaderHandlingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putConsentHeaderHandling",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleHealthcareFhirStoreConsentConfigOutputReference) ResetAccessDeterminationLogConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetAccessDeterminationLogConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleHealthcareFhirStoreConsentConfigOutputReference) ResetAccessEnforced() {
	_jsii_.InvokeVoid(
		g,
		"resetAccessEnforced",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleHealthcareFhirStoreConsentConfigOutputReference) ResetConsentHeaderHandling() {
	_jsii_.InvokeVoid(
		g,
		"resetConsentHeaderHandling",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleHealthcareFhirStoreConsentConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleHealthcareFhirStoreConsentConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

