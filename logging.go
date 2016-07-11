package main


import(
   log "github.com/Sirupsen/logrus"
   "os"
)


func log_init(){
   log.SetFormatter(&log.TextFormatter{})
   log.SetOutput(os.Stderr)
}


func info(msg string){
   log.Info(msg)
}


func warning(msg string){
   log.Warning(msg)
}


func error(msg string){
   log.Error(msg)
}