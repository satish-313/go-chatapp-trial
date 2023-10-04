CREATE TABLE "chatusers" (
    "id" serial  PRIMARY KEY,
    "username" varchar UNIQUE NOT NUll,
    "email" varchar UNIQUE NOT NUll,
    "password" varchar  NOT NUll
);