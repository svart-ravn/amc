package main


import(
   log "github.com/Sirupsen/logrus"
   "os"
)


func log_init(){
   log.SetFormatter(&log.TextFormatter{})
   log.SetOutput(os.Stderr)
}


func debug(v ...interface{}){
   log.Debug(v)
}


func info(v ...interface{}){
   log.Info(v)
}


func warning(v ...interface{}){
   log.Warning(v)
}


func error(v ...interface{}){
   log.Error(v)
}