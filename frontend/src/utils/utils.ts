import { Note } from "@/lib/types";

export function hasChanges(original: Note, updated: Partial<Note>) {
  if (original.heading !== updated.heading) return true;
  if (original.content !== updated.content) return true;

  const originalTags = (original.tags || []).join(",");
  const updatedTags = (updated.tags || []).join(",");
  if (originalTags !== updatedTags) return true;

  return false;
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
        (tag) =>
          typeof tag === "string" && tag.trim() !== "#" && tag.trim() !== ""
      )
    : [];
};

export function extractTagsFromContent(text: string): string[] {
  const tagRegex = /#\w+/g;
  const matches = text.match(tagRegex) || [];
  return matches.map((tag) => tag.slice(1)); // remove the '#' character
}

