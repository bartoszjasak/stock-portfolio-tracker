CREATE TABLE public.users (
    id SERIAL PRIMARY KEY,
    email character varying(255),
    first_name character varying(255),
    last_name character varying(255),
    password character varying(60)
);
    
   INSERT INTO "public"."users"("email","first_name","last_name","password")
VALUES
	(E'bartoszjasak@gmail.com',E'Bartosz',E'Jasak',E'password');
	
CREATE TABLE public.transactions (
    id SERIAL PRIMARY KEY,
    type character varying(255),
    stock_name character varying(255),
    symbol character varying(255),
    price real,
    quantity integer NOT null, 
    date timestamp without time zone,
    user_id SERIAL references public.users(id)
);

   INSERT INTO "public"."transactions"("type", "stock_name", "symbol","price","quantity","date","user_id")
VALUES
	(E'BUY',E'Apple Inc.', E'AAPL', 146.55, 30, E'2022-03-14 00:00:00', 1);
	
   INSERT INTO "public"."transactions"("type", "stock_name","symbol","price","quantity","date","user_id")
VALUES
	(E'BUY',E'Apple Inc.',E'AAPL', 156.55, 10, E'2022-03-14 00:00:00', 1);

   INSERT INTO "public"."transactions"("type", "stock_name","symbol","price","quantity","date","user_id")
VALUES
	(E'BUY',E'Microsoft Inc.',E'MSFT', 246.55, 35, E'2022-03-14 00:00:00', 1);
	
