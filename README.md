# messaggio_test
example programm for messaggio

## Подключение к работающему серверу (http://87.242.118.172:8050) :
у API прописано два роута:
 - POST “/” - для отправки сообщения в body запроса. В ответ приходит 200 OK и JSON {"message":"message was getting","errorbool":false}

 - GET “/getstat” - для получения статистики. В ответ приходит 200 OK и JSON {"all_messages":92,"processed_messages":28}

Пример отправки сообщения cURL:
```
curl --location 'http://87.242.118.172:8050/' \
--header 'Content-Type: text/plain' \
--data 'проверка'
```
Пример получения статистики cURL: 
```
curl --location 'http://87.242.118.172:8050/getstat'
```
## Docker:
 - Клонировать репозиторий 
 - git clone git@github.com:kthucydi/messaggio_test.git
 - Зайти в папку склонированного репозитория
 - запустить make set_env
 - запустить make up
 - для завершения работы: make down

после сборки может ещё минутку подумать:
 - доступ к постгресу по порту localhost:5531 и данным из .env в папке messaggio_test
 - API: localhost:8050
 - доступ к кафке: localhost:8090
 - логи доступны в папках: “/tmp/logs_second_handler” и “/temp/logs_messaggio_api_gate”

 ## Запуск отдельным приложением:
 ### Необходимые условия:
  - GO
  - Kafka
  - Postgres

  прописываем их настройки в файлах .env 
  или
  прописываем их настройки в файлах .env_example и затем запускаем 
  ```
  make set_env
  ```
  ### для запуска приложений:
   - зайти в messagio-test и запустить:
  ```
  make run
  ```

   - зайти в second_hadler и запустить:
  ```
  make run
  ```
   make run - это обертка над go run

### для компиляции:
в тех же папках запускаем
```
make build
```
  