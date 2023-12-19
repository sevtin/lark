package template

var ConfigsYamlTemplate = ParseTemplate(`
name: lark_{{.PackageName}}
server_id: 1
port: 6600
log: "./configs/logger.yaml"
etcd:
  endpoints: ["lark-etcd-01:12379","lark-etcd-02:12381","lark-etcd-03:12383"]
  username:
  password:
  schema: lark
  read_timeout: 5000
  write_timeout: 5000
  dial_timeout: 5000
mysql:
  max_open_conns: 20
  max_idle_conns: 10
  max_lifetime: 120000
  max_idle_time: 7200000
  charset: utf8
  address: "lark-mysql-user:13306"
  db: lark_user
  username: root
  password: lark2022
  debug: false
redis:
  address: ["127.0.0.1:6379"]
  db: 0
  password:
  prefix: "LK:"
  single: true
`)
