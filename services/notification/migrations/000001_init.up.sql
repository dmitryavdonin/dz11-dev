CREATE TABLE "notification" (
    "id" serial primary key,
    "user_id" integer not null,
    "order_id" integer not null,
    "message" varchar,
    "created_at" timestamp not null,
    "modified_at" timestamp not null
);