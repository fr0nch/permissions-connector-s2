package permissions

/*
#include "group.h"
#cgo noescape SetParent
#cgo noescape GetParent
#cgo noescape DumpPermissionsGroup
#cgo noescape GetAllGroups
#cgo noescape HasPermissionGroupExtended
#cgo noescape HasPermissionGroup
#cgo noescape HasParentGroup
#cgo noescape GetPriorityGroup
#cgo noescape AddPermissionGroup
#cgo noescape SetPermissionGroup
#cgo noescape RemovePermissionGroup
#cgo noescape GetOptionGroup
#cgo noescape SetOptionGroup
#cgo noescape GetAllOptionsGroup
#cgo noescape CreateGroup
#cgo noescape DeleteGroup
#cgo noescape GroupExists
#cgo noescape LoadGroups
*/
import "C"
import (
	"errors"
	"github.com/untrustedmodders/go-plugify"
	"reflect"
	"runtime"
	"unsafe"
)

var _ = errors.New("")
var _ = reflect.TypeOf(0)
var _ = runtime.GOOS
var _ = unsafe.Sizeof(0)
var _ = plugify.ApiVersion

// Generated from permissions (group: group)

var _SetParent = func(pluginID int64, childName string, parentName string, dontBroadcast bool) Status {
	var __retVal Status
	__pluginID := C.int64_t(pluginID)
	__childName := plugify.ConstructString(childName)
	__parentName := plugify.ConstructString(parentName)
	__dontBroadcast := C.bool(dontBroadcast)
	plugify.Block{
		Try: func() {
			__retVal = Status(C.SetParent(__pluginID, (*C.String)(unsafe.Pointer(&__childName)), (*C.String)(unsafe.Pointer(&__parentName)), __dontBroadcast))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__childName)
			plugify.DestroyString(&__parentName)
		},
	}.Do()
	return __retVal
}

// SetParent
//
//	@brief Set parent group for child group
//
//	@param pluginID: Identifier of the plugin that calls the method.
//	@param childName: Child group name
//	@param parentName: Parent group name to set
//
//	@return Success, ChildGroupNotFound, ParentGroupNotFound
func SetParent(pluginID int64, childName string, parentName string, dontBroadcast bool) Status {
	return _SetParent(pluginID, childName, parentName, dontBroadcast)
}

var _GetParent = func(groupName string, parentName *string) Status {
	var __retVal Status
	__groupName := plugify.ConstructString(groupName)
	__parentName := plugify.ConstructString(*parentName)
	plugify.Block{
		Try: func() {
			__retVal = Status(C.GetParent((*C.String)(unsafe.Pointer(&__groupName)), (*C.String)(unsafe.Pointer(&__parentName))))
			// Unmarshal - Convert native data to managed data.
			*parentName = plugify.GetStringData[string](&__parentName)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__groupName)
			plugify.DestroyString(&__parentName)
		},
	}.Do()
	return __retVal
}

// GetParent
//
//	@brief Get parent of requested group
//
//	@param groupName: Group name
//	@param parentName: Parent name
//
//	@return Success, ChildGroupNotFound, ParentGroupNotFound
func GetParent(groupName string, parentName *string) Status {
	return _GetParent(groupName, parentName)
}

var _DumpPermissionsGroup = func(name string, perms *[]string) Status {
	var __retVal Status
	__name := plugify.ConstructString(name)
	__perms := plugify.ConstructVectorString(*perms)
	plugify.Block{
		Try: func() {
			__retVal = Status(C.DumpPermissionsGroup((*C.String)(unsafe.Pointer(&__name)), (*C.Vector)(unsafe.Pointer(&__perms))))
			// Unmarshal - Convert native data to managed data.
			plugify.GetVectorDataStringTo(&__perms, perms)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__name)
			plugify.DestroyVectorString(&__perms)
		},
	}.Do()
	return __retVal
}

// DumpPermissionsGroup
//
//	@brief Get permissions of group
//
//	@param name: Group name
//	@param perms: Permissions
//
//	@return Success, GroupNotFound
func DumpPermissionsGroup(name string, perms *[]string) Status {
	return _DumpPermissionsGroup(name, perms)
}

var _GetAllGroups = func() []string {
	var __retVal []string
	var __retVal_native plugify.PlgVector
	plugify.Block{
		Try: func() {
			__native := C.GetAllGroups()
			__retVal_native = *(*plugify.PlgVector)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetVectorDataString[string](&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorString(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// GetAllGroups
//
//	@brief Get all created groups
//
//
//	@return Array of groups
func GetAllGroups() []string {
	return _GetAllGroups()
}

var _HasPermissionGroupExtended = func(name string, perm string, exact bool) Status {
	var __retVal Status
	__name := plugify.ConstructString(name)
	__perm := plugify.ConstructString(perm)
	__exact := C.bool(exact)
	plugify.Block{
		Try: func() {
			__retVal = Status(C.HasPermissionGroupExtended((*C.String)(unsafe.Pointer(&__name)), (*C.String)(unsafe.Pointer(&__perm)), __exact))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__name)
			plugify.DestroyString(&__perm)
		},
	}.Do()
	return __retVal
}

// HasPermissionGroupExtended
//
//	@brief Check if a group has a specific permission.
//
//	@param name: Group name.
//	@param perm: Permission line.
//	@param exact: Checking permission with ignoring wildcards (pass 'false' for default behavior)
//
//	@return Allow, Disallow, PermNotFound, GroupNotFound
func HasPermissionGroupExtended(name string, perm string, exact bool) Status {
	return _HasPermissionGroupExtended(name, perm, exact)
}

var _HasPermissionGroup = func(name string, perm string) Status {
	var __retVal Status
	__name := plugify.ConstructString(name)
	__perm := plugify.ConstructString(perm)
	plugify.Block{
		Try: func() {
			__retVal = Status(C.HasPermissionGroup((*C.String)(unsafe.Pointer(&__name)), (*C.String)(unsafe.Pointer(&__perm))))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__name)
			plugify.DestroyString(&__perm)
		},
	}.Do()
	return __retVal
}

// HasPermissionGroup
//
//	@brief Check if a group has a specific permission.
//
//	@param name: Group name.
//	@param perm: Permission line.
//
//	@return Allow, Disallow, PermNotFound, GroupNotFound
func HasPermissionGroup(name string, perm string) Status {
	return _HasPermissionGroup(name, perm)
}

var _HasParentGroup = func(childName string, parentName string) Status {
	var __retVal Status
	__childName := plugify.ConstructString(childName)
	__parentName := plugify.ConstructString(parentName)
	plugify.Block{
		Try: func() {
			__retVal = Status(C.HasParentGroup((*C.String)(unsafe.Pointer(&__childName)), (*C.String)(unsafe.Pointer(&__parentName))))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__childName)
			plugify.DestroyString(&__parentName)
		},
	}.Do()
	return __retVal
}

// HasParentGroup
//
//	@brief Check if parent_name is a parent group for child_name.
//
//	@param childName: Child group name.
//	@param parentName: Parent group name to check.
//
//	@return Allow, Disallow, ChildGroupNotFound, ParentGroupNotFound
func HasParentGroup(childName string, parentName string) Status {
	return _HasParentGroup(childName, parentName)
}

var _GetPriorityGroup = func(groupName string, priority *int32) Status {
	var __retVal Status
	__groupName := plugify.ConstructString(groupName)
	__priority := C.int32_t(*priority)
	plugify.Block{
		Try: func() {
			__retVal = Status(C.GetPriorityGroup((*C.String)(unsafe.Pointer(&__groupName)), &__priority))
			// Unmarshal - Convert native data to managed data.
			*priority = int32(__priority)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__groupName)
		},
	}.Do()
	return __retVal
}

// GetPriorityGroup
//
//	@brief Get the priority of a group.
//
//	@param groupName: Group name.
//	@param priority: Priority
//
//	@return Success, GroupNotFound
func GetPriorityGroup(groupName string, priority *int32) Status {
	return _GetPriorityGroup(groupName, priority)
}

var _AddPermissionGroup = func(pluginID int64, name string, perm string, dontBroadcast bool) Status {
	var __retVal Status
	__pluginID := C.int64_t(pluginID)
	__name := plugify.ConstructString(name)
	__perm := plugify.ConstructString(perm)
	__dontBroadcast := C.bool(dontBroadcast)
	plugify.Block{
		Try: func() {
			__retVal = Status(C.AddPermissionGroup(__pluginID, (*C.String)(unsafe.Pointer(&__name)), (*C.String)(unsafe.Pointer(&__perm)), __dontBroadcast))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__name)
			plugify.DestroyString(&__perm)
		},
	}.Do()
	return __retVal
}

// AddPermissionGroup
//
//	@brief Add a permission to a group.
//
//	@param pluginID: Identifier of the plugin that calls the method.
//	@param name: Group name.
//	@param perm: Permission line.
//	@param dontBroadcast: If set to `true`, suppresses dispatching of the permission change event to registered GroupPermission listeners. The permission is still applied internally.
//
//	@return Success, GroupNotFound, PermAlreadyGranted
func AddPermissionGroup(pluginID int64, name string, perm string, dontBroadcast bool) Status {
	return _AddPermissionGroup(pluginID, name, perm, dontBroadcast)
}

var _SetPermissionGroup = func(pluginID int64, name string, perm string, dontBroadcast bool) Status {
	var __retVal Status
	__pluginID := C.int64_t(pluginID)
	__name := plugify.ConstructString(name)
	__perm := plugify.ConstructString(perm)
	__dontBroadcast := C.bool(dontBroadcast)
	plugify.Block{
		Try: func() {
			__retVal = Status(C.SetPermissionGroup(__pluginID, (*C.String)(unsafe.Pointer(&__name)), (*C.String)(unsafe.Pointer(&__perm)), __dontBroadcast))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__name)
			plugify.DestroyString(&__perm)
		},
	}.Do()
	return __retVal
}

// SetPermissionGroup
func SetPermissionGroup(pluginID int64, name string, perm string, dontBroadcast bool) Status {
	return _SetPermissionGroup(pluginID, name, perm, dontBroadcast)
}

var _RemovePermissionGroup = func(pluginID int64, name string, perm string, recursiveDeletion bool, dontBroadcast bool) Status {
	var __retVal Status
	__pluginID := C.int64_t(pluginID)
	__name := plugify.ConstructString(name)
	__perm := plugify.ConstructString(perm)
	__recursiveDeletion := C.bool(recursiveDeletion)
	__dontBroadcast := C.bool(dontBroadcast)
	plugify.Block{
		Try: func() {
			__retVal = Status(C.RemovePermissionGroup(__pluginID, (*C.String)(unsafe.Pointer(&__name)), (*C.String)(unsafe.Pointer(&__perm)), __recursiveDeletion, __dontBroadcast))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__name)
			plugify.DestroyString(&__perm)
		},
	}.Do()
	return __retVal
}

// RemovePermissionGroup
//
//	@brief Remove a permission from a group.
//
//	@param pluginID: Identifier of the plugin that calls the method.
//	@param name: Group name.
//	@param perm: Permission line.
//	@param recursiveDeletion: Delete all nested perms.
//
//	@return Success, GroupNotFound, PermNotFound
func RemovePermissionGroup(pluginID int64, name string, perm string, recursiveDeletion bool, dontBroadcast bool) Status {
	return _RemovePermissionGroup(pluginID, name, perm, recursiveDeletion, dontBroadcast)
}

var _GetOptionGroup = func(groupName string, optionName string, value *any) Status {
	var __retVal Status
	__groupName := plugify.ConstructString(groupName)
	__optionName := plugify.ConstructString(optionName)
	__value := plugify.ConstructVariant(*value)
	plugify.Block{
		Try: func() {
			__retVal = Status(C.GetOptionGroup((*C.String)(unsafe.Pointer(&__groupName)), (*C.String)(unsafe.Pointer(&__optionName)), (*C.Variant)(unsafe.Pointer(&__value))))
			// Unmarshal - Convert native data to managed data.
			*value = plugify.GetVariantData(&__value)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__groupName)
			plugify.DestroyString(&__optionName)
			plugify.DestroyVariant(&__value)
		},
	}.Do()
	return __retVal
}

// GetOptionGroup
//
//	@brief Get an option value for a group.
//
//	@param groupName: Group name
//	@param optionName: Option name
//	@param value: Option value
//
//	@return Success, OptionNotFound, GroupNotFound
func GetOptionGroup(groupName string, optionName string, value *any) Status {
	return _GetOptionGroup(groupName, optionName, value)
}

var _SetOptionGroup = func(pluginID int64, groupName string, optionName string, value any, dontBroadcast bool) Status {
	var __retVal Status
	__pluginID := C.int64_t(pluginID)
	__groupName := plugify.ConstructString(groupName)
	__optionName := plugify.ConstructString(optionName)
	__value := plugify.ConstructVariant(value)
	__dontBroadcast := C.bool(dontBroadcast)
	plugify.Block{
		Try: func() {
			__retVal = Status(C.SetOptionGroup(__pluginID, (*C.String)(unsafe.Pointer(&__groupName)), (*C.String)(unsafe.Pointer(&__optionName)), (*C.Variant)(unsafe.Pointer(&__value)), __dontBroadcast))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__groupName)
			plugify.DestroyString(&__optionName)
			plugify.DestroyVariant(&__value)
		},
	}.Do()
	return __retVal
}

// SetOptionGroup
//
//	@brief Set an option value for a group.
//
//	@param pluginID: Identifier of the plugin that calls the method.
//	@param groupName: Group name
//	@param optionName: Option name
//	@param value: Option value.
//
//	@return Success, GroupNotFound
func SetOptionGroup(pluginID int64, groupName string, optionName string, value any, dontBroadcast bool) Status {
	return _SetOptionGroup(pluginID, groupName, optionName, value, dontBroadcast)
}

var _GetAllOptionsGroup = func(groupName string, optionNames *[]string, values *[]any) Status {
	var __retVal Status
	__groupName := plugify.ConstructString(groupName)
	__optionNames := plugify.ConstructVectorString(*optionNames)
	__values := plugify.ConstructVectorVariant(*values)
	plugify.Block{
		Try: func() {
			__retVal = Status(C.GetAllOptionsGroup((*C.String)(unsafe.Pointer(&__groupName)), (*C.Vector)(unsafe.Pointer(&__optionNames)), (*C.Vector)(unsafe.Pointer(&__values))))
			// Unmarshal - Convert native data to managed data.
			plugify.GetVectorDataStringTo(&__optionNames, optionNames)
			plugify.GetVectorDataVariantTo(&__values, values)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__groupName)
			plugify.DestroyVectorString(&__optionNames)
			plugify.DestroyVectorVariant(&__values)
		},
	}.Do()
	return __retVal
}

// GetAllOptionsGroup
//
//	@brief Get all options from group.
//
//	@param groupName: Group name
//	@param optionNames: Array of option names
//	@param values: Array of option values
//
//	@return Success, GroupNotFound
func GetAllOptionsGroup(groupName string, optionNames *[]string, values *[]any) Status {
	return _GetAllOptionsGroup(groupName, optionNames, values)
}

var _CreateGroup = func(pluginID int64, name string, perms []string, priority int32, parent string, dontBroadcast bool) Status {
	var __retVal Status
	__pluginID := C.int64_t(pluginID)
	__name := plugify.ConstructString(name)
	__perms := plugify.ConstructVectorString(perms)
	__priority := C.int32_t(priority)
	__parent := plugify.ConstructString(parent)
	__dontBroadcast := C.bool(dontBroadcast)
	plugify.Block{
		Try: func() {
			__retVal = Status(C.CreateGroup(__pluginID, (*C.String)(unsafe.Pointer(&__name)), (*C.Vector)(unsafe.Pointer(&__perms)), __priority, (*C.String)(unsafe.Pointer(&__parent)), __dontBroadcast))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__name)
			plugify.DestroyVectorString(&__perms)
			plugify.DestroyString(&__parent)
		},
	}.Do()
	return __retVal
}

// CreateGroup
//
//	@brief Create a new group.
//
//	@param pluginID: Identifier of the plugin that calls the method.
//	@param name: Group name.
//	@param perms: Array of permission lines.
//	@param priority: Group priority.
//	@param parent: Parent group name.
//
//	@return Success, GroupAlreadyExist, ParentGroupNotFound
func CreateGroup(pluginID int64, name string, perms []string, priority int32, parent string, dontBroadcast bool) Status {
	return _CreateGroup(pluginID, name, perms, priority, parent, dontBroadcast)
}

var _DeleteGroup = func(pluginID int64, name string, dontBroadcast bool) Status {
	var __retVal Status
	__pluginID := C.int64_t(pluginID)
	__name := plugify.ConstructString(name)
	__dontBroadcast := C.bool(dontBroadcast)
	plugify.Block{
		Try: func() {
			__retVal = Status(C.DeleteGroup(__pluginID, (*C.String)(unsafe.Pointer(&__name)), __dontBroadcast))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__name)
		},
	}.Do()
	return __retVal
}

// DeleteGroup
//
//	@brief Delete a group.
//
//	@param pluginID: Identifier of the plugin that calls the method.
//	@param name: Group name.
//
//	@return Success if deleted; GroupNotFound if group not found.
func DeleteGroup(pluginID int64, name string, dontBroadcast bool) Status {
	return _DeleteGroup(pluginID, name, dontBroadcast)
}

var _GroupExists = func(name string) bool {
	var __retVal bool
	__name := plugify.ConstructString(name)
	plugify.Block{
		Try: func() {
			__retVal = bool(C.GroupExists((*C.String)(unsafe.Pointer(&__name))))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__name)
		},
	}.Do()
	return __retVal
}

// GroupExists
//
//	@brief Check if a group exists.
//
//	@param name: Group name.
//
//	@return True if group exists, false otherwise.
func GroupExists(name string) bool {
	return _GroupExists(name)
}

var _LoadGroups = func(pluginID int64, dontBroadcast bool) Status {
	var __retVal Status
	__pluginID := C.int64_t(pluginID)
	__dontBroadcast := C.bool(dontBroadcast)
	__retVal = Status(C.LoadGroups(__pluginID, __dontBroadcast))
	return __retVal
}

// LoadGroups
//
//	@brief Dispatches a request to load server groups for a plugin.
//
//	@param pluginID: Identifier of the calling plugin.
func LoadGroups(pluginID int64, dontBroadcast bool) Status {
	return _LoadGroups(pluginID, dontBroadcast)
}
