#Messenger

Функционал:
1) Валидация токена/сеанса
2) Регистрация
3) Авторизация
4) Отправка сообщения



API (REST API)

1) Middleware (Token validation)
   Request:
   Response:

2) Регистрация
   POST ${address}/registration
   Request:
   {
   "login" string;
   "password" : string;
   "nickname" : string;
   }
   Response:
   {
   "id" : int,
   }

3) Авторизация
   POST $(address)/login
   Request:
   {
   "login" :string;
   "password":string;
   }

Response:

    {
    "token":string
    }

4) Отправка сообщения
   POST $(adress)/chat
   Request:
   {
   "message" string
   }
   Response: