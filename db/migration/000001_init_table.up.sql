CREATE TABLE "customer_support" (
  "cs_id" UUID PRIMARY KEY NOT NULL,
  "user_id" UUID NOT NULL,
  "order_id" UUID NOT NULL,
  "subject" VARCHAR,
  "message" TEXT,
  "status" VARCHAR NOT NULL,
  "created_at" TIMESTAMP(0) NOT NULL DEFAULT NOW(),
  "updated_at" TIMESTAMP(0)
);
