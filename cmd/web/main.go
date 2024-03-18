package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	// snippets *mysql.SnippetModel
}

func main() {
	addr := flag.String("addr", ":8080", "Сетевой адрес веб-сервера")
	// dsn := flag.String("dsn", "web:ilya2003#@/snippetbox?parseTime=true", "Название PostgreSQL источника данных")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// db, err := openDB(*dsn)
	// if err != nil {
	// 	errorLog.Fatal(err)
	// }

	// defer db.Close()

	// Инициализируем экземпляр mysql.SnippetModel и добавляем его в зависимостях.
	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
		// snippets: &mysql.SnippetModel{DB: db},
	}

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Запуск сервера на %s", *addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}

// func openDB(dsn string) (*sql.DB, error) {
// 	db, err := sql.Open("postgres", dsn)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if err = db.Ping(); err != nil {
// 		return nil, err
// 	}
// 	return db, nil
// }
