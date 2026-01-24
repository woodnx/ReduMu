import { enhance, type UniversalHandler } from "@universal-middleware/core";
import { z } from "zod";
import type { DatabaseClient } from "../db/client";

// Zodスキーマ定義: textは必須で、かつ1文字以上
const CreateTodoSchema = z.object({
  text: z.string().min(1, "Text is required"),
});

// コンテキストの型定義: db が注入されていることを保証
type ContextWithDb = Universal.Context & {
  db: DatabaseClient;
};

export const createTodoHandler: UniversalHandler<ContextWithDb> = enhance(
  async (request, context, _runtime) => {
    let body: unknown;
    try {
      body = await request.json();
    } catch {
      return new Response(JSON.stringify({ error: "Invalid JSON" }), {
        status: 400,
      });
    }

    const result = CreateTodoSchema.safeParse(body);

    if (!result.success) {
      return new Response(
        JSON.stringify({
          error: "Validation Failed",
          details: result.error.flatten(),
        }),
        {
          status: 400,
          headers: { "content-type": "application/json" },
        },
      );
    }

    // ここで newTodo は { text: string } 型として保証される
    const newTodo = result.data;

    // いずれ書く
    // await kyselyQueries.insertTodo(context.db, newTodo.text);

    return new Response(JSON.stringify({ status: "OK" }), {
      status: 200,
      headers: {
        "content-type": "application/json",
      },
    });
  },
  {
    name: "my-app:todo-handler",
    path: `/api/todo/create`,
    method: ["POST"],
    immutable: false,
  },
);
