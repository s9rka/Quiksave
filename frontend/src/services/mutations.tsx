import { useMutation, useQueryClient } from "@tanstack/react-query";
import { createNote, deleteNote } from "./api";
import { Note } from "@/lib/models";

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
            console.log("deleted successfully")
        },
        onSettled: async (_, error) => {
            if (error) {
                console.log(error)
            } else {
                await queryClient.invalidateQueries({queryKey: ["notes"]})
            }

        }
    })
}