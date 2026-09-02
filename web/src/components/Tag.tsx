import type { Tag } from "@/types/tag";
import type { UUID } from "@/types/uuid";
import { X } from "lucide-react";

export function TagsEditor({
  allTags,
  tags,
  setTags,
}: {
  allTags: Tag[];
  tags: UUID[];
  setTags: (tags: UUID[]) => void;
}) {
  return (
    <div className="flex gap-2 h-8">
      {tags.map((id) => (
        <div key={id} className="bg-primary-lightest p-1 rounded-xl">
          {allTags.find((t) => t.id === id)?.name || "[unknown]"}
          <X
            className="ml-0.5 inline align-sub cursor-pointer"
            size={18}
            onClick={() => setTags(tags.filter((t) => t != id))}
          />
        </div>
      ))}
      <select
        onChange={(e) => {
          setTags([...tags, e.target.value]);
        }}
        className="bg-primary-lightest p-1 rounded-xl"
      >
        <option value="default">Add Tag</option>
        {allTags
          .filter((t) => !tags.includes(t.id))
          .map((u) => (
            <option value={u.id} key={u.id}>
              {u.name}
            </option>
          ))}
      </select>
    </div>
  );
}
