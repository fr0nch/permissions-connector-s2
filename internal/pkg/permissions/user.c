#include "shared.h"

PLUGIFY_EXPORT int32_t (*__permissions_DumpPermissions)(uint64_t, Vector*) = NULL;


PLUGIFY_EXPORT int32_t (*__permissions_CanAffectUser)(uint64_t, uint64_t) = NULL;


PLUGIFY_EXPORT int32_t (*__permissions_HasPermissionExtended)(uint64_t, String*, bool, uint32_t*, int64_t*) = NULL;


PLUGIFY_EXPORT int32_t (*__permissions_HasPermission)(uint64_t, String*) = NULL;


PLUGIFY_EXPORT int32_t (*__permissions_HasGroupExtended)(uint64_t, String*, int64_t*) = NULL;


PLUGIFY_EXPORT int32_t (*__permissions_HasGroup)(uint64_t, String*) = NULL;


PLUGIFY_EXPORT int32_t (*__permissions_GetUserGroups)(uint64_t, Vector*) = NULL;


PLUGIFY_EXPORT int32_t (*__permissions_GetImmunity)(uint64_t, int32_t*) = NULL;


PLUGIFY_EXPORT int32_t (*__permissions_SetImmunity)(int64_t, uint64_t, int32_t, bool) = NULL;


PLUGIFY_EXPORT int32_t (*__permissions_AddPermission)(int64_t, uint64_t, String*, int64_t, bool) = NULL;


PLUGIFY_EXPORT int32_t (*__permissions_SetPermission)(int64_t, uint64_t, String*, int64_t, bool) = NULL;


PLUGIFY_EXPORT int32_t (*__permissions_RemovePermission)(int64_t, uint64_t, String*, bool, bool) = NULL;


PLUGIFY_EXPORT int32_t (*__permissions_AddGroup)(int64_t, uint64_t, String*, int64_t, bool) = NULL;


PLUGIFY_EXPORT int32_t (*__permissions_RemoveGroup)(int64_t, uint64_t, String*, bool) = NULL;


PLUGIFY_EXPORT int32_t (*__permissions_GetCookie)(uint64_t, String*, Variant*) = NULL;


PLUGIFY_EXPORT int32_t (*__permissions_SetCookie)(int64_t, uint64_t, String*, Variant*, bool) = NULL;


PLUGIFY_EXPORT int32_t (*__permissions_GetAllCookies)(uint64_t, Vector*, Vector*) = NULL;


PLUGIFY_EXPORT int32_t (*__permissions_CreateUser)(int64_t, uint64_t, int32_t, bool, Vector*) = NULL;


PLUGIFY_EXPORT int32_t (*__permissions_DeleteUser)(int64_t, uint64_t) = NULL;


PLUGIFY_EXPORT uint32_t (*__permissions_UserExists)(uint64_t) = NULL;


PLUGIFY_EXPORT Vector (*__permissions_DumpUsersList)() = NULL;


PLUGIFY_EXPORT int32_t (*__permissions_LoadUser)(int64_t, uint64_t, String*, bool, bool) = NULL;


