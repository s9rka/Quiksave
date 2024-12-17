import { useState } from "react";
import apiClient from "../services/api";

interface AuthResponse {
  accessToken: string;
}

interface LoginCredentials {
  username: string;
  password: string;
}

interface RegisterCredentials {
  username: string;
  email: string;
  password: string;
}

export const useAuthServices = () => {
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const login = async ({ username, password }: LoginCredentials): Promise<string | null> => {
    setLoading(true);
    setError(null);

    try {
      const response = await apiClient.post<AuthResponse>("/login", {
        username,
        password,
      });
      setLoading(false);
      return response.data.accessToken;
    } catch (err: any) {
      setError(err.response?.data?.message || "Login failed.");
      setLoading(false);
      return null;
    }
  };

  const register = async ({
    username,
    email,
    password,
  }: RegisterCredentials): Promise<string | null> => {
    setLoading(true);
    setError(null);

    try {
      const response = await apiClient.post<AuthResponse>("/register", {
        username,
        email,
        password,
      });
      setLoading(false);
      return response.data.accessToken;
    } catch (err: any) {
      setError(err.response?.data?.message || "Registration failed.");
      setLoading(false);
      return null;
    }
  };

  return { login, register, loading, error };
};
