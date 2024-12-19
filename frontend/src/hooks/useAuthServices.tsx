import { useMutation } from "@tanstack/react-query";
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
  // Login Mutation
  const loginMutation = useMutation({
    mutationFn: async ({ username, password }: LoginCredentials): Promise<AuthResponse> => {
      const response = await apiClient.post<AuthResponse>("/login", { username, password });
      return response.data;
    },
  });

  // Register Mutation
  const registerMutation = useMutation({
    mutationFn: async ({ username, email, password }: RegisterCredentials): Promise<AuthResponse> => {
      const response = await apiClient.post<AuthResponse>("/register", {
        username,
        email,
        password,
      });
      return response.data;
    },
  });

  return { loginMutation, registerMutation };
};
