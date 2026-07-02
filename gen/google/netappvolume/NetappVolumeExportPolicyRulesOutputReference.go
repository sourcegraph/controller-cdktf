package netappvolume

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/netappvolume/internal"
)

type NetappVolumeExportPolicyRulesOutputReference interface {
	cdktf.ComplexObject
	AccessType() *string
	SetAccessType(val *string)
	AccessTypeInput() *string
	AllowedClients() *string
	SetAllowedClients(val *string)
	AllowedClientsInput() *string
	AnonUid() *float64
	SetAnonUid(val *float64)
	AnonUidInput() *float64
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() any
	// Experimental.
	SetComplexObjectIndex(val any)
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
	HasRootAccess() *string
	SetHasRootAccess(val *string)
	HasRootAccessInput() *string
	InternalValue() any
	SetInternalValue(val any)
	Kerberos5IReadOnly() any
	SetKerberos5IReadOnly(val any)
	Kerberos5IReadOnlyInput() any
	Kerberos5IReadWrite() any
	SetKerberos5IReadWrite(val any)
	Kerberos5IReadWriteInput() any
	Kerberos5PReadOnly() any
	SetKerberos5PReadOnly(val any)
	Kerberos5PReadOnlyInput() any
	Kerberos5PReadWrite() any
	SetKerberos5PReadWrite(val any)
	Kerberos5PReadWriteInput() any
	Kerberos5ReadOnly() any
	SetKerberos5ReadOnly(val any)
	Kerberos5ReadOnlyInput() any
	Kerberos5ReadWrite() any
	SetKerberos5ReadWrite(val any)
	Kerberos5ReadWriteInput() any
	Nfsv3() any
	SetNfsv3(val any)
	Nfsv3Input() any
	Nfsv4() any
	SetNfsv4(val any)
	Nfsv4Input() any
	SquashMode() *string
	SetSquashMode(val *string)
	SquashModeInput() *string
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
	GetAnyMapAttribute(terraformAttribute *string) *map[string]any
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
	ResetAccessType()
	ResetAllowedClients()
	ResetAnonUid()
	ResetHasRootAccess()
	ResetKerberos5IReadOnly()
	ResetKerberos5IReadWrite()
	ResetKerberos5PReadOnly()
	ResetKerberos5PReadWrite()
	ResetKerberos5ReadOnly()
	ResetKerberos5ReadWrite()
	ResetNfsv3()
	ResetNfsv4()
	ResetSquashMode()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for NetappVolumeExportPolicyRulesOutputReference
type jsiiProxy_NetappVolumeExportPolicyRulesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference) AccessType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference) AccessTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference) AllowedClients() *string {
	var returns *string
	_jsii_.Get(
		j,
		"allowedClients",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference) AllowedClientsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"allowedClientsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference) AnonUid() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"anonUid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference) AnonUidInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"anonUidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference) HasRootAccess() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hasRootAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference) HasRootAccessInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hasRootAccessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference) InternalValue() any {
	var returns any
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference) Kerberos5IReadOnly() any {
	var returns any
	_jsii_.Get(
		j,
		"kerberos5IReadOnly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference) Kerberos5IReadOnlyInput() any {
	var returns any
	_jsii_.Get(
		j,
		"kerberos5IReadOnlyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference) Kerberos5IReadWrite() any {
	var returns any
	_jsii_.Get(
		j,
		"kerberos5IReadWrite",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference) Kerberos5IReadWriteInput() any {
	var returns any
	_jsii_.Get(
		j,
		"kerberos5IReadWriteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference) Kerberos5PReadOnly() any {
	var returns any
	_jsii_.Get(
		j,
		"kerberos5PReadOnly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference) Kerberos5PReadOnlyInput() any {
	var returns any
	_jsii_.Get(
		j,
		"kerberos5PReadOnlyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference) Kerberos5PReadWrite() any {
	var returns any
	_jsii_.Get(
		j,
		"kerberos5PReadWrite",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference) Kerberos5PReadWriteInput() any {
	var returns any
	_jsii_.Get(
		j,
		"kerberos5PReadWriteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference) Kerberos5ReadOnly() any {
	var returns any
	_jsii_.Get(
		j,
		"kerberos5ReadOnly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference) Kerberos5ReadOnlyInput() any {
	var returns any
	_jsii_.Get(
		j,
		"kerberos5ReadOnlyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference) Kerberos5ReadWrite() any {
	var returns any
	_jsii_.Get(
		j,
		"kerberos5ReadWrite",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference) Kerberos5ReadWriteInput() any {
	var returns any
	_jsii_.Get(
		j,
		"kerberos5ReadWriteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference) Nfsv3() any {
	var returns any
	_jsii_.Get(
		j,
		"nfsv3",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference) Nfsv3Input() any {
	var returns any
	_jsii_.Get(
		j,
		"nfsv3Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference) Nfsv4() any {
	var returns any
	_jsii_.Get(
		j,
		"nfsv4",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference) Nfsv4Input() any {
	var returns any
	_jsii_.Get(
		j,
		"nfsv4Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference) SquashMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"squashMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference) SquashModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"squashModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewNetappVolumeExportPolicyRulesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) NetappVolumeExportPolicyRulesOutputReference {
	_init_.Initialize()

	if err := validateNewNetappVolumeExportPolicyRulesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_NetappVolumeExportPolicyRulesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.netappVolume.NetappVolumeExportPolicyRulesOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewNetappVolumeExportPolicyRulesOutputReference_Override(n NetappVolumeExportPolicyRulesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.netappVolume.NetappVolumeExportPolicyRulesOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		n,
	)
}

func (j *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference) SetAccessType(val *string) {
	if err := j.validateSetAccessTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accessType",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference) SetAllowedClients(val *string) {
	if err := j.validateSetAllowedClientsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedClients",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference) SetAnonUid(val *float64) {
	if err := j.validateSetAnonUidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"anonUid",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference) SetHasRootAccess(val *string) {
	if err := j.validateSetHasRootAccessParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hasRootAccess",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference) SetInternalValue(val any) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference) SetKerberos5IReadOnly(val any) {
	if err := j.validateSetKerberos5IReadOnlyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kerberos5IReadOnly",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference) SetKerberos5IReadWrite(val any) {
	if err := j.validateSetKerberos5IReadWriteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kerberos5IReadWrite",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference) SetKerberos5PReadOnly(val any) {
	if err := j.validateSetKerberos5PReadOnlyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kerberos5PReadOnly",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference) SetKerberos5PReadWrite(val any) {
	if err := j.validateSetKerberos5PReadWriteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kerberos5PReadWrite",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference) SetKerberos5ReadOnly(val any) {
	if err := j.validateSetKerberos5ReadOnlyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kerberos5ReadOnly",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference) SetKerberos5ReadWrite(val any) {
	if err := j.validateSetKerberos5ReadWriteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kerberos5ReadWrite",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference) SetNfsv3(val any) {
	if err := j.validateSetNfsv3Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nfsv3",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference) SetNfsv4(val any) {
	if err := j.validateSetNfsv4Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nfsv4",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference) SetSquashMode(val *string) {
	if err := j.validateSetSquashModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"squashMode",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (n *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := n.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		n,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := n.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := n.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		n,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := n.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		n,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := n.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		n,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := n.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		n,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := n.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		n,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := n.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		n,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := n.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		n,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := n.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationForAttribute",
		[]any{property},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference) ResetAccessType() {
	_jsii_.InvokeVoid(
		n,
		"resetAccessType",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference) ResetAllowedClients() {
	_jsii_.InvokeVoid(
		n,
		"resetAllowedClients",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference) ResetAnonUid() {
	_jsii_.InvokeVoid(
		n,
		"resetAnonUid",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference) ResetHasRootAccess() {
	_jsii_.InvokeVoid(
		n,
		"resetHasRootAccess",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference) ResetKerberos5IReadOnly() {
	_jsii_.InvokeVoid(
		n,
		"resetKerberos5IReadOnly",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference) ResetKerberos5IReadWrite() {
	_jsii_.InvokeVoid(
		n,
		"resetKerberos5IReadWrite",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference) ResetKerberos5PReadOnly() {
	_jsii_.InvokeVoid(
		n,
		"resetKerberos5PReadOnly",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference) ResetKerberos5PReadWrite() {
	_jsii_.InvokeVoid(
		n,
		"resetKerberos5PReadWrite",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference) ResetKerberos5ReadOnly() {
	_jsii_.InvokeVoid(
		n,
		"resetKerberos5ReadOnly",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference) ResetKerberos5ReadWrite() {
	_jsii_.InvokeVoid(
		n,
		"resetKerberos5ReadWrite",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference) ResetNfsv3() {
	_jsii_.InvokeVoid(
		n,
		"resetNfsv3",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference) ResetNfsv4() {
	_jsii_.InvokeVoid(
		n,
		"resetNfsv4",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference) ResetSquashMode() {
	_jsii_.InvokeVoid(
		n,
		"resetSquashMode",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference) Resolve(_context cdktf.IResolveContext) any {
	if err := n.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns any

	_jsii_.Invoke(
		n,
		"resolve",
		[]any{_context},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeExportPolicyRulesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
