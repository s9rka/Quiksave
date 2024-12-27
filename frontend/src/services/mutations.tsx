import { useMutation, useQueryClient } from "@tanstack/react-query";
import {
  createNote,
  deleteNote,
  getUser,
  login,
  logout,
  register,
} from "./api";
import { Note } from "@/lib/types";

import { useNavigate } from "react-router-dom";
import { userAtom } from "@/context/UserContext";
import { useSetAtom } from "jotai";

export const useCreateNote = () => {
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
  const setUser = useSetAtom(userAtom)
  return useMutation({
    mutationFn: login,
    onSuccess: async () => {
      console.log("Login success!");
      const response = await getUser();
      if (response?.data) {
        setUser(response.data);
      }

    },
    onError: (error) => {
      console.error("Login failed:", error.message);
    },
  });
};

export const useRegister = () => {
  return useMutation({
    mutationFn: register,
    onSuccess: () => {
      console.log("Register success, logging in...");
    },
    onError: (error) => {
      console.error("Registering failed:", error.message);
    },
  });
};

export const useLogout = () => {
  const navigate = useNavigate();
  const setUser = useSetAtom(userAtom);

  return useMutation({
    mutationFn: logout,
    onSuccess: () => {
      console.log("User logged out successfully");
      setUser(null);
      navigate("/");
    },
    onError: (error) => {
      console.log("User logout error, ", error);
    },
  });
};
