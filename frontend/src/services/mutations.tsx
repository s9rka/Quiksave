import { useMutation, useQueryClient } from "@tanstack/react-query";
import { createNote, deleteNote, login, register } from "./api";
import { Note } from "@/lib/models";
import { useSetAtom } from "jotai/react";
import { authTokenAtom, usernameAtom } from "@/store/auth";
import { useNavigate } from "react-router-dom";

export const useCreateNote = () => {
  const queryClient = useQueryClient();
  return useMutation({
    mutationFn: (data: Note) => createNote(data),
    onMutate: () => {
      console.log("mutate");
    },
    onError: () => {
      console.log("error");
    },

    onSuccess: () => {
      console.log("success");
    },
    onSettled: async (_, error) => {
      console.log("settled");
      if (error) {
        console.log(error);
      } else {
        await queryClient.invalidateQueries({ queryKey: ["notes"] });
      }
    },
  });
};

export const useDeleteNote = () => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (id: number) => deleteNote(id),
    onSuccess: () => {
      console.log("deleted successfully");
    },
    onSettled: async (_, error) => {
      if (error) {
        console.log(error);
      } else {
        await queryClient.invalidateQueries({ queryKey: ["notes"] });
      }
    },
  });
};

export const useLogin = () => {
  const navigate = useNavigate();
  const setAuthToken = useSetAtom(authTokenAtom);
  const setUsername = useSetAtom(usernameAtom);
  return useMutation({
    mutationFn: login,
    onSuccess: (data, variables) => {
      setAuthToken(data);
      setUsername(variables.username);
      console.log("Login success, authToken: ", data);
      navigate(`/${variables.username}`);

    },
    onError: (error) => {
      console.error("Login failed:", error.message);
    },
  });
};

export const useRegister = () => {
  const setAuthToken = useSetAtom(authTokenAtom);
  const setUsername = useSetAtom(usernameAtom);
  return useMutation({
    mutationFn: register,
    onSuccess: (data, variables) => {
      setAuthToken(data);
      setUsername(variables.username);
      console.log("Register success, logging in..., authToken: ", data);
    },
    onError: (error) => {
      console.error("Registering failed:", error.message);
    },
  });
};
