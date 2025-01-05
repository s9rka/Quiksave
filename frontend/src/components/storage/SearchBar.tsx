import { useState } from "react";
import { Input } from "../ui/input";

interface SearchBarProps {
  onSearch: (query: string) => void;
}

export const SearchBar = ({ onSearch }: SearchBarProps) => {
  const [searchQuery, setSearchQuery] = useState("");

  const handleInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const value = e.target.value;
    setSearchQuery(value);
    onSearch(value);
  };

  return (
    <Input
      type="text"
      placeholder="Search notes..."
      value={searchQuery}
      onChange={handleInputChange}
      className="w-fit px-4 py-2 text-sm bg-transparent focus:outline-none"
    />
  );
};
