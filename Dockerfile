FROM 085551151339.dkr.ecr.eu-west-1.amazonaws.com/base_beego:builder as builder #https://github.com/gtforge/docker_files/blob/master/golang/multistage/Dockerfile.builder
FROM 085551151339.dkr.ecr.eu-west-1.amazonaws.com/base_beego:runtime #https://github.com/gtforge/docker_files/blob/master/golang/multistage/Dockerfile.runtime
