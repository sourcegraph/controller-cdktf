package vmwareenginecluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google/vmwareenginecluster/internal"
)

type VmwareengineClusterDatastoreMountConfigOutputReference interface {
	cdktn.ComplexObject
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
	DatastoreNetwork() VmwareengineClusterDatastoreMountConfigDatastoreNetworkOutputReference
	DatastoreNetworkInput() *VmwareengineClusterDatastoreMountConfigDatastoreNetwork
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
	PutDatastoreNetwork(value *VmwareengineClusterDatastoreMountConfigDatastoreNetwork)
	ResetAccessMode()
	ResetIgnoreColocation()
	ResetNfsVersion()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for VmwareengineClusterDatastoreMountConfigOutputReference
type jsiiProxy_VmwareengineClusterDatastoreMountConfigOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_VmwareengineClusterDatastoreMountConfigOutputReference) AccessMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwareengineClusterDatastoreMountConfigOutputReference) AccessModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwareengineClusterDatastoreMountConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwareengineClusterDatastoreMountConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwareengineClusterDatastoreMountConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwareengineClusterDatastoreMountConfigOutputReference) Datastore() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datastore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwareengineClusterDatastoreMountConfigOutputReference) DatastoreInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datastoreInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwareengineClusterDatastoreMountConfigOutputReference) DatastoreNetwork() VmwareengineClusterDatastoreMountConfigDatastoreNetworkOutputReference {
	var returns VmwareengineClusterDatastoreMountConfigDatastoreNetworkOutputReference
	_jsii_.Get(
		j,
		"datastoreNetwork",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwareengineClusterDatastoreMountConfigOutputReference) DatastoreNetworkInput() *VmwareengineClusterDatastoreMountConfigDatastoreNetwork {
	var returns *VmwareengineClusterDatastoreMountConfigDatastoreNetwork
	_jsii_.Get(
		j,
		"datastoreNetworkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwareengineClusterDatastoreMountConfigOutputReference) FileShare() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileShare",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwareengineClusterDatastoreMountConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwareengineClusterDatastoreMountConfigOutputReference) IgnoreColocation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreColocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwareengineClusterDatastoreMountConfigOutputReference) IgnoreColocationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreColocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwareengineClusterDatastoreMountConfigOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwareengineClusterDatastoreMountConfigOutputReference) NfsVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nfsVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwareengineClusterDatastoreMountConfigOutputReference) NfsVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nfsVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwareengineClusterDatastoreMountConfigOutputReference) Servers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"servers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwareengineClusterDatastoreMountConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwareengineClusterDatastoreMountConfigOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewVmwareengineClusterDatastoreMountConfigOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) VmwareengineClusterDatastoreMountConfigOutputReference {
	_init_.Initialize()

	if err := validateNewVmwareengineClusterDatastoreMountConfigOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_VmwareengineClusterDatastoreMountConfigOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.vmwareengineCluster.VmwareengineClusterDatastoreMountConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewVmwareengineClusterDatastoreMountConfigOutputReference_Override(v VmwareengineClusterDatastoreMountConfigOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.vmwareengineCluster.VmwareengineClusterDatastoreMountConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		v,
	)
}

func (j *jsiiProxy_VmwareengineClusterDatastoreMountConfigOutputReference)SetAccessMode(val *string) {
	if err := j.validateSetAccessModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accessMode",
		val,
	)
}

func (j *jsiiProxy_VmwareengineClusterDatastoreMountConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_VmwareengineClusterDatastoreMountConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_VmwareengineClusterDatastoreMountConfigOutputReference)SetDatastore(val *string) {
	if err := j.validateSetDatastoreParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"datastore",
		val,
	)
}

func (j *jsiiProxy_VmwareengineClusterDatastoreMountConfigOutputReference)SetIgnoreColocation(val interface{}) {
	if err := j.validateSetIgnoreColocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ignoreColocation",
		val,
	)
}

func (j *jsiiProxy_VmwareengineClusterDatastoreMountConfigOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VmwareengineClusterDatastoreMountConfigOutputReference)SetNfsVersion(val *string) {
	if err := j.validateSetNfsVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nfsVersion",
		val,
	)
}

func (j *jsiiProxy_VmwareengineClusterDatastoreMountConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VmwareengineClusterDatastoreMountConfigOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (v *jsiiProxy_VmwareengineClusterDatastoreMountConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VmwareengineClusterDatastoreMountConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := v.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VmwareengineClusterDatastoreMountConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := v.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		v,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VmwareengineClusterDatastoreMountConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := v.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		v,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VmwareengineClusterDatastoreMountConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := v.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		v,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VmwareengineClusterDatastoreMountConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := v.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		v,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VmwareengineClusterDatastoreMountConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := v.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		v,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VmwareengineClusterDatastoreMountConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := v.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		v,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VmwareengineClusterDatastoreMountConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := v.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		v,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VmwareengineClusterDatastoreMountConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := v.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		v,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VmwareengineClusterDatastoreMountConfigOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VmwareengineClusterDatastoreMountConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := v.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VmwareengineClusterDatastoreMountConfigOutputReference) PutDatastoreNetwork(value *VmwareengineClusterDatastoreMountConfigDatastoreNetwork) {
	if err := v.validatePutDatastoreNetworkParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putDatastoreNetwork",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VmwareengineClusterDatastoreMountConfigOutputReference) ResetAccessMode() {
	_jsii_.InvokeVoid(
		v,
		"resetAccessMode",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VmwareengineClusterDatastoreMountConfigOutputReference) ResetIgnoreColocation() {
	_jsii_.InvokeVoid(
		v,
		"resetIgnoreColocation",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VmwareengineClusterDatastoreMountConfigOutputReference) ResetNfsVersion() {
	_jsii_.InvokeVoid(
		v,
		"resetNfsVersion",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VmwareengineClusterDatastoreMountConfigOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := v.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		v,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VmwareengineClusterDatastoreMountConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

