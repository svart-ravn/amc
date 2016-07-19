package main;


import (
   "encoding/json"
   "io/ioutil"
   "reflect"
   "strings"
   "os"
   "bufio"
)



// ---------------------------------------------------------------------------
type Config struct{
   ClusterName string
   Name string
   Tag string
   User string
   Version float64 
}


type Cluster struct{
   AmbariUrl string
   AmbariPassword string
   AmbariUser string
   Name string
}


type Property struct{
   Key string
   Value string
}



// ---------------------------------------------------------------------------
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




// ---------------------------------------------------------------------------
type Patterns struct{
   ReplaceWith string
   What []string
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



// --------------------------------------------------------------------------------------------------------------
func fillStruct(data map[string]interface{}, result interface{}){
   t := reflect.ValueOf(result).Elem()
   for k_lower, v := range data {
      k := strings.Title(k_lower)
      val := t.FieldByName(k)
      val.Set(reflect.ValueOf(v))
   }
}


// ----------------------------------------------------------------
func getClusterByName(clusters []Cluster, name string)(Cluster){
   for _, cl := range clusters{
      if cl.Name == name {
         return cl
      }
   }

   return Cluster{}
}



// ------------------------------------------------------------------------------
func uploadMatchingFolder(filename string, matchedPatterns []Patterns) ([]Patterns){
   if file, err := os.Open("configs/" + filename); err == nil {
      defer file.Close()

      scanner := bufio.NewScanner(file)
      for scanner.Scan() {
         data := strings.Split(scanner.Text(), "=")
         matchedPatterns = append(matchedPatterns, Patterns{data[0], strings.Split(data[1], ",")})
      }     
   }

   return matchedPatterns
}



// ------------------------------------------------------------------------------
func uploadListFromFile(filename string) []string{
   result := make([]string, 0)
   if file, err := os.Open("configs/" + filename); err == nil {
      defer file.Close()

      scanner := bufio.NewScanner(file)
      for scanner.Scan() {
         result = append(result, scanner.Text())
      }     
   }

   return result
}
