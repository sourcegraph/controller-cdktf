package googlecomputeregionbackendservice

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecomputeregionbackendservice/internal"
)

type GoogleComputeRegionBackendServiceTlsSettingsOutputReference interface {
	cdktn.ComplexObject
	AuthenticationConfig() *string
	SetAuthenticationConfig(val *string)
	AuthenticationConfigInput() *string
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
	InternalValue() *GoogleComputeRegionBackendServiceTlsSettings
	SetInternalValue(val *GoogleComputeRegionBackendServiceTlsSettings)
	Sni() *string
	SetSni(val *string)
	SniInput() *string
	SubjectAltNames() GoogleComputeRegionBackendServiceTlsSettingsSubjectAltNamesList
	SubjectAltNamesInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable
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
	InterpolationAsList() cdktn.IResolvable
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	PutSubjectAltNames(value interface{})
	ResetAuthenticationConfig()
	ResetSni()
	ResetSubjectAltNames()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleComputeRegionBackendServiceTlsSettingsOutputReference
type jsiiProxy_GoogleComputeRegionBackendServiceTlsSettingsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceTlsSettingsOutputReference) AuthenticationConfig() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticationConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceTlsSettingsOutputReference) AuthenticationConfigInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticationConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceTlsSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceTlsSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceTlsSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceTlsSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceTlsSettingsOutputReference) InternalValue() *GoogleComputeRegionBackendServiceTlsSettings {
	var returns *GoogleComputeRegionBackendServiceTlsSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceTlsSettingsOutputReference) Sni() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sni",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceTlsSettingsOutputReference) SniInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sniInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceTlsSettingsOutputReference) SubjectAltNames() GoogleComputeRegionBackendServiceTlsSettingsSubjectAltNamesList {
	var returns GoogleComputeRegionBackendServiceTlsSettingsSubjectAltNamesList
	_jsii_.Get(
		j,
		"subjectAltNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceTlsSettingsOutputReference) SubjectAltNamesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"subjectAltNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceTlsSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceTlsSettingsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleComputeRegionBackendServiceTlsSettingsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleComputeRegionBackendServiceTlsSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleComputeRegionBackendServiceTlsSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleComputeRegionBackendServiceTlsSettingsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleComputeRegionBackendService.GoogleComputeRegionBackendServiceTlsSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleComputeRegionBackendServiceTlsSettingsOutputReference_Override(g GoogleComputeRegionBackendServiceTlsSettingsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleComputeRegionBackendService.GoogleComputeRegionBackendServiceTlsSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceTlsSettingsOutputReference)SetAuthenticationConfig(val *string) {
	if err := j.validateSetAuthenticationConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authenticationConfig",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceTlsSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceTlsSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceTlsSettingsOutputReference)SetInternalValue(val *GoogleComputeRegionBackendServiceTlsSettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceTlsSettingsOutputReference)SetSni(val *string) {
	if err := j.validateSetSniParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sni",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceTlsSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionBackendServiceTlsSettingsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleComputeRegionBackendServiceTlsSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeRegionBackendServiceTlsSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleComputeRegionBackendServiceTlsSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeRegionBackendServiceTlsSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleComputeRegionBackendServiceTlsSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleComputeRegionBackendServiceTlsSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleComputeRegionBackendServiceTlsSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleComputeRegionBackendServiceTlsSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleComputeRegionBackendServiceTlsSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleComputeRegionBackendServiceTlsSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleComputeRegionBackendServiceTlsSettingsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeRegionBackendServiceTlsSettingsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeRegionBackendServiceTlsSettingsOutputReference) PutSubjectAltNames(value interface{}) {
	if err := g.validatePutSubjectAltNamesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSubjectAltNames",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeRegionBackendServiceTlsSettingsOutputReference) ResetAuthenticationConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetAuthenticationConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionBackendServiceTlsSettingsOutputReference) ResetSni() {
	_jsii_.InvokeVoid(
		g,
		"resetSni",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionBackendServiceTlsSettingsOutputReference) ResetSubjectAltNames() {
	_jsii_.InvokeVoid(
		g,
		"resetSubjectAltNames",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionBackendServiceTlsSettingsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := g.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeRegionBackendServiceTlsSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

