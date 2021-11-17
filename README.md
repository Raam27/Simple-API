# API for Keycloak SPI
## This API help migrate old database in the system to keycloak database using Keycloak SPI from (https://github.com/daniel-frak/keycloak-user-migration)

## need .env file for JWT secret and config.json in config folder for DB config

# Different users table :

### Change the user and keycloak model inside model folder to your database user table model

Change json return model :

# Change user.go inside api folder and keycloak model

### Current api still compare plain input to saved password in database, if you use bcrypt you can change the router.go inside route folder
```
e.POST("users/:username", api.Authenticate, middlewares.IsAuthenticated)
```
## Change api.Authenticate to controllers.CheckLogin and change the return to
```
return c.NoContent(http.StatusOK)
```
### folder structure based on https://github.com/markpenaranda/echo-gorm