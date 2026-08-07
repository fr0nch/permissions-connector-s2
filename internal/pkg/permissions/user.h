#pragma once

#include "shared.h"

extern int32_t (*__permissions_DumpPermissions)(uint64_t, Vector*);

static int32_t DumpPermissions(uint64_t targetID, Vector* perms) {
	return __permissions_DumpPermissions(targetID, perms);
}

extern int32_t (*__permissions_CanAffectUser)(uint64_t, uint64_t);

static int32_t CanAffectUser(uint64_t actorID, uint64_t targetID) {
	return __permissions_CanAffectUser(actorID, targetID);
}

extern int32_t (*__permissions_HasPermissionExtended)(uint64_t, String*, bool, uint32_t*, int64_t*);

static int32_t HasPermissionExtended(uint64_t targetID, String* perm, bool exact, uint32_t* permSource, int64_t* timestamp) {
	return __permissions_HasPermissionExtended(targetID, perm, exact, permSource, timestamp);
}

extern int32_t (*__permissions_HasPermission)(uint64_t, String*);

static int32_t HasPermission(uint64_t targetID, String* perm) {
	return __permissions_HasPermission(targetID, perm);
}

extern int32_t (*__permissions_HasGroupExtended)(uint64_t, String*, int64_t*);

static int32_t HasGroupExtended(uint64_t targetID, String* groupName, int64_t* timestamp) {
	return __permissions_HasGroupExtended(targetID, groupName, timestamp);
}

extern int32_t (*__permissions_HasGroup)(uint64_t, String*);

static int32_t HasGroup(uint64_t targetID, String* groupName) {
	return __permissions_HasGroup(targetID, groupName);
}

extern int32_t (*__permissions_GetUserGroups)(uint64_t, Vector*);

static int32_t GetUserGroups(uint64_t targetID, Vector* outGroups) {
	return __permissions_GetUserGroups(targetID, outGroups);
}

extern int32_t (*__permissions_GetImmunity)(uint64_t, int32_t*);

static int32_t GetImmunity(uint64_t targetID, int32_t* immunity) {
	return __permissions_GetImmunity(targetID, immunity);
}

extern int32_t (*__permissions_SetImmunity)(int64_t, uint64_t, int32_t, bool);

static int32_t SetImmunity(int64_t pluginID, uint64_t targetID, int32_t immunity, bool dontBroadcast) {
	return __permissions_SetImmunity(pluginID, targetID, immunity, dontBroadcast);
}

extern int32_t (*__permissions_AddPermission)(int64_t, uint64_t, String*, int64_t, bool);

static int32_t AddPermission(int64_t pluginID, uint64_t targetID, String* perm, int64_t timestamp, bool dontBroadcast) {
	return __permissions_AddPermission(pluginID, targetID, perm, timestamp, dontBroadcast);
}

extern int32_t (*__permissions_SetPermission)(int64_t, uint64_t, String*, int64_t, bool);

static int32_t SetPermission(int64_t pluginID, uint64_t targetID, String* perm, int64_t timestamp, bool dontBroadcast) {
	return __permissions_SetPermission(pluginID, targetID, perm, timestamp, dontBroadcast);
}

extern int32_t (*__permissions_RemovePermission)(int64_t, uint64_t, String*, bool, bool);

static int32_t RemovePermission(int64_t pluginID, uint64_t targetID, String* perm, bool recursiveDeletion, bool dontBroadcast) {
	return __permissions_RemovePermission(pluginID, targetID, perm, recursiveDeletion, dontBroadcast);
}

extern int32_t (*__permissions_AddGroup)(int64_t, uint64_t, String*, int64_t, bool);

static int32_t AddGroup(int64_t pluginID, uint64_t targetID, String* groupName, int64_t timestamp, bool dontBroadcast) {
	return __permissions_AddGroup(pluginID, targetID, groupName, timestamp, dontBroadcast);
}

extern int32_t (*__permissions_RemoveGroup)(int64_t, uint64_t, String*, bool);

static int32_t RemoveGroup(int64_t pluginID, uint64_t targetID, String* groupName, bool dontBroadcast) {
	return __permissions_RemoveGroup(pluginID, targetID, groupName, dontBroadcast);
}

extern int32_t (*__permissions_GetCookie)(uint64_t, String*, Variant*);

static int32_t GetCookie(uint64_t targetID, String* name, Variant* value) {
	return __permissions_GetCookie(targetID, name, value);
}

extern int32_t (*__permissions_SetCookie)(int64_t, uint64_t, String*, Variant*, bool);

static int32_t SetCookie(int64_t pluginID, uint64_t targetID, String* name, Variant* cookie, bool dontBroadcast) {
	return __permissions_SetCookie(pluginID, targetID, name, cookie, dontBroadcast);
}

extern int32_t (*__permissions_GetAllCookies)(uint64_t, Vector*, Vector*);

static int32_t GetAllCookies(uint64_t targetID, Vector* names, Vector* values) {
	return __permissions_GetAllCookies(targetID, names, values);
}

extern int32_t (*__permissions_CreateUser)(int64_t, uint64_t, int32_t, bool, Vector*);

static int32_t CreateUser(int64_t pluginID, uint64_t targetID, int32_t immunity, bool offline, Vector* groupsList) {
	return __permissions_CreateUser(pluginID, targetID, immunity, offline, groupsList);
}

extern int32_t (*__permissions_DeleteUser)(int64_t, uint64_t);

static int32_t DeleteUser(int64_t pluginID, uint64_t targetID) {
	return __permissions_DeleteUser(pluginID, targetID);
}

extern uint32_t (*__permissions_UserExists)(uint64_t);

static uint32_t UserExists(uint64_t targetID) {
	return __permissions_UserExists(targetID);
}

extern Vector (*__permissions_DumpUsersList)();

static Vector DumpUsersList() {
	return __permissions_DumpUsersList();
}

extern int32_t (*__permissions_LoadUser)(int64_t, uint64_t, String*, bool, bool);

static int32_t LoadUser(int64_t pluginID, uint64_t targetID, String* username, bool offline, bool dontBroadcast) {
	return __permissions_LoadUser(pluginID, targetID, username, offline, dontBroadcast);
}

