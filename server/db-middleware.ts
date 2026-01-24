import { db, type DatabaseClient } from "../db/client";
import { enhance, type UniversalMiddleware } from "@universal-middleware/core";

declare global {
  namespace Universal {
    interface Context {
      db: DatabaseClient;
    }
  }
}

/**
 * Add the `db` object to the context.
 */
export const dbMiddleware: UniversalMiddleware = enhance(
  async (_request, context, _runtime) => {
    return {
      ...context,
      db,
    };
  },
  {
    name: "my-app:db-middleware",
    immutable: false,
  },
);
