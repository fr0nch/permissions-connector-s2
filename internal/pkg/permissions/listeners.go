package permissions

/*
#include "listeners.h"
#cgo noescape OnGroupCreateStorage_Register
#cgo noescape OnGroupCreateStorage_Unregister
#cgo noescape OnGroupCreate_Register
#cgo noescape OnGroupCreate_Unregister
#cgo noescape OnGroupDeleteStorage_Register
#cgo noescape OnGroupDeleteStorage_Unregister
#cgo noescape OnGroupDelete_Register
#cgo noescape OnGroupDelete_Unregister
#cgo noescape OnGroupExpiration_Register
#cgo noescape OnGroupExpiration_Unregister
#cgo noescape OnGroupOptionChangeStorage_Register
#cgo noescape OnGroupOptionChangeStorage_Unregister
#cgo noescape OnGroupOptionChange_Register
#cgo noescape OnGroupOptionChange_Unregister
#cgo noescape OnGroupPermissionChangeStorage_Register
#cgo noescape OnGroupPermissionChangeStorage_Unregister
#cgo noescape OnGroupPermissionChange_Register
#cgo noescape OnGroupPermissionChange_Unregister
#cgo noescape OnGroupsLoad_Register
#cgo noescape OnGroupsLoad_Unregister
#cgo noescape OnPermissionExpiration_Register
#cgo noescape OnPermissionExpiration_Unregister
#cgo noescape OnSetParentStorage_Register
#cgo noescape OnSetParentStorage_Unregister
#cgo noescape OnSetParent_Register
#cgo noescape OnSetParent_Unregister
#cgo noescape OnUserCookieChangeStorage_Register
#cgo noescape OnUserCookieChangeStorage_Unregister
#cgo noescape OnUserCookieChange_Register
#cgo noescape OnUserCookieChange_Unregister
#cgo noescape OnUserCreateStorage_Register
#cgo noescape OnUserCreateStorage_Unregister
#cgo noescape OnUserCreate_Register
#cgo noescape OnUserCreate_Unregister
#cgo noescape OnUserDeleteStorage_Register
#cgo noescape OnUserDeleteStorage_Unregister
#cgo noescape OnUserDelete_Register
#cgo noescape OnUserDelete_Unregister
#cgo noescape OnUserGroupChangeStorage_Register
#cgo noescape OnUserGroupChangeStorage_Unregister
#cgo noescape OnUserGroupChange_Register
#cgo noescape OnUserGroupChange_Unregister
#cgo noescape OnUserImmunityChangeStorage_Register
#cgo noescape OnUserImmunityChangeStorage_Unregister
#cgo noescape OnUserImmunityChange_Register
#cgo noescape OnUserImmunityChange_Unregister
#cgo noescape OnUserLoaded_Register
#cgo noescape OnUserLoaded_Unregister
#cgo noescape OnUserPermissionChangeStorage_Register
#cgo noescape OnUserPermissionChangeStorage_Unregister
#cgo noescape OnUserPermissionChange_Register
#cgo noescape OnUserPermissionChange_Unregister
#cgo noescape OnUserRequest_Register
#cgo noescape OnUserRequest_Unregister
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

// Generated from permissions (group: listeners)

var _OnGroupCreateStorage_Register = func(callback GroupCreateStorageCallback) Status {
	var __retVal Status
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	__retVal = Status(C.OnGroupCreateStorage_Register(__callback))
	return __retVal
}

// OnGroupCreateStorage_Register
//
//	@brief Registers a listener for the OnGroupCreateStorage event. Callback invoked after a group is successfully created.
//
//	@param callback: The callback to register.
func OnGroupCreateStorage_Register(callback GroupCreateStorageCallback) Status {
	return _OnGroupCreateStorage_Register(callback)
}

var _OnGroupCreateStorage_Unregister = func(callback GroupCreateStorageCallback) Status {
	var __retVal Status
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	__retVal = Status(C.OnGroupCreateStorage_Unregister(__callback))
	return __retVal
}

// OnGroupCreateStorage_Unregister
//
//	@brief Unregisters a previously registered listener for the OnGroupCreateStorage event.
//
//	@param callback: The callback to unregister.
func OnGroupCreateStorage_Unregister(callback GroupCreateStorageCallback) Status {
	return _OnGroupCreateStorage_Unregister(callback)
}

var _OnGroupCreate_Register = func(callback GroupCreateCallback) Status {
	var __retVal Status
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	__retVal = Status(C.OnGroupCreate_Register(__callback))
	return __retVal
}

// OnGroupCreate_Register
//
//	@brief Registers a listener for the OnGroupCreate event. Callback invoked after a group is successfully created.
//
//	@param callback: The callback to register.
func OnGroupCreate_Register(callback GroupCreateCallback) Status {
	return _OnGroupCreate_Register(callback)
}

var _OnGroupCreate_Unregister = func(callback GroupCreateCallback) Status {
	var __retVal Status
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	__retVal = Status(C.OnGroupCreate_Unregister(__callback))
	return __retVal
}

// OnGroupCreate_Unregister
//
//	@brief Unregisters a previously registered listener for the OnGroupCreate event.
//
//	@param callback: The callback to unregister.
func OnGroupCreate_Unregister(callback GroupCreateCallback) Status {
	return _OnGroupCreate_Unregister(callback)
}

var _OnGroupDeleteStorage_Register = func(callback GroupDeleteStorageCallback) Status {
	var __retVal Status
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	__retVal = Status(C.OnGroupDeleteStorage_Register(__callback))
	return __retVal
}

// OnGroupDeleteStorage_Register
//
//	@brief Registers a listener for the OnGroupDeleteStorage event. Callback invoked before a group is deleted.
//
//	@param callback: The callback to register.
func OnGroupDeleteStorage_Register(callback GroupDeleteStorageCallback) Status {
	return _OnGroupDeleteStorage_Register(callback)
}

var _OnGroupDeleteStorage_Unregister = func(callback GroupDeleteStorageCallback) Status {
	var __retVal Status
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	__retVal = Status(C.OnGroupDeleteStorage_Unregister(__callback))
	return __retVal
}

// OnGroupDeleteStorage_Unregister
//
//	@brief Unregisters a previously registered listener for the OnGroupDeleteStorage event.
//
//	@param callback: The callback to unregister.
func OnGroupDeleteStorage_Unregister(callback GroupDeleteStorageCallback) Status {
	return _OnGroupDeleteStorage_Unregister(callback)
}

var _OnGroupDelete_Register = func(callback GroupDeleteCallback) Status {
	var __retVal Status
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	__retVal = Status(C.OnGroupDelete_Register(__callback))
	return __retVal
}

// OnGroupDelete_Register
//
//	@brief Registers a listener for the OnGroupDelete event. Callback invoked before a group is deleted.
//
//	@param callback: The callback to register.
func OnGroupDelete_Register(callback GroupDeleteCallback) Status {
	return _OnGroupDelete_Register(callback)
}

var _OnGroupDelete_Unregister = func(callback GroupDeleteCallback) Status {
	var __retVal Status
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	__retVal = Status(C.OnGroupDelete_Unregister(__callback))
	return __retVal
}

// OnGroupDelete_Unregister
//
//	@brief Unregisters a previously registered listener for the OnGroupDelete event.
//
//	@param callback: The callback to unregister.
func OnGroupDelete_Unregister(callback GroupDeleteCallback) Status {
	return _OnGroupDelete_Unregister(callback)
}

var _OnGroupExpiration_Register = func(callback GroupExpirationCallback) Status {
	var __retVal Status
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	__retVal = Status(C.OnGroupExpiration_Register(__callback))
	return __retVal
}

// OnGroupExpiration_Register
//
//	@brief Registers a listener for the OnGroupExpiration event. Callback invoked when a group in user has been expired.
//
//	@param callback: The callback to register.
func OnGroupExpiration_Register(callback GroupExpirationCallback) Status {
	return _OnGroupExpiration_Register(callback)
}

var _OnGroupExpiration_Unregister = func(callback GroupExpirationCallback) Status {
	var __retVal Status
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	__retVal = Status(C.OnGroupExpiration_Unregister(__callback))
	return __retVal
}

// OnGroupExpiration_Unregister
//
//	@brief Unregisters a previously registered listener for the OnGroupExpiration event.
//
//	@param callback: The callback to unregister.
func OnGroupExpiration_Unregister(callback GroupExpirationCallback) Status {
	return _OnGroupExpiration_Unregister(callback)
}

var _OnGroupOptionChangeStorage_Register = func(callback GroupOptionStorageCallback) Status {
	var __retVal Status
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	__retVal = Status(C.OnGroupOptionChangeStorage_Register(__callback))
	return __retVal
}

// OnGroupOptionChangeStorage_Register
//
//	@brief Registers a listener for the OnGroupOptionChangeStorage event. Callback invoked when an option value is set for a group.
//
//	@param callback: The callback to register.
func OnGroupOptionChangeStorage_Register(callback GroupOptionStorageCallback) Status {
	return _OnGroupOptionChangeStorage_Register(callback)
}

var _OnGroupOptionChangeStorage_Unregister = func(callback GroupOptionStorageCallback) Status {
	var __retVal Status
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	__retVal = Status(C.OnGroupOptionChangeStorage_Unregister(__callback))
	return __retVal
}

// OnGroupOptionChangeStorage_Unregister
//
//	@brief Unregisters a previously registered listener for the OnGroupOptionChangeStorage event.
//
//	@param callback: The callback to unregister.
func OnGroupOptionChangeStorage_Unregister(callback GroupOptionStorageCallback) Status {
	return _OnGroupOptionChangeStorage_Unregister(callback)
}

var _OnGroupOptionChange_Register = func(callback GroupOptionCallback) Status {
	var __retVal Status
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	__retVal = Status(C.OnGroupOptionChange_Register(__callback))
	return __retVal
}

// OnGroupOptionChange_Register
//
//	@brief Registers a listener for the OnGroupOptionChange event. Callback invoked when an option value is set for a group.
//
//	@param callback: The callback to register.
func OnGroupOptionChange_Register(callback GroupOptionCallback) Status {
	return _OnGroupOptionChange_Register(callback)
}

var _OnGroupOptionChange_Unregister = func(callback GroupOptionCallback) Status {
	var __retVal Status
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	__retVal = Status(C.OnGroupOptionChange_Unregister(__callback))
	return __retVal
}

// OnGroupOptionChange_Unregister
//
//	@brief Unregisters a previously registered listener for the OnGroupOptionChange event.
//
//	@param callback: The callback to unregister.
func OnGroupOptionChange_Unregister(callback GroupOptionCallback) Status {
	return _OnGroupOptionChange_Unregister(callback)
}

var _OnGroupPermissionChangeStorage_Register = func(callback GroupPermissionStorageCallback) Status {
	var __retVal Status
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	__retVal = Status(C.OnGroupPermissionChangeStorage_Register(__callback))
	return __retVal
}

// OnGroupPermissionChangeStorage_Register
//
//	@brief Registers a listener for the OnGroupPermissionChangeStorage event. Callback invoked when a permission is added or removed from a group.
//
//	@param callback: The callback to register.
func OnGroupPermissionChangeStorage_Register(callback GroupPermissionStorageCallback) Status {
	return _OnGroupPermissionChangeStorage_Register(callback)
}

var _OnGroupPermissionChangeStorage_Unregister = func(callback GroupPermissionStorageCallback) Status {
	var __retVal Status
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	__retVal = Status(C.OnGroupPermissionChangeStorage_Unregister(__callback))
	return __retVal
}

// OnGroupPermissionChangeStorage_Unregister
//
//	@brief Unregisters a previously registered listener for the OnGroupPermissionChangeStorage event.
//
//	@param callback: The callback to unregister.
func OnGroupPermissionChangeStorage_Unregister(callback GroupPermissionStorageCallback) Status {
	return _OnGroupPermissionChangeStorage_Unregister(callback)
}

var _OnGroupPermissionChange_Register = func(callback GroupPermissionCallback) Status {
	var __retVal Status
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	__retVal = Status(C.OnGroupPermissionChange_Register(__callback))
	return __retVal
}

// OnGroupPermissionChange_Register
//
//	@brief Registers a listener for the OnGroupPermissionChange event. Callback invoked when a permission is added or removed from a group.
//
//	@param callback: The callback to register.
func OnGroupPermissionChange_Register(callback GroupPermissionCallback) Status {
	return _OnGroupPermissionChange_Register(callback)
}

var _OnGroupPermissionChange_Unregister = func(callback GroupPermissionCallback) Status {
	var __retVal Status
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	__retVal = Status(C.OnGroupPermissionChange_Unregister(__callback))
	return __retVal
}

// OnGroupPermissionChange_Unregister
//
//	@brief Unregisters a previously registered listener for the OnGroupPermissionChange event.
//
//	@param callback: The callback to unregister.
func OnGroupPermissionChange_Unregister(callback GroupPermissionCallback) Status {
	return _OnGroupPermissionChange_Unregister(callback)
}

var _OnGroupsLoad_Register = func(callback LoadGroupsCallback) Status {
	var __retVal Status
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	__retVal = Status(C.OnGroupsLoad_Register(__callback))
	return __retVal
}

// OnGroupsLoad_Register
//
//	@brief Registers a listener for the OnGroupsLoad event. Called when the core requests loading of server groups.
//
//	@param callback: The callback to register.
func OnGroupsLoad_Register(callback LoadGroupsCallback) Status {
	return _OnGroupsLoad_Register(callback)
}

var _OnGroupsLoad_Unregister = func(callback LoadGroupsCallback) Status {
	var __retVal Status
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	__retVal = Status(C.OnGroupsLoad_Unregister(__callback))
	return __retVal
}

// OnGroupsLoad_Unregister
//
//	@brief Unregisters a previously registered listener for the OnGroupsLoad event.
//
//	@param callback: The callback to unregister.
func OnGroupsLoad_Unregister(callback LoadGroupsCallback) Status {
	return _OnGroupsLoad_Unregister(callback)
}

var _OnPermissionExpiration_Register = func(callback PermExpirationCallback) Status {
	var __retVal Status
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	__retVal = Status(C.OnPermissionExpiration_Register(__callback))
	return __retVal
}

// OnPermissionExpiration_Register
//
//	@brief Registers a listener for the OnPermissionExpiration event. Callback invoked when a permission in user has been expired.
//
//	@param callback: The callback to register.
func OnPermissionExpiration_Register(callback PermExpirationCallback) Status {
	return _OnPermissionExpiration_Register(callback)
}

var _OnPermissionExpiration_Unregister = func(callback PermExpirationCallback) Status {
	var __retVal Status
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	__retVal = Status(C.OnPermissionExpiration_Unregister(__callback))
	return __retVal
}

// OnPermissionExpiration_Unregister
//
//	@brief Unregisters a previously registered listener for the OnPermissionExpiration event.
//
//	@param callback: The callback to unregister.
func OnPermissionExpiration_Unregister(callback PermExpirationCallback) Status {
	return _OnPermissionExpiration_Unregister(callback)
}

var _OnSetParentStorage_Register = func(callback SetParentStorageCallback) Status {
	var __retVal Status
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	__retVal = Status(C.OnSetParentStorage_Register(__callback))
	return __retVal
}

// OnSetParentStorage_Register
//
//	@brief Registers a listener for the OnSetParentStorage event. Callback invoked when a parent group is set for a child group.
//
//	@param callback: The callback to register.
func OnSetParentStorage_Register(callback SetParentStorageCallback) Status {
	return _OnSetParentStorage_Register(callback)
}

var _OnSetParentStorage_Unregister = func(callback SetParentStorageCallback) Status {
	var __retVal Status
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	__retVal = Status(C.OnSetParentStorage_Unregister(__callback))
	return __retVal
}

// OnSetParentStorage_Unregister
//
//	@brief Unregisters a previously registered listener for the OnSetParentStorage event.
//
//	@param callback: The callback to unregister.
func OnSetParentStorage_Unregister(callback SetParentStorageCallback) Status {
	return _OnSetParentStorage_Unregister(callback)
}

var _OnSetParent_Register = func(callback SetParentCallback) Status {
	var __retVal Status
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	__retVal = Status(C.OnSetParent_Register(__callback))
	return __retVal
}

// OnSetParent_Register
//
//	@brief Registers a listener for the OnSetParent event. Callback invoked when a parent group is set for a child group.
//
//	@param callback: The callback to register.
func OnSetParent_Register(callback SetParentCallback) Status {
	return _OnSetParent_Register(callback)
}

var _OnSetParent_Unregister = func(callback SetParentCallback) Status {
	var __retVal Status
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	__retVal = Status(C.OnSetParent_Unregister(__callback))
	return __retVal
}

// OnSetParent_Unregister
//
//	@brief Unregisters a previously registered listener for the OnSetParent event.
//
//	@param callback: The callback to unregister.
func OnSetParent_Unregister(callback SetParentCallback) Status {
	return _OnSetParent_Unregister(callback)
}

var _OnUserCookieChangeStorage_Register = func(callback UserCookieStorageCallback) Status {
	var __retVal Status
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	__retVal = Status(C.OnUserCookieChangeStorage_Register(__callback))
	return __retVal
}

// OnUserCookieChangeStorage_Register
//
//	@brief Registers a listener for the OnUserCookieChangeStorage event. Callback invoked when a cookie is set for a user.
//
//	@param callback: The callback to register.
func OnUserCookieChangeStorage_Register(callback UserCookieStorageCallback) Status {
	return _OnUserCookieChangeStorage_Register(callback)
}

var _OnUserCookieChangeStorage_Unregister = func(callback UserCookieStorageCallback) Status {
	var __retVal Status
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	__retVal = Status(C.OnUserCookieChangeStorage_Unregister(__callback))
	return __retVal
}

// OnUserCookieChangeStorage_Unregister
//
//	@brief Unregisters a previously registered listener for the OnUserCookieChangeStorage event.
//
//	@param callback: The callback to unregister.
func OnUserCookieChangeStorage_Unregister(callback UserCookieStorageCallback) Status {
	return _OnUserCookieChangeStorage_Unregister(callback)
}

var _OnUserCookieChange_Register = func(callback UserCookieCallback) Status {
	var __retVal Status
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	__retVal = Status(C.OnUserCookieChange_Register(__callback))
	return __retVal
}

// OnUserCookieChange_Register
//
//	@brief Registers a listener for the OnUserCookieChange event. Callback invoked when a cookie is set for a user.
//
//	@param callback: The callback to register.
func OnUserCookieChange_Register(callback UserCookieCallback) Status {
	return _OnUserCookieChange_Register(callback)
}

var _OnUserCookieChange_Unregister = func(callback UserCookieCallback) Status {
	var __retVal Status
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	__retVal = Status(C.OnUserCookieChange_Unregister(__callback))
	return __retVal
}

// OnUserCookieChange_Unregister
//
//	@brief Unregisters a previously registered listener for the OnUserCookieChange event.
//
//	@param callback: The callback to unregister.
func OnUserCookieChange_Unregister(callback UserCookieCallback) Status {
	return _OnUserCookieChange_Unregister(callback)
}

var _OnUserCreateStorage_Register = func(callback UserCreateStorageCallback) Status {
	var __retVal Status
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	__retVal = Status(C.OnUserCreateStorage_Register(__callback))
	return __retVal
}

// OnUserCreateStorage_Register
//
//	@brief Registers a listener for the OnUserCreateStorage event. Callback invoked after a user is successfully created.
//
//	@param callback: The callback to register.
func OnUserCreateStorage_Register(callback UserCreateStorageCallback) Status {
	return _OnUserCreateStorage_Register(callback)
}

var _OnUserCreateStorage_Unregister = func(callback UserCreateStorageCallback) Status {
	var __retVal Status
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	__retVal = Status(C.OnUserCreateStorage_Unregister(__callback))
	return __retVal
}

// OnUserCreateStorage_Unregister
//
//	@brief Unregisters a previously registered listener for the OnUserCreateStorage event.
//
//	@param callback: The callback to unregister.
func OnUserCreateStorage_Unregister(callback UserCreateStorageCallback) Status {
	return _OnUserCreateStorage_Unregister(callback)
}

var _OnUserCreate_Register = func(callback UserCreateCallback) Status {
	var __retVal Status
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	__retVal = Status(C.OnUserCreate_Register(__callback))
	return __retVal
}

// OnUserCreate_Register
//
//	@brief Registers a listener for the OnUserCreate event. Callback invoked after a user is successfully created.
//
//	@param callback: The callback to register.
func OnUserCreate_Register(callback UserCreateCallback) Status {
	return _OnUserCreate_Register(callback)
}

var _OnUserCreate_Unregister = func(callback UserCreateCallback) Status {
	var __retVal Status
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	__retVal = Status(C.OnUserCreate_Unregister(__callback))
	return __retVal
}

// OnUserCreate_Unregister
//
//	@brief Unregisters a previously registered listener for the OnUserCreate event.
//
//	@param callback: The callback to unregister.
func OnUserCreate_Unregister(callback UserCreateCallback) Status {
	return _OnUserCreate_Unregister(callback)
}

var _OnUserDeleteStorage_Register = func(callback UserDeleteStorageCallback) Status {
	var __retVal Status
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	__retVal = Status(C.OnUserDeleteStorage_Register(__callback))
	return __retVal
}

// OnUserDeleteStorage_Register
//
//	@brief Registers a listener for the OnUserDeleteStorage event. Callback invoked before a user is deleted.
//
//	@param callback: The callback to register.
func OnUserDeleteStorage_Register(callback UserDeleteStorageCallback) Status {
	return _OnUserDeleteStorage_Register(callback)
}

var _OnUserDeleteStorage_Unregister = func(callback UserDeleteStorageCallback) Status {
	var __retVal Status
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	__retVal = Status(C.OnUserDeleteStorage_Unregister(__callback))
	return __retVal
}

// OnUserDeleteStorage_Unregister
//
//	@brief Unregisters a previously registered listener for the OnUserDeleteStorage event.
//
//	@param callback: The callback to unregister.
func OnUserDeleteStorage_Unregister(callback UserDeleteStorageCallback) Status {
	return _OnUserDeleteStorage_Unregister(callback)
}

var _OnUserDelete_Register = func(callback UserDeleteCallback) Status {
	var __retVal Status
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	__retVal = Status(C.OnUserDelete_Register(__callback))
	return __retVal
}

// OnUserDelete_Register
//
//	@brief Registers a listener for the OnUserDelete event. Callback invoked before a user is deleted.
//
//	@param callback: The callback to register.
func OnUserDelete_Register(callback UserDeleteCallback) Status {
	return _OnUserDelete_Register(callback)
}

var _OnUserDelete_Unregister = func(callback UserDeleteCallback) Status {
	var __retVal Status
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	__retVal = Status(C.OnUserDelete_Unregister(__callback))
	return __retVal
}

// OnUserDelete_Unregister
//
//	@brief Unregisters a previously registered listener for the OnUserDelete event.
//
//	@param callback: The callback to unregister.
func OnUserDelete_Unregister(callback UserDeleteCallback) Status {
	return _OnUserDelete_Unregister(callback)
}

var _OnUserGroupChangeStorage_Register = func(callback UserGroupStorageCallback) Status {
	var __retVal Status
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	__retVal = Status(C.OnUserGroupChangeStorage_Register(__callback))
	return __retVal
}

// OnUserGroupChangeStorage_Register
//
//	@brief Registers a listener for the OnUserGroupChangeStorage event. Callback invoked when a group is added or removed from a user.
//
//	@param callback: The callback to register.
func OnUserGroupChangeStorage_Register(callback UserGroupStorageCallback) Status {
	return _OnUserGroupChangeStorage_Register(callback)
}

var _OnUserGroupChangeStorage_Unregister = func(callback UserGroupStorageCallback) Status {
	var __retVal Status
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	__retVal = Status(C.OnUserGroupChangeStorage_Unregister(__callback))
	return __retVal
}

// OnUserGroupChangeStorage_Unregister
//
//	@brief Unregisters a previously registered listener for the OnUserGroupChangeStorage event.
//
//	@param callback: The callback to unregister.
func OnUserGroupChangeStorage_Unregister(callback UserGroupStorageCallback) Status {
	return _OnUserGroupChangeStorage_Unregister(callback)
}

var _OnUserGroupChange_Register = func(callback UserGroupCallback) Status {
	var __retVal Status
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	__retVal = Status(C.OnUserGroupChange_Register(__callback))
	return __retVal
}

// OnUserGroupChange_Register
//
//	@brief Registers a listener for the OnUserGroupChange event. Callback invoked when a group is added or removed from a user.
//
//	@param callback: The callback to register.
func OnUserGroupChange_Register(callback UserGroupCallback) Status {
	return _OnUserGroupChange_Register(callback)
}

var _OnUserGroupChange_Unregister = func(callback UserGroupCallback) Status {
	var __retVal Status
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	__retVal = Status(C.OnUserGroupChange_Unregister(__callback))
	return __retVal
}

// OnUserGroupChange_Unregister
//
//	@brief Unregisters a previously registered listener for the OnUserGroupChange event.
//
//	@param callback: The callback to unregister.
func OnUserGroupChange_Unregister(callback UserGroupCallback) Status {
	return _OnUserGroupChange_Unregister(callback)
}

var _OnUserImmunityChangeStorage_Register = func(callback UserImmunityStorageCallback) Status {
	var __retVal Status
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	__retVal = Status(C.OnUserImmunityChangeStorage_Register(__callback))
	return __retVal
}

// OnUserImmunityChangeStorage_Register
//
//	@brief Registers a listener for the OnUserImmunityChangeStorage event. Callback invoked when immunity is set for a user.
//
//	@param callback: The callback to register.
func OnUserImmunityChangeStorage_Register(callback UserImmunityStorageCallback) Status {
	return _OnUserImmunityChangeStorage_Register(callback)
}

var _OnUserImmunityChangeStorage_Unregister = func(callback UserImmunityStorageCallback) Status {
	var __retVal Status
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	__retVal = Status(C.OnUserImmunityChangeStorage_Unregister(__callback))
	return __retVal
}

// OnUserImmunityChangeStorage_Unregister
//
//	@brief Unregisters a previously registered listener for the OnUserImmunityChangeStorage event.
//
//	@param callback: The callback to unregister.
func OnUserImmunityChangeStorage_Unregister(callback UserImmunityStorageCallback) Status {
	return _OnUserImmunityChangeStorage_Unregister(callback)
}

var _OnUserImmunityChange_Register = func(callback UserImmunityCallback) Status {
	var __retVal Status
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	__retVal = Status(C.OnUserImmunityChange_Register(__callback))
	return __retVal
}

// OnUserImmunityChange_Register
//
//	@brief Registers a listener for the OnUserImmunityChange event. Callback invoked when immunity is set for a user.
//
//	@param callback: The callback to register.
func OnUserImmunityChange_Register(callback UserImmunityCallback) Status {
	return _OnUserImmunityChange_Register(callback)
}

var _OnUserImmunityChange_Unregister = func(callback UserImmunityCallback) Status {
	var __retVal Status
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	__retVal = Status(C.OnUserImmunityChange_Unregister(__callback))
	return __retVal
}

// OnUserImmunityChange_Unregister
//
//	@brief Unregisters a previously registered listener for the OnUserImmunityChange event.
//
//	@param callback: The callback to unregister.
func OnUserImmunityChange_Unregister(callback UserImmunityCallback) Status {
	return _OnUserImmunityChange_Unregister(callback)
}

var _OnUserLoaded_Register = func(callback UserLoadedCallback) Status {
	var __retVal Status
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	__retVal = Status(C.OnUserLoaded_Register(__callback))
	return __retVal
}

// OnUserLoaded_Register
//
//	@brief Registers a listener for the OnUserLoaded event. Called when a user's data has been fully loaded.
//
//	@param callback: The callback to register.
func OnUserLoaded_Register(callback UserLoadedCallback) Status {
	return _OnUserLoaded_Register(callback)
}

var _OnUserLoaded_Unregister = func(callback UserLoadedCallback) Status {
	var __retVal Status
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	__retVal = Status(C.OnUserLoaded_Unregister(__callback))
	return __retVal
}

// OnUserLoaded_Unregister
//
//	@brief Unregisters a previously registered listener for the OnUserLoaded event.
//
//	@param callback: The callback to unregister.
func OnUserLoaded_Unregister(callback UserLoadedCallback) Status {
	return _OnUserLoaded_Unregister(callback)
}

var _OnUserPermissionChangeStorage_Register = func(callback UserPermissionStorageCallback) Status {
	var __retVal Status
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	__retVal = Status(C.OnUserPermissionChangeStorage_Register(__callback))
	return __retVal
}

// OnUserPermissionChangeStorage_Register
//
//	@brief Registers a listener for the OnUserPermissionChangeStorage event. Callback invoked when a permission is added, removed, or replaced for a user.
//
//	@param callback: The callback to register.
func OnUserPermissionChangeStorage_Register(callback UserPermissionStorageCallback) Status {
	return _OnUserPermissionChangeStorage_Register(callback)
}

var _OnUserPermissionChangeStorage_Unregister = func(callback UserPermissionStorageCallback) Status {
	var __retVal Status
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	__retVal = Status(C.OnUserPermissionChangeStorage_Unregister(__callback))
	return __retVal
}

// OnUserPermissionChangeStorage_Unregister
//
//	@brief Unregisters a previously registered listener for the OnUserPermissionChangeStorage event.
//
//	@param callback: The callback to unregister.
func OnUserPermissionChangeStorage_Unregister(callback UserPermissionStorageCallback) Status {
	return _OnUserPermissionChangeStorage_Unregister(callback)
}

var _OnUserPermissionChange_Register = func(callback UserPermissionCallback) Status {
	var __retVal Status
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	__retVal = Status(C.OnUserPermissionChange_Register(__callback))
	return __retVal
}

// OnUserPermissionChange_Register
//
//	@brief Registers a listener for the OnUserPermissionChange event. Callback invoked when a permission is added, removed, or replaced for a user.
//
//	@param callback: The callback to register.
func OnUserPermissionChange_Register(callback UserPermissionCallback) Status {
	return _OnUserPermissionChange_Register(callback)
}

var _OnUserPermissionChange_Unregister = func(callback UserPermissionCallback) Status {
	var __retVal Status
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	__retVal = Status(C.OnUserPermissionChange_Unregister(__callback))
	return __retVal
}

// OnUserPermissionChange_Unregister
//
//	@brief Unregisters a previously registered listener for the OnUserPermissionChange event.
//
//	@param callback: The callback to unregister.
func OnUserPermissionChange_Unregister(callback UserPermissionCallback) Status {
	return _OnUserPermissionChange_Unregister(callback)
}

var _OnUserRequest_Register = func(callback UserRequestCallback) Status {
	var __retVal Status
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	__retVal = Status(C.OnUserRequest_Register(__callback))
	return __retVal
}

// OnUserRequest_Register
//
//	@brief Registers a listener for the OnUserRequest event. Called when a user data load is requested.
//
//	@param callback: The callback to register.
func OnUserRequest_Register(callback UserRequestCallback) Status {
	return _OnUserRequest_Register(callback)
}

var _OnUserRequest_Unregister = func(callback UserRequestCallback) Status {
	var __retVal Status
	__callback := plugify.GetFunctionPointerForDelegate(callback)
	__retVal = Status(C.OnUserRequest_Unregister(__callback))
	return __retVal
}

// OnUserRequest_Unregister
//
//	@brief Unregisters a previously registered listener for the OnUserRequest event.
//
//	@param callback: The callback to unregister.
func OnUserRequest_Unregister(callback UserRequestCallback) Status {
	return _OnUserRequest_Unregister(callback)
}
