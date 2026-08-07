package permissions

/*
#include "user.h"
#cgo noescape DumpPermissions
#cgo noescape CanAffectUser
#cgo noescape HasPermissionExtended
#cgo noescape HasPermission
#cgo noescape HasGroupExtended
#cgo noescape HasGroup
#cgo noescape GetUserGroups
#cgo noescape GetImmunity
#cgo noescape SetImmunity
#cgo noescape AddPermission
#cgo noescape SetPermission
#cgo noescape RemovePermission
#cgo noescape AddGroup
#cgo noescape RemoveGroup
#cgo noescape GetCookie
#cgo noescape SetCookie
#cgo noescape GetAllCookies
#cgo noescape CreateUser
#cgo noescape DeleteUser
#cgo noescape UserExists
#cgo noescape DumpUsersList
#cgo noescape LoadUser
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

// Generated from permissions (group: user)

var _DumpPermissions = func(targetID uint64, perms *[]string) Status {
	var __retVal Status
	__targetID := C.uint64_t(targetID)
	__perms := plugify.ConstructVectorString(*perms)
	plugify.Block{
		Try: func() {
			__retVal = Status(C.DumpPermissions(__targetID, (*C.Vector)(unsafe.Pointer(&__perms))))
			// Unmarshal - Convert native data to managed data.
			plugify.GetVectorDataStringTo(&__perms, perms)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorString(&__perms)
		},
	}.Do()
	return __retVal
}

// DumpPermissions
//
//	@brief Get permissions of user
//
//	@param targetID: Player ID
//	@param perms: Permissions
//
//	@return Success, TargetUserNotFound
func DumpPermissions(targetID uint64, perms *[]string) Status {
	return _DumpPermissions(targetID, perms)
}

var _CanAffectUser = func(actorID uint64, targetID uint64) Status {
	var __retVal Status
	__actorID := C.uint64_t(actorID)
	__targetID := C.uint64_t(targetID)
	__retVal = Status(C.CanAffectUser(__actorID, __targetID))
	return __retVal
}

// CanAffectUser
//
//	@brief Check players immunity or groups priority
//
//	@param actorID: Player performing the action
//	@param targetID: Player receiving the action
//
//	@return Allow, Disallow, ActorUserNotFound, or TargetUserNotFound
func CanAffectUser(actorID uint64, targetID uint64) Status {
	return _CanAffectUser(actorID, targetID)
}

var _HasPermissionExtended = func(targetID uint64, perm string, exact bool, permSource *PermSource, timestamp *int64) Status {
	var __retVal Status
	__targetID := C.uint64_t(targetID)
	__perm := plugify.ConstructString(perm)
	__exact := C.bool(exact)
	__permSource := C.uint32_t(*permSource)
	__timestamp := C.int64_t(*timestamp)
	plugify.Block{
		Try: func() {
			__retVal = Status(C.HasPermissionExtended(__targetID, (*C.String)(unsafe.Pointer(&__perm)), __exact, &__permSource, &__timestamp))
			// Unmarshal - Convert native data to managed data.
			*permSource = PermSource(__permSource)
			*timestamp = int64(__timestamp)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__perm)
		},
	}.Do()
	return __retVal
}

// HasPermissionExtended
//
//	@brief Check if a user has a specific permission.
//
//	@param targetID: Player ID.
//	@param perm: Permission line.
//	@param exact: Checking permission with ignoring wildcards (pass 'false' for default behavior).
//	@param permSource: Permission source.
//	@param timestamp: Permission timestamp.
//
//	@return Allow, Disallow, PermNotFound, TargetUserNotFound
func HasPermissionExtended(targetID uint64, perm string, exact bool, permSource *PermSource, timestamp *int64) Status {
	return _HasPermissionExtended(targetID, perm, exact, permSource, timestamp)
}

var _HasPermission = func(targetID uint64, perm string) Status {
	var __retVal Status
	__targetID := C.uint64_t(targetID)
	__perm := plugify.ConstructString(perm)
	plugify.Block{
		Try: func() {
			__retVal = Status(C.HasPermission(__targetID, (*C.String)(unsafe.Pointer(&__perm))))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__perm)
		},
	}.Do()
	return __retVal
}

// HasPermission
//
//	@brief Check if a user has a specific permission.
//
//	@param targetID: Player ID.
//	@param perm: Permission line.
//
//	@return Allow, Disallow, PermNotFound, TargetUserNotFound
func HasPermission(targetID uint64, perm string) Status {
	return _HasPermission(targetID, perm)
}

var _HasGroupExtended = func(targetID uint64, groupName string, timestamp *int64) Status {
	var __retVal Status
	__targetID := C.uint64_t(targetID)
	__groupName := plugify.ConstructString(groupName)
	__timestamp := C.int64_t(*timestamp)
	plugify.Block{
		Try: func() {
			__retVal = Status(C.HasGroupExtended(__targetID, (*C.String)(unsafe.Pointer(&__groupName)), &__timestamp))
			// Unmarshal - Convert native data to managed data.
			*timestamp = int64(__timestamp)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__groupName)
		},
	}.Do()
	return __retVal
}

// HasGroupExtended
//
//	@brief Check if a user belongs to a specific group (directly or via parent groups).
//
//	@param targetID: Player ID.
//	@param groupName: Group name.
//	@param timestamp: Group timestamp.
//
//	@return PermanentGroup, TemporalGroup, GroupNotDefined, TargetUserNotFound, GroupNotFound
func HasGroupExtended(targetID uint64, groupName string, timestamp *int64) Status {
	return _HasGroupExtended(targetID, groupName, timestamp)
}

var _HasGroup = func(targetID uint64, groupName string) Status {
	var __retVal Status
	__targetID := C.uint64_t(targetID)
	__groupName := plugify.ConstructString(groupName)
	plugify.Block{
		Try: func() {
			__retVal = Status(C.HasGroup(__targetID, (*C.String)(unsafe.Pointer(&__groupName))))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__groupName)
		},
	}.Do()
	return __retVal
}

// HasGroup
//
//	@brief Check if a user belongs to a specific group (directly or via parent groups).
//
//	@param targetID: Player ID.
//	@param groupName: Group name.
//
//	@return PermanentGroup, TemporalGroup, GroupNotDefined, TargetUserNotFound, GroupNotFound
func HasGroup(targetID uint64, groupName string) Status {
	return _HasGroup(targetID, groupName)
}

var _GetUserGroups = func(targetID uint64, outGroups *[]string) Status {
	var __retVal Status
	__targetID := C.uint64_t(targetID)
	__outGroups := plugify.ConstructVectorString(*outGroups)
	plugify.Block{
		Try: func() {
			__retVal = Status(C.GetUserGroups(__targetID, (*C.Vector)(unsafe.Pointer(&__outGroups))))
			// Unmarshal - Convert native data to managed data.
			plugify.GetVectorDataStringTo(&__outGroups, outGroups)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorString(&__outGroups)
		},
	}.Do()
	return __retVal
}

// GetUserGroups
//
//	@brief Get user groups.
//
//	@param targetID: Player ID.
//	@param outGroups: Groups
//
//	@return Success, TargetUserNotFound
func GetUserGroups(targetID uint64, outGroups *[]string) Status {
	return _GetUserGroups(targetID, outGroups)
}

var _GetImmunity = func(targetID uint64, immunity *int32) Status {
	var __retVal Status
	__targetID := C.uint64_t(targetID)
	__immunity := C.int32_t(*immunity)
	__retVal = Status(C.GetImmunity(__targetID, &__immunity))
	// Unmarshal - Convert native data to managed data.
	*immunity = int32(__immunity)
	return __retVal
}

// GetImmunity
//
//	@brief Get the immunity level of a user.
//
//	@param targetID: Player ID.
//	@param immunity: Immunity
//
//	@return Success, TargetUserNotFound
func GetImmunity(targetID uint64, immunity *int32) Status {
	return _GetImmunity(targetID, immunity)
}

var _SetImmunity = func(pluginID int64, targetID uint64, immunity int32, dontBroadcast bool) Status {
	var __retVal Status
	__pluginID := C.int64_t(pluginID)
	__targetID := C.uint64_t(targetID)
	__immunity := C.int32_t(immunity)
	__dontBroadcast := C.bool(dontBroadcast)
	__retVal = Status(C.SetImmunity(__pluginID, __targetID, __immunity, __dontBroadcast))
	return __retVal
}

// SetImmunity
//
//	@brief Set the immunity level of a user.
//
//	@param pluginID: Identifier of the plugin that calls the method.
//	@param targetID: Player ID.
//	@param immunity: Immunity.
//
//	@return Success, TargetUserNotFound
func SetImmunity(pluginID int64, targetID uint64, immunity int32, dontBroadcast bool) Status {
	return _SetImmunity(pluginID, targetID, immunity, dontBroadcast)
}

var _AddPermission = func(pluginID int64, targetID uint64, perm string, timestamp int64, dontBroadcast bool) Status {
	var __retVal Status
	__pluginID := C.int64_t(pluginID)
	__targetID := C.uint64_t(targetID)
	__perm := plugify.ConstructString(perm)
	__timestamp := C.int64_t(timestamp)
	__dontBroadcast := C.bool(dontBroadcast)
	plugify.Block{
		Try: func() {
			__retVal = Status(C.AddPermission(__pluginID, __targetID, (*C.String)(unsafe.Pointer(&__perm)), __timestamp, __dontBroadcast))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__perm)
		},
	}.Do()
	return __retVal
}

// AddPermission
//
//	@brief Add a permission to a user.
//
//	@param pluginID: Identifier of the plugin that calls the method.
//	@param targetID: Player ID.
//	@param perm: Permission line.
//	@param timestamp: Permission duration
//	@param dontBroadcast: If set to `true`, suppresses dispatching of the permission change event to registered UserPermission listeners. The permission is still applied internally.
//
//	@return Success, TargetUserNotFound, PermAlreadyGranted
func AddPermission(pluginID int64, targetID uint64, perm string, timestamp int64, dontBroadcast bool) Status {
	return _AddPermission(pluginID, targetID, perm, timestamp, dontBroadcast)
}

var _SetPermission = func(pluginID int64, targetID uint64, perm string, timestamp int64, dontBroadcast bool) Status {
	var __retVal Status
	__pluginID := C.int64_t(pluginID)
	__targetID := C.uint64_t(targetID)
	__perm := plugify.ConstructString(perm)
	__timestamp := C.int64_t(timestamp)
	__dontBroadcast := C.bool(dontBroadcast)
	plugify.Block{
		Try: func() {
			__retVal = Status(C.SetPermission(__pluginID, __targetID, (*C.String)(unsafe.Pointer(&__perm)), __timestamp, __dontBroadcast))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__perm)
		},
	}.Do()
	return __retVal
}

// SetPermission
//
//	@brief Set a permission to a user.
//
//	@param pluginID: Identifier of the plugin that calls the method.
//	@param targetID: Player ID.
//	@param perm: Permission line.
//	@param timestamp: Permission duration
//	@param dontBroadcast: If set to `true`, suppresses dispatching of the permission change event to registered UserPermission listeners. The permission is still applied internally.
//
//	@return Success, TargetUserNotFound, PermAlreadyGranted
func SetPermission(pluginID int64, targetID uint64, perm string, timestamp int64, dontBroadcast bool) Status {
	return _SetPermission(pluginID, targetID, perm, timestamp, dontBroadcast)
}

var _RemovePermission = func(pluginID int64, targetID uint64, perm string, recursiveDeletion bool, dontBroadcast bool) Status {
	var __retVal Status
	__pluginID := C.int64_t(pluginID)
	__targetID := C.uint64_t(targetID)
	__perm := plugify.ConstructString(perm)
	__recursiveDeletion := C.bool(recursiveDeletion)
	__dontBroadcast := C.bool(dontBroadcast)
	plugify.Block{
		Try: func() {
			__retVal = Status(C.RemovePermission(__pluginID, __targetID, (*C.String)(unsafe.Pointer(&__perm)), __recursiveDeletion, __dontBroadcast))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__perm)
		},
	}.Do()
	return __retVal
}

// RemovePermission
//
//	@brief Remove a permission from a user.
//
//	@param pluginID: Identifier of the plugin that calls the method.
//	@param targetID: Player ID.
//	@param perm: Permission line.
//	@param recursiveDeletion: Delete all nested perms.
//
//	@return Success, TargetUserNotFound, PermNotFound
func RemovePermission(pluginID int64, targetID uint64, perm string, recursiveDeletion bool, dontBroadcast bool) Status {
	return _RemovePermission(pluginID, targetID, perm, recursiveDeletion, dontBroadcast)
}

var _AddGroup = func(pluginID int64, targetID uint64, groupName string, timestamp int64, dontBroadcast bool) Status {
	var __retVal Status
	__pluginID := C.int64_t(pluginID)
	__targetID := C.uint64_t(targetID)
	__groupName := plugify.ConstructString(groupName)
	__timestamp := C.int64_t(timestamp)
	__dontBroadcast := C.bool(dontBroadcast)
	plugify.Block{
		Try: func() {
			__retVal = Status(C.AddGroup(__pluginID, __targetID, (*C.String)(unsafe.Pointer(&__groupName)), __timestamp, __dontBroadcast))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__groupName)
		},
	}.Do()
	return __retVal
}

// AddGroup
//
//	@brief Add a group to a user.
//
//	@param pluginID: Identifier of the plugin that calls the method.
//	@param targetID: Player ID.
//	@param groupName: Group name.
//	@param timestamp: Group duration.
//	@param dontBroadcast: If set to `true`, suppresses dispatching of the group change event to registered UserGroup listeners. The group is still applied internally.
//
//	@return Success, TargetUserNotFound, GroupNotFound, GroupAlreadyExist
func AddGroup(pluginID int64, targetID uint64, groupName string, timestamp int64, dontBroadcast bool) Status {
	return _AddGroup(pluginID, targetID, groupName, timestamp, dontBroadcast)
}

var _RemoveGroup = func(pluginID int64, targetID uint64, groupName string, dontBroadcast bool) Status {
	var __retVal Status
	__pluginID := C.int64_t(pluginID)
	__targetID := C.uint64_t(targetID)
	__groupName := plugify.ConstructString(groupName)
	__dontBroadcast := C.bool(dontBroadcast)
	plugify.Block{
		Try: func() {
			__retVal = Status(C.RemoveGroup(__pluginID, __targetID, (*C.String)(unsafe.Pointer(&__groupName)), __dontBroadcast))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__groupName)
		},
	}.Do()
	return __retVal
}

// RemoveGroup
//
//	@brief Remove a group from a user.
//
//	@param pluginID: Identifier of the plugin that calls the method.
//	@param targetID: Player ID.
//	@param groupName: Group name.
//
//	@return Success, TargetUserNotFound, ChildGroupNotFound, ParentGroupNotFound
func RemoveGroup(pluginID int64, targetID uint64, groupName string, dontBroadcast bool) Status {
	return _RemoveGroup(pluginID, targetID, groupName, dontBroadcast)
}

var _GetCookie = func(targetID uint64, name string, value *any) Status {
	var __retVal Status
	__targetID := C.uint64_t(targetID)
	__name := plugify.ConstructString(name)
	__value := plugify.ConstructVariant(*value)
	plugify.Block{
		Try: func() {
			__retVal = Status(C.GetCookie(__targetID, (*C.String)(unsafe.Pointer(&__name)), (*C.Variant)(unsafe.Pointer(&__value))))
			// Unmarshal - Convert native data to managed data.
			*value = plugify.GetVariantData(&__value)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__name)
			plugify.DestroyVariant(&__value)
		},
	}.Do()
	return __retVal
}

// GetCookie
//
//	@brief Get a cookie value for a user.
//
//	@param targetID: Player ID.
//	@param name: Cookie name.
//	@param value: Cookie value.
//
//	@return Success, TargetUserNotFound, CookieNotFound
func GetCookie(targetID uint64, name string, value *any) Status {
	return _GetCookie(targetID, name, value)
}

var _SetCookie = func(pluginID int64, targetID uint64, name string, cookie any, dontBroadcast bool) Status {
	var __retVal Status
	__pluginID := C.int64_t(pluginID)
	__targetID := C.uint64_t(targetID)
	__name := plugify.ConstructString(name)
	__cookie := plugify.ConstructVariant(cookie)
	__dontBroadcast := C.bool(dontBroadcast)
	plugify.Block{
		Try: func() {
			__retVal = Status(C.SetCookie(__pluginID, __targetID, (*C.String)(unsafe.Pointer(&__name)), (*C.Variant)(unsafe.Pointer(&__cookie)), __dontBroadcast))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__name)
			plugify.DestroyVariant(&__cookie)
		},
	}.Do()
	return __retVal
}

// SetCookie
//
//	@brief Set a cookie value for a user.
//
//	@param pluginID: Identifier of the plugin that calls the method.
//	@param targetID: Player ID.
//	@param name: Cookie name.
//	@param cookie: Cookie value.
//	@param dontBroadcast: If set to `true`, suppresses dispatching of the cookie change event to registered UserSetCookie listeners. The cookie is still applied internally.
//
//	@return Success, TargetUserNotFound
func SetCookie(pluginID int64, targetID uint64, name string, cookie any, dontBroadcast bool) Status {
	return _SetCookie(pluginID, targetID, name, cookie, dontBroadcast)
}

var _GetAllCookies = func(targetID uint64, names *[]string, values *[]any) Status {
	var __retVal Status
	__targetID := C.uint64_t(targetID)
	__names := plugify.ConstructVectorString(*names)
	__values := plugify.ConstructVectorVariant(*values)
	plugify.Block{
		Try: func() {
			__retVal = Status(C.GetAllCookies(__targetID, (*C.Vector)(unsafe.Pointer(&__names)), (*C.Vector)(unsafe.Pointer(&__values))))
			// Unmarshal - Convert native data to managed data.
			plugify.GetVectorDataStringTo(&__names, names)
			plugify.GetVectorDataVariantTo(&__values, values)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorString(&__names)
			plugify.DestroyVectorVariant(&__values)
		},
	}.Do()
	return __retVal
}

// GetAllCookies
//
//	@brief Get all cookies from user.
//
//	@param targetID: Player ID.
//	@param names: Array of cookie names
//	@param values: Array of cookie values
//
//	@return Success, TargetUserNotFound
func GetAllCookies(targetID uint64, names *[]string, values *[]any) Status {
	return _GetAllCookies(targetID, names, values)
}

var _CreateUser = func(pluginID int64, targetID uint64, immunity int32, offline bool, groupsList []string) Status {
	var __retVal Status
	__pluginID := C.int64_t(pluginID)
	__targetID := C.uint64_t(targetID)
	__immunity := C.int32_t(immunity)
	__offline := C.bool(offline)
	__groupsList := plugify.ConstructVectorString(groupsList)
	plugify.Block{
		Try: func() {
			__retVal = Status(C.CreateUser(__pluginID, __targetID, __immunity, __offline, (*C.Vector)(unsafe.Pointer(&__groupsList))))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorString(&__groupsList)
		},
	}.Do()
	return __retVal
}

// CreateUser
//
//	@brief Create a new user.
//
//	@param pluginID: Identifier of the plugin that calls the method.
//	@param targetID: Player ID.
//	@param immunity: User immunity (set -1 to return highest group priority).
//	@param offline: Create as fake player.
//	@param groupsList: Array of groups to inherit ("group timestamp").
//
//	@return Success, UserAlreadyExist, GroupNotFound, ChildGroupNotFound
func CreateUser(pluginID int64, targetID uint64, immunity int32, offline bool, groupsList []string) Status {
	return _CreateUser(pluginID, targetID, immunity, offline, groupsList)
}

var _DeleteUser = func(pluginID int64, targetID uint64) Status {
	var __retVal Status
	__pluginID := C.int64_t(pluginID)
	__targetID := C.uint64_t(targetID)
	__retVal = Status(C.DeleteUser(__pluginID, __targetID))
	return __retVal
}

// DeleteUser
//
//	@brief Delete a user.
//
//	@param pluginID: Identifier of the plugin that calls the method.
//	@param targetID: Player ID.
//
//	@return Success, TargetUserNotFound
func DeleteUser(pluginID int64, targetID uint64) Status {
	return _DeleteUser(pluginID, targetID)
}

var _UserExists = func(targetID uint64) PlayerState {
	var __retVal PlayerState
	__targetID := C.uint64_t(targetID)
	__retVal = PlayerState(C.UserExists(__targetID))
	return __retVal
}

// UserExists
//
//	@brief Check if a user exists.
//
//	@param targetID: Player ID.
//
//	@return PlayerState::NotFound, PlayerState::Online, PlayerState::Offline
func UserExists(targetID uint64) PlayerState {
	return _UserExists(targetID)
}

var _DumpUsersList = func() []uint64 {
	var __retVal []uint64
	var __retVal_native plugify.PlgVector
	plugify.Block{
		Try: func() {
			__native := C.DumpUsersList()
			__retVal_native = *(*plugify.PlgVector)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetVectorDataUInt64[uint64](&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorUInt64(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// DumpUsersList
//
//	@brief Returns a list of IDs for all players registered in the core.
//
//
//	@return A vector containing all registered player IDs.
func DumpUsersList() []uint64 {
	return _DumpUsersList()
}

var _LoadUser = func(pluginID int64, targetID uint64, username string, offline bool, dontBroadcast bool) Status {
	var __retVal Status
	__pluginID := C.int64_t(pluginID)
	__targetID := C.uint64_t(targetID)
	__username := plugify.ConstructString(username)
	__offline := C.bool(offline)
	__dontBroadcast := C.bool(dontBroadcast)
	plugify.Block{
		Try: func() {
			__retVal = Status(C.LoadUser(__pluginID, __targetID, (*C.String)(unsafe.Pointer(&__username)), __offline, __dontBroadcast))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__username)
		},
	}.Do()
	return __retVal
}

// LoadUser
//
//	@brief Dispatches a request to load user data.
//
//	@param pluginID: Identifier of the calling plugin.
//	@param targetID: PlayerID of the user to be loaded.
//	@param username: The user's current username. Intended for synchronizing the username with external storage (e.g. updating an existing record or setting it during initial user creation).
//	@param offline: Indicates whether the user's data was loaded without user presence on server.
func LoadUser(pluginID int64, targetID uint64, username string, offline bool, dontBroadcast bool) Status {
	return _LoadUser(pluginID, targetID, username, offline, dontBroadcast)
}
