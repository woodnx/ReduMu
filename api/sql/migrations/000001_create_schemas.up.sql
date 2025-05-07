CREATE TABLE "user" (
    "id" UUID PRIMARY KEY,
    "name" VARCHAR(20) NOT NULL,
    "password" VARCHAR(80) NOT NULL,
    "role" VARCHAR(80) NOT NULL,
    "created" TIMESTAMP(6) NOT NULL,
    "modified" TIMESTAMP(6) NOT NULL,
    CONSTRAINT uix_name UNIQUE ("name")
);

CREATE TABLE "task" (
    "id" UUID PRIMARY KEY,
    "user_id" UUID NOT NULL,
    "title" VARCHAR(128) NOT NULL,
    "status" VARCHAR(20) NOT NULL,
    "deadline" TIMESTAMP(6) NOT NULL,
    "created" TIMESTAMP(6) NOT NULL,
    "modified" TIMESTAMP(6) NOT NULL,
    CONSTRAINT fk_user_id FOREIGN KEY ("user_id")
        REFERENCES "user" (id)
        ON DELETE RESTRICT
        ON UPDATE RESTRICT
);
