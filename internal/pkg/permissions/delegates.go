package permissions

import "github.com/untrustedmodders/go-plugify"

var _ = plugify.ApiVersion

// Generated from permissions

// GroupCreateStorageCallback - Callback invoked after a group is successfully created.
type GroupCreateStorageCallback func(pluginID int64, name string, perms []string, priority int32, parent string) bool

// GroupCreateCallback - Callback invoked after a group is successfully created.
type GroupCreateCallback func(pluginID int64, name string, perms []string, priority int32, parent string)

// GroupDeleteStorageCallback - Callback invoked before a group is deleted.
type GroupDeleteStorageCallback func(pluginID int64, name string) bool

// GroupDeleteCallback - Callback invoked before a group is deleted.
type GroupDeleteCallback func(pluginID int64, name string)

// GroupExpirationCallback - Callback invoked when a group in user has been expired.
type GroupExpirationCallback func(targetID uint64, group string)

// GroupOptionStorageCallback - Callback invoked when an option value is set for a group.
type GroupOptionStorageCallback func(pluginID int64, groupName string, optionName string, value any) bool

// GroupOptionCallback - Callback invoked when an option value is set for a group.
type GroupOptionCallback func(pluginID int64, groupName string, optionName string, value any)

// GroupPermissionStorageCallback - Callback invoked when a permission is added or removed from a group.
type GroupPermissionStorageCallback func(pluginID int64, action Action, groupName string, perm string, oldState Status, newState Status) bool

// GroupPermissionCallback - Callback invoked when a permission is added or removed from a group.
type GroupPermissionCallback func(pluginID int64, action Action, groupName string, perm string, oldState Status, newState Status)

// LoadGroupsCallback - Called when the core requests loading of server groups.
type LoadGroupsCallback func(pluginID int64) bool

// PermExpirationCallback - Callback invoked when a permission in user has been expired.
type PermExpirationCallback func(targetID uint64, perm string, state Status)

// SetParentStorageCallback - Callback invoked when a parent group is set for a child group.
type SetParentStorageCallback func(pluginID int64, childName string, parentName string) bool

// SetParentCallback - Callback invoked when a parent group is set for a child group.
type SetParentCallback func(pluginID int64, childName string, parentName string)

// UserCookieStorageCallback - Callback invoked when a cookie is set for a user.
type UserCookieStorageCallback func(pluginID int64, targetID uint64, name string, cookie any) bool

// UserCookieCallback - Callback invoked when a cookie is set for a user.
type UserCookieCallback func(pluginID int64, targetID uint64, name string, cookie any)

// UserCreateStorageCallback - Callback invoked after a user is successfully created.
type UserCreateStorageCallback func(pluginID int64, targetID uint64, immunity int32, offline bool, groupNames []string) bool

// UserCreateCallback - Callback invoked after a user is successfully created.
type UserCreateCallback func(pluginID int64, targetID uint64, immunity int32, offline bool, groupNames []string)

// UserDeleteStorageCallback - Callback invoked before a user is deleted.
type UserDeleteStorageCallback func(pluginID int64, targetID uint64) bool

// UserDeleteCallback - Callback invoked before a user is deleted.
type UserDeleteCallback func(pluginID int64, targetID uint64)

// UserGroupStorageCallback - Callback invoked when a group is added or removed from a user.
type UserGroupStorageCallback func(pluginID int64, action Action, targetID uint64, group string, oldTimestamp int64, newTimestamp int64) bool

// UserGroupCallback - Callback invoked when a group is added or removed from a user.
type UserGroupCallback func(pluginID int64, action Action, targetID uint64, group string, oldTimestamp int64, newTimestamp int64)

// UserImmunityStorageCallback - Callback invoked when immunity is set for a user.
type UserImmunityStorageCallback func(pluginID int64, targetID uint64, immunity int32) bool

// UserImmunityCallback - Callback invoked when immunity is set for a user.
type UserImmunityCallback func(pluginID int64, targetID uint64, immunity int32)

// UserLoadedCallback - Called when a user's data has been fully loaded.
type UserLoadedCallback func(pluginID int64, targetID uint64, playerState PlayerState)

// UserPermissionStorageCallback - Callback invoked when a permission is added, removed, or replaced for a user.
type UserPermissionStorageCallback func(pluginID int64, action Action, targetID uint64, perm string, oldState Status, newState Status, oldTimestamp int64, newTimestamp int64) bool

// UserPermissionCallback - Callback invoked when a permission is added, removed, or replaced for a user.
type UserPermissionCallback func(pluginID int64, action Action, targetID uint64, perm string, oldState Status, newState Status, oldTimestamp int64, newTimestamp int64)

// UserRequestCallback - Called when a user data load is requested.
type UserRequestCallback func(pluginID int64, targetID uint64, username string, offline bool) bool
