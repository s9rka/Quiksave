import { Card } from "../ui/card";
import { Badge } from "../ui/badge";
import WarningDialogue from "../layout/WarningDialogue";
import { filterValidTags } from "@/utils/utils";
import { Note } from "@/lib/types";

interface NoteCardProps {
    note: Note,
    openNote: (id: number) => void,
    handleDeleteNote: (id: number) => void,
}

export const NoteCard = ({ note, openNote, handleDeleteNote }:NoteCardProps) => {
  const validTags = filterValidTags(note.tags);

  return (
    <Card className="relative p-2 mb-2 bg-white/50">
      <div className="px-2" onClick={() => openNote(note.id)}>
        <h3 className="my-1">{note.heading}</h3>
        <p
          className="line-clamp-2 whitespace-pre-line"
          title={note.content}
        >
          {note.content}
        </p>
      </div>
      <div className="flex gap-2 items-center justify-start mt-2">
        {validTags.length > 0 && (
          validTags.map((tag) => (
            <Badge key={tag} variant="secondary">
              #{tag.trim()}
            </Badge>
          ))
        )}
      </div>
      <WarningDialogue
        noteID={note.id}
        deleteNote={handleDeleteNote}
      />
    </Card>
  );
};