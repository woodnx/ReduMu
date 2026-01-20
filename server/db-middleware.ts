import { dbKysely } from "../database/kysely/db";
import { enhance, type UniversalMiddleware } from "@universal-middleware/core";

declare global {
  namespace Universal {
    interface Context {
      db: ReturnType<typeof dbKysely>;
    }
  }
}

// Note: You can also directly use your server, instead of defining a universal middleware. (Vike's scaffolder uses https://github.com/magne4000/universal-middleware to simplify its internal logic.)
/**
 * Add the `db` object to the context.
 */
export const dbMiddleware: UniversalMiddleware = enhance(
  // The context we add here is automatically merged into pageContext
  async (_request, context, _runtime) => {
    const db = dbKysely();

    return {
      ...context,
      // Sets pageContext.db
      db: db,
    };
  },
  {
    name: "my-app:db-middleware",
    immutable: false,
  },
);
