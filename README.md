# Kafka 

Implementing Kafka using Go and Docker.

## How To Run
1. Open terminal and run command:
    ```
    docker network create kafka
    ```
2. Pull and run zookeeper docker :
    ```
    docker run --net=kafka -d --name=zookeeper -e ZOOKEEPER_CLIENT_PORT=2181 confluentinc/cp-zookeeper:4.1.0
    ```
3. Pull and run kafka docker :
    ```
    docker run --net=kafka -d -p 9092:9092 --name=kafka -e KAFKA_ZOOKEEPER_CONNECT=zookeeper:2181 -e KAFKA_ADVERTISED_LISTENERS=PLAINTEXT://kafka:9092 -e KAFKA_OFFSETS_TOPIC_REPLICATION_FACTOR=1 confluentinc/cp-kafka:4.1.0
    ```   
4. Go to DNS Server configuration :
    ```
    sudo vim /etc/hosts 
    ```
    Set DNS Kafka at file 'hosts':
    ```
    127.0.0.1 kafka
    ```
5. Run Consumer :
    ```
    go run cmd/consumer/main.go
    ```    
6. Run Producer or Publisher :
    ```
    go run cmd/publisher/main.go
    ```

