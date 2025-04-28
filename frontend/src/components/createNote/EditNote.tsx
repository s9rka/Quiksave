import { useParams } from "react-router-dom";
import CreateForm from "./CreateNote";
import { useNote } from "@/services/queries";
import { useVault } from "@/context/VaultContext";

export const EditNotePage = () => {
    const { noteId, vaultId } = useParams<{ noteId: string; vaultId: string }>();
    const noteIdNum = Number(noteId);
    const vaultIdNum = Number(vaultId);

    if (!vaultIdNum) {
        return <p>No vault selected</p>;
    }

    const { data: note, isLoading, isError } = useNote(noteIdNum, vaultIdNum);
  
    if (isLoading) return <p>Loading note...</p>;
    if (isError) return <p>Error loading note.</p>;
  
    return <CreateForm initialNote={note} />;
};