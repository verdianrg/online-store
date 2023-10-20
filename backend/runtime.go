package onlinestore

import (
	"log"
	"net/http"
	"online-store/config"
	"os"

	"github.com/go-openapi/errors"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func NewRuntime() *Runtime {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	db := config.InitDB()
	rt := &Runtime{
		ErrorLog: log.New(os.Stderr, "[ERROR] ", log.Ldate|log.Ltime|log.Lshortfile),
		InfoLog:  log.New(os.Stdout, "[INFO] ", log.Ldate|log.Ltime|log.Lshortfile),
		Logger:   log.New(os.Stdout, "[Online-Store-Service] ", log.Ldate|log.Ltime|log.Lshortfile),

		Db: db,
	}

	return rt
}

type Runtime struct {
	ErrorLog *log.Logger
	InfoLog  *log.Logger
	*log.Logger

	Db *gorm.DB
}

func (r *Runtime) Info() *log.Logger {
	return r.InfoLog
}

func (r *Runtime) Error() *log.Logger {
	return r.ErrorLog
}

func (rt *Runtime) Debugf(format string, args ...interface{}) {
	rt.Printf("[DEBUG] "+format, args...)
}

func (rt *Runtime) Infof(format string, args ...interface{}) {
	rt.Printf("[INFO] "+format, args...)
}

func (rt *Runtime) Warnf(format string, args ...interface{}) {
	rt.Printf("[WARN] "+format, args...)
}

func (rt *Runtime) Errorf(format string, args ...interface{}) {
	rt.Printf("[ERROR] "+format, args...)
}

func (r *Runtime) SetError(code int, msg string, args ...interface{}) error {
	return errors.New(int32(code), msg, args...)
}

func (r *Runtime) GetError(err error) errors.Error {
	if v, ok := err.(errors.Error); ok {
		return v
	}

	return errors.New(http.StatusInternalServerError, err.Error())
}

func (r *Runtime) DB() *gorm.DB {
	return r.Db
}
