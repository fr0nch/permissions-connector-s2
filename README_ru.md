[![English](https://img.shields.io/badge/English-%F0%9F%87%AC%F0%9F%87%A7-blue?style=for-the-badge)](README.md)

# permissions-connector-s2

Плагин для [Plugify](https://github.com/untrustedmodders/plugify), который связывает плагин [permissions](https://github.com/untrustedmodders/plugify-plugin-permissions) с сервером CS2 через [s2sdk](https://github.com/untrustedmodders/plugify-plugin-s2sdk): загружает группы и права игроков при входе, уведомляет игроков о выданных, снятых и истёкших правах и группах и добавляет консольные/чат-команды для управления пользователями и группами.

## Команды

Игрок указывается через `#userid`, `$steamid64`, `&account_id` или ник. Длительность в единицах из `time_mode` (по умолчанию секунды).

<details open>
<summary><b>Общие</b></summary>

| Команда | Право | Описание |
|---|---|---|
| `perms` | - | Список доступных команд |
| `perms_reload_cfg` | `permissions.reload_config` | Перезагрузить конфиг |
| `perms_reload_translations` | `permissions.reload_translations` | Перезагрузить переводы |

</details>

<details>
<summary><b>Пользователи</b></summary>

| Команда | Право | Описание |
|---|---|---|
| `perms_users` | `permissions.user.list` | Список игроков на сервере |
| `perms_user_info <игрок>` | `permissions.user.info` | Вся информация об игроке |
| `perms_user_permissions <игрок>` | `permissions.user.permission.info` | Права и ограничения игрока |
| `perms_user_groups <игрок>` | `permissions.user.group.info` | Группы игрока |
| `perms_user_check_permission <игрок> <право>` | `permissions.user.permission.check` | Проверить право или ограничение |
| `perms_user_check_group <игрок> <группа>` | `permissions.user.group.check` | Проверить, состоит ли игрок в группе |
| `perms_user_add_permission <игрок> <право> <длительность>` | `permissions.user.permission.add` | Добавить право или ограничение |
| `perms_user_set_permission <игрок> <право> <длительность>` | `permissions.user.permission.set` | Установить право или ограничение |
| `perms_user_remove_permission <игрок> <право>` | `permissions.user.permission.remove` | Снять право или ограничение |
| `perms_user_add_group <игрок> <группа> <длительность>` | `permissions.user.group.add` | Добавить группу |
| `perms_user_set_group <игрок> <группа> <длительность>` | `permissions.user.group.set` | Установить группу |
| `perms_user_remove_group <игрок> <группа>` | `permissions.user.group.remove` | Снять группу |

</details>

<details>
<summary><b>Группы</b></summary>

| Команда | Право | Описание |
|---|---|---|
| `perms_groups` | `permissions.listgroups` | Список групп |
| `perms_group_info <группа>` | `permissions.group.info` | Информация о группе |
| `perms_group_permissions <группа>` | `permissions.group.permissions` | Права группы |
| `perms_group_check_permission <группа> <право>` | `permissions.group.permission.check` | Проверить право или ограничение группы |
| `perms_creategroup <группа> <приоритет> <родитель> <права...>` | `permissions.creategroup` | Создать группу (права через пробел) |
| `perms_deletegroup <группа>` | `permissions.deletegroup` | Удалить группу |
| `perms_group_parent_set <дочерняя> <родитель>` | `permissions.group.parent.set` | Установить родительскую группу |
| `perms_group_parent_info <группа>` | `permissions.group.parent.info` | Показать родительскую группу |
| `perms_group_parent_check <дочерняя> <родитель>` | `permissions.group.parent.check` | Проверить, является ли группа родителем другой |
| `perms_group_set_option <группа> <опция> <значение>` | `permissions.group.option.set` | Установить опцию группы |
| `perms_group_option <группа> <опция>` | `permissions.group.option` | Показать опцию группы |
| `perms_group_add_permission <группа> <право>` | `permissions.group.permission.add` | Добавить право |
| `perms_group_set_permission <группа> <право>` | `permissions.group.permission.set` | Установить право |
| `perms_group_remove_permission <группа> <право>` | `permissions.group.permission.remove` | Снять право |

</details>

## Настройка

Конфиг создаётся при первом запуске: `configs/permissions/connector.yml`.

| Параметр | Описание |
|---|---|
| `notify_mode` | Куда уведомлять игроков об изменении прав и групп: `console`, `chat`, `hint` |
| `hint_duration` | Длительность hint-уведомления в секундах |
| `time_mode` | Единица аргумента `<длительность>` в `perms_user_add/set_permission` и `perms_user_add/set_group`: `0` секунды, `1` минуты, `2` часы, `3` дни |
| `abbreviate_time` | Сокращённые единицы (`5 мин.`) вместо полных слов |
| `format_timestamp` | Формат даты в [нотации Go](https://pkg.go.dev/time#pkg-constants), например `2006-01-02 15:04:05` |
| `display_active_groups` | Показывать игроку его активные группы при входе |
| `hidden_groups` | Группы, скрытые от игроков (например, `Default`) |

## Переводы

`permissions-connector-s2.yaml` (английский и русский) поставляется рядом с манифестом плагина. Чтобы переопределить его, положите свою копию в `translations/`, `configs/` или `data/` Plugify. Перезагрузка - `perms_reload_translations`.

## Зависимости

- [permissions](https://github.com/untrustedmodders/plugify-plugin-permissions) `>= 1.3.0`
- [s2sdk](https://github.com/untrustedmodders/plugify-plugin-s2sdk)
- [translations](https://github.com/untrustedmodders/plugify-plugin-translations)

Группы и пользователи хранятся плагином хранилища, например [permissions-storage](https://github.com/fr0nch/permissions-storage).

## Установка

Из conda-канала:
```sh
conda install -c https://fr0nch.github.io/permissions-connector-s2/ permissions-connector-s2
```

## Сборка

Нативно (нужны Go и `jq`):
```sh
./build.sh [debug|release]
```

Через Docker (образ собирается автоматически при первом запуске):
```sh
./build.sh [debug|release] --docker
```

Результат в обоих случаях: `build/<mode>/libpermissions-connector-s2.so`, `build/<mode>/permissions-connector-s2.pplugin` и `build/<mode>/permissions-connector-s2.yaml`.

## Лицензия

Распространяется под [лицензией MIT](LICENSE).
