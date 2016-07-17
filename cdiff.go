package main;


// import (
//    "encoding/json"
//    "fmt"
// )


type ConfigProperty struct{
   ClusterName string
   ConfigName string
   Properties []Property
}



// type ConfigInfo struct {
//    Name string
//    Configs []Config
// }


// type PropertyInfo struct{
//    Name string
//    Properties []Property
// }


type PropertyValue struct{
   ClusterName string
   Value string
}

type PropertyInfo struct{
   ConfigName string
   PropName string
   Values []PropertyValue
}




// -----------------------------------------------------------------------------------
func getDiffBetweenClusters(){

   // configsInfo := mergingConfigsIntoOnce()
   var propertiesInfo []PropertyInfo

   for _, cl := range cmdParameters.Clusters {
      configs := getClusterConfigs(cl)
      for _, cfg := range configs {
         properties := getConfigProperties(cfg, cl)
         // warning(properties)
         propertiesInfo = mergeProperties(propertiesInfo, properties, cfg.Name, cl.Name)
      }
   }

   compareProperties(propertiesInfo)
}


// -----------------------------------------------------------------------------------
func mergeProperties(propertiesInfo []PropertyInfo, properties []Property, ConfigName string, ClusterName string)([]PropertyInfo){
   // for _, prop := range properties {

   // }

   return propertiesInfo
}



func compareProperties(propertiesInfo []PropertyInfo) {
   info("compare them finally")
}

// // -----------------------------------------------------------------------------------
// func mergingConfigsIntoOnce() ([]ConfigInfo){
//    var configsInfo []ConfigInfo

//    for _, cl := range cmdParameters.Clusters {
//       tmp := getClusterConfigs(cl)
//       for _, cfg := range tmp {
//          index := -1
//          for i, config := range configsInfo {
//             if config.Name == cfg.Name {
//                index = i
//                break
//             } 
//          }

//          if (index == -1){
//             configsInfo = append(configsInfo, ConfigInfo{Name: cfg.Name, Configs: []Config{cfg}})
//          } else {
//             configsInfo[index].Configs = append(configsInfo[index].Configs, cfg)
//          }
//       }
//    }   

//    return configsInfo
// }


// func getProperties(config Config)([]Property){
//    var properties []Property

//    for _, configInfo := range config.Info {
//       jsonData := sendAmbariRequest(getClusterByName(clusters, config.ClusterName), "/api/v1/clusters/" + configInfo.ClusterName + "/configurations?type=" + config.Name + "&tag=" + configInfo.Tag)

//       var data map[string]interface{}
//       json.Unmarshal(jsonData, &data)

//       if len(data["items"].([]interface{})) == 0 {
//          log.Warning("Seems there is no config for this env...")
//          continue
//       }

//       if data["items"].([]interface{})[0].(map[string]interface{})["properties"] != nil {
//          rawProps := data["items"].([]interface{})[0].(map[string]interface{})["properties"].(map[string]interface{})
//          properties = composePropertiesList(rawProps, properties, configInfo.ClusterName)
//       }
//    }

//    return properties
// }


// // -----------------------------------------------------------------------------------
// func compareConfigs(configs []Config){
//    info(configs)
// }
