export interface Note {
  id: number;
  title: string;
  content: string;
  created_at: string;
  tags: string[];
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