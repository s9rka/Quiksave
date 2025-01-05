import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@/components/ui/dialog";
import { Button } from "../ui/button";
import { Ellipsis, Trash } from "lucide-react";
import { Popover, PopoverContent, PopoverTrigger } from "../ui/popover";

interface WarningProps {
  noteID: number;
  deleteNote: (noteID: number) => void;
}

const WarningDialogue = ({ noteID, deleteNote }: WarningProps) => {
  return (
    <div>
      <Dialog >
        <Popover >
          <PopoverTrigger className="absolute top-2 right-4">
            <Ellipsis className="text-stone-500" />
          </PopoverTrigger>
          <PopoverContent><DialogTrigger className="p-2 bg-red-50"><Trash/></DialogTrigger></PopoverContent>
        </Popover>
        
        <DialogContent>
          <DialogHeader>
            <DialogTitle>Are you absolutely sure?</DialogTitle>
            <DialogDescription>
              This action cannot be undone. This will permanently delete your
              note and remove your data from our servers.
            </DialogDescription>
          </DialogHeader>
          <Button onClick={() => deleteNote(noteID)} variant="ghost">
            Delete
          </Button>
        </DialogContent>
      </Dialog>
    </div>
  );
};

export default WarningDialogue;
