# GettSkeleton golang
![alt text](https://cdn-images-2.medium.com/max/1200/1*AemYIFm92tl5RW9nBzNSAw.jpeg "")

Easily set-up a new golang micro service

Just run the following in 
####GOPATH/src
```bash
curl -s -L http://bit.ly/2boMMLW | bash -s YOUR_SERVICE_NAME
```

and it will:
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
bee run -gendoc=true
```
glide for dependency manager
Troubleshooting with glide:
1. rm -rf ~/.glide/cache
2. rm glide.*
3. glide init && glide install

Code style:
1. golint + go vet + gofmt => your best friends.
2. gometalinter => your other best friend.
3. https://github.com/gtforge/services_common_go => your family 
4. Don't use beego orm, Use gettStorages.DB(work with gorm)
5. Write comments in controllers for generate swagger
7. If you don't need db or redis => remove dbconf.yml or redis.yml
8. Don't be harry potter, less magic more simples.

Setting up a new service?
If there is something basic that is missing from this starter template, please add it to the template so everyone could enjoy

To edit the install script:
https://gist.github.com/sergeylanzman/4e246edacc5c87aae85534500683cba8
