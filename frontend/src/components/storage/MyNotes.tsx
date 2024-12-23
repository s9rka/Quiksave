import { useDeleteNote } from "@/services/mutations";
import { useNotes } from "@/services/queries";

export const Notes = () => {
  const notesQuery = useNotes();
  const deleteNoteMutation = useDeleteNote();

  const handleDeleteNote = async (id: number) => {
    if (!id) {
      console.error("Invalid note ID:", id);
      return;
    }

    try {
      await deleteNoteMutation.mutateAsync(id);
      console.log(`Note with ID ${id} deleted successfully!`);
      
    } catch (error) {
      console.error("Error deleting note:", error);
    }
  };

  if (notesQuery.isPending) return <p>Loading notes...</p>;
  if (notesQuery.isError)
    return <p>Error loading notes: {notesQuery.error.message}</p>;

  const notes = notesQuery.data || [];

  return (
    <>
      <p>Fetch status (function status): {notesQuery.fetchStatus}</p>
      <p>Query data status: {notesQuery.status}</p>
      <ul>
        {notes.length === 0 && <p>No notes available</p>}
        {notes.map((note) => (
          <li key={note.id}>
            <h1>{note.id}</h1>
            <h3>{note.title}</h3>
            <p>{note.content}</p>
            <p>Created At: {note.created_at}</p>
            <p>Tags: {note.tags.join(", ")}</p>
            <button onClick={() => handleDeleteNote(note.id)}>Delete</button>
          </li>
        ))}
      </ul>
    </>
  );
};
