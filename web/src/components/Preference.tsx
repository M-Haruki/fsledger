import { Plus } from "lucide-react";

export function AddBtn({
  className = "",
  onAdd,
}: {
  className: string;
  onAdd: () => void;
}) {
  return (
    <button
      onClick={onAdd}
      className={`rounded-xl p-1 boder bg-primary-lighter hover:bg-primary-light ${className}`}
    >
      <Plus size={30} strokeWidth={3} />
    </button>
  );
}
