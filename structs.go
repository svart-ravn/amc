package main;

import (
   "encoding/json"
   "io/ioutil"
)


type Cluster struct{
   AmbariUrl string
   AmbariPassword string
   AmbariUser string
   Name string
}


// --------------------------------------------------------------------------------------------------------------
func readConfigFile(configFile string)([]Cluster){
   file, err := ioutil.ReadFile(configFile)
   if err != nil {
      error("Cannot read file: " + configFile)
   }

   var cl []Cluster
   if json.Unmarshal(file, &cl) != nil {
      error("Cannot unmarshal cluster config file")
   }

   return cl
}