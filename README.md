## Инструкция по запуску
Можно склонировать репозиторий развернуть БД и сервер с готовыми настройками из .env локально (БД на порту 5432, сервер на порту 8080) при помощи:
```
docker compose build
docker compose up
```
Собирают приложение из двух контейнеров (сделал через compose специально для того, чтобы удобнее было локально развернуть, на гитлабе же подключение к внешней бд
и использование только одного Dockerfile для сервера). Параметры запуска указываются в .env файле.

На БД накатываются миграции, создающие все таблички, а также инсертятся все данные в таблички employee, organization, organization_responsible
(данные взял с БД развернутой на кластере).

Пара тонких моментов:
- в переменной POSTGRES_CONN хостом указывается psql, потому что такое имя контейнера при сборке и докер объединяет их в одну сеть, снаружи же все доступно по данным из .env.
- миграции локально накатываются, если в файле /app/migrate.go в переменной pgUrl отключается sslmode; для подключения к БД на кластере эту строку я убирал.
- в случае если надо развернуть не через compose, а как говорилось в задании, то нужно настроить .env для подключении к внешней БД, создать образ и запустить контейнер + убрать sslmode=disable (о чем говорилось выше) + возможно понадобится убрать из миграций создание таблиц, которые уже были в БД на кластере, а также вставку в них значений.
  
## Используемый стек
- Go
- PostgreSQL и pgx библиотека для работы с БД
- Фрэймворк Echo
- Для миграций golang-migrate
- go-playground/validator для валидации данных запросов

## Немного примеров запросов
- Создание тендера ![image](https://github.com/user-attachments/assets/ebd529e2-c0d1-4ca2-8d68-7bc22a49f2c2)
  
- Изменение статуса тендера ![image](https://github.com/user-attachments/assets/8ca91ed8-7f67-455b-aed9-3cb08afef8d5)

- Редактирование тендера (причем за другого ответственного) ![image](https://github.com/user-attachments/assets/19f3ca72-8289-4acd-bf27-42e68e3a93af)

## P.S.
- Задание было довольно интересное, иногда, правда, некоторые тонкие моменты реализации (особенно с правами доступа к сущностям) были не совсем ясны, делал по ощущениям. Поэтому большое спасибо за предоставленную возможность, надеюсь на успешный исход после проверки :)

- В целом это был, наверное, первый более менее объемный проект на Go, и даже после Python язык показался очень приятным, но из-за новизны некоторые моменты, которые можно улучшить, обнаруживались в
момент столкновения с последствиями, однако в будущем буду учитывать подобное.

- Постарался поддержать более менее чистую архитектуру, структуру директорий, возможность использования другой бд, другой реализации сервисов и подобного за счет использования интерфейсов, где это было уместно.

- Добавил Graceful Shutdown, чтобы дать сервису корректно завершить свою работу, освободив ресурсы.

- Как мне кажется, получилось неплохо настроить конфигурацию, отделив ее от непосредственно самой работы приложения

- Внедрил систему логов на сервере при помощи logrus (уровень логирования выставляется в конфиге), а также логирование запросов в файл при помощи функционала Echo.

- При большем объеме времени точно стоило лучше разобраться с валидатором, расширив его функционал под собственные нужды, а также воспользовавшись уже имевшимися возможностями, о которых я мог не знать.

Контакты:
- tg: https://t.me/yaQvadrat
- mail: yaqvadrat@mail.ru
