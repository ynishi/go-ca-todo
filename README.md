# go-ca-todo
A sample of clean architecture in go.

I referenced below and make this, very thanks.
* [The Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)
* go-clean-arch https://github.com/bxcodec/go-clean-arch

This sample is,
* Domain is simply struct based and public interface just use from Entity.
* Parser and Presenter (Interface Adapter layer) are implemented as function chain.
* It values independence of Domain, Entity is top of domain package and contains multiple application from shared logic.
* DB access repository is not implemented.
* Not support context, error, logger yet.
```
.
├── _examples -- build and run sample with shell.
│   ├── cmd.sh
│   └── server.sh
├── build.sh -- build script.
├── cmd -- for cli, interface adapter and framework layer
│   └── todo
│       ├── adapter -- customised interface adapters for cli
│       │   └── adapter.go
│       ├── main.go -- for cli
│       └── todo -- excutable may be built
├── internal -- sample repository.
│   └── pkg
│       └── repository
│           └── repository.go
├── pkg -- for domain and common interface adapter
│   └── todo
│       ├── adapter -- interface adapter layer
│       │   └── adapter.go
│       ├── interacter -- use case interacter layer
│       │   ├── interacter.go
│       │   └── interacter_test.go
│       ├── todo.go -- entity layer
│       └── todo_test.go
└── web -- for web server, Frameworks layer
    └── todo
        ├── adapter -- customised interface adapter for server
        │   └── adapter.go
        ├── main.go -- for server
        └── todo -- server excutable may be built
```

## How to run
* run shell script in _examples, it build and run sample command.

## build
```
sh build.sh
```

## LICENSE
MIT, See LICENSE.
