import { useNotes } from "@/services/queries";

export const Notes = () => {
  const notesQuery = useNotes()
  
  if (notesQuery.isPending) return <p>Loading notes...</p>;
  if (notesQuery.isError)
    return <p>Error loading notes: {notesQuery.error.message}</p>;

  return (
    <ul>
      {notesQuery.isFetched && !notesQuery.data && <p>no notes</p>}
      {notesQuery.data?.map((note) => (
        <li key={note.id}>
          {note.title}
          <p>{note.content}</p>
          <p>{note.created_at}</p>
        </li>
      ))}
    </ul>
  );
};
