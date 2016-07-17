package main;


import (
   "net/http"
   "crypto/tls"
   "io/ioutil"
   "encoding/json"
   "fmt"
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
func getConfigProperties(cfg Config, cl Cluster)([]Property){
   var properties []Property


   jsonData := sendAmbariRequest(cl, "/configurations?type=" + cfg.Name + "&tag=" + cfg.Tag)

   var data map[string]interface{}
   json.Unmarshal(jsonData, &data)

   if len(data["items"].([]interface{})) == 0 {
      warning("Seems there is no config", cfg.Name, " for this env...")
      return properties
   }

   if data["items"].([]interface{})[0].(map[string]interface{})["properties"] != nil {
      rawProps := data["items"].([]interface{})[0].(map[string]interface{})["properties"].(map[string]interface{})
      properties = composePropertiesList(rawProps)
   }


   return properties
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



// -----------------------------------------------------------------------------------------------------------------------
func composePropertiesList(data map[string]interface{})([]Property){
   var result []Property
   for key, value := range data{
      // info(key, value)
      v := fmt.Sprint(value)

      result = append(result, Property{Key: key, Value: v})

   }

   return result
}