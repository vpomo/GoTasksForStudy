go clean -modcache
go mod init main
go get github.com/IBM/sarama

pip3 install podman-compose
podman-compose up -d

В одном терминале запустите consumer:
go run consumer.go

В другом терминале запустите producer:
go run producer.go

-----------------------------------------
Особенности реализации:

Producer:
Отправляет 5 сообщений с интервалом 1 секунда
Использует синхронного продюсера
Логирует успешную отправку сообщений

Consumer:
Использует Consumer Group API
Поддерживает graceful shutdown
Автоматически коммитит offset
Выводит полную информацию о полученных сообщениях
-----------------------------------------

Для обработки ошибок добавьте:
// Для producer
config.Producer.Retry.Max = 5
config.Producer.Retry.Backoff = 1000 * time.Millisecond

// Для consumer
config.Consumer.Offsets.Initial = sarama.OffsetOldest
config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRoundRobin

