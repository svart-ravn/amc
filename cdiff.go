package main;



// -----------------------------------------------------------------------------------
func getDiffBetweenClusters(){

   removeUnusedClusters()

   for _, cl := range clusterList {
      info(getClusterConfigs(cl))
   }

//    var configList []ConfigInfo
//    for _, cl := range clusterList {
//       configList = mergeConfigs(cl, configList)
//    }

//    for _, cfg := range configList {
//       properties := getConfigProperties(cfg, clusterList)
//       compareProperties(properties)
//    }
}



// -----------------------------------------------------------------------------------
func getConfigProperties(cfg Config, ){

}


// -----------------------------------------------------------------------------------
func removeUnusedClusters(){

}



// -----------------------------------------------------------------------------------
// func mergeConfigs(cl Cluster, configList []Config)([]Config){
//    info("cluster name: ", cl.Name)
//    var jsonData []byte
//    jsonData = sendAmbariRequest(cl, "/api/v1/clusters/" + cluster.Name + "?fields=Clusters/desired_configs")

//    var data map[string]interface{}
//    json.Unmarshal(jsonData, &data)
//    desiredConfigs := data["Clusters"].(map[string]interface{})["desired_configs"].(map[string]interface{})

//    return composeConfigList(desiredConfigs, configs, cluster.Name)




//    return configList
// }