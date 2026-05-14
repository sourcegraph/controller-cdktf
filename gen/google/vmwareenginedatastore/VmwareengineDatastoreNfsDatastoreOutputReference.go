package vmwareenginedatastore

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/vmwareenginedatastore/internal"
)

type VmwareengineDatastoreNfsDatastoreOutputReference interface {
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
	GoogleFileService() VmwareengineDatastoreNfsDatastoreGoogleFileServiceOutputReference
	GoogleFileServiceInput() *VmwareengineDatastoreNfsDatastoreGoogleFileService
	InternalValue() *VmwareengineDatastoreNfsDatastore
	SetInternalValue(val *VmwareengineDatastoreNfsDatastore)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	ThirdPartyFileService() VmwareengineDatastoreNfsDatastoreThirdPartyFileServiceOutputReference
	ThirdPartyFileServiceInput() *VmwareengineDatastoreNfsDatastoreThirdPartyFileService
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
	PutGoogleFileService(value *VmwareengineDatastoreNfsDatastoreGoogleFileService)
	PutThirdPartyFileService(value *VmwareengineDatastoreNfsDatastoreThirdPartyFileService)
	ResetGoogleFileService()
	ResetThirdPartyFileService()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for VmwareengineDatastoreNfsDatastoreOutputReference
type jsiiProxy_VmwareengineDatastoreNfsDatastoreOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_VmwareengineDatastoreNfsDatastoreOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwareengineDatastoreNfsDatastoreOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwareengineDatastoreNfsDatastoreOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwareengineDatastoreNfsDatastoreOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwareengineDatastoreNfsDatastoreOutputReference) GoogleFileService() VmwareengineDatastoreNfsDatastoreGoogleFileServiceOutputReference {
	var returns VmwareengineDatastoreNfsDatastoreGoogleFileServiceOutputReference
	_jsii_.Get(
		j,
		"googleFileService",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwareengineDatastoreNfsDatastoreOutputReference) GoogleFileServiceInput() *VmwareengineDatastoreNfsDatastoreGoogleFileService {
	var returns *VmwareengineDatastoreNfsDatastoreGoogleFileService
	_jsii_.Get(
		j,
		"googleFileServiceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwareengineDatastoreNfsDatastoreOutputReference) InternalValue() *VmwareengineDatastoreNfsDatastore {
	var returns *VmwareengineDatastoreNfsDatastore
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwareengineDatastoreNfsDatastoreOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwareengineDatastoreNfsDatastoreOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwareengineDatastoreNfsDatastoreOutputReference) ThirdPartyFileService() VmwareengineDatastoreNfsDatastoreThirdPartyFileServiceOutputReference {
	var returns VmwareengineDatastoreNfsDatastoreThirdPartyFileServiceOutputReference
	_jsii_.Get(
		j,
		"thirdPartyFileService",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwareengineDatastoreNfsDatastoreOutputReference) ThirdPartyFileServiceInput() *VmwareengineDatastoreNfsDatastoreThirdPartyFileService {
	var returns *VmwareengineDatastoreNfsDatastoreThirdPartyFileService
	_jsii_.Get(
		j,
		"thirdPartyFileServiceInput",
		&returns,
	)
	return returns
}


func NewVmwareengineDatastoreNfsDatastoreOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) VmwareengineDatastoreNfsDatastoreOutputReference {
	_init_.Initialize()

	if err := validateNewVmwareengineDatastoreNfsDatastoreOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_VmwareengineDatastoreNfsDatastoreOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.vmwareengineDatastore.VmwareengineDatastoreNfsDatastoreOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewVmwareengineDatastoreNfsDatastoreOutputReference_Override(v VmwareengineDatastoreNfsDatastoreOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.vmwareengineDatastore.VmwareengineDatastoreNfsDatastoreOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		v,
	)
}

func (j *jsiiProxy_VmwareengineDatastoreNfsDatastoreOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_VmwareengineDatastoreNfsDatastoreOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_VmwareengineDatastoreNfsDatastoreOutputReference)SetInternalValue(val *VmwareengineDatastoreNfsDatastore) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VmwareengineDatastoreNfsDatastoreOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VmwareengineDatastoreNfsDatastoreOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (v *jsiiProxy_VmwareengineDatastoreNfsDatastoreOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VmwareengineDatastoreNfsDatastoreOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (v *jsiiProxy_VmwareengineDatastoreNfsDatastoreOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := v.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VmwareengineDatastoreNfsDatastoreOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (v *jsiiProxy_VmwareengineDatastoreNfsDatastoreOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (v *jsiiProxy_VmwareengineDatastoreNfsDatastoreOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (v *jsiiProxy_VmwareengineDatastoreNfsDatastoreOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (v *jsiiProxy_VmwareengineDatastoreNfsDatastoreOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (v *jsiiProxy_VmwareengineDatastoreNfsDatastoreOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (v *jsiiProxy_VmwareengineDatastoreNfsDatastoreOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (v *jsiiProxy_VmwareengineDatastoreNfsDatastoreOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VmwareengineDatastoreNfsDatastoreOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := v.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VmwareengineDatastoreNfsDatastoreOutputReference) PutGoogleFileService(value *VmwareengineDatastoreNfsDatastoreGoogleFileService) {
	if err := v.validatePutGoogleFileServiceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putGoogleFileService",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VmwareengineDatastoreNfsDatastoreOutputReference) PutThirdPartyFileService(value *VmwareengineDatastoreNfsDatastoreThirdPartyFileService) {
	if err := v.validatePutThirdPartyFileServiceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putThirdPartyFileService",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VmwareengineDatastoreNfsDatastoreOutputReference) ResetGoogleFileService() {
	_jsii_.InvokeVoid(
		v,
		"resetGoogleFileService",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VmwareengineDatastoreNfsDatastoreOutputReference) ResetThirdPartyFileService() {
	_jsii_.InvokeVoid(
		v,
		"resetThirdPartyFileService",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VmwareengineDatastoreNfsDatastoreOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := v.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		v,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VmwareengineDatastoreNfsDatastoreOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

