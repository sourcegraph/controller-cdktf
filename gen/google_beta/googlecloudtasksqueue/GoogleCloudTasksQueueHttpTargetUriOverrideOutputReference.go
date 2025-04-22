package googlecloudtasksqueue

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecloudtasksqueue/internal"
)

type GoogleCloudTasksQueueHttpTargetUriOverrideOutputReference interface {
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
	// Experimental.
	Fqn() *string
	Host() *string
	SetHost(val *string)
	HostInput() *string
	InternalValue() *GoogleCloudTasksQueueHttpTargetUriOverride
	SetInternalValue(val *GoogleCloudTasksQueueHttpTargetUriOverride)
	PathOverride() GoogleCloudTasksQueueHttpTargetUriOverridePathOverrideOutputReference
	PathOverrideInput() *GoogleCloudTasksQueueHttpTargetUriOverridePathOverride
	Port() *string
	SetPort(val *string)
	PortInput() *string
	QueryOverride() GoogleCloudTasksQueueHttpTargetUriOverrideQueryOverrideOutputReference
	QueryOverrideInput() *GoogleCloudTasksQueueHttpTargetUriOverrideQueryOverride
	Scheme() *string
	SetScheme(val *string)
	SchemeInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UriOverrideEnforceMode() *string
	SetUriOverrideEnforceMode(val *string)
	UriOverrideEnforceModeInput() *string
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
	PutPathOverride(value *GoogleCloudTasksQueueHttpTargetUriOverridePathOverride)
	PutQueryOverride(value *GoogleCloudTasksQueueHttpTargetUriOverrideQueryOverride)
	ResetHost()
	ResetPathOverride()
	ResetPort()
	ResetQueryOverride()
	ResetScheme()
	ResetUriOverrideEnforceMode()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleCloudTasksQueueHttpTargetUriOverrideOutputReference
type jsiiProxy_GoogleCloudTasksQueueHttpTargetUriOverrideOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleCloudTasksQueueHttpTargetUriOverrideOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudTasksQueueHttpTargetUriOverrideOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudTasksQueueHttpTargetUriOverrideOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudTasksQueueHttpTargetUriOverrideOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudTasksQueueHttpTargetUriOverrideOutputReference) Host() *string {
	var returns *string
	_jsii_.Get(
		j,
		"host",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudTasksQueueHttpTargetUriOverrideOutputReference) HostInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudTasksQueueHttpTargetUriOverrideOutputReference) InternalValue() *GoogleCloudTasksQueueHttpTargetUriOverride {
	var returns *GoogleCloudTasksQueueHttpTargetUriOverride
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudTasksQueueHttpTargetUriOverrideOutputReference) PathOverride() GoogleCloudTasksQueueHttpTargetUriOverridePathOverrideOutputReference {
	var returns GoogleCloudTasksQueueHttpTargetUriOverridePathOverrideOutputReference
	_jsii_.Get(
		j,
		"pathOverride",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudTasksQueueHttpTargetUriOverrideOutputReference) PathOverrideInput() *GoogleCloudTasksQueueHttpTargetUriOverridePathOverride {
	var returns *GoogleCloudTasksQueueHttpTargetUriOverridePathOverride
	_jsii_.Get(
		j,
		"pathOverrideInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudTasksQueueHttpTargetUriOverrideOutputReference) Port() *string {
	var returns *string
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudTasksQueueHttpTargetUriOverrideOutputReference) PortInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"portInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudTasksQueueHttpTargetUriOverrideOutputReference) QueryOverride() GoogleCloudTasksQueueHttpTargetUriOverrideQueryOverrideOutputReference {
	var returns GoogleCloudTasksQueueHttpTargetUriOverrideQueryOverrideOutputReference
	_jsii_.Get(
		j,
		"queryOverride",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudTasksQueueHttpTargetUriOverrideOutputReference) QueryOverrideInput() *GoogleCloudTasksQueueHttpTargetUriOverrideQueryOverride {
	var returns *GoogleCloudTasksQueueHttpTargetUriOverrideQueryOverride
	_jsii_.Get(
		j,
		"queryOverrideInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudTasksQueueHttpTargetUriOverrideOutputReference) Scheme() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scheme",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudTasksQueueHttpTargetUriOverrideOutputReference) SchemeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schemeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudTasksQueueHttpTargetUriOverrideOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudTasksQueueHttpTargetUriOverrideOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudTasksQueueHttpTargetUriOverrideOutputReference) UriOverrideEnforceMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uriOverrideEnforceMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudTasksQueueHttpTargetUriOverrideOutputReference) UriOverrideEnforceModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uriOverrideEnforceModeInput",
		&returns,
	)
	return returns
}


func NewGoogleCloudTasksQueueHttpTargetUriOverrideOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleCloudTasksQueueHttpTargetUriOverrideOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleCloudTasksQueueHttpTargetUriOverrideOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleCloudTasksQueueHttpTargetUriOverrideOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleCloudTasksQueue.GoogleCloudTasksQueueHttpTargetUriOverrideOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleCloudTasksQueueHttpTargetUriOverrideOutputReference_Override(g GoogleCloudTasksQueueHttpTargetUriOverrideOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleCloudTasksQueue.GoogleCloudTasksQueueHttpTargetUriOverrideOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleCloudTasksQueueHttpTargetUriOverrideOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudTasksQueueHttpTargetUriOverrideOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudTasksQueueHttpTargetUriOverrideOutputReference)SetHost(val *string) {
	if err := j.validateSetHostParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"host",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudTasksQueueHttpTargetUriOverrideOutputReference)SetInternalValue(val *GoogleCloudTasksQueueHttpTargetUriOverride) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudTasksQueueHttpTargetUriOverrideOutputReference)SetPort(val *string) {
	if err := j.validateSetPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudTasksQueueHttpTargetUriOverrideOutputReference)SetScheme(val *string) {
	if err := j.validateSetSchemeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scheme",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudTasksQueueHttpTargetUriOverrideOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudTasksQueueHttpTargetUriOverrideOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudTasksQueueHttpTargetUriOverrideOutputReference)SetUriOverrideEnforceMode(val *string) {
	if err := j.validateSetUriOverrideEnforceModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"uriOverrideEnforceMode",
		val,
	)
}

func (g *jsiiProxy_GoogleCloudTasksQueueHttpTargetUriOverrideOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudTasksQueueHttpTargetUriOverrideOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleCloudTasksQueueHttpTargetUriOverrideOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleCloudTasksQueueHttpTargetUriOverrideOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleCloudTasksQueueHttpTargetUriOverrideOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleCloudTasksQueueHttpTargetUriOverrideOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleCloudTasksQueueHttpTargetUriOverrideOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleCloudTasksQueueHttpTargetUriOverrideOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleCloudTasksQueueHttpTargetUriOverrideOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleCloudTasksQueueHttpTargetUriOverrideOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleCloudTasksQueueHttpTargetUriOverrideOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudTasksQueueHttpTargetUriOverrideOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleCloudTasksQueueHttpTargetUriOverrideOutputReference) PutPathOverride(value *GoogleCloudTasksQueueHttpTargetUriOverridePathOverride) {
	if err := g.validatePutPathOverrideParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPathOverride",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCloudTasksQueueHttpTargetUriOverrideOutputReference) PutQueryOverride(value *GoogleCloudTasksQueueHttpTargetUriOverrideQueryOverride) {
	if err := g.validatePutQueryOverrideParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putQueryOverride",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCloudTasksQueueHttpTargetUriOverrideOutputReference) ResetHost() {
	_jsii_.InvokeVoid(
		g,
		"resetHost",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudTasksQueueHttpTargetUriOverrideOutputReference) ResetPathOverride() {
	_jsii_.InvokeVoid(
		g,
		"resetPathOverride",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudTasksQueueHttpTargetUriOverrideOutputReference) ResetPort() {
	_jsii_.InvokeVoid(
		g,
		"resetPort",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudTasksQueueHttpTargetUriOverrideOutputReference) ResetQueryOverride() {
	_jsii_.InvokeVoid(
		g,
		"resetQueryOverride",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudTasksQueueHttpTargetUriOverrideOutputReference) ResetScheme() {
	_jsii_.InvokeVoid(
		g,
		"resetScheme",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudTasksQueueHttpTargetUriOverrideOutputReference) ResetUriOverrideEnforceMode() {
	_jsii_.InvokeVoid(
		g,
		"resetUriOverrideEnforceMode",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudTasksQueueHttpTargetUriOverrideOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleCloudTasksQueueHttpTargetUriOverrideOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

