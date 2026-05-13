package googlevmwareenginecluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlevmwareenginecluster/internal"
)

type GoogleVmwareengineClusterDatastoreMountConfigOutputReference interface {
	cdktf.ComplexObject
	AccessMode() *string
	SetAccessMode(val *string)
	AccessModeInput() *string
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
	Datastore() *string
	SetDatastore(val *string)
	DatastoreInput() *string
	DatastoreNetwork() GoogleVmwareengineClusterDatastoreMountConfigDatastoreNetworkOutputReference
	DatastoreNetworkInput() *GoogleVmwareengineClusterDatastoreMountConfigDatastoreNetwork
	FileShare() *string
	// Experimental.
	Fqn() *string
	IgnoreColocation() interface{}
	SetIgnoreColocation(val interface{})
	IgnoreColocationInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	NfsVersion() *string
	SetNfsVersion(val *string)
	NfsVersionInput() *string
	Servers() *[]*string
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
	PutDatastoreNetwork(value *GoogleVmwareengineClusterDatastoreMountConfigDatastoreNetwork)
	ResetAccessMode()
	ResetIgnoreColocation()
	ResetNfsVersion()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleVmwareengineClusterDatastoreMountConfigOutputReference
type jsiiProxy_GoogleVmwareengineClusterDatastoreMountConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleVmwareengineClusterDatastoreMountConfigOutputReference) AccessMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareengineClusterDatastoreMountConfigOutputReference) AccessModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareengineClusterDatastoreMountConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareengineClusterDatastoreMountConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareengineClusterDatastoreMountConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareengineClusterDatastoreMountConfigOutputReference) Datastore() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datastore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareengineClusterDatastoreMountConfigOutputReference) DatastoreInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datastoreInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareengineClusterDatastoreMountConfigOutputReference) DatastoreNetwork() GoogleVmwareengineClusterDatastoreMountConfigDatastoreNetworkOutputReference {
	var returns GoogleVmwareengineClusterDatastoreMountConfigDatastoreNetworkOutputReference
	_jsii_.Get(
		j,
		"datastoreNetwork",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareengineClusterDatastoreMountConfigOutputReference) DatastoreNetworkInput() *GoogleVmwareengineClusterDatastoreMountConfigDatastoreNetwork {
	var returns *GoogleVmwareengineClusterDatastoreMountConfigDatastoreNetwork
	_jsii_.Get(
		j,
		"datastoreNetworkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareengineClusterDatastoreMountConfigOutputReference) FileShare() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileShare",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareengineClusterDatastoreMountConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareengineClusterDatastoreMountConfigOutputReference) IgnoreColocation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreColocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareengineClusterDatastoreMountConfigOutputReference) IgnoreColocationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreColocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareengineClusterDatastoreMountConfigOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareengineClusterDatastoreMountConfigOutputReference) NfsVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nfsVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareengineClusterDatastoreMountConfigOutputReference) NfsVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nfsVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareengineClusterDatastoreMountConfigOutputReference) Servers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"servers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareengineClusterDatastoreMountConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareengineClusterDatastoreMountConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleVmwareengineClusterDatastoreMountConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GoogleVmwareengineClusterDatastoreMountConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleVmwareengineClusterDatastoreMountConfigOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleVmwareengineClusterDatastoreMountConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleVmwareengineCluster.GoogleVmwareengineClusterDatastoreMountConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGoogleVmwareengineClusterDatastoreMountConfigOutputReference_Override(g GoogleVmwareengineClusterDatastoreMountConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleVmwareengineCluster.GoogleVmwareengineClusterDatastoreMountConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GoogleVmwareengineClusterDatastoreMountConfigOutputReference)SetAccessMode(val *string) {
	if err := j.validateSetAccessModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accessMode",
		val,
	)
}

func (j *jsiiProxy_GoogleVmwareengineClusterDatastoreMountConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleVmwareengineClusterDatastoreMountConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleVmwareengineClusterDatastoreMountConfigOutputReference)SetDatastore(val *string) {
	if err := j.validateSetDatastoreParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"datastore",
		val,
	)
}

func (j *jsiiProxy_GoogleVmwareengineClusterDatastoreMountConfigOutputReference)SetIgnoreColocation(val interface{}) {
	if err := j.validateSetIgnoreColocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ignoreColocation",
		val,
	)
}

func (j *jsiiProxy_GoogleVmwareengineClusterDatastoreMountConfigOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleVmwareengineClusterDatastoreMountConfigOutputReference)SetNfsVersion(val *string) {
	if err := j.validateSetNfsVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nfsVersion",
		val,
	)
}

func (j *jsiiProxy_GoogleVmwareengineClusterDatastoreMountConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleVmwareengineClusterDatastoreMountConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleVmwareengineClusterDatastoreMountConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVmwareengineClusterDatastoreMountConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleVmwareengineClusterDatastoreMountConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleVmwareengineClusterDatastoreMountConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleVmwareengineClusterDatastoreMountConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleVmwareengineClusterDatastoreMountConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleVmwareengineClusterDatastoreMountConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleVmwareengineClusterDatastoreMountConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleVmwareengineClusterDatastoreMountConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleVmwareengineClusterDatastoreMountConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleVmwareengineClusterDatastoreMountConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVmwareengineClusterDatastoreMountConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleVmwareengineClusterDatastoreMountConfigOutputReference) PutDatastoreNetwork(value *GoogleVmwareengineClusterDatastoreMountConfigDatastoreNetwork) {
	if err := g.validatePutDatastoreNetworkParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDatastoreNetwork",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleVmwareengineClusterDatastoreMountConfigOutputReference) ResetAccessMode() {
	_jsii_.InvokeVoid(
		g,
		"resetAccessMode",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVmwareengineClusterDatastoreMountConfigOutputReference) ResetIgnoreColocation() {
	_jsii_.InvokeVoid(
		g,
		"resetIgnoreColocation",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVmwareengineClusterDatastoreMountConfigOutputReference) ResetNfsVersion() {
	_jsii_.InvokeVoid(
		g,
		"resetNfsVersion",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVmwareengineClusterDatastoreMountConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleVmwareengineClusterDatastoreMountConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

