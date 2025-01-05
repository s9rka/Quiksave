import { Rabbit } from "lucide-react";

export const NoNotesPlaceholder = () => (
  <div className="mt-16 px-20 py-10 select-none text-stone-700 flex flex-col items-center">
    <Rabbit className=" -ml-2 text-stone-500" />
    <h3>No notes yet</h3>
  </div>
);