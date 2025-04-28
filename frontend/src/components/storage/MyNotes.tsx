import { useDeleteNote } from "@/services/mutations";
import { useNotes, useTags } from "@/services/queries";
import { useNavigate, useParams } from "react-router-dom";
import { NoNotesPlaceholder } from "@/components/storage/NoNotes";
import { NoteCard } from "@/components/storage/NoteCard";
import { TagsDropdown } from "@/components/storage/TagsDropdown";
import { useState } from "react";
import { SearchBar } from "./SearchBar";

export const Notes = () => {
  const params = useParams();
  const vaultId = params.id;
  const notesQuery = useNotes(Number(vaultId));
  const tagsQuery = useTags();
  const deleteNoteMutation = useDeleteNote();
  const navigate = useNavigate();

  const [selectedTag, setSelectedTag] = useState<string | null>(null);
  const [searchQuery, setSearchQuery] = useState("");

  const openNote = (id: number) => navigate(`/vault/${vaultId}/note/${id}`);

  const handleDeleteNote = async (id: number) => {
    if (!id || !vaultId) {
      console.error("Invalid note ID or vault ID:", id, vaultId);
      return;
    }

    try {
      await deleteNoteMutation.mutateAsync({ id, vaultId: Number(vaultId) });
      console.log(`Note with ID ${id} deleted successfully!`);
    } catch (error) {
      console.error("Error deleting note:", error);
    }
  };

  if (!vaultId) {
    return <p>No vault selected</p>;
  }

  if (notesQuery.isPending || tagsQuery.isPending) return <p>Loading...</p>;
  if (notesQuery.isError)
    return <p>Error loading notes: {notesQuery.error.message}</p>;
  if (tagsQuery.isError)
    return <p>Error loading tags: {tagsQuery.error.message}</p>;

  const notes = notesQuery.data || [];
  const tags = tagsQuery.data || [];

  const filteredNotes = notes
    .filter((note) => (selectedTag ? note.tags.includes(selectedTag) : true))
    .filter((note) =>
      note.content.toLowerCase().includes(searchQuery.toLowerCase())
    );

  return (
    <div>
      <div className="flex items-center justify-between pt-8">
        <TagsDropdown
          tags={tags}
          selectedTag={selectedTag}
          setSelectedTag={setSelectedTag}
        />
        <SearchBar onSearch={setSearchQuery} />
      </div>
      <ul className="py-4">
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
