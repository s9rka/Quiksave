import { useQuery, useQueryClient } from "@tanstack/react-query"
import { getNoteById, getNoteIds, getNotes, getTags } from "./api"
import { Note } from "@/lib/types"

export const useNotesIds = () => {
    return useQuery({
        queryKey: ["noteIds"],
        queryFn: getNoteIds,
        refetchOnReconnect: false,
    })
}

export const useNotes = () => {
    return useQuery({
        queryKey: ["notes"],
        queryFn: getNotes
    })
}


export const useNote = (id: number) => {
  return useQuery<Note, Error>({
    queryKey: ["notes", id],
    queryFn: () => getNoteById(id), // Pass a function, not the result of a function call
    enabled: !!id, // Ensure the query only runs if `id` is valid
  });
};


export const useTags = () => {
  return useQuery({
    queryKey: ["tags"],
    queryFn: getTags
  })
}