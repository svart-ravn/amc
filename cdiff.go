package main;



type ConfigProperty struct{
   ClusterName string
   ConfigName string
   Properties []Property
}



type ConfigInfo struct {
   Name string
   Configs []Config
}



// -----------------------------------------------------------------------------------
func getDiffBetweenClusters(){

   var configs []ConfigInfo

   for _, cl := range cmdParameters.Clusters {
      tmp := getClusterConfigs(cl)
      for _, cfg := range tmp {
         index := -1
         for i, config := range configs {
            if config.Name == cfg.Name {
               index = i
               break
            } 
         }

         if (index == -1){
            configs = append(configs, ConfigInfo{Name: cfg.Name, Configs: []Config{cfg}})
         } else {
            configs[index].Configs = append(configs[index].Configs, cfg)
         }
      }
   }

   // info("---------------------------------")

   // for _, ci := range configs{
   //    info(ci)
   //    info("")
   // }

   // removeUnusedClusters()

   // var configs []Config

   // for _, cl := range cmdParameters.Clusters {
   //    tmp := getClusterConfigs(cl)

   //    // info(configs)
   //    // var ConfigProperties []ConfigProperty{ClusterName: cl.Name, ConfigName: }
   //    // for _, cfg := range configs {
   //    //    info(cfg)
   //    //    info(getConfigProperties(cfg, cl))
   //    //    info("-----------------------")
   //    // }
   // }

}



// -----------------------------------------------------------------------------------

