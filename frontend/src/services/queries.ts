import { useQuery, useQueryClient } from "@tanstack/react-query"
import { getNoteById, getNoteIds, getNotes, getTags, getVaults, getVaultById } from "./api"
import { Note, Vault } from "@/lib/types"

export const useNotesIds = (vaultId: number) => {
    return useQuery({
        queryKey: ["noteIds", vaultId],
        queryFn: () => getNoteIds(vaultId),
        enabled: !!vaultId,
        refetchOnReconnect: false,
    })
}

export const useNotes = (vaultId: number) => {
    return useQuery({
        queryKey: ["notes", vaultId],
        queryFn: () => getNotes(vaultId),
        enabled: !!vaultId
    })
}

export const useNote = (id: number, vaultId: number) => {
  return useQuery<Note, Error>({
    queryKey: ["notes", id, vaultId],
    queryFn: () => getNoteById(id, vaultId),
    enabled: !!id && !!vaultId,
  });
};

export const useTags = () => {
  return useQuery({
    queryKey: ["tags"],
    queryFn: getTags
  })
}

export const useVaults = () => {
    return useQuery<Vault[]>({
        queryKey: ['vaults'],
        queryFn: getVaults,
    });
};

export const useVault = (vaultId: number) => {
    return useQuery<Vault>({
        queryKey: ['vault', vaultId],
        queryFn: () => getVaultById(vaultId),
        enabled: !!vaultId,
    });
};