#pragma once

#include "shared.h"

extern int32_t (*__permissions_SetParent)(int64_t, String*, String*, bool);

static int32_t SetParent(int64_t pluginID, String* childName, String* parentName, bool dontBroadcast) {
	return __permissions_SetParent(pluginID, childName, parentName, dontBroadcast);
}

extern int32_t (*__permissions_GetParent)(String*, String*);

static int32_t GetParent(String* groupName, String* parentName) {
	return __permissions_GetParent(groupName, parentName);
}

extern int32_t (*__permissions_DumpPermissionsGroup)(String*, Vector*);

static int32_t DumpPermissionsGroup(String* name, Vector* perms) {
	return __permissions_DumpPermissionsGroup(name, perms);
}

extern Vector (*__permissions_GetAllGroups)();

static Vector GetAllGroups() {
	return __permissions_GetAllGroups();
}

extern int32_t (*__permissions_HasPermissionGroupExtended)(String*, String*, bool);

static int32_t HasPermissionGroupExtended(String* name, String* perm, bool exact) {
	return __permissions_HasPermissionGroupExtended(name, perm, exact);
}

extern int32_t (*__permissions_HasPermissionGroup)(String*, String*);

static int32_t HasPermissionGroup(String* name, String* perm) {
	return __permissions_HasPermissionGroup(name, perm);
}

extern int32_t (*__permissions_HasParentGroup)(String*, String*);

static int32_t HasParentGroup(String* childName, String* parentName) {
	return __permissions_HasParentGroup(childName, parentName);
}

extern int32_t (*__permissions_GetPriorityGroup)(String*, int32_t*);

static int32_t GetPriorityGroup(String* groupName, int32_t* priority) {
	return __permissions_GetPriorityGroup(groupName, priority);
}

extern int32_t (*__permissions_AddPermissionGroup)(int64_t, String*, String*, bool);

static int32_t AddPermissionGroup(int64_t pluginID, String* name, String* perm, bool dontBroadcast) {
	return __permissions_AddPermissionGroup(pluginID, name, perm, dontBroadcast);
}

extern int32_t (*__permissions_SetPermissionGroup)(int64_t, String*, String*, bool);

static int32_t SetPermissionGroup(int64_t pluginID, String* name, String* perm, bool dontBroadcast) {
	return __permissions_SetPermissionGroup(pluginID, name, perm, dontBroadcast);
}

extern int32_t (*__permissions_RemovePermissionGroup)(int64_t, String*, String*, bool, bool);

static int32_t RemovePermissionGroup(int64_t pluginID, String* name, String* perm, bool recursiveDeletion, bool dontBroadcast) {
	return __permissions_RemovePermissionGroup(pluginID, name, perm, recursiveDeletion, dontBroadcast);
}

extern int32_t (*__permissions_GetOptionGroup)(String*, String*, Variant*);

static int32_t GetOptionGroup(String* groupName, String* optionName, Variant* value) {
	return __permissions_GetOptionGroup(groupName, optionName, value);
}

extern int32_t (*__permissions_SetOptionGroup)(int64_t, String*, String*, Variant*, bool);

static int32_t SetOptionGroup(int64_t pluginID, String* groupName, String* optionName, Variant* value, bool dontBroadcast) {
	return __permissions_SetOptionGroup(pluginID, groupName, optionName, value, dontBroadcast);
}

extern int32_t (*__permissions_GetAllOptionsGroup)(String*, Vector*, Vector*);

static int32_t GetAllOptionsGroup(String* groupName, Vector* optionNames, Vector* values) {
	return __permissions_GetAllOptionsGroup(groupName, optionNames, values);
}

extern int32_t (*__permissions_CreateGroup)(int64_t, String*, Vector*, int32_t, String*, bool);

static int32_t CreateGroup(int64_t pluginID, String* name, Vector* perms, int32_t priority, String* parent, bool dontBroadcast) {
	return __permissions_CreateGroup(pluginID, name, perms, priority, parent, dontBroadcast);
}

extern int32_t (*__permissions_DeleteGroup)(int64_t, String*, bool);

static int32_t DeleteGroup(int64_t pluginID, String* name, bool dontBroadcast) {
	return __permissions_DeleteGroup(pluginID, name, dontBroadcast);
}

extern bool (*__permissions_GroupExists)(String*);

static bool GroupExists(String* name) {
	return __permissions_GroupExists(name);
}

extern int32_t (*__permissions_LoadGroups)(int64_t, bool);

static int32_t LoadGroups(int64_t pluginID, bool dontBroadcast) {
	return __permissions_LoadGroups(pluginID, dontBroadcast);
}

