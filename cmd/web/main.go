package main

import (
	"crypto/tls"
	"database/sql"
	"flag"
	"html/template"
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"

	// Import the models package that we just created. You need to prefix this with
	// whatever module path you set up back in chapter 02.01 (Project Setup and Creating
	// a Module) so that the import statement looks like this:
	// "{your-module-path}/internal/models". If you can't remember what module path you
	// used, you can find it at the top of the go.mod file.
	"example.com/snippetbox/internal/models"
	"github.com/alexedwards/scs/mysqlstore"
	"github.com/alexedwards/scs/v2"
	"github.com/go-playground/form/v4"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type application struct {
	debug    bool
	logger   *slog.Logger
	snippets models.SnippetModelInterface // Use our new interface type.

	users models.UserModelInterface
	// Use our new interface type.
	templateCache  map[string]*template.Template
	formDecoder    *form.Decoder
	sessionManager *scs.SessionManager
}

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
	}))
	err := godotenv.Load()
	if err != nil {
		logger.Error("env file not found")
	}
	defaultPort := os.Getenv("app_port")
	defaultDsn := os.Getenv("database_dsn")
	addr := flag.String("addr", ":"+defaultPort, "HTTP network address")
	dsn := flag.String("dsn", defaultDsn, "MySQL data source name")
	debug := flag.Bool("debug", false, "Enable debug mode")
	flag.Parse()

	log.Println(dsn)

	db, err := openDB(*dsn)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	defer db.Close()
	// Initialize a models.SnippetModel instance containing the connection pool
	// and add it to the application dependencies.

	// Initialize a new template cache...
	templateCache, err := newTemplateCache()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	// And add it to the application dependencies.

	formDecoder := form.NewDecoder()
	sessionManager := scs.New()
	sessionManager.Store = mysqlstore.New(db)
	sessionManager.Lifetime = 12 * time.Hour
	sessionManager.Cookie.Secure = true
	// And add the session manager to our application dependencies.
	app := &application{
		debug:          *debug,
		logger:         logger,
		snippets:       &models.SnippetModel{DB: db},
		templateCache:  templateCache,
		formDecoder:    formDecoder,
		sessionManager: sessionManager,
		users:          &models.UserModel{DB: db},
	}
	tlsConfig := &tls.Config{
		CurvePreferences: []tls.CurveID{tls.X25519, tls.CurveP256},
	}
	// Set the server's TLSConfig field to use the tlsConfig variable we just
	// created.
	srv := &http.Server{
		Addr:      *addr,
		Handler:   app.routes(),
		ErrorLog:  slog.NewLogLogger(logger.Handler(), slog.LevelError),
		TLSConfig: tlsConfig,
		// Add Idle, Read and Write timeouts to the server.
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	logger.Info("starting server", "addr", srv.Addr)
	// Call the ListenAndServe() method on our new http.Server struct to start
	// the server.
	err = srv.ListenAndServeTLS("./tls/cert.pem", "./tls/key.pem")

	logger.Error(err.Error())
	os.Exit(1)
}

// The openDB() function wraps sql.Open() and returns a sql.DB connection pool
// for a given DSN.
func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}
	return db, nil
}
