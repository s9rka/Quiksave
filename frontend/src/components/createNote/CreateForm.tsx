import { Note } from "@/lib/types";
import { useCreateNote } from "@/services/mutations";
import { SubmitHandler, useForm } from "react-hook-form";
import { Input } from "../ui/input";

const CreateForm = () => {
  const { register, handleSubmit } = useForm<Note>();
  // const [tags, setTags] = useState<string[]>([])

  const createNoteMutation = useCreateNote();

  const handleCreateNote: SubmitHandler<Note> = (data) => {
    createNoteMutation.mutate(data);
    // createNoteMutation.mutate({ ...data, tags }); // Send tags array
  };
  return (
    <div>
      <h1>Create Note</h1>
      <form onSubmit={handleSubmit(handleCreateNote)}>
        <Input placeholder="Title" {...register("title")} />
        <Input placeholder="Content" {...register("content")} />

        <Input
          type="submit"
          disabled={createNoteMutation.isPending}
          value={createNoteMutation.isPending ? "Creating..." : "Create note"}
        />
      </form>
    </div>
  );
};

export default CreateForm;
