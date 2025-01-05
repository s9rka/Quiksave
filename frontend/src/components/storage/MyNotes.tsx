import { useDeleteNote } from "@/services/mutations";
import { useNotes, useTags } from "@/services/queries";
import { useNavigate } from "react-router-dom";
import { NoNotesPlaceholder } from "@/components/storage/NoNotes";
import { NoteCard } from "@/components/storage/NoteCard";
import { Tag } from "@/lib/types";
import { useState } from "react";

export const Notes = () => {
  const notesQuery = useNotes();
  const tagsQuery = useTags();
  const deleteNoteMutation = useDeleteNote();
  const navigate = useNavigate();

  const [selectedTag, setSelectedTag] = useState<string | null>(null);


  const openNote = (id: number) => navigate(`/note/${id}`);

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

  if (notesQuery.isPending || tagsQuery.isPending) return <p>Loading...</p>;
  if (notesQuery.isError)
    return <p>Error loading notes: {notesQuery.error.message}</p>;
  if (tagsQuery.isError)
    return <p>Error loading tags: {tagsQuery.error.message}</p>;

  const notes = notesQuery.data || [];
  const tags = tagsQuery.data || [];

  const filteredNotes = selectedTag
    ? notes.filter((note) => note.tags.includes(selectedTag))
    : notes;


  return (
    
      <div>
      {tags.map((tag: Tag) => (
          <button
            key={tag}
            className={`tag ${selectedTag === tag ? "bg-slate-400" : ""}`}
            onClick={() => setSelectedTag(tag)}
          >
            {tag}
          </button>
        ))}
      
      <ul>
      {filteredNotes.length === 0 ? (
        <NoNotesPlaceholder />
      ) : (
        filteredNotes.map((note) => (
          <li key={note.id}>
            <NoteCard
              note={note}
              openNote={openNote}
              handleDeleteNote={handleDeleteNote}
            />
          </li>
        ))
      )}
    </ul>
    </div>
  );
};