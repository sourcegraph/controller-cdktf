package backupdrrestoreworkload

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google/backupdrrestoreworkload/internal"
)

type BackupDrRestoreWorkloadDiskRestorePropertiesGuestOsFeatureList interface {
	cdktn.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	WrapsSet() *bool
	// Experimental.
	SetWrapsSet(val *bool)
	// Creating an iterator for this complex list.
	//
	// The list will be converted into a map with the mapKeyAttributeName as the key.
	// Experimental.
	AllWithMapKey(mapKeyAttributeName *string) cdktn.DynamicListTerraformIterator
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) BackupDrRestoreWorkloadDiskRestorePropertiesGuestOsFeatureOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BackupDrRestoreWorkloadDiskRestorePropertiesGuestOsFeatureList
type jsiiProxy_BackupDrRestoreWorkloadDiskRestorePropertiesGuestOsFeatureList struct {
	internal.Type__cdktnComplexList
}

func (j *jsiiProxy_BackupDrRestoreWorkloadDiskRestorePropertiesGuestOsFeatureList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadDiskRestorePropertiesGuestOsFeatureList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadDiskRestorePropertiesGuestOsFeatureList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadDiskRestorePropertiesGuestOsFeatureList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadDiskRestorePropertiesGuestOsFeatureList) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadDiskRestorePropertiesGuestOsFeatureList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewBackupDrRestoreWorkloadDiskRestorePropertiesGuestOsFeatureList(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) BackupDrRestoreWorkloadDiskRestorePropertiesGuestOsFeatureList {
	_init_.Initialize()

	if err := validateNewBackupDrRestoreWorkloadDiskRestorePropertiesGuestOsFeatureListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_BackupDrRestoreWorkloadDiskRestorePropertiesGuestOsFeatureList{}

	_jsii_.Create(
		"@cdktn/provider-google.backupDrRestoreWorkload.BackupDrRestoreWorkloadDiskRestorePropertiesGuestOsFeatureList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewBackupDrRestoreWorkloadDiskRestorePropertiesGuestOsFeatureList_Override(b BackupDrRestoreWorkloadDiskRestorePropertiesGuestOsFeatureList, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.backupDrRestoreWorkload.BackupDrRestoreWorkloadDiskRestorePropertiesGuestOsFeatureList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		b,
	)
}

func (j *jsiiProxy_BackupDrRestoreWorkloadDiskRestorePropertiesGuestOsFeatureList)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BackupDrRestoreWorkloadDiskRestorePropertiesGuestOsFeatureList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BackupDrRestoreWorkloadDiskRestorePropertiesGuestOsFeatureList)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_BackupDrRestoreWorkloadDiskRestorePropertiesGuestOsFeatureList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (b *jsiiProxy_BackupDrRestoreWorkloadDiskRestorePropertiesGuestOsFeatureList) AllWithMapKey(mapKeyAttributeName *string) cdktn.DynamicListTerraformIterator {
	if err := b.validateAllWithMapKeyParameters(mapKeyAttributeName); err != nil {
		panic(err)
	}
	var returns cdktn.DynamicListTerraformIterator

	_jsii_.Invoke(
		b,
		"allWithMapKey",
		[]interface{}{mapKeyAttributeName},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupDrRestoreWorkloadDiskRestorePropertiesGuestOsFeatureList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupDrRestoreWorkloadDiskRestorePropertiesGuestOsFeatureList) Get(index *float64) BackupDrRestoreWorkloadDiskRestorePropertiesGuestOsFeatureOutputReference {
	if err := b.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns BackupDrRestoreWorkloadDiskRestorePropertiesGuestOsFeatureOutputReference

	_jsii_.Invoke(
		b,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupDrRestoreWorkloadDiskRestorePropertiesGuestOsFeatureList) Resolve(context cdktn.IResolveContext) interface{} {
	if err := b.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		b,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupDrRestoreWorkloadDiskRestorePropertiesGuestOsFeatureList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

