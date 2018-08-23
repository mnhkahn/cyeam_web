// Package app
package app

// InitRouter ...
func InitRouter() {
	usr, pwd, _ := GetConfigAuth()
	Handle("/debug/goapp", BacisAuthHandler(&Got{H: GoAppHandler}, usr, pwd))
	Handle("/debug/router", &Got{H: DebugRouter})
	Handle("/debug/log/level", Got{LogLevelHandler})
}
