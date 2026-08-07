#include "shared.h"

PLUGIFY_EXPORT int32_t (*__permissions_SetParent)(int64_t, String*, String*, bool) = NULL;


PLUGIFY_EXPORT int32_t (*__permissions_GetParent)(String*, String*) = NULL;


PLUGIFY_EXPORT int32_t (*__permissions_DumpPermissionsGroup)(String*, Vector*) = NULL;


PLUGIFY_EXPORT Vector (*__permissions_GetAllGroups)() = NULL;


PLUGIFY_EXPORT int32_t (*__permissions_HasPermissionGroupExtended)(String*, String*, bool) = NULL;


PLUGIFY_EXPORT int32_t (*__permissions_HasPermissionGroup)(String*, String*) = NULL;


PLUGIFY_EXPORT int32_t (*__permissions_HasParentGroup)(String*, String*) = NULL;


PLUGIFY_EXPORT int32_t (*__permissions_GetPriorityGroup)(String*, int32_t*) = NULL;


PLUGIFY_EXPORT int32_t (*__permissions_AddPermissionGroup)(int64_t, String*, String*, bool) = NULL;


PLUGIFY_EXPORT int32_t (*__permissions_SetPermissionGroup)(int64_t, String*, String*, bool) = NULL;


PLUGIFY_EXPORT int32_t (*__permissions_RemovePermissionGroup)(int64_t, String*, String*, bool, bool) = NULL;


PLUGIFY_EXPORT int32_t (*__permissions_GetOptionGroup)(String*, String*, Variant*) = NULL;


PLUGIFY_EXPORT int32_t (*__permissions_SetOptionGroup)(int64_t, String*, String*, Variant*, bool) = NULL;


PLUGIFY_EXPORT int32_t (*__permissions_GetAllOptionsGroup)(String*, Vector*, Vector*) = NULL;


PLUGIFY_EXPORT int32_t (*__permissions_CreateGroup)(int64_t, String*, Vector*, int32_t, String*, bool) = NULL;


PLUGIFY_EXPORT int32_t (*__permissions_DeleteGroup)(int64_t, String*, bool) = NULL;


PLUGIFY_EXPORT bool (*__permissions_GroupExists)(String*) = NULL;


PLUGIFY_EXPORT int32_t (*__permissions_LoadGroups)(int64_t, bool) = NULL;


