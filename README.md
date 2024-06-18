# data-api-cis
Data-api CIS

** Run -> docker compose up --build 

* Postman collection

{
	"info": {
		"_postman_id": "c74e61b1-b64f-4444-91cb-fefb8b8f7724",
		"name": "[ENTREVISTA] CIS - Go",
		"schema": "https://schema.getpostman.com/json/collection/v2.1.0/collection.json",
		"_exporter_id": "17815795"
	},
	"item": [
		{
			"name": "User",
			"item": [
				{
					"name": "[GET] -  User",
					"request": {
						"method": "GET",
						"header": [],
						"url": {
							"raw": "http://localhost:8080/users/1/",
							"protocol": "http",
							"host": [
								"localhost"
							],
							"port": "8080",
							"path": [
								"users",
								"1",
								""
							]
						}
					},
					"response": []
				},
				{
					"name": "[POST] - User",
					"request": {
						"method": "POST",
						"header": [],
						"body": {
							"mode": "raw",
							"raw": "{\r\n  \"email\": \"exemplo@dominio.com\",\r\n  \"username\": \"usuarioexemplo\",\r\n  \"password\": \"senhasegura\",\r\n  \"first_name\": \"Nome\",\r\n  \"last_name\": \"Sobrenome\"\r\n}\r\n",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://localhost:8080/users",
							"protocol": "http",
							"host": [
								"localhost"
							],
							"port": "8080",
							"path": [
								"users"
							]
						}
					},
					"response": []
				},
				{
					"name": "[DELETE] -  User",
					"request": {
						"method": "DELETE",
						"header": [],
						"url": {
							"raw": "http://localhost:8080/users/1/",
							"protocol": "http",
							"host": [
								"localhost"
							],
							"port": "8080",
							"path": [
								"users",
								"1",
								""
							]
						}
					},
					"response": []
				}
			]
		}
	]
}