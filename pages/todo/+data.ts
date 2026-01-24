// https://vike.dev/data

import type { PageContextServer } from "vike/types";

export type Data = Awaited<ReturnType<typeof data>>;

export async function data(_pageContext: PageContextServer) {
  // あとで，Controller 層として書く
  // const todoItemsInitial = await kyselyQueries.getAllTodos(_pageContext.db);
  // return { todoItemsInitial };
}
