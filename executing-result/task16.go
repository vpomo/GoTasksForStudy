package main

import (
	"flag"
	"fmt"
)

var port int

func init() {
	flag.IntVar(&port, "port", 8000, "port number")
}

func main() {
	flag.Parse()
	fmt.Println("port", port)
}

/**
Что необходимо добавить в init() чтобы спарсить аргумент терминала - port?

- flag.ParseString(&port, flag.PORT, "port number")
- flag.IpVar(&port, "port")
- flag.StringVar(&port, "port", "8000", "port number")
+ flag.IntVar(&port, "port", 8000, "port number")
- в init нельзя добалять парсинг аргументов
- Не знаю
*/
