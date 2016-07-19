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
         info(cfg)
         properties := getConfigProperties(cfg, cl)
         // warning(properties)
         propertiesInfo = mergeProperties(propertiesInfo, properties, cfg.Name, cl.Name)
      }
   }

   compareProperties(propertiesInfo)
}


// -----------------------------------------------------------------------------------
func mergeProperties(propertiesInfo []PropertyInfo, properties []Property, configName string, clusterName string)([]PropertyInfo){
   for _, prop := range properties {
      index := -1
      propValue := PropertyValue{ClusterName: clusterName, Value: prop.Value}
      for i, propInfo := range propertiesInfo {
         if propInfo.PropName == prop.Key && propInfo.ConfigName == configName {
            propertiesInfo[i].Values = append(propertiesInfo[i].Values, propValue)
            index = i
            break
         }
      }

      if index == -1 {
         v := PropertyInfo{ConfigName: configName, PropName: prop.Key, Values: []PropertyValue{propValue}}
         propertiesInfo = append(propertiesInfo, v)
      }
   }

   return propertiesInfo
}



func compareProperties(propertiesInfo []PropertyInfo) {
   info("compare them finally")
//    amountOfClusters := len(cmdParameters.Clusters)

// // func processingProperties(properties []Property, amountOfClusters int, config Config){
//    var matchedPatterns []Patterns

//    matchedPatterns = uploadMatchingFolder("_default_", matchedPatterns)
//    // matchedPatterns = uploadMatchingFolder(config.Name, matchedPatterns)

//    ignoredPatterns := uploadListFromFile("_ignores_")
//    warning("ignorant properties: ", ignoredPatterns)
//    warning("matching patterns: ", matchedPatterns)

//    for _, propInfo := range propertiesInfo {
//       if tryToFindIgnoredPatterns(propInfo.Name, ignoredPatterns) {
//          log.Debug("Ignoring property: ", propInfo.Name)
//          continue
//       }

//       // if strings.Contains(propInfo.Info[0].Value, "\n") && input.CompareConfigProps == false {
//       //    log.Debug("Skipping config property: ", propInfo.Name)
//       //    continue
//       // }


//       // hasTheSameValues := tryToFndTheDifference(propInfo, matchedPatterns)

//       // if amountOfClusters != len(propInfo.Info) {
//       //    if input.NoLackOfData == false {
//       //       outputPropertyDiff("[LACK OF DATA] ", propInfo, config.Name)
//       //    }
//       // } else if !hasTheSameValues {
//       //    outputPropertyDiff("[NOT THE SAME] ", propInfo, config.Name)
//       // }

//    }
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
