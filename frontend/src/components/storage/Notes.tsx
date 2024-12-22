import { useDeleteNote } from "@/services/mutations";
import { useNotes } from "@/services/queries";

export const Notes = () => {
  const notesQuery = useNotes();
  const deleteNoteMutation = useDeleteNote();

  const handleDeleteNote = async (id: number) => {
    await deleteNoteMutation.mutateAsync(id);
    console.log("deleted successfully!")
  };

  if (notesQuery.isPending) return <p>Loading notes...</p>;
  if (notesQuery.isError)
    return <p>Error loading notes: {notesQuery.error.message}</p>;

  return (
    <>
      Fetchstatus (function status): {notesQuery.fetchStatus}
      <p>Query data status: {notesQuery.status}</p>
      <ul>
        {notesQuery.isFetched && !notesQuery.data && <p>No notes</p>}
        {notesQuery.data?.map((note) => (
          <li key={note.id}>
            {note.title}
            <p>{note.content}</p>
            <p>{note.created_at}</p>
            <p>{note.tags.join(", ")}</p>
              <button onClick={() => handleDeleteNote(note.id!)}>Delete</button>
          </li>
        ))}
      </ul>
    </>
  );
};
