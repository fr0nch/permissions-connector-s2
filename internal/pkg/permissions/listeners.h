#pragma once

#include "shared.h"

extern int32_t (*__permissions_OnGroupCreateStorage_Register)(void*);

static int32_t OnGroupCreateStorage_Register(void* callback) {
	return __permissions_OnGroupCreateStorage_Register(callback);
}

extern int32_t (*__permissions_OnGroupCreateStorage_Unregister)(void*);

static int32_t OnGroupCreateStorage_Unregister(void* callback) {
	return __permissions_OnGroupCreateStorage_Unregister(callback);
}

extern int32_t (*__permissions_OnGroupCreate_Register)(void*);

static int32_t OnGroupCreate_Register(void* callback) {
	return __permissions_OnGroupCreate_Register(callback);
}

extern int32_t (*__permissions_OnGroupCreate_Unregister)(void*);

static int32_t OnGroupCreate_Unregister(void* callback) {
	return __permissions_OnGroupCreate_Unregister(callback);
}

extern int32_t (*__permissions_OnGroupDeleteStorage_Register)(void*);

static int32_t OnGroupDeleteStorage_Register(void* callback) {
	return __permissions_OnGroupDeleteStorage_Register(callback);
}

extern int32_t (*__permissions_OnGroupDeleteStorage_Unregister)(void*);

static int32_t OnGroupDeleteStorage_Unregister(void* callback) {
	return __permissions_OnGroupDeleteStorage_Unregister(callback);
}

extern int32_t (*__permissions_OnGroupDelete_Register)(void*);

static int32_t OnGroupDelete_Register(void* callback) {
	return __permissions_OnGroupDelete_Register(callback);
}

extern int32_t (*__permissions_OnGroupDelete_Unregister)(void*);

static int32_t OnGroupDelete_Unregister(void* callback) {
	return __permissions_OnGroupDelete_Unregister(callback);
}

extern int32_t (*__permissions_OnGroupExpiration_Register)(void*);

static int32_t OnGroupExpiration_Register(void* callback) {
	return __permissions_OnGroupExpiration_Register(callback);
}

extern int32_t (*__permissions_OnGroupExpiration_Unregister)(void*);

static int32_t OnGroupExpiration_Unregister(void* callback) {
	return __permissions_OnGroupExpiration_Unregister(callback);
}

extern int32_t (*__permissions_OnGroupOptionChangeStorage_Register)(void*);

static int32_t OnGroupOptionChangeStorage_Register(void* callback) {
	return __permissions_OnGroupOptionChangeStorage_Register(callback);
}

extern int32_t (*__permissions_OnGroupOptionChangeStorage_Unregister)(void*);

static int32_t OnGroupOptionChangeStorage_Unregister(void* callback) {
	return __permissions_OnGroupOptionChangeStorage_Unregister(callback);
}

extern int32_t (*__permissions_OnGroupOptionChange_Register)(void*);

static int32_t OnGroupOptionChange_Register(void* callback) {
	return __permissions_OnGroupOptionChange_Register(callback);
}

extern int32_t (*__permissions_OnGroupOptionChange_Unregister)(void*);

static int32_t OnGroupOptionChange_Unregister(void* callback) {
	return __permissions_OnGroupOptionChange_Unregister(callback);
}

extern int32_t (*__permissions_OnGroupPermissionChangeStorage_Register)(void*);

static int32_t OnGroupPermissionChangeStorage_Register(void* callback) {
	return __permissions_OnGroupPermissionChangeStorage_Register(callback);
}

extern int32_t (*__permissions_OnGroupPermissionChangeStorage_Unregister)(void*);

static int32_t OnGroupPermissionChangeStorage_Unregister(void* callback) {
	return __permissions_OnGroupPermissionChangeStorage_Unregister(callback);
}

extern int32_t (*__permissions_OnGroupPermissionChange_Register)(void*);

static int32_t OnGroupPermissionChange_Register(void* callback) {
	return __permissions_OnGroupPermissionChange_Register(callback);
}

extern int32_t (*__permissions_OnGroupPermissionChange_Unregister)(void*);

static int32_t OnGroupPermissionChange_Unregister(void* callback) {
	return __permissions_OnGroupPermissionChange_Unregister(callback);
}

extern int32_t (*__permissions_OnGroupsLoad_Register)(void*);

static int32_t OnGroupsLoad_Register(void* callback) {
	return __permissions_OnGroupsLoad_Register(callback);
}

extern int32_t (*__permissions_OnGroupsLoad_Unregister)(void*);

static int32_t OnGroupsLoad_Unregister(void* callback) {
	return __permissions_OnGroupsLoad_Unregister(callback);
}

extern int32_t (*__permissions_OnPermissionExpiration_Register)(void*);

static int32_t OnPermissionExpiration_Register(void* callback) {
	return __permissions_OnPermissionExpiration_Register(callback);
}

extern int32_t (*__permissions_OnPermissionExpiration_Unregister)(void*);

static int32_t OnPermissionExpiration_Unregister(void* callback) {
	return __permissions_OnPermissionExpiration_Unregister(callback);
}

extern int32_t (*__permissions_OnSetParentStorage_Register)(void*);

static int32_t OnSetParentStorage_Register(void* callback) {
	return __permissions_OnSetParentStorage_Register(callback);
}

extern int32_t (*__permissions_OnSetParentStorage_Unregister)(void*);

static int32_t OnSetParentStorage_Unregister(void* callback) {
	return __permissions_OnSetParentStorage_Unregister(callback);
}

extern int32_t (*__permissions_OnSetParent_Register)(void*);

static int32_t OnSetParent_Register(void* callback) {
	return __permissions_OnSetParent_Register(callback);
}

extern int32_t (*__permissions_OnSetParent_Unregister)(void*);

static int32_t OnSetParent_Unregister(void* callback) {
	return __permissions_OnSetParent_Unregister(callback);
}

extern int32_t (*__permissions_OnUserCookieChangeStorage_Register)(void*);

static int32_t OnUserCookieChangeStorage_Register(void* callback) {
	return __permissions_OnUserCookieChangeStorage_Register(callback);
}

extern int32_t (*__permissions_OnUserCookieChangeStorage_Unregister)(void*);

static int32_t OnUserCookieChangeStorage_Unregister(void* callback) {
	return __permissions_OnUserCookieChangeStorage_Unregister(callback);
}

extern int32_t (*__permissions_OnUserCookieChange_Register)(void*);

static int32_t OnUserCookieChange_Register(void* callback) {
	return __permissions_OnUserCookieChange_Register(callback);
}

extern int32_t (*__permissions_OnUserCookieChange_Unregister)(void*);

static int32_t OnUserCookieChange_Unregister(void* callback) {
	return __permissions_OnUserCookieChange_Unregister(callback);
}

extern int32_t (*__permissions_OnUserCreateStorage_Register)(void*);

static int32_t OnUserCreateStorage_Register(void* callback) {
	return __permissions_OnUserCreateStorage_Register(callback);
}

extern int32_t (*__permissions_OnUserCreateStorage_Unregister)(void*);

static int32_t OnUserCreateStorage_Unregister(void* callback) {
	return __permissions_OnUserCreateStorage_Unregister(callback);
}

extern int32_t (*__permissions_OnUserCreate_Register)(void*);

static int32_t OnUserCreate_Register(void* callback) {
	return __permissions_OnUserCreate_Register(callback);
}

extern int32_t (*__permissions_OnUserCreate_Unregister)(void*);

static int32_t OnUserCreate_Unregister(void* callback) {
	return __permissions_OnUserCreate_Unregister(callback);
}

extern int32_t (*__permissions_OnUserDeleteStorage_Register)(void*);

static int32_t OnUserDeleteStorage_Register(void* callback) {
	return __permissions_OnUserDeleteStorage_Register(callback);
}

extern int32_t (*__permissions_OnUserDeleteStorage_Unregister)(void*);

static int32_t OnUserDeleteStorage_Unregister(void* callback) {
	return __permissions_OnUserDeleteStorage_Unregister(callback);
}

extern int32_t (*__permissions_OnUserDelete_Register)(void*);

static int32_t OnUserDelete_Register(void* callback) {
	return __permissions_OnUserDelete_Register(callback);
}

extern int32_t (*__permissions_OnUserDelete_Unregister)(void*);

static int32_t OnUserDelete_Unregister(void* callback) {
	return __permissions_OnUserDelete_Unregister(callback);
}

extern int32_t (*__permissions_OnUserGroupChangeStorage_Register)(void*);

static int32_t OnUserGroupChangeStorage_Register(void* callback) {
	return __permissions_OnUserGroupChangeStorage_Register(callback);
}

extern int32_t (*__permissions_OnUserGroupChangeStorage_Unregister)(void*);

static int32_t OnUserGroupChangeStorage_Unregister(void* callback) {
	return __permissions_OnUserGroupChangeStorage_Unregister(callback);
}

extern int32_t (*__permissions_OnUserGroupChange_Register)(void*);

static int32_t OnUserGroupChange_Register(void* callback) {
	return __permissions_OnUserGroupChange_Register(callback);
}

extern int32_t (*__permissions_OnUserGroupChange_Unregister)(void*);

static int32_t OnUserGroupChange_Unregister(void* callback) {
	return __permissions_OnUserGroupChange_Unregister(callback);
}

extern int32_t (*__permissions_OnUserImmunityChangeStorage_Register)(void*);

static int32_t OnUserImmunityChangeStorage_Register(void* callback) {
	return __permissions_OnUserImmunityChangeStorage_Register(callback);
}

extern int32_t (*__permissions_OnUserImmunityChangeStorage_Unregister)(void*);

static int32_t OnUserImmunityChangeStorage_Unregister(void* callback) {
	return __permissions_OnUserImmunityChangeStorage_Unregister(callback);
}

extern int32_t (*__permissions_OnUserImmunityChange_Register)(void*);

static int32_t OnUserImmunityChange_Register(void* callback) {
	return __permissions_OnUserImmunityChange_Register(callback);
}

extern int32_t (*__permissions_OnUserImmunityChange_Unregister)(void*);

static int32_t OnUserImmunityChange_Unregister(void* callback) {
	return __permissions_OnUserImmunityChange_Unregister(callback);
}

extern int32_t (*__permissions_OnUserLoaded_Register)(void*);

static int32_t OnUserLoaded_Register(void* callback) {
	return __permissions_OnUserLoaded_Register(callback);
}

extern int32_t (*__permissions_OnUserLoaded_Unregister)(void*);

static int32_t OnUserLoaded_Unregister(void* callback) {
	return __permissions_OnUserLoaded_Unregister(callback);
}

extern int32_t (*__permissions_OnUserPermissionChangeStorage_Register)(void*);

static int32_t OnUserPermissionChangeStorage_Register(void* callback) {
	return __permissions_OnUserPermissionChangeStorage_Register(callback);
}

extern int32_t (*__permissions_OnUserPermissionChangeStorage_Unregister)(void*);

static int32_t OnUserPermissionChangeStorage_Unregister(void* callback) {
	return __permissions_OnUserPermissionChangeStorage_Unregister(callback);
}

extern int32_t (*__permissions_OnUserPermissionChange_Register)(void*);

static int32_t OnUserPermissionChange_Register(void* callback) {
	return __permissions_OnUserPermissionChange_Register(callback);
}

extern int32_t (*__permissions_OnUserPermissionChange_Unregister)(void*);

static int32_t OnUserPermissionChange_Unregister(void* callback) {
	return __permissions_OnUserPermissionChange_Unregister(callback);
}

extern int32_t (*__permissions_OnUserRequest_Register)(void*);

static int32_t OnUserRequest_Register(void* callback) {
	return __permissions_OnUserRequest_Register(callback);
}

extern int32_t (*__permissions_OnUserRequest_Unregister)(void*);

static int32_t OnUserRequest_Unregister(void* callback) {
	return __permissions_OnUserRequest_Unregister(callback);
}

