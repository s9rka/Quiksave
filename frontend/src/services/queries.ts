import { useQuery } from "@tanstack/react-query"
import { getNoteIds, getNotes } from "./api"

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