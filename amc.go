package main


import (
   "flag"
   "os"
   "strings"
   "io/ioutil"
   "encoding/json"
   "time"
)


var cmdParameters = CommandLineParameters{LogLevel: "DEBUG", NoTimestamp: false};


// ----------------------------------------------------------------------------------------------------------
func parse_arguments(){
   info(cmdParameters)


   flag.StringVar(&cmdParameters.LogLevel, "log-level", "WARN", "log level")
   flag.BoolVar(&cmdParameters.ToConsole, "console", false, "output to console")
   flag.StringVar(&cmdParameters.OutputFolder, "output", "", "output folder")
   flag.BoolVar(&cmdParameters.NoTimestamp, "no-timestamp", false, "skip timestamp in folder name")
   flag.BoolVar(&cmdParameters.NoLackOfData, "no-lack", false, "do not print '[lack of data]' messages")
   flag.BoolVar(&cmdParameters.CompareConfigProps, "compare-config-props", false, "compare configs, i.e. multiline properties")

   flag.StringVar(&cmdParameters.ClustersFilter, "clusters", "*", "list of clusters")
   flag.StringVar(&cmdParameters.ConfigsFilter, "configs", "*", "list of HDP configs to be compared")

   flag.StringVar(&cmdParameters.ConfigFile, "config", "", "config file with defaults (configs/clusters.cfg by defaults)")
   flag.StringVar(&cmdParameters.MatchingPatterns, "mpatterns", "", "folder with matching patterns")

   var ShowHelp = flag.Bool("help", false, "show help")

   flag.Parse()

   info(cmdParameters)
   info(*ShowHelp)

   if *ShowHelp == true {
      flag.PrintDefaults()
      os.Exit(0)
   }

   cmdParameters.LogLevel = strings.ToLower(cmdParameters.LogLevel)

   if cmdParameters.ConfigFile == "" {
      cmdParameters.ConfigFile = "configs/clusters.cfg"
      warning("--config option was missed. Trying to use default config: " + cmdParameters.ConfigFile)
   }

   file, err := ioutil.ReadFile(cmdParameters.ConfigFile)
   if err != nil {
      error("Error reading file: " + cmdParameters.ConfigFile)
   }

   var clusterList []Cluster
   if json.Unmarshal(file, &clusterList) != nil {
      error("Cannot unmarshal cluster config file: " + cmdParameters.ConfigFile)
   }

   if cmdParameters.ClustersFilter != "*" {
      for _, v := range strings.Split(cmdParameters.ClustersFilter, ",") {
         for _, cl := range clusterList {
            if cl.Name == v {
               cmdParameters.Clusters = append(cmdParameters.Clusters, cl)
            }
         }
      }
   } else {
      cmdParameters.Clusters = clusterList
   }

   if ! cmdParameters.NoTimestamp && cmdParameters.OutputFolder != "" {
      cmdParameters.OutputFolder = cmdParameters.OutputFolder + "/" + time.Now().Format("2006-01-02_15-04-05")
   }
}


// ----------------------------------------------------------------------------------------------------------
func init(){
   cmdParameters.Action = os.Args[1]
   os.Args = append(os.Args[:1], os.Args[2:]...)

   parse_arguments()
}



func main(){

   // TODO:
   // 
   // 1. git diff between configs
   // 
   //        amc cdiff --clusters-list UAT,PROD --configs-list HDFS,HIVE       
   // 
   // 2. find changes for selected variable(s)
   // 
   //        amc log --cluster UAT --parameter dfs.namenode.thread.count
   // 
   // 3. diff between 2 version of selected configs
   // 
   //        amc vdiff --version1 v55|head --version2 v54|-1
   // 
   // 4. ?
   // 


   if cmdParameters.Action == "cdiff" {
      getDiffBetweenClusters()
   } else if cmdParameters.Action == "vdiff" {
      info("Going to check difference between two versions")
   } else if cmdParameters.Action == "log" {
      info("Going to check how parameter was changed")
   } else {
      error("First argument should specify action (cdiff, vdiff, log) while we have: " + cmdParameters.Action)
   }


   info("Completed. OK!")
   info(cmdParameters)
}