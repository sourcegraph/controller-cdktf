package googledatastreamstream

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googledatastreamstream/internal"
)

type GoogleDatastreamStreamDestinationConfigOutputReference interface {
	cdktf.ComplexObject
	BigqueryDestinationConfig() GoogleDatastreamStreamDestinationConfigBigqueryDestinationConfigOutputReference
	BigqueryDestinationConfigInput() *GoogleDatastreamStreamDestinationConfigBigqueryDestinationConfig
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
	DestinationConnectionProfile() *string
	SetDestinationConnectionProfile(val *string)
	DestinationConnectionProfileInput() *string
	// Experimental.
	Fqn() *string
	GcsDestinationConfig() GoogleDatastreamStreamDestinationConfigGcsDestinationConfigOutputReference
	GcsDestinationConfigInput() *GoogleDatastreamStreamDestinationConfigGcsDestinationConfig
	InternalValue() *GoogleDatastreamStreamDestinationConfig
	SetInternalValue(val *GoogleDatastreamStreamDestinationConfig)
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
	PutBigqueryDestinationConfig(value *GoogleDatastreamStreamDestinationConfigBigqueryDestinationConfig)
	PutGcsDestinationConfig(value *GoogleDatastreamStreamDestinationConfigGcsDestinationConfig)
	ResetBigqueryDestinationConfig()
	ResetGcsDestinationConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleDatastreamStreamDestinationConfigOutputReference
type jsiiProxy_GoogleDatastreamStreamDestinationConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleDatastreamStreamDestinationConfigOutputReference) BigqueryDestinationConfig() GoogleDatastreamStreamDestinationConfigBigqueryDestinationConfigOutputReference {
	var returns GoogleDatastreamStreamDestinationConfigBigqueryDestinationConfigOutputReference
	_jsii_.Get(
		j,
		"bigqueryDestinationConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamDestinationConfigOutputReference) BigqueryDestinationConfigInput() *GoogleDatastreamStreamDestinationConfigBigqueryDestinationConfig {
	var returns *GoogleDatastreamStreamDestinationConfigBigqueryDestinationConfig
	_jsii_.Get(
		j,
		"bigqueryDestinationConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamDestinationConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamDestinationConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamDestinationConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamDestinationConfigOutputReference) DestinationConnectionProfile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationConnectionProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamDestinationConfigOutputReference) DestinationConnectionProfileInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationConnectionProfileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamDestinationConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamDestinationConfigOutputReference) GcsDestinationConfig() GoogleDatastreamStreamDestinationConfigGcsDestinationConfigOutputReference {
	var returns GoogleDatastreamStreamDestinationConfigGcsDestinationConfigOutputReference
	_jsii_.Get(
		j,
		"gcsDestinationConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamDestinationConfigOutputReference) GcsDestinationConfigInput() *GoogleDatastreamStreamDestinationConfigGcsDestinationConfig {
	var returns *GoogleDatastreamStreamDestinationConfigGcsDestinationConfig
	_jsii_.Get(
		j,
		"gcsDestinationConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamDestinationConfigOutputReference) InternalValue() *GoogleDatastreamStreamDestinationConfig {
	var returns *GoogleDatastreamStreamDestinationConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamDestinationConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamDestinationConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleDatastreamStreamDestinationConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleDatastreamStreamDestinationConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDatastreamStreamDestinationConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDatastreamStreamDestinationConfigOutputReference{}

	_jsii_.Create(
		"google-beta.googleDatastreamStream.GoogleDatastreamStreamDestinationConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleDatastreamStreamDestinationConfigOutputReference_Override(g GoogleDatastreamStreamDestinationConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"google-beta.googleDatastreamStream.GoogleDatastreamStreamDestinationConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleDatastreamStreamDestinationConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDatastreamStreamDestinationConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDatastreamStreamDestinationConfigOutputReference)SetDestinationConnectionProfile(val *string) {
	if err := j.validateSetDestinationConnectionProfileParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationConnectionProfile",
		val,
	)
}

func (j *jsiiProxy_GoogleDatastreamStreamDestinationConfigOutputReference)SetInternalValue(val *GoogleDatastreamStreamDestinationConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDatastreamStreamDestinationConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDatastreamStreamDestinationConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleDatastreamStreamDestinationConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDatastreamStreamDestinationConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDatastreamStreamDestinationConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDatastreamStreamDestinationConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDatastreamStreamDestinationConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDatastreamStreamDestinationConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDatastreamStreamDestinationConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDatastreamStreamDestinationConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDatastreamStreamDestinationConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDatastreamStreamDestinationConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDatastreamStreamDestinationConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDatastreamStreamDestinationConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDatastreamStreamDestinationConfigOutputReference) PutBigqueryDestinationConfig(value *GoogleDatastreamStreamDestinationConfigBigqueryDestinationConfig) {
	if err := g.validatePutBigqueryDestinationConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putBigqueryDestinationConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDatastreamStreamDestinationConfigOutputReference) PutGcsDestinationConfig(value *GoogleDatastreamStreamDestinationConfigGcsDestinationConfig) {
	if err := g.validatePutGcsDestinationConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putGcsDestinationConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDatastreamStreamDestinationConfigOutputReference) ResetBigqueryDestinationConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetBigqueryDestinationConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDatastreamStreamDestinationConfigOutputReference) ResetGcsDestinationConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetGcsDestinationConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDatastreamStreamDestinationConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleDatastreamStreamDestinationConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

