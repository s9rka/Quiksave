import { Note } from "@/lib/models";
import {privateClient} from "./apiClient";

export const getNoteIds = async () => {
  const response = (await privateClient.get<Note[]>("/get-notes")).data.map(
    (note) => note.id
  );

  return response;
};

export const getNotes = async () => {
  return (await privateClient.get<Note[]>("/get-notes")).data;
};

export const createNote = async (data: Note) => {
  await privateClient.post("/create-note", data)
}

export const deleteNote = async (id: number) => {
  await privateClient.delete(`/note/${id}`)
}