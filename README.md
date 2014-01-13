RaysGo
======

A  generic blog system implemented in Go!

* === Under development, learning process!
* === It's my first development experience with Go! It's exciting to learn and use Go!

## Installation
Please make sure you have installed Go (v>=1.1) and set up your GOPATH.
* Get the project
```
$ go get github.com/Raysmond/RaysGo
```

* Switch to the project
```
$ cd $GOPATH/src/github.com/Raysmond/RaysGo
```

* Build and run
```
$ go build 
$ ./app
```
Then, you can view the website in the browser: `http://localhost:8080`. Currently, the application uses MySQL for storage and when you run `./app` command, the corresponding database will be created and initialized according to the configuration file [app.conf](https://github.com/Raysmond/RaysGo/blob/master/conf/app.conf).


## Goal
A common blog system supports following features:
* User registration and login, authorization, and etc..
* A user can publish posts and comments
* A user can manage his/her own posts, comments, profile, and etc...
* Administrators can manage all contents within the whole blog system
* Basic statistic analytical tools
  * content view counters
  * content rating
  * content counting for administration
  * etc
* Content editor types: `text`, `html`, `markdown`
* Other functions in a blog system

## References
* [beego](https://github.com/astaxie/beego) : An open-source, high-performance web framework for the Go
* [bootstrap](http://getbootstrap.com/) : A front-end framework
* [xorm](https://github.com/lunny/xorm) : A Simple and Powerful ORM for Go
* [Blackfriday]() : a Markdown processor implemented in Go


