# mx-bank
A payment gateway system

docker exec -it tx-svc-kafka-1 bash

CONSOLE PRODUCER:

kafka-console-producer  --bootstrap-server=localhost:9092 --topic=transaction
kafka-console-producer  --bootstrap-server=localhost:29092 --topic=transaction

CONSOLE CONSUMER:

kafka-console-consumer --bootstrap-server=localhost:9092 --topic=transaction_result
kafka-console-consumer --bootstrap-server=localhost:29092 --topic=transaction_result

SAMPLE MESSAGE:

{ "id": "1234", "account_id": "1", "credit_card_number": "1234123412341234", "credit_card_name": "Example", "credit_card_expiration_montH": 12, "credit_card_expiration_year": 2025, "credit_card_cvv": 123, "amount": 1200 }




echo $(set -o pipefail && go test ./... | { grep -v 'no test files'; true; } && set +o pipefail) | sed 's/ok /\nok /g'