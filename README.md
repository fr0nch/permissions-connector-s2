[![Русский](https://img.shields.io/badge/Русский-%F0%9F%87%B7%F0%9F%87%BA-green?style=for-the-badge)](README_ru.md)

# permissions-connector-s2

A [Plugify](https://github.com/untrustedmodders/plugify) plugin that connects the [permissions](https://github.com/untrustedmodders/plugify-plugin-permissions) plugin to a CS2 server through [s2sdk](https://github.com/untrustedmodders/plugify-plugin-s2sdk): it loads groups and players' permissions when they join, notifies players about granted, removed and expired permissions and groups, and adds console/chat commands to manage users and groups.

## Commands

Players are targeted by `#userid`, `$steamid64`, `&account_id` or name. Durations are in the unit set by `time_mode` (seconds by default).

<details>
<summary><b>General</b></summary>

| Command | Permission | Description |
|---|---|---|
| `perms` | - | List of available commands |
| `perms_reload_cfg` | `permissions.reload_config` | Reload the config |
| `perms_reload_translations` | `permissions.reload_translations` | Reload the translations |

</details>

<details>
<summary><b>Users</b></summary>

| Command | Permission | Description |
|---|---|---|
| `perms_users` | `permissions.user.list` | List of players on the server |
| `perms_user_info <player>` | `permissions.user.info` | All information about a player |
| `perms_user_permissions <player>` | `permissions.user.permission.info` | Player's permissions and restrictions |
| `perms_user_groups <player>` | `permissions.user.group.info` | Player's groups |
| `perms_user_check_permission <player> <permission>` | `permissions.user.permission.check` | Check a permission or restriction |
| `perms_user_check_group <player> <group>` | `permissions.user.group.check` | Check whether a player is in a group |
| `perms_user_add_permission <player> <permission> <duration>` | `permissions.user.permission.add` | Add a permission or restriction |
| `perms_user_set_permission <player> <permission> <duration>` | `permissions.user.permission.set` | Set a permission or restriction |
| `perms_user_remove_permission <player> <permission>` | `permissions.user.permission.remove` | Remove a permission or restriction |
| `perms_user_add_group <player> <group> <duration>` | `permissions.user.group.add` | Add a group |
| `perms_user_set_group <player> <group> <duration>` | `permissions.user.group.set` | Set a group |
| `perms_user_remove_group <player> <group>` | `permissions.user.group.remove` | Remove a group |

</details>

<details>
<summary><b>Groups</b></summary>

| Command | Permission | Description |
|---|---|---|
| `perms_groups` | `permissions.listgroups` | List of groups |
| `perms_group_info <group>` | `permissions.group.info` | Information about a group |
| `perms_group_permissions <group>` | `permissions.group.permissions` | Group's permissions |
| `perms_group_check_permission <group> <permission>` | `permissions.group.permission.check` | Check a group's permission or restriction |
| `perms_creategroup <group> <priority> <parent> <permissions...>` | `permissions.creategroup` | Create a group (permissions are space separated) |
| `perms_deletegroup <group>` | `permissions.deletegroup` | Delete a group |
| `perms_group_parent_set <child> <parent>` | `permissions.group.parent.set` | Set the parent group |
| `perms_group_parent_info <group>` | `permissions.group.parent.info` | Show the parent group |
| `perms_group_parent_check <child> <parent>` | `permissions.group.parent.check` | Check whether a group is a parent of another |
| `perms_group_set_option <group> <option> <value>` | `permissions.group.option.set` | Set a group option |
| `perms_group_option <group> <option>` | `permissions.group.option` | Show a group option |
| `perms_group_add_permission <group> <permission>` | `permissions.group.permission.add` | Add a permission |
| `perms_group_set_permission <group> <permission>` | `permissions.group.permission.set` | Set a permission |
| `perms_group_remove_permission <group> <permission>` | `permissions.group.permission.remove` | Remove a permission |

</details>

## Configuration

The config is created on the first start: `configs/permissions/connector.yml`.

| Setting | Description |
|---|---|
| `notify_mode` | Where players are notified about permission and group changes: `console`, `chat`, `hint` |
| `hint_duration` | Hint notification duration in seconds |
| `time_mode` | Unit of the `<duration>` argument in `perms_user_add/set_permission` and `perms_user_add/set_group`: `0` seconds, `1` minutes, `2` hours, `3` days |
| `abbreviate_time` | Abbreviated units (`5 min.`) instead of full words |
| `format_timestamp` | Date format in [Go layout](https://pkg.go.dev/time#pkg-constants), e.g. `2006-01-02 15:04:05` |
| `display_active_groups` | Show a player's active groups when they join |
| `hidden_groups` | Groups hidden from players (e.g. `Default`) |

## Translations

`permissions-connector-s2.yaml` (English and Russian) is shipped next to the plugin manifest. To override it, put your copy into `translations/`, `configs/` or `data/` of Plugify. Reload with `perms_reload_translations`.

## Dependencies

- [permissions](https://github.com/untrustedmodders/plugify-plugin-permissions) `>= 1.3.0`
- [s2sdk](https://github.com/untrustedmodders/plugify-plugin-s2sdk)
- [translations](https://github.com/untrustedmodders/plugify-plugin-translations)

Groups and users are stored by a storage plugin, e.g. [permissions-storage](https://github.com/fr0nch/permissions-storage).

## Installation

From the conda channel:
```sh
conda install -c https://fr0nch.github.io/permissions-connector-s2/ permissions-connector-s2
```

## Building

Natively (requires Go and `jq`):
```sh
./build.sh [debug|release]
```

Via Docker (the image is built automatically on first run):
```sh
./build.sh [debug|release] --docker
```

Output in both cases: `build/<mode>/libpermissions-connector-s2.so`, `build/<mode>/permissions-connector-s2.pplugin` and `build/<mode>/permissions-connector-s2.yaml`.

## License

Licensed under the [MIT License](LICENSE).
