import { AuthResponse, CreateNote, Note, Tag, Vault, VaultFormData } from "@/lib/types";
import { privateClient, publicClient } from "./apiClient";
import { LoginCredentials, RegisterCredentials } from "@/lib/types";

export const getNoteIds = async (vaultId: number) => {
  const response = (await privateClient.get<Note[]>(`/get-notes?vaultId=${vaultId}`)).data.map(
    (note) => note.id
  );

  return response;
};

export const getNotes = async (vaultId: number) => {
  return (await privateClient.get<Note[]>(`/get-notes?vaultId=${vaultId}`)).data;
};

export async function createNote(note: CreateNote) {
  return ((await privateClient.post("/create-note", note)).data);
}

export async function editNote(note: Note) {
  return await privateClient.put(`/note/${note.id}?vaultId=${note.vaultId}`, note);
}

export const getNoteById = async (id: number, vaultId: number): Promise<Note> => {
  const response = await privateClient.get(`/note/${id}?vaultId=${vaultId}`);
  return response.data;
};

export const deleteNote = async (id: number, vaultId: number) => {
  return (await privateClient.delete(`/note/${id}?vaultId=${vaultId}`));
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

export const createVault = async (vaultData: VaultFormData) => {
    const response = await privateClient.post<{ vaultID: number }>('/create-vault', vaultData);
    return response.data;
};

export const getVaults = async () => {
    const response = await privateClient.get<Vault[]>('/get-vaults');
    return response.data;
};

export const getVaultById = async (vaultId: number) => {
    const response = await privateClient.get<Vault>(`/vault/${vaultId}`);
    return response.data;
};