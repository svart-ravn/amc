package main;


import ("fmt")



// -----------------------------------------------------------------------------------
func getDiffBetweenClusters(){

   removeUnusedClusters()

   var configList []ConfigInfo
   for _, cl := range clusterList {
      configList = mergeConfigs(cl, configList)
   }

   // info("get_config_diff")
   // fmt.Println(cmdParameters)
}



// -----------------------------------------------------------------------------------
func removeUnusedClusters(){

}



// -----------------------------------------------------------------------------------
func mergeConfigs(cl Cluster, configList []ConfigInfo)([]ConfigInfo){
   return configList
}