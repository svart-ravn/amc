package main;


import (
   "net/http"
   "crypto/tls"
   "io/ioutil"
)


// ------------------------------------------------------------------------------------------
func request(cl Cluster, url string) ([]byte){
   // log.Debug("sending request: ", cluster.AmbariUrl + url)


   tr := &http.Transport{
      TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
   }

   req, _ := http.NewRequest("GET", cl.AmbariUrl + url, nil); 
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
func get_configs(){
}



// ------------------------------------------------------------------------------------------
func get_config_values(){
}



// ------------------------------------------------------------------------------------------
func get_config_versions(){
}

