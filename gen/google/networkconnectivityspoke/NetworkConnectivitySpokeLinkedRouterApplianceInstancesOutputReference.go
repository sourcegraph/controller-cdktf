package networkconnectivityspoke

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/networkconnectivityspoke/internal"
)

type NetworkConnectivitySpokeLinkedRouterApplianceInstancesOutputReference interface {
	cdktf.ComplexObject
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
	ExcludeExportRanges() *[]*string
	SetExcludeExportRanges(val *[]*string)
	ExcludeExportRangesInput() *[]*string
	ExcludeImportRanges() *[]*string
	SetExcludeImportRanges(val *[]*string)
	ExcludeImportRangesInput() *[]*string
	// Experimental.
	Fqn() *string
	IncludeExportRanges() *[]*string
	SetIncludeExportRanges(val *[]*string)
	IncludeExportRangesInput() *[]*string
	IncludeImportRanges() *[]*string
	SetIncludeImportRanges(val *[]*string)
	IncludeImportRangesInput() *[]*string
	Instances() NetworkConnectivitySpokeLinkedRouterApplianceInstancesInstancesList
	InstancesInput() any
	InternalValue() *NetworkConnectivitySpokeLinkedRouterApplianceInstances
	SetInternalValue(val *NetworkConnectivitySpokeLinkedRouterApplianceInstances)
	SiteToSiteDataTransfer() any
	SetSiteToSiteDataTransfer(val any)
	SiteToSiteDataTransferInput() any
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
	PutInstances(value any)
	ResetExcludeExportRanges()
	ResetExcludeImportRanges()
	ResetIncludeExportRanges()
	ResetIncludeImportRanges()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for NetworkConnectivitySpokeLinkedRouterApplianceInstancesOutputReference
type jsiiProxy_NetworkConnectivitySpokeLinkedRouterApplianceInstancesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_NetworkConnectivitySpokeLinkedRouterApplianceInstancesOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivitySpokeLinkedRouterApplianceInstancesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivitySpokeLinkedRouterApplianceInstancesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivitySpokeLinkedRouterApplianceInstancesOutputReference) ExcludeExportRanges() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludeExportRanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivitySpokeLinkedRouterApplianceInstancesOutputReference) ExcludeExportRangesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludeExportRangesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivitySpokeLinkedRouterApplianceInstancesOutputReference) ExcludeImportRanges() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludeImportRanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivitySpokeLinkedRouterApplianceInstancesOutputReference) ExcludeImportRangesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludeImportRangesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivitySpokeLinkedRouterApplianceInstancesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivitySpokeLinkedRouterApplianceInstancesOutputReference) IncludeExportRanges() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includeExportRanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivitySpokeLinkedRouterApplianceInstancesOutputReference) IncludeExportRangesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includeExportRangesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivitySpokeLinkedRouterApplianceInstancesOutputReference) IncludeImportRanges() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includeImportRanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivitySpokeLinkedRouterApplianceInstancesOutputReference) IncludeImportRangesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includeImportRangesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivitySpokeLinkedRouterApplianceInstancesOutputReference) Instances() NetworkConnectivitySpokeLinkedRouterApplianceInstancesInstancesList {
	var returns NetworkConnectivitySpokeLinkedRouterApplianceInstancesInstancesList
	_jsii_.Get(
		j,
		"instances",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivitySpokeLinkedRouterApplianceInstancesOutputReference) InstancesInput() any {
	var returns any
	_jsii_.Get(
		j,
		"instancesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivitySpokeLinkedRouterApplianceInstancesOutputReference) InternalValue() *NetworkConnectivitySpokeLinkedRouterApplianceInstances {
	var returns *NetworkConnectivitySpokeLinkedRouterApplianceInstances
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivitySpokeLinkedRouterApplianceInstancesOutputReference) SiteToSiteDataTransfer() any {
	var returns any
	_jsii_.Get(
		j,
		"siteToSiteDataTransfer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivitySpokeLinkedRouterApplianceInstancesOutputReference) SiteToSiteDataTransferInput() any {
	var returns any
	_jsii_.Get(
		j,
		"siteToSiteDataTransferInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivitySpokeLinkedRouterApplianceInstancesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivitySpokeLinkedRouterApplianceInstancesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewNetworkConnectivitySpokeLinkedRouterApplianceInstancesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) NetworkConnectivitySpokeLinkedRouterApplianceInstancesOutputReference {
	_init_.Initialize()

	if err := validateNewNetworkConnectivitySpokeLinkedRouterApplianceInstancesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_NetworkConnectivitySpokeLinkedRouterApplianceInstancesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.networkConnectivitySpoke.NetworkConnectivitySpokeLinkedRouterApplianceInstancesOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewNetworkConnectivitySpokeLinkedRouterApplianceInstancesOutputReference_Override(n NetworkConnectivitySpokeLinkedRouterApplianceInstancesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.networkConnectivitySpoke.NetworkConnectivitySpokeLinkedRouterApplianceInstancesOutputReference",
		[]any{terraformResource, terraformAttribute},
		n,
	)
}

func (j *jsiiProxy_NetworkConnectivitySpokeLinkedRouterApplianceInstancesOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_NetworkConnectivitySpokeLinkedRouterApplianceInstancesOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_NetworkConnectivitySpokeLinkedRouterApplianceInstancesOutputReference) SetExcludeExportRanges(val *[]*string) {
	if err := j.validateSetExcludeExportRangesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludeExportRanges",
		val,
	)
}

func (j *jsiiProxy_NetworkConnectivitySpokeLinkedRouterApplianceInstancesOutputReference) SetExcludeImportRanges(val *[]*string) {
	if err := j.validateSetExcludeImportRangesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludeImportRanges",
		val,
	)
}

func (j *jsiiProxy_NetworkConnectivitySpokeLinkedRouterApplianceInstancesOutputReference) SetIncludeExportRanges(val *[]*string) {
	if err := j.validateSetIncludeExportRangesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeExportRanges",
		val,
	)
}

func (j *jsiiProxy_NetworkConnectivitySpokeLinkedRouterApplianceInstancesOutputReference) SetIncludeImportRanges(val *[]*string) {
	if err := j.validateSetIncludeImportRangesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeImportRanges",
		val,
	)
}

func (j *jsiiProxy_NetworkConnectivitySpokeLinkedRouterApplianceInstancesOutputReference) SetInternalValue(val *NetworkConnectivitySpokeLinkedRouterApplianceInstances) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_NetworkConnectivitySpokeLinkedRouterApplianceInstancesOutputReference) SetSiteToSiteDataTransfer(val any) {
	if err := j.validateSetSiteToSiteDataTransferParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"siteToSiteDataTransfer",
		val,
	)
}

func (j *jsiiProxy_NetworkConnectivitySpokeLinkedRouterApplianceInstancesOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_NetworkConnectivitySpokeLinkedRouterApplianceInstancesOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (n *jsiiProxy_NetworkConnectivitySpokeLinkedRouterApplianceInstancesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkConnectivitySpokeLinkedRouterApplianceInstancesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
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

func (n *jsiiProxy_NetworkConnectivitySpokeLinkedRouterApplianceInstancesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (n *jsiiProxy_NetworkConnectivitySpokeLinkedRouterApplianceInstancesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (n *jsiiProxy_NetworkConnectivitySpokeLinkedRouterApplianceInstancesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (n *jsiiProxy_NetworkConnectivitySpokeLinkedRouterApplianceInstancesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (n *jsiiProxy_NetworkConnectivitySpokeLinkedRouterApplianceInstancesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (n *jsiiProxy_NetworkConnectivitySpokeLinkedRouterApplianceInstancesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (n *jsiiProxy_NetworkConnectivitySpokeLinkedRouterApplianceInstancesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (n *jsiiProxy_NetworkConnectivitySpokeLinkedRouterApplianceInstancesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (n *jsiiProxy_NetworkConnectivitySpokeLinkedRouterApplianceInstancesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkConnectivitySpokeLinkedRouterApplianceInstancesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (n *jsiiProxy_NetworkConnectivitySpokeLinkedRouterApplianceInstancesOutputReference) PutInstances(value any) {
	if err := n.validatePutInstancesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putInstances",
		[]any{value},
	)
}

func (n *jsiiProxy_NetworkConnectivitySpokeLinkedRouterApplianceInstancesOutputReference) ResetExcludeExportRanges() {
	_jsii_.InvokeVoid(
		n,
		"resetExcludeExportRanges",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkConnectivitySpokeLinkedRouterApplianceInstancesOutputReference) ResetExcludeImportRanges() {
	_jsii_.InvokeVoid(
		n,
		"resetExcludeImportRanges",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkConnectivitySpokeLinkedRouterApplianceInstancesOutputReference) ResetIncludeExportRanges() {
	_jsii_.InvokeVoid(
		n,
		"resetIncludeExportRanges",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkConnectivitySpokeLinkedRouterApplianceInstancesOutputReference) ResetIncludeImportRanges() {
	_jsii_.InvokeVoid(
		n,
		"resetIncludeImportRanges",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkConnectivitySpokeLinkedRouterApplianceInstancesOutputReference) Resolve(_context cdktf.IResolveContext) any {
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

func (n *jsiiProxy_NetworkConnectivitySpokeLinkedRouterApplianceInstancesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
