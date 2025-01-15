import { openDB, IDBPDatabase } from "idb";
import { Note } from "@/lib/types";

const DB_NAME = "myNotesDB";
const DB_VERSION = 1;
const STORE_NAME = "draftNotes";

async function getDB(): Promise<IDBPDatabase> {
  return  await openDB(DB_NAME, DB_VERSION, {
    upgrade(db) {
      if (!db.objectStoreNames.contains(STORE_NAME)) {
        db.createObjectStore(STORE_NAME, { keyPath: "id" });
      }
    },
  });
}

export async function saveDraft(
    id: number,
    data: Partial<Note>
  ): Promise<void> {
    const db = await getDB();
    await db.put(STORE_NAME, { id, ...data });
  }
  
  export async function getDraft(
    id: number
  ): Promise<Partial<Note> | undefined> {
    if (!id) return undefined;
    const db = await getDB();
    return (await db.get(STORE_NAME, id)) as Partial<Note> | undefined;
  }
  

export async function deleteDraft(id: number): Promise<void> {
  const db = await getDB();
  await db.delete(STORE_NAME, id);
}
