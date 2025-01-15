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
    queryFn: () => getNoteById(id),
    enabled: !!id,
  });
};


export const useTags = () => {
  return useQuery({
    queryKey: ["tags"],
    queryFn: getTags
  })
}