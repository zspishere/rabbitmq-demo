
# Go code for RabbitMQ tutorials


Here you can find Go code examples from [RabbitMQ tutorials](https://www.rabbitmq.com/getstarted.html).


## Requirements

To run this code you need [Go RabbitMQ client](https://github.com/streadway/amqp):

    go get github.com/streadway/amqp

And run a rabbitmq server:

    docker-compose up -d

## Code

Code examples are executed via `go run`:

[Tutorial one: "Hello World!"](https://www.rabbitmq.com/tutorial-one-go.html):

    go run demo1/send.go
    go run demo1/receive.go

[Tutorial two: Work Queues](https://www.rabbitmq.com/tutorial-two-go.html):

    go run demo2/new_task.go hello world
    go run demo2/worker.go

[Tutorial three: Publish/Subscribe](https://www.rabbitmq.com/tutorial-three-go.html)

    go run demo3/receive_logs.go
    go run demo3/emit_log.go hello world

[Tutorial four: Routing](https://www.rabbitmq.com/tutorial-four-go.html)

    go run demo4/receive_logs_direct.go info warn
    go run demo4/emit_log_direct.go warn "a warning"

[Tutorial five: Topics](https://www.rabbitmq.com/tutorial-five-go.html)

    go run demo5/receive_logs_topic.go "kern.*" "*.critical"
    go run demo5/emit_log_topic.go kern.critical "A critical kernel error"

[Tutorial six: RPC](https://www.rabbitmq.com/tutorial-six-go.html)

    go run demo6/rpc_server.go
    go run demo6/rpc_client.go 10

To learn more, see [Go RabbitMQ client](https://github.com/streadway/amqp).