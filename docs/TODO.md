# Tik-shop TODO

Here are some subtasks.

## Route service

### route API

- [ ] Admin API.
    - [ ] thrift API.
    - [ ] functions just place in route.

### route middleware

- [ ] Auth Middleware.
    - [ ] write JWT Token shared code.
    - [ ] decide weather using User service to auth JWT or just place them in route.
    - [ ] implement JWT Token Auth logic for user and add to middleware.
    - [ ] same auth to admin, middleware.

### connect route with RPC service

- [ ] Register each RPC service to etcd.
    - [X] register with fixed port.
    - [ ] allow using Environment variable to listen ports(see below).
- [X] Get and init each RPC connection in route.
    - [X] get and init with service name in etcd.

## DB Related logic

### Connect to DB

- [x] write struct to store in gorm.
- [x] connect to DB and do some manual test.
- [x] warp a model unit for rpc service unit to call. Better build with multiple isolated model for each unit, and connect to same db sharing(? maybe shouldn't) connection.
- [x] let each rpc unit connect to db.

## ENV Management

- [ ] Need to decide how to manage env variables, like each rpc service unit's listen ports, or where to find db for gorm.