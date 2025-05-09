import { Note, CreateNote } from "@/lib/types";
import { useCreateNote, useUpdateNote } from "@/services/mutations";
import { useEffect, useRef, useState } from "react";
import { useForm } from "react-hook-form";
import { hasChanges } from "@/utils/utils";
import { Label } from "../ui/label";
import logo from "@/assets/logo.svg";
import { Loader2 } from "lucide-react";
import { useVault } from "@/context/VaultContext";

type CreateFormProps = {
  initialNote?: Note;
};

const CreateForm = ({ initialNote }: CreateFormProps) => {
  const { vaultId } = useVault();

  if (!vaultId) {
    return <div>No vault selected</div>;
  }

  const { register, watch, setValue } = useForm<Partial<Note>>({
    defaultValues: initialNote ?? {
      heading: "",
      content: "",
      tags: [],
      vaultId,
    },
  });

  const createNoteMutation = useCreateNote();
  const updateNoteMutation = useUpdateNote();

  const debounceTimerRef = useRef<NodeJS.Timeout | null>(null);
  const [lastSaveTime, setLastSaveTime] = useState<number>(Date.now());
  const [previousData, setPreviousData] = useState<Note | undefined>(
    initialNote
  );

  const heading = watch("heading");
  const content = watch("content") || "";

  const extractTags = (text: string): string[] => {
    const tagRegex = /#\w+/g;
    if (!tagRegex.test(text)) return [];
    return Array.from(text.match(tagRegex) || []).map((tag) => tag.slice(1)); // Remove `#`
  };

  const isBlank = () => !heading?.trim() && !content.trim();
  const isEditMode = !!previousData?.id;
  const [isSaving, setIsSaving] = useState(false);
  const [hasUnsavedChanges, setHasUnsavedChanges] = useState(false);

  useEffect(() => {
    if (isBlank()) {
      setHasUnsavedChanges(false);
      return;
    }

    const tags = extractTags(content);
    if (tags.length > 0) {
      setValue("tags", tags);
    }

    const hasModified = !previousData || hasChanges(previousData, { heading, content, tags });
    setHasUnsavedChanges(hasModified);

    if (debounceTimerRef.current) {
      clearTimeout(debounceTimerRef.current);
    }

    debounceTimerRef.current = setTimeout(() => {
      autoSave();
    }, 2000);

    const now = Date.now();
    if (now - lastSaveTime > 20000) {
      if (debounceTimerRef.current) {
        clearTimeout(debounceTimerRef.current);
      }
      autoSave();
    }
  }, [heading, content]);

  const autoSave = () => {
    if (isBlank()) return;

    const tags = extractTags(content);
    const currentData = {
      ...previousData,
      heading,
      content,
      tags,
      vaultId,
    } as Partial<Note>;

    if (previousData && !hasChanges(previousData, currentData)) {
      return;
    }

    setIsSaving(true); // Start saving
    if (isEditMode) {
      updateNoteMutation.mutate(currentData as Note, {
        onSuccess: () => {
          setPreviousData(currentData as Note);
          setLastSaveTime(Date.now());
          setHasUnsavedChanges(false);
        },
        onSettled: () => {
          setIsSaving(false); // Stop saving
        },
      });
    } else {
      createNoteMutation.mutate(currentData as CreateNote, {
        onSuccess: (response) => {
          if (response?.noteID) {
            const newNoteData = {
              ...currentData,
              id: response.noteID,
            } as Note;

            setValue("id", response.noteID);
            setPreviousData(newNoteData);
          }
          setLastSaveTime(Date.now());
          setHasUnsavedChanges(false);
        },
        onSettled: () => {
          setIsSaving(false); // Stop saving
        },
      });
    }
  };

  return (
    <div className="mx-auto max-w-md px-4 py-10 max-h-screen">
      <div className="flex flex-row justify-between items-center">
        <img className="h-6 mb-2" src={logo} alt="logo" />
        <div className="flex items-center space-x-2">
          {isSaving && (
            <Loader2 className="h-5 w-5 animate-spin text-green-500" />
          )}
          <Label
            className={`px-2 py-1 rounded text-sm ${
              isSaving
                ? "text-stone-800"
                : hasUnsavedChanges
                ? "text-red-800"
                : "text-green-800"
            }`}
          >
            {isSaving
              ? "Saving..."
              : hasUnsavedChanges
              ? "Unsaved Changes"
              : "Saved"}
          </Label>
        </div>
      </div>
      <form className="flex flex-col space-y-4">
        <div className="flex justify-between items-center">
          <input
            type="text"
            placeholder="Heading"
            {...register("heading")}
            className="bg-transparent focus:outline-none text-lg placeholder-gray-500"
          />
        </div>
        <textarea
          placeholder="Write here (use #tag to add tags)"
          {...register("content")}
          rows={16}
          className="w-full bg-transparent focus:outline-none resize-none text-base placeholder-gray-500"
        />
        <div className="flex flex-wrap space-x-2">
          {watch("tags")?.map((tag, index) => (
            <Label
              key={index}
              className="bg-blue-100 text-blue-800 px-2 py-1 rounded text-sm"
            >
              #{tag}
            </Label>
          ))}
        </div>
      </form>
    </div>
  );
};

export default CreateForm;
