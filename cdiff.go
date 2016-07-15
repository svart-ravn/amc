package main;



// -----------------------------------------------------------------------------------
func getDiffBetweenClusters(){

   removeUnusedClusters()

   for _, cl := range clusterList {
      configs := getClusterConfigs(cl)

      for _, cfg := range configs {
         info(cfg)
         info(getConfigProperties(cfg, cl))
         info("-----------------------")
      }
   }

}



// -----------------------------------------------------------------------------------
// func getConfigProperties(cfg Config, ){

// }


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