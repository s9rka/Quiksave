import { useMutation, useQueryClient } from "@tanstack/react-query";
import { CreateNote, Note } from "@/lib/types";
import { useNavigate } from "react-router-dom";
import { userAtom } from "@/context/UserContext";
import { useSetAtom } from "jotai";
import {
  createNote,
  editNote,
  deleteNote,
  getUser,
  login,
  logout,
  register,
} from "./api";

export const useCreateNote = () => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (data: CreateNote) => createNote(data),
    onMutate: () => {
      console.log("mutate");
    },
    onError: () => {
      console.log("error");
    },

    onSuccess: () => {
      console.log("success");
      queryClient.invalidateQueries({ queryKey: ["notes"] });
      queryClient.invalidateQueries({ queryKey: ["tags"] });
    },
    onSettled: async (_, error) => {
      console.log("settled");
      if (error) {
        console.log(error);
      }
    },
  });
};

export const useUpdateNote = () => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: async (data: Note) => {
      const response = await editNote(data);
      return response;
    },
    onMutate: () => {
      console.log("update mutate");
    },
    onError: () => {
      console.log("update error");
    },
    onSuccess: () => {
      console.log("update success");
      queryClient.invalidateQueries({ queryKey: ["notes"] });
      queryClient.invalidateQueries({ queryKey: ["tags"] });
    },
    onSettled: async (_, error) => {
      console.log("update settled");
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
        queryClient.invalidateQueries({ queryKey: ["tags"] });
      }
    },
  });
};

export const useLogin = () => {
  const setUser = useSetAtom(userAtom);
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
