# SSH

## Check Format

```yaml
- name:
  release:
    org: compscore
    repo: ssh
    tag: latest
  credentials:
    username:
    password:
  target:
  command:
  expectedOutput:
  weight:
```

## Parameters

|    parameter     |          path           |   type   | default  | required | description                                     |
| :--------------: | :---------------------: | :------: | :------: | :------: | :---------------------------------------------- |
|      `name`      |         `.name`         | `string` |   `""`   |  `true`  | `name of check (must be unique)`                |
|      `org`       |     `.release.org`      | `string` |   `""`   |  `true`  | `organization that check repository belongs to` |
|      `repo`      |     `.release.repo`     | `string` |   `""`   |  `true`  | `repository of the check`                       |
|      `tag`       |     `.release.tag`      | `string` | `latest` | `false`  | `tagged version of check`                       |
|    `username`    | `.credentials.username` | `string` |   `""`   | `false`  | `username of ssh user`                          |
|    `password`    | `.credentials.password` | `string` |   `""`   | `false`  | `default password of ssh user`                  |
|     `target`     |        `.target`        | `string` |   `""`   |  `true`  | `network target of ssh server`                  |
|    `command`     |       `.command`        | `string` |   `""`   | `false`  | `command to run on ssh server`                  |
| `expectedOutput` |    `.expectedOutput`    | `string` |   `""`   | `false`  | `expected output of provided command`           |
|     `weight`     |        `.weight`        |  `int`   |   `0`    |  `true`  | `amount of points a successful check is worth`  |

## Examples

```yaml
- name: host_a-ssh
  release:
    org: compscore
    repo: ssh
    tag: latest
  credentials:
    username: ubuntu
    password: changeme
  target: 10.{{ .Team }}.1.1:22
  command: whoami
  expectedOutput: ubuntu
  weight: 2
```
