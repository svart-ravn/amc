package main;


import (
   "strings"
)


type ConfigProperty struct{
   ClusterName string
   ConfigName string
   Properties []Property
}



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

   var propertiesInfo []PropertyInfo

   for _, cl := range cmdParameters.Clusters {
      configs := getClusterConfigs(cl)
      warning("configs: ", configs)
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
   var matchedPatterns []Patterns
   ignoredPatterns := uploadListFromFile("_ignores_")

   configName := ""
   for _, propsInfo := range propertiesInfo {
      if propsInfo.ConfigName != configName {
         configName = propsInfo.ConfigName
         matchedPatterns = nil
         matchedPatterns = uploadMatchingFolder("_defaults_", matchedPatterns)
         matchedPatterns = uploadMatchingFolder(configName, matchedPatterns)
      }

      // ignored props
      if tryToFindIgnoredPatterns(propsInfo.PropName, ignoredPatterns) {
         debug("Ignoring property: ", propsInfo.PropName)
         continue
      }

      // skipping comments if required
      if strings.Contains(propsInfo.Values[0].Value, "\n") && cmdParameters.CompareConfigProps == false {
         debug("Skipping config property: ", propsInfo.PropName)
         continue
      }

      if ! cmdParameters.NoLackOfData && (len(propsInfo.Values) != len(cmdParameters.Clusters) ){
         warning("lack of data: ", propsInfo)
      } else {
         hasTheSameValues := tryToFndTheDifference(propsInfo, matchedPatterns)

         if hasTheSameValues {
            info(propsInfo)
         } else {
            error(propsInfo)
         }
      }
   }
}



func tryToFindIgnoredPatterns(propName string, ignoredPatterns []string) (bool){
   for _, v := range ignoredPatterns {
      if v == propName {
         return true
      }
   }
   return false
}



func tryToFndTheDifference(propsInfo PropertyInfo, matchedPatterns []Patterns) (bool){
   hasTheSameValues := true

   var standard string
   for ind, v := range propsInfo.Values {
      value := applyMatchedPatterns(v.Value, matchedPatterns, ind)
      if standard == "" {
         standard = value
      } else{
         if standard != value {
            hasTheSameValues = false
            break
         }
      }


   }

   return hasTheSameValues
}


func applyMatchedPatterns(value string, matchedPatterns []Patterns, ind int) (string){

   for _, pattern := range matchedPatterns{
      for _,v := range pattern.What {
         value = strings.Replace(value, v, pattern.ReplaceWith, -1)
      }
   }

   // debug
   // if strings.Contains(value, "0") && ind % 2 == 1 {
   //    return "xzy"
   // } else {
   return value
   // }
}