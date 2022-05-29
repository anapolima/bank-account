# Bank Account

Sistema de contas bancárias. Funcionalidades:
- Depósito, taxa de 1% sobre o valor depositado
- Saque, taxa fixa de R$4,00
- Transferência, taxa fixa de R$1,00
- Consulta de extrato, gratuito
- Criação de conta, gratuito

## Executando o projeto
### Requisitos
- Docker
- Docker Compose

Após clonar o repositório, em um terminal, para *primeira execução* do projeto execute o comando
```shell
make run-build
```

Após, pode-se utilizar o comando
```shell
make run
```

Para resetar o projeto, excluindo todos os dados do banco de dados, execute o comando
```shell
make reset
```

Sempre que uma request for executada, o conteúdo de sua response será um objeto com os campos `data` e `messages`, se uma request for executada com erro, `messages` conterá um array de strings com as mensagens de erro obtidas durante o processamento da requisição, caso contrário, `data` possuirá os dados de sucesso da request

## Criando uma conta
Para criar uma conta, execute uma request do tipo `POST` para o endpoint `/create-account` enviando os seguintes dados
```json
{
	"name": string,
	"document": string,
	"birthdate": string,
	"accountPassword": string
}
```

Exemplo de request:
```
curl --request POST \
  --url http://localhost:8080/create-account \
  --header 'Content-Type: application/json' \
  --data '{
	"name": "Account Owner",
	"document": "76763008007",
	"birthdate": "1997-12-01",
	"accountPassword": "2222"
}'
```

Exemplo de response:
```json
{
	"data": {
		"agencyNumber": 4570,
		"agencyVerificationCode": 8,
		"accountNumber": 985684,
		"accountVerificationCode": 7,
		"owner": "Account Owner",
		"document": "76763008007",
		"birthdate": "1997-11-21"
	},
	"messages": []
}
```

## Fazendo um depósito
Para fazer um depósito, execute uma request do tipo `POST` para o endpoint `/deposit` enviando os seguintes dados
```json
{
	"account": {
		"agencyNumber": number,
		"agencyVerificationCode": number,
		"accountNumber": number,
		"accountVerificationCode": number,
		"document": string
	},
	"value": number
}
```

Exemplo de request:
```
curl --request POST \
  --url http://localhost:8080/deposit \
  --header 'Content-Type: application/json' \
  --data '{
	"account": {
		"agencyNumber": 4570,
		"agencyVerificationCode": 8,
		"accountNumber": 985684,
		"accountVerificationCode": 7,
		"document": "76763008007"
	},
	"value": 100
}'
```

Exemplo de response:
```json
{
	"data": {
		"transactionId": "7b556cf8-c03c-4223-988d-b9e5dabf1576",
		"type": "deposit",
		"value": 100,
		"date": "2022-05-29T01:28:04.44379Z",
		"account": {
			"agencyNumber": 4570,
			"agencyVerificationCode": 8,
			"accountNumber": 985684,
			"accountVerificationCode": 7,
			"owner": "Account Owner",
			"document": "76763008007"
		}
	},
	"messages": []
}
```

## Fazendo um saque
Para fazer um saque, execute uma request do tipo `POST` para o endpoint `/draft` enviando os seguintes dados
```json
{
	"account": {
		"agencyNumber": number,
		"agencyVerificationCode": number,
		"accountNumber": number,
		"accountVerificationCode": number,
		"document": string,
		"accountPassword": string
	},
	"value": number
}
```

Exemplo de request:
```
curl --request POST \
  --url http://localhost:8080/draft \
  --header 'Content-Type: application/json' \
  --data '{
	"account": {
		"agencyNumber": 4570,
		"agencyVerificationCode": 8,
		"accountNumber": 985684,
		"accountVerificationCode": 7,
		"document": "76763008007",
		"accountPassword": "2222"
	},
	"value": 25.55
}'
```

Exemplo de response:
```json
{
	"data": {
		"transactionId": "4fdf83d3-f644-4d1f-893e-bc0d28bf2bcc",
		"type": "draft",
		"value": 25.55,
		"date": "2022-05-29T01:30:14.478597Z",
		"account": {
			"agencyNumber": 4570,
			"agencyVerificationCode": 8,
			"accountNumber": 985684,
			"accountVerificationCode": 7,
			"owner": "Account Owner",
			"document": "76763008007"
		}
	},
	"messages": []
}
```

## Fazendo uma transferência
Para fazer uma transferência entre contas, execute uma request do tipo `POST` para o endpoint `/transfer` enviando os seguintes dados
```json
{
	"originAccount": {
		"agencyNumber": number,
		"agencyVerificationCode": number,
		"accountNumber": number,
		"accountVerificationCode": number,
		"document": string,
		"accountPassword": string
	},
	"destinyAccount": {
		"agencyNumber": number,
		"agencyVerificationCode": number,
		"accountNumber": number,
		"accountVerificationCode": number,
		"document": string
	},
	"value": number
}
```

Exemplo de request:
```
curl --request POST \
  --url http://localhost:8080/transfer \
  --header 'Content-Type: application/json' \
  --data '{
	"originAccount": {
		"agencyNumber": 4570,
		"agencyVerificationCode": 8,
		"accountNumber": 985684,
		"accountVerificationCode": 7,
		"document": "76763008007",
		"accountPassword": "2222"
	},
	"destinyAccount": {
		"agencyNumber": 9619,
		"agencyVerificationCode": 3,
		"accountNumber": 749881,
		"accountVerificationCode": 8,
		"document": "96001284059"
	},
	"value": 15.51
}'
```

Exemplo de response:
```json
{
	"data": {
		"transactionId": "700e89dd-db8e-4d68-b284-74f8c96554e0",
		"type": "transfer",
		"value": 15.51,
		"date": "2022-05-29T01:33:42.435212Z",
		"originAccount": {
			"agencyNumber": 4570,
			"agencyVerificationCode": 8,
			"accountNumber": 985684,
			"accountVerificationCode": 7,
			"document": "76763008007"
		},
		"destinyAccount": {
			"agencyNumber": 9619,
			"agencyVerificationCode": 3,
			"accountNumber": 749881,
			"accountVerificationCode": 8,
			"document": "96001284059"
		}
	},
	"messages": []
}
```

## Consultando o extrato
Para consultar o extrato de uma conta, execute uma request do tipo `GET` para o endpoint `/extract` enviando os seguintes dados
```http
agencyNumber=
agencyVerificationCode=
accountNumber=
accountVerificationCode=
document=
```

Exemplo de request:
```
curl --request GET \
  --url 'http://localhost:8080/extract?agencyNumber=4570&agencyVerificationCode=8&accountNumber=985684&accountVerificationCode=7&document=76763008007'
```

Exemplo de response:
```json
{
	"data": {
		"agencyNumber": 4570,
		"agencyVerificationCode": 8,
		"accountNumber": 985684,
		"accountVerificationCode": 7,
		"owner": "Account Owner",
		"document": "76763008007",
		"birthdate": "1997-11-21T00:00:00Z",
		"balance": 52.94,
		"transactions": [
			{
				"transactionId": "9a30e6d4-d035-4a41-9d8c-c0832f4cd6b9",
				"type": "transfer fee",
				"value": -1,
				"date": "2022-05-29T01:33:42.453601Z"
			},
			{
				"transactionId": "700e89dd-db8e-4d68-b284-74f8c96554e0",
				"type": "transfer",
				"value": -15.51,
				"date": "2022-05-29T01:33:42.435212Z"
			},
			{
				"transactionId": "9303f20d-2228-4b9c-8ffe-3ac731a8abbe",
				"type": "draft fee",
				"value": -4,
				"date": "2022-05-29T01:30:14.498643Z"
			},
			{
				"transactionId": "4fdf83d3-f644-4d1f-893e-bc0d28bf2bcc",
				"type": "draft",
				"value": -25.55,
				"date": "2022-05-29T01:30:14.478597Z"
			},
			{
				"transactionId": "5d27d6df-69e5-43be-be2f-7cd42e010d0e",
				"type": "deposit fee",
				"value": -1,
				"date": "2022-05-29T01:28:04.466963Z"
			},
			{
				"transactionId": "7b556cf8-c03c-4223-988d-b9e5dabf1576",
				"type": "deposit",
				"value": 100,
				"date": "2022-05-29T01:28:04.44379Z"
			}
		]
	},
	"messages": []
}
```
