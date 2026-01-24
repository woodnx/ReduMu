import type { dbKysely } from "./db/client";

declare global {
  namespace Vike {
    interface PageContextServer {
      db: ReturnType<typeof dbKysely>;
    }
  }
}
