import { DropdownMenu, DropdownMenuContent, DropdownMenuItem, DropdownMenuTrigger } from "@/components/ui/dropdown-menu";
import { Tag } from "@/lib/types";
import { ChevronDown } from "lucide-react";

interface TagsDropdownProps {
  tags: Tag[];
  selectedTag: string | null;
  setSelectedTag: (tag: string | null) => void;
}

export const TagsDropdown = ({ tags, selectedTag, setSelectedTag }: TagsDropdownProps) => {
  return (
    <DropdownMenu>
      <DropdownMenuTrigger asChild>
        <button className="px-4 py-2 border rounded-md flex gap-2">
          #{selectedTag || "All"}
          <ChevronDown className="w-4 text-stone-600"/>
        </button>
      </DropdownMenuTrigger>
      <DropdownMenuContent className="ml-4 ">
        <DropdownMenuItem onClick={() => setSelectedTag(null)}>
          #All
        </DropdownMenuItem>
        {tags.map((tag) => (
          <DropdownMenuItem key={tag} onClick={() => setSelectedTag(tag)}>
            #{tag}
          </DropdownMenuItem>
        ))}
      </DropdownMenuContent>
    </DropdownMenu>
  );
};
