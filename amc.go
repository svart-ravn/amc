package main


import (
   "flag"
   "os"
   "strings"
   "io/ioutil"
   "encoding/json"
   "time"

   "fmt"
)


var cmdParameters = CommandLineParameters{LogLevel: "DEBUG", NoTimestamp: false};



// ----------------------------------------------------------------------------------------------------------
func parse_arguments(){
   flag.StringVar(&cmdParameters.LogLevel, "log-level", "INFO", "log level")
   flag.BoolVar(&cmdParameters.ToConsole, "console", false, "output to console")
   flag.StringVar(&cmdParameters.OutputFolder, "output", "", "output folder")
   flag.BoolVar(&cmdParameters.NoTimestamp, "no-timestamp", false, "skip timestamp in folder name")
   flag.BoolVar(&cmdParameters.NoLackOfData, "no-lack", false, "do not print '[lack of data]' messages")
   flag.BoolVar(&cmdParameters.CompareConfigProps, "compare-config-props", false, "compare configs, i.e. multiline properties")

   flag.StringVar(&cmdParameters.ClustersFilter, "clusters", "*", "list of clusters")

   flag.StringVar(&cmdParameters.ConfigFile, "config", "", "config file with defaults")
   flag.StringVar(&cmdParameters.MatchingPatterns, "mpatterns", "", "folder with matching patterns")

   flag.StringVar(&cmdParameters.ConfigsFilter, "configs", "*", "list of HDP configs to be compared")

   var ShowHelp = flag.Bool("help", false, "show help")
   flag.Parse()

   if *ShowHelp == true {
      flag.PrintDefaults()
      os.Exit(0)
   }


   cmdParameters.LogLevel = strings.ToLower(cmdParameters.LogLevel)

   file, err := ioutil.ReadFile(cmdParameters.ConfigFile)
   if err != nil {
      error("Error reading file: " + cmdParameters.ConfigFile)
   }

   var cl []Cluster
   if json.Unmarshal(file, &cl) != nil {
      error("Cannot unmarshal cluster config file: ", cmdParameters.ConfigFile)
   }

   if cmdParameters.ClustersFilter != "*" {
      for _, v := range strings.Split(cmdParameters.ClustersFilter, ",") {
         for _, cl := range cl {
            if cl.Name == v {
               cmdParameters.Clusters = append(cmdParameters.Clusters, cl)
            }
         }
      }
   } else {
      cmdParameters.Clusters = cl
   }

   if ! cmdParameters.NoTimestamp && cmdParameters.OutputFolder != "" {
      cmdParameters.OutputFolder = cmdParameters.OutputFolder + "/" + time.Now().Format("2006-01-02_15-04-05")
   }
}


// ----------------------------------------------------------------------------------------------------------
func init(){
   parse_arguments()
}


func main(){

   // TODO:
   // 
   // 1. git diff between clusters
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


   fmt.Println(cmdParameters)

   info("Completed. OK!")
}