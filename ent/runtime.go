// Code generated by entc, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/thuanvh/golicenseserver/ent/customer"
	"github.com/thuanvh/golicenseserver/ent/customerlicense"
	"github.com/thuanvh/golicenseserver/ent/license"
	"github.com/thuanvh/golicenseserver/ent/schema"
	"github.com/thuanvh/golicenseserver/ent/todo"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	customerFields := schema.Customer{}.Fields()
	_ = customerFields
	// customerDescName is the schema descriptor for name field.
	customerDescName := customerFields[0].Descriptor()
	// customer.NameValidator is a validator for the "name" field. It is called by the builders before save.
	customer.NameValidator = customerDescName.Validators[0].(func(string) error)
	// customerDescAddress is the schema descriptor for address field.
	customerDescAddress := customerFields[1].Descriptor()
	// customer.AddressValidator is a validator for the "address" field. It is called by the builders before save.
	customer.AddressValidator = customerDescAddress.Validators[0].(func(string) error)
	customerlicenseFields := schema.CustomerLicense{}.Fields()
	_ = customerlicenseFields
	// customerlicenseDescLicenseCode is the schema descriptor for license_code field.
	customerlicenseDescLicenseCode := customerlicenseFields[0].Descriptor()
	// customerlicense.LicenseCodeValidator is a validator for the "license_code" field. It is called by the builders before save.
	customerlicense.LicenseCodeValidator = customerlicenseDescLicenseCode.Validators[0].(func(string) error)
	// customerlicenseDescActive is the schema descriptor for active field.
	customerlicenseDescActive := customerlicenseFields[1].Descriptor()
	// customerlicense.DefaultActive holds the default value on creation for the active field.
	customerlicense.DefaultActive = customerlicenseDescActive.Default.(bool)
	// customerlicenseDescStartDate is the schema descriptor for start_date field.
	customerlicenseDescStartDate := customerlicenseFields[2].Descriptor()
	// customerlicense.DefaultStartDate holds the default value on creation for the start_date field.
	customerlicense.DefaultStartDate = customerlicenseDescStartDate.Default.(func() time.Time)
	// customerlicenseDescEndDate is the schema descriptor for end_date field.
	customerlicenseDescEndDate := customerlicenseFields[3].Descriptor()
	// customerlicense.DefaultEndDate holds the default value on creation for the end_date field.
	customerlicense.DefaultEndDate = customerlicenseDescEndDate.Default.(func() time.Time)
	// customerlicenseDescCPU is the schema descriptor for cpu field.
	customerlicenseDescCPU := customerlicenseFields[4].Descriptor()
	// customerlicense.DefaultCPU holds the default value on creation for the cpu field.
	customerlicense.DefaultCPU = customerlicenseDescCPU.Default.(int)
	// customerlicenseDescStorage is the schema descriptor for storage field.
	customerlicenseDescStorage := customerlicenseFields[5].Descriptor()
	// customerlicense.DefaultStorage holds the default value on creation for the storage field.
	customerlicense.DefaultStorage = customerlicenseDescStorage.Default.(int)
	// customerlicenseDescNumberOfNodes is the schema descriptor for number_of_nodes field.
	customerlicenseDescNumberOfNodes := customerlicenseFields[6].Descriptor()
	// customerlicense.DefaultNumberOfNodes holds the default value on creation for the number_of_nodes field.
	customerlicense.DefaultNumberOfNodes = customerlicenseDescNumberOfNodes.Default.(int)
	licenseFields := schema.License{}.Fields()
	_ = licenseFields
	// licenseDescName is the schema descriptor for name field.
	licenseDescName := licenseFields[0].Descriptor()
	// license.NameValidator is a validator for the "name" field. It is called by the builders before save.
	license.NameValidator = licenseDescName.Validators[0].(func(string) error)
	// licenseDescDuration is the schema descriptor for duration field.
	licenseDescDuration := licenseFields[1].Descriptor()
	// license.DurationValidator is a validator for the "duration" field. It is called by the builders before save.
	license.DurationValidator = licenseDescDuration.Validators[0].(func(int) error)
	// licenseDescCPU is the schema descriptor for cpu field.
	licenseDescCPU := licenseFields[2].Descriptor()
	// license.DefaultCPU holds the default value on creation for the cpu field.
	license.DefaultCPU = licenseDescCPU.Default.(int)
	// licenseDescStorage is the schema descriptor for storage field.
	licenseDescStorage := licenseFields[3].Descriptor()
	// license.DefaultStorage holds the default value on creation for the storage field.
	license.DefaultStorage = licenseDescStorage.Default.(int)
	// licenseDescNumberOfNodes is the schema descriptor for number_of_nodes field.
	licenseDescNumberOfNodes := licenseFields[4].Descriptor()
	// license.DefaultNumberOfNodes holds the default value on creation for the number_of_nodes field.
	license.DefaultNumberOfNodes = licenseDescNumberOfNodes.Default.(int)
	todoFields := schema.Todo{}.Fields()
	_ = todoFields
	// todoDescText is the schema descriptor for text field.
	todoDescText := todoFields[0].Descriptor()
	// todo.TextValidator is a validator for the "text" field. It is called by the builders before save.
	todo.TextValidator = todoDescText.Validators[0].(func(string) error)
	// todoDescCreatedAt is the schema descriptor for created_at field.
	todoDescCreatedAt := todoFields[1].Descriptor()
	// todo.DefaultCreatedAt holds the default value on creation for the created_at field.
	todo.DefaultCreatedAt = todoDescCreatedAt.Default.(func() time.Time)
	// todoDescPriority is the schema descriptor for priority field.
	todoDescPriority := todoFields[3].Descriptor()
	// todo.DefaultPriority holds the default value on creation for the priority field.
	todo.DefaultPriority = todoDescPriority.Default.(int)
}