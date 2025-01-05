import { AuthResponse, CreateNote, Note, Tag } from "@/lib/types";
import { privateClient, publicClient } from "./apiClient";
import { LoginCredentials, RegisterCredentials } from "@/lib/types";

export const getNoteIds = async () => {
  const response = (await privateClient.get<Note[]>("/get-notes")).data.map(
    (note) => note.id
  );

  return response;
};

export const getNotes = async () => {
  return (await privateClient.get<Note[]>("/get-notes")).data;
};

export async function createNote(note: CreateNote) {
  return ((await privateClient.post("/create-note", note)).data);
}

export async function editNote(note: Note) {
  return await privateClient.put(`/note/${note.id}`, note);
}

export const getNoteById = async (id: number): Promise<Note> => {
  const response = await privateClient.get(`/note/${id}`);
  return response.data;
};


export const deleteNote = async (id: number) => {
  return (await privateClient.delete(`/note/${id}`));
};

export const login = async ({ username, password }: LoginCredentials) => {
  const response = await publicClient.post<AuthResponse>("/login", {
    username,
    password,
  });
  return response.status;
};

export const register = async ({username, email, password}: RegisterCredentials) => {
  const response = await publicClient.post<AuthResponse>("/register", {
    username, email, password
  });
  return response.data.accessToken;
}

export const logout = async () => {
  return await privateClient.post("/logout")
}

export const getUser = async () => {
  return await privateClient.get("/me")
}

export const getTags = async (): Promise<Tag[]> => {
  return (await privateClient.get('/tags')).data
}