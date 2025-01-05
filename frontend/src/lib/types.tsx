export interface CreateNote {
  heading: string;
  content: string;
  tags: string[];
}

export interface Note extends CreateNote {
  id: number;               
  created_at: string;
  last_edit: string;
}

export interface LoginCredentials {
  username: string;
  password: string;
}

export interface RegisterCredentials {
  username: string;
  email: string;
  password: string;
}

export interface AuthResponse {
  accessToken: string;
}

export interface User {
  id: number;
    username: string;
}

export type Tag = string;
