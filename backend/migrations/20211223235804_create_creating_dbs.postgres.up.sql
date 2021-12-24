CREATE TABLE IF NOT EXISTS "meal" (
                        "id" SERIAL PRIMARY KEY,
                        "name" varchar(70) UNIQUE NOT NULL,
                        "created_at" timestamp NOT NULL DEFAULT (now()),
                        "deleted_at" timestamp NOT NULL DEFAULT (now()),
                        "updated_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE IF NOT EXISTS "sandwich" (
                            "id" SERIAL PRIMARY KEY,
                            "name" varchar(70) UNIQUE NOT NULL,
                            "created_at" timestamp NOT NULL DEFAULT (now()),
                            "deleted_at" timestamp NOT NULL DEFAULT (now()),
                            "updated_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE IF NOT EXISTS "meal_sandwichs" (
                                  "id" SERIAL PRIMARY KEY,
                                  "meal_id" integer NOT NULL,
                                  "sandwich_id" integer NOT NULL,
                                  "created_at" timestamp NOT NULL DEFAULT (now()),
                                  "deleted_at" timestamp NOT NULL DEFAULT (now()),
                                  "updated_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE IF NOT EXISTS "drink" (
                         "id" SERIAL PRIMARY KEY,
                         "name" varchar(70) UNIQUE NOT NULL,
                         "created_at" timestamp NOT NULL DEFAULT (now()),
                         "deleted_at" timestamp NOT NULL DEFAULT (now()),
                         "updated_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE IF NOT EXISTS "meal_drink" (
                              "id" SERIAL PRIMARY KEY,
                              "meal_id" integer NOT NULL,
                              "drink_id" integer NOT NULL,
                              "created_at" timestamp NOT NULL DEFAULT (now()),
                              "deleted_at" timestamp NOT NULL DEFAULT (now()),
                              "updated_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE IF NOT EXISTS "dessert" (
                           "id" SERIAL PRIMARY KEY,
                           "name" varchar(70) UNIQUE NOT NULL,
                           "created_at" timestamp NOT NULL DEFAULT (now()),
                           "deleted_at" timestamp NOT NULL DEFAULT (now()),
                           "updated_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE IF NOT EXISTS "meal_desserts" (
                                 "id" SERIAL PRIMARY KEY,
                                 "meal_id" integer NOT NULL,
                                 "dessert_id" integer NOT NULL,
                                 "created_at" timestamp NOT NULL DEFAULT (now()),
                                 "deleted_at" timestamp NOT NULL DEFAULT (now()),
                                 "updated_at" timestamp NOT NULL DEFAULT (now())
);

ALTER TABLE "meal_sandwichs" ADD FOREIGN KEY ("meal_id") REFERENCES "meal" ("id");

ALTER TABLE "meal_sandwichs" ADD FOREIGN KEY ("sandwich_id") REFERENCES "sandwich" ("id");

ALTER TABLE "meal_drink" ADD FOREIGN KEY ("meal_id") REFERENCES "meal" ("id");

ALTER TABLE "meal_drink" ADD FOREIGN KEY ("drink_id") REFERENCES "drink" ("id");

ALTER TABLE "meal_desserts" ADD FOREIGN KEY ("meal_id") REFERENCES "meal" ("id");

ALTER TABLE "meal_desserts" ADD FOREIGN KEY ("dessert_id") REFERENCES "dessert" ("id");

CREATE INDEX ON "meal" ("name");

CREATE INDEX ON "sandwich" ("name");

CREATE UNIQUE INDEX ON "meal_sandwichs" ("meal_id", "sandwich_id");

CREATE INDEX ON "drink" ("name");

CREATE UNIQUE INDEX ON "meal_drink" ("meal_id", "drink_id");

CREATE INDEX ON "dessert" ("name");

CREATE UNIQUE INDEX ON "meal_desserts" ("meal_id", "dessert_id");