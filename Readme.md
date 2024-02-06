### Build Plugin
#### Proje çalışırken plugins dizini altındaki tüm .so uzantıları çalıştırarak çıktıyı tablo olarak ekrana bastırmaktadır
#### go build -buildmode=plugin -o ./plugins/hazelcast.so ./hazelcast/hazelcast.go

### Run Plugin
#### go run main.go
#### Ekran çıktısı örneği
```
+---------+-------------------------+
| Tool    | Hazelcast               |
| Check   | OK                      |
| Brokers | broker1,broker2,broker3 |
| Status  | Up                      |
+---------+-------------------------+
+---------+-------------------------+
| Tool    | Kafka                   |
| Check   | OK                      |
| Brokers | broker1,broker2,broker3 |
| Status  | Up                      |
+---------+-------------------------+
```