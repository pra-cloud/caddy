package main

import (
        "fmt"
        "net/http"
)

func main() {

        // allowedDomains := []string{"caddy.hyperzod.xyz", "devops1.hyperzod.xyz"}

        http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

                if r.Method == "GET" {

                        query := r.URL.Query()

                        requestDomain := query.Get("domain")

                        if requestDomain == "" {
                                w.WriteHeader(http.StatusNotFound)
                                return
                        }

                        // Perform domain validation against your API
                        apiURL := fmt.Sprintf("https://api.hyperzod.dev/public/v1/tenant/validate?x-tenant=%s", requestDomain)
                        resp, err := http.Get(apiURL)
                        if err != nil {
                                w.WriteHeader(http.StatusInternalServerError)
                                fmt.Fprintf(w, "Failed to validate domain: %v", err)
                                return
                        }

                        defer resp.Body.Close()

                        if resp.StatusCode == http.StatusOK {
                                w.WriteHeader(http.StatusOK)
                                fmt.Fprintf(w, "Domain is validated: %s", requestDomain)
                        } else {
                                w.WriteHeader(http.StatusForbidden)
                                fmt.Fprintf(w, "Domain validation failed: %s", requestDomain)
                        }
                }
                })
                // added for ssl key store
                {
                        storage redis {
                            address   "localhost:6379"     # Address of the Redis server
                            password  ""                   # Password for Redis (if any)
                            db        0                    # Redis database number
                            key_prefix "caddy"             # Prefix for keys (optional)
                            timeout   5                    # Redis connection timeout (seconds)
                        }
                    }


                fmt.Println("TLS check server is running on port 5555")
                http.ListenAndServe(":5555", nil)






}
