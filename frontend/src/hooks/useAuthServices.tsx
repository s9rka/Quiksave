import { useMutation } from "@tanstack/react-query";
import { publicClient } from "../services/apiClient";
import { authTokenAtom, usernameAtom } from "@/store/auth";
import { useSetAtom } from "jotai/react";

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

const login = async ({ username, password }: LoginCredentials) => {
  const response = await publicClient.post<AuthResponse>("/login", {
    username,
    password,
  });
  return response.data;
};

const register = async ({
  username,
  email,
  password,
}: RegisterCredentials): Promise<AuthResponse> => {
  const response = await publicClient.post<AuthResponse>("/register", {
    username,
    email,
    password,
  });
  return response.data;
};

export const useAuthServices = () => {
  const setAuthToken = useSetAtom(authTokenAtom);
  const setUsername = useSetAtom(usernameAtom);
  
  // Login Mutation
  const loginMutation = useMutation({
    mutationFn: login,
    onSuccess: (data, variables) => {
      setAuthToken(data.accessToken);
      setUsername(variables.username);
    },
  });

  // Register Mutation
  const registerMutation = useMutation({
    mutationFn: register,
    onSuccess: (data) => {
      setAuthToken(data.accessToken);
    },
  });

  return { loginMutation, registerMutation };
};