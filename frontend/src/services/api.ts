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
