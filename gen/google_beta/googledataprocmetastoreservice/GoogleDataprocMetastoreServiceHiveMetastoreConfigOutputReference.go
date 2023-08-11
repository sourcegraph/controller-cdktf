package googledataprocmetastoreservice

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googledataprocmetastoreservice/internal"
)

type GoogleDataprocMetastoreServiceHiveMetastoreConfigOutputReference interface {
	cdktf.ComplexObject
	AuxiliaryVersions() GoogleDataprocMetastoreServiceHiveMetastoreConfigAuxiliaryVersionsList
	AuxiliaryVersionsInput() interface{}
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
	ConfigOverrides() *map[string]*string
	SetConfigOverrides(val *map[string]*string)
	ConfigOverridesInput() *map[string]*string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	EndpointProtocol() *string
	SetEndpointProtocol(val *string)
	EndpointProtocolInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleDataprocMetastoreServiceHiveMetastoreConfig
	SetInternalValue(val *GoogleDataprocMetastoreServiceHiveMetastoreConfig)
	KerberosConfig() GoogleDataprocMetastoreServiceHiveMetastoreConfigKerberosConfigOutputReference
	KerberosConfigInput() *GoogleDataprocMetastoreServiceHiveMetastoreConfigKerberosConfig
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
	PutAuxiliaryVersions(value interface{})
	PutKerberosConfig(value *GoogleDataprocMetastoreServiceHiveMetastoreConfigKerberosConfig)
	ResetAuxiliaryVersions()
	ResetConfigOverrides()
	ResetEndpointProtocol()
	ResetKerberosConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleDataprocMetastoreServiceHiveMetastoreConfigOutputReference
type jsiiProxy_GoogleDataprocMetastoreServiceHiveMetastoreConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleDataprocMetastoreServiceHiveMetastoreConfigOutputReference) AuxiliaryVersions() GoogleDataprocMetastoreServiceHiveMetastoreConfigAuxiliaryVersionsList {
	var returns GoogleDataprocMetastoreServiceHiveMetastoreConfigAuxiliaryVersionsList
	_jsii_.Get(
		j,
		"auxiliaryVersions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocMetastoreServiceHiveMetastoreConfigOutputReference) AuxiliaryVersionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"auxiliaryVersionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocMetastoreServiceHiveMetastoreConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocMetastoreServiceHiveMetastoreConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocMetastoreServiceHiveMetastoreConfigOutputReference) ConfigOverrides() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"configOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocMetastoreServiceHiveMetastoreConfigOutputReference) ConfigOverridesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"configOverridesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocMetastoreServiceHiveMetastoreConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocMetastoreServiceHiveMetastoreConfigOutputReference) EndpointProtocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointProtocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocMetastoreServiceHiveMetastoreConfigOutputReference) EndpointProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointProtocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocMetastoreServiceHiveMetastoreConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocMetastoreServiceHiveMetastoreConfigOutputReference) InternalValue() *GoogleDataprocMetastoreServiceHiveMetastoreConfig {
	var returns *GoogleDataprocMetastoreServiceHiveMetastoreConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocMetastoreServiceHiveMetastoreConfigOutputReference) KerberosConfig() GoogleDataprocMetastoreServiceHiveMetastoreConfigKerberosConfigOutputReference {
	var returns GoogleDataprocMetastoreServiceHiveMetastoreConfigKerberosConfigOutputReference
	_jsii_.Get(
		j,
		"kerberosConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocMetastoreServiceHiveMetastoreConfigOutputReference) KerberosConfigInput() *GoogleDataprocMetastoreServiceHiveMetastoreConfigKerberosConfig {
	var returns *GoogleDataprocMetastoreServiceHiveMetastoreConfigKerberosConfig
	_jsii_.Get(
		j,
		"kerberosConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocMetastoreServiceHiveMetastoreConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocMetastoreServiceHiveMetastoreConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocMetastoreServiceHiveMetastoreConfigOutputReference) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocMetastoreServiceHiveMetastoreConfigOutputReference) VersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionInput",
		&returns,
	)
	return returns
}


func NewGoogleDataprocMetastoreServiceHiveMetastoreConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleDataprocMetastoreServiceHiveMetastoreConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDataprocMetastoreServiceHiveMetastoreConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDataprocMetastoreServiceHiveMetastoreConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDataprocMetastoreService.GoogleDataprocMetastoreServiceHiveMetastoreConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleDataprocMetastoreServiceHiveMetastoreConfigOutputReference_Override(g GoogleDataprocMetastoreServiceHiveMetastoreConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDataprocMetastoreService.GoogleDataprocMetastoreServiceHiveMetastoreConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleDataprocMetastoreServiceHiveMetastoreConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocMetastoreServiceHiveMetastoreConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocMetastoreServiceHiveMetastoreConfigOutputReference)SetConfigOverrides(val *map[string]*string) {
	if err := j.validateSetConfigOverridesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"configOverrides",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocMetastoreServiceHiveMetastoreConfigOutputReference)SetEndpointProtocol(val *string) {
	if err := j.validateSetEndpointProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endpointProtocol",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocMetastoreServiceHiveMetastoreConfigOutputReference)SetInternalValue(val *GoogleDataprocMetastoreServiceHiveMetastoreConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocMetastoreServiceHiveMetastoreConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocMetastoreServiceHiveMetastoreConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocMetastoreServiceHiveMetastoreConfigOutputReference)SetVersion(val *string) {
	if err := j.validateSetVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"version",
		val,
	)
}

func (g *jsiiProxy_GoogleDataprocMetastoreServiceHiveMetastoreConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataprocMetastoreServiceHiveMetastoreConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDataprocMetastoreServiceHiveMetastoreConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDataprocMetastoreServiceHiveMetastoreConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDataprocMetastoreServiceHiveMetastoreConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDataprocMetastoreServiceHiveMetastoreConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDataprocMetastoreServiceHiveMetastoreConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDataprocMetastoreServiceHiveMetastoreConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDataprocMetastoreServiceHiveMetastoreConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDataprocMetastoreServiceHiveMetastoreConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDataprocMetastoreServiceHiveMetastoreConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataprocMetastoreServiceHiveMetastoreConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDataprocMetastoreServiceHiveMetastoreConfigOutputReference) PutAuxiliaryVersions(value interface{}) {
	if err := g.validatePutAuxiliaryVersionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAuxiliaryVersions",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataprocMetastoreServiceHiveMetastoreConfigOutputReference) PutKerberosConfig(value *GoogleDataprocMetastoreServiceHiveMetastoreConfigKerberosConfig) {
	if err := g.validatePutKerberosConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putKerberosConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataprocMetastoreServiceHiveMetastoreConfigOutputReference) ResetAuxiliaryVersions() {
	_jsii_.InvokeVoid(
		g,
		"resetAuxiliaryVersions",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocMetastoreServiceHiveMetastoreConfigOutputReference) ResetConfigOverrides() {
	_jsii_.InvokeVoid(
		g,
		"resetConfigOverrides",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocMetastoreServiceHiveMetastoreConfigOutputReference) ResetEndpointProtocol() {
	_jsii_.InvokeVoid(
		g,
		"resetEndpointProtocol",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocMetastoreServiceHiveMetastoreConfigOutputReference) ResetKerberosConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetKerberosConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocMetastoreServiceHiveMetastoreConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleDataprocMetastoreServiceHiveMetastoreConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

