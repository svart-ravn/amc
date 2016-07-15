package main;


import (
   "net/http"
   "crypto/tls"
   "io/ioutil"
   "encoding/json"
)


// ------------------------------------------------------------------------------------------
func sendAmbariRequest(cl Cluster, url string) ([]byte){
   // log.Debug("sending request: ", cluster.AmbariUrl + url)


   tr := &http.Transport{
      TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
   }

   req, _ := http.NewRequest("GET", cl.AmbariUrl + "/api/v1/clusters/" + cl.Name + url, nil); 
   req.SetBasicAuth(cl.AmbariUser, cl.AmbariPassword)
   req.Header.Set("X-Requested-By", "Ambari")

   client := http.Client{Transport: tr}
   resp, err := client.Do(req)
   // if err != nil {
   //    log.Panic("Cannot process request:", err)
   // }

   body, err := ioutil.ReadAll(resp.Body)
   resp.Body.Close()
   if err != nil {
      panic(err)
   }

   return body 
}



// ------------------------------------------------------------------------------------------
func getClusterConfigs(cl Cluster) ([]Config){
   info("cluster name: ", cl.Name)
   var jsonData []byte
   jsonData = sendAmbariRequest(cl, "?fields=Clusters/desired_configs")

   var data map[string]interface{}
   json.Unmarshal(jsonData, &data)
   desiredConfigs := data["Clusters"].(map[string]interface{})["desired_configs"].(map[string]interface{})

   return composeConfigList(desiredConfigs, cl.Name)
}



// ------------------------------------------------------------------------------------------
func composeConfigList(data map[string]interface{}, clusterName string)([]Config){
   var configs []Config;

   for configName := range data {
      cfg := &Config{Name: configName, ClusterName: clusterName}
      fillStruct(data[configName].(map[string]interface{}), cfg)
      configs = append(configs, *cfg)
   }

   return configs
}


// // ------------------------------------------------------------------------------------------
// func get_config_values(){
// }



// // ------------------------------------------------------------------------------------------
// func get_config_versions(){
// }

