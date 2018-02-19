# GettSkeleton golang
![alt text](https://cdn-images-2.medium.com/max/1200/1*AemYIFm92tl5RW9nBzNSAw.jpeg "")

_(Note: To update the skeleton script make a pull request to [init_new_service.sh](https://github.com/gtforge/gett-skeleton-go/blob/install/init_new_service.sh) )_

Easily set-up a new golang micro service

Just run the following in 
####GOPATH/src
```bash
source <(curl -s -L https://git.io/vHkjM) YOUR_SERVICE_NAME
```

After the script finishes, you can run `./install.sh` to create a database, run migrations and seed users

What the script does:
   1. Clone the latest version of this repository, using beego.
   2. Name your service appropriately
   3. Seed basic data
   4. Redis configuration
   5. DB configuration
   6. Make Dockerfile
   7. Make PIZZA!
  
Project plan
   1. conf => configuration files
   2. controllers => all controllers with base controller embedded.
   3. lib => business logic.
   4. models => objects with methods
   5. routes => routes + filter + other webserver and beego settings
   6. tests => api tests for service.

For run in dev env use https://github.com/beego/bee

```
go get github.com/beego/bee
bee run -gendoc=true
```


go dep for dependency manager https://github.com/golang/dep
```
go get -u github.com/golang/dep/cmd/dep
```

Code style:
   1. golint + go vet + gofmt => your best friends.
   2. gometalinter => your other best friend.
   3. https://github.com/gtforge/services_common_go => your family 
   4. Don't use beego orm, Use gettStorages.DB(work with gorm)
   5. Write comments in controllers for generate swagger
   6. If you don't need db or redis => remove dbconf.yml or redis.yml
   7. Don't be harry potter, less magic makes things simple.
Setting up a new service?
If there is something basic that is missing from this starter template, please add it to the template so everyone could enjoy

### Target readme file (Template) is located inside [README_TEMPLATE.md](README_TEMPLATE.md)  
After creating the skeleton it will become the project's README.md file

Baseline requirements: 
   fill the 2 jsons in db folder according your service:
   tablesDataHandling.json ==> write inside which tables' data is irrelevant for local environments
   tableConnections.json ==> write which one from the main entities you are referencing in your table
