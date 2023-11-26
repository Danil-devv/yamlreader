# YAML READER
#### Это имплементация одного из подпунктов домашнего задания Академии Бэкенда Тинькофф (поток 2023-2024)

Использование:

YAML:
```yaml
server_configs:
  host: "localhost"
  port: "18080"
  timeout: 1s
```
main.go:
```go
r, err := yamlreader.New("example/config.yaml")
if err != nil {
    log.Fatalf("cannot open and load .yaml file: %s", err)
}

type ServerConfig struct {
    Host    string        `yaml:"host"`
    Port    string        `yaml:"port"`
    Timeout time.Duration `yaml:"timeout"`
}
conf := struct {
    ServerConf ServerConfig `yaml:"server_configs"`
}{}

err = r.ReadYAML(&conf)
if err != nil {
    log.Fatalf("cannot read .yaml file: %s", err)
}
fmt.Println(conf.ServerConf.Host) // -> localhost
fmt.Println(conf.ServerConf.Port) // -> 18080
fmt.Println(conf.ServerConf.Timeout) // -> 1s
```
