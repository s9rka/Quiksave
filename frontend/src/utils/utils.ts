import { Note } from "@/lib/types";

export function hasChanges(oldData: Partial<Note>, newData: Partial<Note>): boolean {
    return (
      oldData.heading !== newData.heading ||
      oldData.content !== newData.content ||
      JSON.stringify(oldData.tags || []) !== JSON.stringify(newData.tags || [])
    );
  }
  
  export const formatDate = (isoString: string) => {
    const date = new Date(isoString);
    return new Intl.DateTimeFormat("en-US", {
      dateStyle: "medium",
      timeStyle: "short",
    }).format(date);
  };

  export const filterValidTags = (tags: any[]): string[] => {
    return Array.isArray(tags)
      ? tags.filter(
          (tag) => typeof tag === "string" && tag.trim() !== "#" && tag.trim() !== ""
        )
      : [];
  };