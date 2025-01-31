### Внутренности Go Runtime
1. Как работает планировщик (scheduler) в Go?
Ответ: Планировщик Go использует модель M:N , где M — потоки ОС (машины), G — горутины, P — процессоры (логические контексты). Планировщик распределяет горутины по потокам, обеспечивая конкурентность.

2. Что такое work-stealing в планировщике Go?
- Ответ: Work-stealing позволяет незанятым потокам (P) "воровать" горутины из очередей других потоков, чтобы равномерно распределить нагрузку.

3. Как избежать false sharing в Go?
- Ответ: Используйте padding в структурах для выравнивания данных по кэш-линиям:
```go
type Data struct {
    value int64
    _     [56]byte // padding для выравнивания (64 байта)
}
```

4. Что такое escape analysis и как им управлять?
- Ответ: Escape analysis определяет, выживает ли переменная после завершения функции. Используйте go build -gcflags="-m" для анализа. Чтобы избежать аллокаций в куче, уменьшайте использование указателей и замыканий.

5. Как работает сборщик мусора (GC) в Go?
- Ответ: GC использует триколорную маркировку (tricolor mark-and-sweep) с concurrent фазами. Он останавливает программу (STW) только на короткие периоды.

6. Как настроить параметры GC (например, GOGC)?
- Ответ: Переменная GOGC задает процент роста кучи перед запуском GC. По умолчанию 100. Установка GOGC=off отключает GC.

7. Что такое write barrier в GC?
- Ответ: Write barrier — это механизм, который отслеживает изменения указателей во время маркировки, чтобы гарантировать корректность работы GC.

8. Как использовать runtime.SetFinalizer?
- Ответ: Функция SetFinalizer(obj, func) вызывает финализатор для obj перед сборкой мусора. Используется для освобождения внешних ресурсов (например, дескрипторов файлов).

9. Как работает runtime.Gosched()?
- Ответ: Gosched() передает управление планировщику, позволяя другим горутинам выполниться. Не гарантирует переключение контекста.

11. Как использовать runtime.LockOSThread?
- Ответ: LockOSThread() привязывает горутину к текущему потоку ОС. Используется в GUI, CGO или при работе с thread-local storage.

12. Что такое runtime.NumGoroutine и как её интерпретировать?
- Ответ: Функция возвращает текущее количество горутин. Высокие значения могут указывать на утечки (например, незавершенные горутины).

13. Как работает runtime.MemStats?
- Ответ: MemStats предоставляет статистику по использованию памяти:
```go
    var stats runtime.MemStats
    runtime.ReadMemStats(&stats)
    fmt.Println(stats.HeapAlloc)
```

14. Что такое runtime.Trace и как им пользоваться?
- Ответ: Пакет runtime/trace позволяет записывать трассировку выполнения программы для анализа производительности и горутин.

15. Как использовать runtime/pprof для профилирования блокировок?
- Ответ:
```go
import _ "net/http/pprof"
go http.ListenAndServe(":6060", nil)
```
Затем:
```bash
go tool pprof http://localhost:6060/debug/pprof/block
```

16. Что такое GODEBUG-переменные (например, GODEBUG=gctrace=1)?
- Ответ: GODEBUG включает отладочную информацию. Например, gctrace=1 выводит логи GC, schedtrace=1000 — логи планировщика каждую секунду.

### Продвинутая конкурентность

17. Как реализовать lock-free структуру данных на Go?
- Ответ: Используйте пакет sync/atomic для операций CAS (Compare-And-Swap):
```go
type Counter struct {
    value int64
}
func (c *Counter) Increment() {
    atomic.AddInt64(&c.value, 1)
}
```

18. Что такое channel pipeline и как его реализовать?
- Ответ: Pipeline — цепочка обработки данных через каналы. Пример:
```go
    func stage(in <-chan int) <-chan int {
        out := make(chan int)
        go func() {
            for n := range in {
                out <- n * 2
            }
			close(out)
        }()
        return out
    }
```

19. Как реализовать fan-out/fan-in паттерн?
- Ответ: Fan-out — несколько горутин читают из одного канала. Fan-in — объединение данных из нескольких каналов в один:
```go
// Fan-in
    func merge(channels ...<-chan int) <-chan int {
        var wg sync.WaitGroup
        out := make(chan int)
        for _, ch := range channels {
            wg.Add(1)
            go func(c <-chan int) {
                for n := range c {
					out <- n
				}
                wg.Done()
            }(ch)
       }
	   go func() { wg.Wait(); close(out) }()
		return out
    }
```

20. Как избежать утечек горутин?
- Ответ: Всегда закрывайте каналы, используйте context для отмены, отслеживайте завершение горутин через sync.WaitGroup.

21. Что такое sync.Cond и когда его использовать?
- Ответ: sync.Cond — примитив для блокировки горутин до наступления условия. Используется редко, например, для реализации пулов с ожиданием.

22. Как использовать sync.Once для ленивой инициализации?
- Ответ:
```go
    var (
        instance *Singleton
        once     sync.Once
    )
    func GetInstance() *Singleton {
		once.Do(func() {
			instance = &Singleton{}
		})
        return instance
    }
```

23. Как реализовать таймаут для операции с каналом?
- Ответ: Используйте time.After в select:
```go
select {
    case <-ch:
    // данные получены
    case <-time.After(1 * time.Second):
    // таймаут
}
```

24. Что такое errgroup и как его применять?
- Ответ: errgroup (пакет golang.org/x/sync/errgroup) управляет группой горутин, возвращая первую ошибку:
```go
g, ctx := errgroup.WithContext(ctx)
g.Go(func() error { return doWork(ctx) })
if err := g.Wait(); err != nil { /* обработка */ }
```

25. Как реализовать пул воркеров с динамическим размером?
- Ответ: Используйте канал для управления количеством активных воркеров:
```go
    tasks := make(chan Task)
    var wg sync.WaitGroup
    for i := 0; i < maxWorkers; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            for task := range tasks {
                process(task)
            }
        }()
    }
    close(tasks)
    wg.Wait()
```

26. Как использовать atomic.Value для lock-free чтения/записи?
- Ответ: atomic.Value хранит указатель, позволяя атомарно обновлять данные:
```go
var config atomic.Value
config.Store(newConfig)
current := config.Load().(*Config)
```

27. Что такое livelock и как его избежать?
- Ответ: Livelock — состояние, когда горутины постоянно взаимодействуют без прогресса. Решение: добавить случайные задержки или перепроектировать логику.

28. Как реализовать graceful shutdown сервера?
- Ответ: Перехватывайте сигналы и закрывайте ресурсы через context:
```go
ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
defer cancel()
server.Shutdown(ctx)
```

29. Как использовать sync.Pool для снижения аллокаций?
- Ответ: sync.Pool кэширует объекты для повторного использования:
```go
var pool = sync.Pool{
    New: func() interface{} { return new(Buffer) },
}
buf := pool.Get().(*Buffer)
defer pool.Put(buf)
```

30. Что такое non-blocking операции с каналами?
- Ответ: Используйте select с default для неблокирующих операций:
```go
select {
case ch <- data:
    // отправлено
default:
    // буфер заполнен
}
```

31. Как реализовать priority queue с использованием каналов?
- Ответ: Используйте несколько каналов с приоритетами и select:
```go
select {
case task := <-highPriorityChan:
    processHigh(task)
case task := <-lowPriorityChan:
    processLow(task)
}
```

### Оптимизация и производительность

32. Как использовать pprof для анализа блокировок?
- Ответ: Запустите профилирование блокировок:
```bash
go tool pprof http://localhost:6060/debug/pprof/block
```

33. Как интерпретировать flame graph в Go?
- Ответ: Flame graph визуализирует время выполнения функций. Горизонтальные полосы — функции, ширина — время. Используйте go tool pprof -http=:8080 profile.pb.gz.

34. Как оптимизировать аллокации в горячем цикле?
- Ответ: Выносите аллокации за цикл, используйте sync.Pool, избегайте замыканий.

35. Что такое inline-функции и как ими управлять?
- Ответ: Inline подставляет тело функции вместо вызова. Управляйте через комментарии //go:noinline или флаги компилятора (-gcflags="-l").

36. Как уменьшить размер бинарного файла?
- Ответ: Используйте:
```bash
go build -ldflags="-s -w" -trimpath
upx --best binary
```

37. Как оптимизировать использование памяти структурами?
- Ответ: Упорядочивайте поля по убыванию размера для выравнивания:
```go
type Struct struct {
    a int64  // 8 байт
    b int32  // 4 байта
    c bool   // 1 байт
    // padding 3 байта
}
```

38. Как использовать go:noescape директиву?
- Ответ: Директива указывает компилятору, что указатель не "убегает" из функции, позволяя оптимизировать аллокации:
```go
//go:noescape
func myFunc(buf *byte)
```

39. Что такое escape analysis и как его использовать для оптимизации?
- Ответ: Анализ определяет, выделяется ли переменная в куче. Используйте -gcflags="-m" для проверки. Оптимизируйте, уменьшая указатели и замыкания.

40. Как использовать unsafe.Pointer для низкоуровневых операций?
- Ответ: Пример преобразования slice в массив:
```go
slice := []int{1, 2, 3}
arrPtr := (*[3]int)(unsafe.Pointer(&slice[0]))
```

41. Как оптимизировать работу с сетью (например, TCP параметры)?
- Ответ: Настройте net.Dialer:
```go
dialer := &net.Dialer{
    KeepAlive: 30 * time.Second,
    Timeout:   5 * time.Second,
}
```

42. Как использовать io.Writer с буферизацией?
- Ответ: Оберните в bufio.Writer:
```go
writer := bufio.NewWriter(file)
defer writer.Flush()
```

43. Как избежать contention в sync.Mutex?
- Ответ: Используйте шардинг (разделение данных на части) или sync.RWMutex для read-heavy нагрузок.

44. Как использовать SIMD инструкции в Go?
- Ответ: В Go нет прямой поддержки SIMD, но можно использовать ассемблерные вставки или оптимизировать код компилятором (автовекторизация).

45. Как профилировать память в рантайме?
- Ответ: Используйте runtime.ReadMemStats или pprof:
```go
f, _ := os.Create("heap.pprof")
pprof.WriteHeapProfile(f)
f.Close()
```

46. Как использовать бенчмарки для сравнения алгоритмов?
- Ответ: Пишите тесты с Benchmark и используйте testing.B.ReportAllocs():
```go
    func BenchmarkSort(b *testing.B) {
        data := generateData()
        b.ResetTimer()
        for i := 0; i < b.N; i++ {
            sort.Slice(data, func(i, j int) bool { return data[i] < data[j] })
        }
    }
```

### Безопасность (10 вопросов)
47. Как предотвратить SQL инъекции в Go?
- Ответ: Используйте подготовленные запросы (Prepare, Exec) или ORM, избегайте конкатенации строк.

48. Как безопасно хранить пароли?
- Ответ: Хешируйте с солью, используйте golang.org/x/crypto/bcrypt:
```go
hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
```

49. Как реализовать JWT аутентификацию?
- Ответ: Используйте github.com/golang-jwt/jwt:
```go
token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"user": "admin"})
signed, _ := token.SignedString([]byte("secret"))
```

50. Как защитить API от CSRF?
- Ответ: Используйте токены CSRF, проверяйте заголовок Origin, устанавливайте SameSite для кук.

51. Как настроить HTTPS в Go?
- Ответ: Используйте ListenAndServeTLS:
```go
err := http.ListenAndServeTLS(":443", "cert.pem", "key.pem", nil)
```

52. Как реализовать rate limiting?
- Ответ: Используйте golang.org/x/time/rate:
```go
limiter := rate.NewLimiter(rate.Every(time.Second), 10)
if !limiter.Allow() { /* отклонить запрос */ }
```

53. Как защититься от XSS атак?
- Ответ: Экранируйте вывод с помощью html/template, устанавливайте заголовок Content-Security-Policy.

54. Как безопасно работать с файловыми загрузками?
- Ответ: Ограничивайте типы MIME, размер файлов, сохраняйте вне корня веб-сервера, сканируйте антивирусом.

55. Как реализовать проверку входных данных?
- Ответ: Используйте библиотеки валидации (например, github.com/go-playground/validator), проверяйте диапазоны и форматы.

56. Как использовать OAuth2 в Go?
- Ответ: Используйте golang.org/x/oauth2:
```go
conf := &oauth2.Config{ClientID: "id", ClientSecret: "secret"}
url := conf.AuthCodeURL("state")
// перенаправьте пользователя на url
```

### Облачные технологии и микросервисы

57. Как реализовать health checks в микросервисе?
- Ответ: Создайте эндпоинт /health, возвращающий статус 200 при готовности.

58. Как настроить трейсинг в микросервисах?
- Ответ: Используйте OpenTelemetry или Jaeger:
```go
tracer := otel.Tracer("service-name")
ctx, span := tracer.Start(ctx, "operation")
defer span.End
```

