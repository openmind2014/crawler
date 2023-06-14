# crawler

## Basic Usage

    docker run -d -p 9200:9200 elasticsearch

    go run distributed/persist/server/itemsaver.go -port 9000
    go run distributed/worker/server/worker.go -port 9001
    go run distributed/worker/server/worker.go -port 9002
    go run distributed/worker/server/worker.go -port 9003
    go run distributed/worker/server/worker.go -port 9004

    go run main.go -itemsaver_host=":9000" -worker_hosts=":9001,:9002,:9003,:9004"

    http://localhost:9200/dating_profile/
