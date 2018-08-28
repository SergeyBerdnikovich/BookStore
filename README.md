# GettSkeleton for Go projects

![alt text](https://cdn-images-2.medium.com/max/1200/1*AemYIFm92tl5RW9nBzNSAw.jpeg "")

_(Note: To update the skeleton script make a pull request to [init_new_service.sh](https://github.com/gtforge/gett-skeleton-go/blob/master/init_new_service.sh) )_

Easily set up a new Golang microservice.

Just run the following in `$GOPATH/src` folder:

```bash
source <(curl -s -L https://git.io/vx70A) YOUR_SERVICE_NAME
```

After the script finishes, you can run `./install.sh` to create a database, run migrations and seed users.

What the script does:

1. Clones the latest version of this repository, using `beego`.
2. Names your service appropriately
3. Seeds basic data
4. Setups Redis configuration
5. Setups DB configuration (Postgres)
6. Makes Dockerfile
7. Makes PIZZA!
  
Project plan:

1. conf => configuration files
2. controllers => all controllers with base controller embedded.
3. lib => business logic.
4. models => objects with methods
5. routes => routes + filter + other webserver and beego settings
6. tests => api tests for service.

To run in `dev` env, use the [Bee](https://github.com/beego/bee) tool:

```bash
go get github.com/beego/bee
bee run -gendoc=true
```

Use [dep](https://github.com/golang/dep) tool for dependency management:

```bash
go get -u github.com/golang/dep/cmd/dep
```

Code style:

1. `golint` + `go vet` + `gofmt`/`goimports` => your best friends
2. `gometalinter` => your other best friend
3. <https://github.com/gtforge/services_common_go> => your family
4. Don't use beego ORM - use `gettStorages.DB` instead (through [Gorm](https://github.com/jinzhu/gorm))
5. Write comments in controllers for proper `swagger` generation
6. If you don't need Postgres or Redis => remove `dbconf.yml` or `redis.yml`
7. Don't be Harry Potter: less magic makes things simple.

Setting up a new service?

If there is something basic that is missing from this starter template, please add it to the template so everyone could enjoy

**Target README file template is located inside [README_TEMPLATE.md](README_TEMPLATE.md)**

After creating the skeleton it will become the project's README.md file

Baseline requirements:

Fill the 2 JSONs in `db` folder accordingly in your service:

- `tablesDataHandling.json` ==> write inside which tables' data is irrelevant for local environments
- `tableConnections.json` ==> write which one from the main entities you are referencing in your table
