Давай разберём все способы от простых к сложным.

## 1. Базовые способы создания ошибок

### `errors.New()` - самый простой способ
```go
import "errors"

err := errors.New("что-то пошло не так")
```

### `fmt.Errorf()` - с форматированием
```go
import "fmt"

err := fmt.Errorf("пользователь %s не найден", userID)
```

## 2. Ошибки с контекстом (Sentinel errors)

Часто используются для сравнения:

```go
var (
    ErrNotFound = errors.New("entity not found")
    ErrInvalidInput = errors.New("invalid input")
    ErrPermissionDenied = errors.New("permission denied")
)

// Использование
if err == ErrNotFound {
    // обработка специфической ошибки
}
```

## 3. Кастомные типы ошибок (для богатого контекста)

```go
type ValidationError struct {
    Field   string
    Value   interface{}
    Message string
}

func (e *ValidationError) Error() string {
    return fmt.Sprintf("validation failed on field %s: %s", e.Field, e.Message)
}

// Использование
func ValidateAge(age int) error {
    if age < 0 {
        return &ValidationError{
            Field:   "age",
            Value:   age,
            Message: "age cannot be negative",
        }
    }
    return nil
}

// Type assertion для получения деталей
if err, ok := err.(*ValidationError); ok {
    fmt.Printf("Field: %s, Value: %v\n", err.Field, err.Value)
}
```

## 4. Ошибки с кодом (для API)

```go
type HTTPError struct {
    Code    int
    Message string
    Err     error // внутренняя ошибка
}

func (e *HTTPError) Error() string {
    if e.Err != nil {
        return fmt.Sprintf("HTTP %d: %s (%v)", e.Code, e.Message, e.Err)
    }
    return fmt.Sprintf("HTTP %d: %s", e.Code, e.Message)
}

func (e *HTTPError) Unwrap() error {
    return e.Err
}

// Использование
return &HTTPError{
    Code:    404,
    Message: "user not found",
    Err:     ErrNotFound,
}
```

## 5. Оборачивание ошибок (Go 1.13+)

```go
// Базовое оборачивание
if err := someFunc(); err != nil {
    return fmt.Errorf("someFunc failed: %w", err)
}

// Проверка обёрнутых ошибок
if errors.Is(err, ErrNotFound) {
    // ошибка ErrNotFound где-то в цепочке
}

// Получение конкретного типа
var httpErr *HTTPError
if errors.As(err, &httpErr) {
    fmt.Printf("HTTP code: %d\n", httpErr.Code)
}
```

## 6. Практический пример (сложный случай)

```go
package main

import (
    "errors"
    "fmt"
)

// Определяем иерархию ошибок
type DomainError struct {
    Domain  string
    Op      string  // операция, вызвавшая ошибку
    Err     error   // вложенная ошибка
}

func (e *DomainError) Error() string {
    return fmt.Sprintf("[%s] %s: %v", e.Domain, e.Op, e.Err)
}

func (e *DomainError) Unwrap() error {
    return e.Err
}

// Специфичные ошибки
type NotFoundError struct {
    DomainError
    EntityType string
    EntityID   interface{}
}

func (e *NotFoundError) Error() string {
    return fmt.Sprintf("%s %v not found in domain %s", 
        e.EntityType, e.EntityID, e.Domain)
}

// Фабрика для создания ошибок
func NewNotFoundError(domain, op, entityType string, entityID interface{}) error {
    return &NotFoundError{
        DomainError: DomainError{
            Domain: domain,
            Op:     op,
            Err:    errors.New("not found"),
        },
        EntityType: entityType,
        EntityID:   entityID,
    }
}

// Использование
func GetUser(domain, userID string) error {
    // ... логика
    return NewNotFoundError(domain, "GetUser", "User", userID)
}

func main() {
    err := GetUser("accounts", "123")
    
    // Проверка на not found
    var notFound *NotFoundError
    if errors.As(err, &notFound) {
        fmt.Printf("Entity %s with ID %v not found\n", 
            notFound.EntityType, notFound.EntityID)
    }
}
```

## 7. Продвинутые паттерны

### Ошибки со стектрейсом (для отладки)
```go
import "runtime"

type TraceableError struct {
    msg   string
    cause error
    stack []uintptr
}

func NewTraceableError(msg string, cause error) *TraceableError {
    pc := make([]uintptr, 10)
    runtime.Callers(2, pc)
    
    return &TraceableError{
        msg:   msg,
        cause: cause,
        stack: pc,
    }
}

func (e *TraceableError) Error() string {
    if e.cause != nil {
        return fmt.Sprintf("%s: %v", e.msg, e.cause)
    }
    return e.msg
}

func (e *TraceableError) Unwrap() error {
    return e.cause
}
```

## Советы для собеседования:

1. **Всегда используйте `%w` для оборачивания**, если нужно проверить ошибку выше
2. **Sentinel errors** (var ErrSomething) хороши для известных состояний
3. **Кастомные типы** используйте, когда нужно передать дополнительный контекст
4. **Не злоупотребляйте типами** - иногда простого `fmt.Errorf` достаточно
5. **Всегда думайте о том**, как ошибка будет обрабатываться выше

**Q: ** Расскажи про errors.Is() и errors.As()
Отличный вопрос! Это ключевые функции для работы с ошибками в современном Go.

## errors.Is() - проверка на конкретную ошибку

`errors.Is()` проверяет, находится ли целевая ошибка где-то в цепочке обёрнутых ошибок.

### Базовое использование

```go
var ErrNotFound = errors.New("not found")
var ErrPermission = errors.New("permission denied")

func GetUser(id int) error {
    if id <= 0 {
        return fmt.Errorf("invalid id: %w", ErrNotFound)
    }
    return nil
}

func main() {
    err := GetUser(-1)
    
    // Проверяем, является ли err или любая обёрнутая ошибка ErrNotFound
    if errors.Is(err, ErrNotFound) {
        fmt.Println("This is a not found error (or wraps it)")
    }
}
```

### Как это работает под капотом

`errors.Is()` делает следующее:
1. Сравнивает ошибку с целью через `==`
2. Если ошибка реализует интерфейс `Is(error) bool`, вызывает этот метод
3. Рекурсивно разворачивает ошибку через `Unwrap()` и повторяет проверку

```go
// Пример кастомной реализации Is()
type TemporaryError struct {
    Err error
}

func (e *TemporaryError) Error() string {
    return fmt.Sprintf("temporary error: %v", e.Err)
}

func (e *TemporaryError) Is(target error) bool {
    // Считаем любую временную ошибку эквивалентной target,
    // если target - тоже временная ошибка
    _, ok := target.(*TemporaryError)
    return ok
}

func (e *TemporaryError) Unwrap() error {
    return e.Err
}
```

## errors.As() - проверка на тип ошибки

`errors.As()` находит первую ошибку в цепочке, которая соответствует целевому типу, и присваивает её значение.

### Базовое использование

```go
type ValidationError struct {
    Field string
    Value interface{}
}

func (e *ValidationError) Error() string {
    return fmt.Sprintf("invalid field %s: %v", e.Field, e.Value)
}

func Process(data map[string]interface{}) error {
    if val, ok := data["age"]; ok {
        if age, ok := val.(int); ok && age < 0 {
            return &ValidationError{Field: "age", Value: age}
        }
    }
    return nil
}

func main() {
    err := Process(map[string]interface{}{"age": -5})
    
    var validationErr *ValidationError
    if errors.As(err, &validationErr) {
        fmt.Printf("Validation failed: field=%s, value=%v\n", 
            validationErr.Field, validationErr.Value)
    }
}
```

### Важные нюансы использования errors.As()

```go
// Правильно - передаём указатель на указатель
var target *MyError
if errors.As(err, &target) {
    // используем target
}

// Неправильно - так не сработает!
var target MyError
if errors.As(err, target) { // паника или неверная работа
    // ...
}

// Можно использовать для интерфейсов
var temp interface { Temporary() bool }
if errors.As(err, &temp) {
    if temp.Temporary() {
        // временная ошибка
    }
}
```

## Сравнение Is() и As()

```go
// Демонстрация различий
type NetworkError struct {
    Code    int
    Message string
    Err     error
}

func (e *NetworkError) Error() string {
    return fmt.Sprintf("network error %d: %s", e.Code, e.Message)
}

func (e *NetworkError) Unwrap() error {
    return e.Err
}

var ErrTimeout = errors.New("timeout")

func main() {
    err := &NetworkError{
        Code:    408,
        Message: "request timeout",
        Err:     ErrTimeout,
    }
    
    // errors.Is - проверяет значение
    if errors.Is(err, ErrTimeout) {
        fmt.Println("It's a timeout (or wraps timeout)")
    }
    
    // errors.As - проверяет тип и извлекает значение
    var netErr *NetworkError
    if errors.As(err, &netErr) {
        fmt.Printf("Network error: code=%d, message=%s\n", 
            netErr.Code, netErr.Message)
        
        // Можем получить вложенную ошибку
        if errors.Is(netErr.Err, ErrTimeout) {
            fmt.Println("Underlying cause is timeout")
        }
    }
}
```


## Паттерны использования

### 1. Проверка нескольких типов ошибок

```go
func handleError(err error) {
    switch {
    case errors.Is(err, ErrNotFound):
        // 404 handling
    case errors.Is(err, ErrPermission):
        // 403 handling
    default:
        var timeoutErr *TimeoutError
        if errors.As(err, &timeoutErr) {
            // retry logic
        }
    }
}
```

### 2. Извлечение деталей из ошибки

```go
type HTTPError interface {
    StatusCode() int
}

func respondWithError(w http.ResponseWriter, err error) {
    var httpErr HTTPError
    if errors.As(err, &httpErr) {
        w.WriteHeader(httpErr.StatusCode())
    } else {
        w.WriteHeader(http.StatusInternalServerError)
    }
}
```

## Производительность и best practices

```go
// Плохо - создаём новую переменную каждый раз
if errors.As(err, &new(ValidationError)) {
    // ...
}

// Хорошо - объявляем один раз
var valErr *ValidationError
if errors.As(err, &valErr) {
    // ...
}

// Для частых проверок можно кешировать
var ErrTemp = &TemporaryError{} // sentinel для Is()

// Использование с пользовательскими типами
type codedError interface {
    Code() int
}

var coded codedError
if errors.As(err, &coded) {
    fmt.Println("Error code:", coded.Code())
}
```

## Что важно помнить на собеседовании:

1. **errors.Is()** - для проверки **значений** ошибок (sentinel errors)
2. **errors.As()** - для проверки **типов** и извлечения данных
3. Обе функции работают рекурсивно по всей цепочке `Unwrap()`
4. Можно реализовать свои `Is()` и `As()` для кастомной логики
5. Всегда передавайте **двойной указатель** в `errors.As()`
6. Появились в Go 1.13, до этого использовались type assertions и сравнения

Эти функции делают обработку ошибок в Go намного более гибкой и элегантной!


