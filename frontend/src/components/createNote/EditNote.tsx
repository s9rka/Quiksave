import { useParams } from "react-router-dom";
import CreateForm from "./CreateForm";
import { useNote } from "@/services/queries";

export const EditNotePage = () => {
    const { id } = useParams();
    const noteId = Number(id);

    const { data: note, isLoading, isError } = useNote(Number(noteId));
  
    if (isLoading) return <p>Loading note...</p>;
    if (isError) return <p>Error loading note.</p>;
  
    return <CreateForm initialNote={note} />;
  };