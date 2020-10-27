# gin_recovery
Gin-gonic recovery middleware with handler function and google error reporting

``` go get -u github.com/ihassanyousaf/gin_recovery ```

## Examples
In order to initialize GOOGLE ERROR REPORTING to your gin service you need to pass path of service account file in GOOGLE_APPLICATION_CREDENTIALS 
```` 
     package server     

     import (
     	"github.com/gin-gonic/gin"
     	"github.com/ihassanyousaf/gin_recovery"
     	"net/http"
     )
     
     func Server() *gin.Engine {
     	
        server := gin.New()
     
    
     	gin_recovery.New(gin_recovery.RecoveryConfig{
            ErrorReporting: &gin_recovery.ErrorReportingConfig{
     			ProjectId:   "PROJECT_ID",
     			ServiceName: "SERVICE_NAME",
     		}})
     	server.Use(gin_recovery.Recovery())
     	
     	//service health point
     	server.GET("/", func(c *gin.Context) {
     		c.JSON(http.StatusOK, gin.H{
     			"message": "healthy",
     		})
     	})

        server.GET("/panic" func(c *gin.Context) {
            panic("this is a test panic for gin_recovery")
        } 

        return server
     } 
````

Initialization of recovery middleware without GOOGLE ERROR REPORTING

```` 
package server     

     import (
     	"github.com/gin-gonic/gin"
     	"github.com/ihassanyousaf/gin_recovery"
     	"net/http"
     )
     
     func Server() *gin.Engine {
     	
        server := gin.New()
     
    
     	gin_recovery.New(gin_recovery.RecoveryConfig{})
     	server.Use(gin_recovery.Recovery())
     	
     	//service health point
     	server.GET("/", func(c *gin.Context) {
     		c.JSON(http.StatusOK, gin.H{
     			"message": "healthy",
     		})
     	}) 

        server.GET("/panic" func(c *gin.Context) {
            panic("this is a test panic for gin_recovery")
        }

        return server
     }
````