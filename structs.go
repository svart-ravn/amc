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



type CommandLineParameters struct{
   Action string
   ConfigFile string
   MatchingPatterns string
   OutputFolder string
   LogLevel string
   ClustersFilter string
   ConfigsFilter string
   Clusters []Cluster
   ToConsole bool
   NoTimestamp bool
   NoLackOfData bool
   CompareConfigProps bool
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